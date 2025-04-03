package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type VirtualMachineRequest struct {
	Name   string `json:"name"`
	Cores  int    `json:"cores"`
	Ram    int    `json:"ram"`
	OS     string `json:"os"`
	Status string `json:"status"`
}

const (
	StatusRunning  = "running"
	StatusStopped  = "stopped"
	StatusPaused   = "paused"
	StatusFailed   = "failed"
	StatusCreating = "creating"
)

var ValidStatuses = []string{
	StatusRunning,
	StatusStopped,
	StatusPaused,
	StatusFailed,
	StatusCreating,
}

func (vm *VirtualMachineRequest) Validate() error {
	if strings.TrimSpace(vm.Name) == "" {
		return errors.New("Name is required")
	}

	if vm.Cores <= 0 {
		return errors.New("Cores must be a positive number")
	}

	if vm.Ram <= 0 {
		return errors.New("Ram must be a positive number")
	}

	if strings.TrimSpace(vm.OS) == "" {
		return errors.New("OS is required")
	}

	if strings.TrimSpace(vm.Status) == "" {
		return errors.New("Status is required")
	}

	if !isValidStatus(vm.Status) {
		return fmt.Errorf("invalid status: %s. Valid statuses are: %s",
			vm.Status, strings.Join(ValidStatuses, ", "))
	}

	return nil
}

func isValidStatus(status string) bool {
	for _, validStatus := range ValidStatuses {
		if strings.EqualFold(status, validStatus) {
			return true
		}
	}
	return false
}

type VirtualMachineResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Cores     int       `json:"cores"`
	RAM       int       `json:"ram"`
	OS        string    `json:"os"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (vm *VirtualMachineRequest) ToResponse(id string, createAt, updateAt time.Time) VirtualMachineResponse {
	return VirtualMachineResponse{
		ID:        id,
		Name:      vm.Name,
		Cores:     vm.Cores,
		RAM:       vm.Ram,
		OS:        vm.OS,
		Status:    vm.Status,
		CreatedAt: createAt,
		UpdatedAt: updateAt,
	}
}

func (vm *VirtualMachineRequest) UpdateStatus(newStatus string) error {
	if !isValidStatus(newStatus) {
		return fmt.Errorf("Invalid status: %s. Valid statuses are: %s",
			newStatus, strings.Join(ValidStatuses, ", "))
	}
	vm.Status = newStatus
	return nil
}

func GenerateEventMessage(eventType, vmId string) string {
	return fmt.Sprintf("The vm with id %s has been %s", vmId, eventType)
}
