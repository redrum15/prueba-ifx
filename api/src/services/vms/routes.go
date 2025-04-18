package vms

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/redrum15/prueba/src/db/models"
	"github.com/redrum15/prueba/src/db/querys"
	"github.com/redrum15/prueba/src/handlers"
	"github.com/redrum15/prueba/src/utils"
)

func CreateVM(w http.ResponseWriter, r *http.Request) {
	var vmReq utils.VirtualMachineRequest

	err := json.NewDecoder(r.Body).Decode(&vmReq)
	if err != nil {
		handlers.SendJSONError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := vmReq.Validate(); err != nil {
		handlers.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	vm := models.VirtualMachine{
		ID:     uuid.New(),
		Name:   vmReq.Name,
		Cores:  vmReq.Cores,
		Ram:    vmReq.Ram,
		OS:     vmReq.OS,
		Status: vmReq.Status,
	}

	_, err = querys.CreateVM(&vm)
	if err != nil {
		handlers.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	handlers.GetBroadcast() <- handlers.VMEvent{EventType: "created", Message: utils.GenerateEventMessage("created", vm.ID.String())}

	handlers.SendJSONSuccess(w, vmReq.ToResponse(vm.ID.String(), vm.CreatedAt, vm.UpdatedAt), http.StatusOK)
}

func ListVMS(w http.ResponseWriter, r *http.Request) {
	var vms []models.VirtualMachine
	var vmResponses []utils.VirtualMachineResponse

	_, err := querys.ListVMS(&vms)
	if err != nil {
		handlers.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, vm := range vms {
		vmResponses = append(vmResponses, utils.VirtualMachineResponse{
			ID:        vm.ID.String(),
			Name:      vm.Name,
			Cores:     vm.Cores,
			RAM:       vm.Ram,
			OS:        vm.OS,
			Status:    vm.Status,
			CreatedAt: vm.CreatedAt,
			UpdatedAt: vm.UpdatedAt,
		})
	}

	handlers.SendJSONSuccess(w, vmResponses, http.StatusOK)
}

func DetailVM(w http.ResponseWriter, r *http.Request) {
	vmId := chi.URLParam(r, "id")

	vm, err := querys.GetVM(vmId)
	if err != nil {
		handlers.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	vmResponse := utils.VirtualMachineResponse{
		ID:        vm.ID.String(),
		Name:      vm.Name,
		Cores:     vm.Cores,
		RAM:       vm.Ram,
		OS:        vm.OS,
		Status:    vm.Status,
		CreatedAt: vm.CreatedAt,
		UpdatedAt: vm.UpdatedAt,
	}

	handlers.SendJSONSuccess(w, vmResponse, http.StatusOK)
}

func DeleteVM(w http.ResponseWriter, r *http.Request) {
	vmId := chi.URLParam(r, "id")

	err := querys.DeleteVM(vmId)
	if err != nil {
		handlers.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	handlers.GetBroadcast() <- handlers.VMEvent{EventType: "deleted", Message: utils.GenerateEventMessage("deleted", vmId)}

	handlers.SendJSONSuccess(w, nil, http.StatusNoContent)
}

func UpdateVM(w http.ResponseWriter, r *http.Request) {
	vmId := chi.URLParam(r, "id")

	var vmRequest utils.VirtualMachineRequest

	err := json.NewDecoder(r.Body).Decode(&vmRequest)
	if err != nil {
		handlers.SendJSONError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := vmRequest.Validate(); err != nil {
		handlers.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	vm, err := querys.UpdateVMByID(vmId, vmRequest)
	if err != nil {
		handlers.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	handlers.GetBroadcast() <- handlers.VMEvent{EventType: "updated", Message: utils.GenerateEventMessage("updated", vmId)}

	handlers.SendJSONSuccess(w, vmRequest.ToResponse(vm.ID.String(), vm.CreatedAt, vm.UpdatedAt), http.StatusOK)
}
