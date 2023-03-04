package db

import (
	"enerbit/prueba/models"
	"errors"

	"gorm.io/gorm"
)

func checkAdd(dbConn *gorm.DB, address string) error {
	result := dbConn.Where("address = ? AND is_active = true", address).Limit(1).Find(&models.Datameds{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return errors.New("solo puede haber una instalación por predio activa")
	}

	return nil
}

func checkCreate(dbConn *gorm.DB, createData models.NewEMeterModel) error {
	var data models.Datameds

	result := dbConn.Where("serial = ? AND brand =?", createData.Serial, createData.Brand).Limit(1).Find(&data)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return errors.New("el serial-marca proporcionado ya existe en el sistema")
	}

	err := checkAdd(dbConn, createData.Address)
	if err != nil {
		return err
	}

	return nil
}

func checkDelete(dbConn *gorm.DB, id string) error {
	var data models.Datameds

	result := dbConn.First(&data, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if data.IsActive {
		return errors.New("no se puede eliminar un registro activo actualmente")
	}

	return nil
}
