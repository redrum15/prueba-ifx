package querys

import (
	"errors"

	"github.com/redrum15/prueba/src/db"
	"github.com/redrum15/prueba/src/db/models"
	"github.com/redrum15/prueba/src/utils"
	"gorm.io/gorm"
)

func CreateVM(vm *models.VirtualMachine) (*gorm.DB, error) {

	result := db.GetDBInstance().Create(&vm)

	if result.Error != nil {
		return nil, errors.New("couldn't create virtual machine not found")
	}

	return result, nil

}

func GetVM(id string) (*models.VirtualMachine, error) {

	var vm models.VirtualMachine

	result := db.GetDBInstance().First(&vm, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vm, nil

}

func ListVMS(vms *[]models.VirtualMachine) (*[]models.VirtualMachine, error) {
	if err := db.GetDBInstance().Order("created_at DESC").Find(&vms).Error; err != nil {
		return nil, err
	}
	return vms, nil
}

func DeleteVM(id string) error {
	if err := db.GetDBInstance().Unscoped().Delete(&models.VirtualMachine{}, "id = ?", id).Error; err != nil {
		return errors.New("couldn't delete virtual machine")
	}
	return nil
}

func UpdateVMByID(id string, updatedVM utils.VirtualMachineRequest) (*models.VirtualMachine, error) {
	var vm models.VirtualMachine

	result := db.GetDBInstance().First(&vm, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	vm.Name = updatedVM.Name
	vm.Cores = updatedVM.Cores
	vm.Ram = updatedVM.Ram
	vm.OS = updatedVM.OS
	vm.Status = updatedVM.Status

	result = db.GetDBInstance().Save(&vm)
	if result.Error != nil {
		return nil, result.Error
	}

	return &vm, nil
}
