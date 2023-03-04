package db

import (
	"enerbit/prueba/models"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresConnection
func PostgresConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(MakeDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}

// MakeDSN
func MakeDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASS"), os.Getenv("DB"), os.Getenv("PORT"))
}

// Actions

// ListEMeter
func ListEMeter() ([]models.Datameds, error) {
	var data []models.Datameds

	dbConn, err := gorm.Open(postgres.Open(MakeDSN()), &gorm.Config{})
	if err != nil {
		return data, err
	}

	result := dbConn.Find(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

// LastEMeter
func LastEMeter(serial string, brand string) (models.Datameds, error) {
	var data models.Datameds

	dbConn, err := gorm.Open(postgres.Open(MakeDSN()), &gorm.Config{})
	if err != nil {
		return data, err
	}

	result := dbConn.Where("serial = ? AND brand = ?", serial, brand).Order("created_at").Last(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

// ListDisableEMeter
func ListDisableEMeter() ([]models.Datameds, error) {
	var data []models.Datameds

	dbConn, err := gorm.Open(postgres.Open(MakeDSN()), &gorm.Config{})
	if err != nil {
		return data, err
	}

	result := dbConn.Where("is_active = ?", false).Find(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

// NewEMeter
func NewEMeter(data models.NewEMeterModel) error {
	dbConn, err := gorm.Open(postgres.Open(MakeDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	err = checkCreate(dbConn, data)
	if err != nil {
		return err
	}

	createAt, err := time.Parse("2006-01-02", data.InstallationDate)
	if err != nil {
		return err
	}

	dbData := models.Datameds{
		ID:               uuid.New().String(),
		Brand:            data.Brand,
		Address:          data.Address,
		InstallationDate: createAt,
		Serial:           data.Serial,
		Lines:            data.Lines,
		IsActive:         true,
		CreatedAt:        time.Now(),
	}

	result := dbConn.Create(&dbData)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteEMeter
func DeleteEMeter(id string) error {
	dbConn, err := gorm.Open(postgres.Open(MakeDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	err = checkDelete(dbConn, id)
	if err != nil {
		return err
	}

	data := models.Datameds{
		ID: id,
	}

	result := dbConn.Delete(&data)
	if result.Error != nil {
		return err
	}

	return nil
}

// UpdateEMeterStatus
func UpdateEMeterStatus(id string, status bool) error {
	dbConn, err := gorm.Open(postgres.Open(MakeDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	// Si se va a deshabilitar el servicio simplemente lo desactivo solo si actualmente se encuentra activo
	if !status {
		data := models.Datameds{
			ID: id,
		}

		result := dbConn.Model(&data).Where("is_active = ?", true).Updates(map[string]interface{}{"is_active": false, "retirement_date": time.Now()})
		if result.Error != nil {
			return result.Error
		}

		return nil
	}

	// Para activar un servicio primero verifico que no exista un predio con un servicio activo

	var data models.Datameds

	result := dbConn.First(&data, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	err = checkAdd(dbConn, data.Address)
	if err != nil {
		return err
	}

	result = dbConn.Model(&data).Update("is_active", true)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
