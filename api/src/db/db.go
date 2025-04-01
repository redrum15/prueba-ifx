package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	configs "github.com/redrum15/prueba/src/config"
	"github.com/redrum15/prueba/src/db/models"
	"github.com/redrum15/prueba/src/utils"
)

var dbInstance *gorm.DB

func Init() {
	var err error

	dbInstance, err = gorm.Open(postgres.Open(configs.Envs.DBUrl), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	createTables()
	createUsers()
}

func createTables() {
	dbInstance.AutoMigrate(&models.User{})
	dbInstance.AutoMigrate(&models.VirtualMachine{})
}

func createUsers() {
	pass, err := utils.HashPassword(configs.Envs.AdminPasword)
	if err != nil {
		panic("unable to hast password")
	}

	dbInstance.FirstOrCreate(&models.User{
		Email:    "admin@test.com",
		Type:     "admin",
		Password: pass,
	}, models.User{Email: "admin@test.com"})

	dbInstance.FirstOrCreate(&models.User{
		Email:    "client@test.com",
		Type:     "client",
		Password: pass,
	}, models.User{Email: "client@test.com"})
}

func GetDBInstance() *gorm.DB {
	return dbInstance
}
