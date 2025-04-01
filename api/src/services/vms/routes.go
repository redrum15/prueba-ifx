package vms

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/redrum15/prueba/src/db/models"
	"github.com/redrum15/prueba/src/db/querys"
	"github.com/redrum15/prueba/src/utils"
)

func CreateVM(w http.ResponseWriter, r *http.Request) {
	var vmReq utils.VirtualMachineRequest

	err := json.NewDecoder(r.Body).Decode(&vmReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := vmReq.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vmReq.ToResponse(vm.ID.String(), vm.CreatedAt, vm.UpdatedAt)); err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
}

func ListVMS(w http.ResponseWriter, r *http.Request) {
	var vms []models.VirtualMachine

	_, err := querys.ListVMS(&vms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vms); err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
}

func DetailVM(w http.ResponseWriter, r *http.Request) {
	vmId := chi.URLParam(r, "id")

	vm, err := querys.GetVM(vmId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vmResponse); err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

}

func DeleteVM(w http.ResponseWriter, r *http.Request) {
	vmId := chi.URLParam(r, "id")

	err := querys.DeleteVM(vmId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func UpdateVM(w http.ResponseWriter, r *http.Request) {
	vmId := chi.URLParam(r, "id")

	var vmRequest utils.VirtualMachineRequest

	err := json.NewDecoder(r.Body).Decode(&vmRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := vmRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vm, err := querys.UpdateVMByID(vmId, vmRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vmRequest.ToResponse(vm.ID.String(), vm.CreatedAt, vm.UpdatedAt)); err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
}
