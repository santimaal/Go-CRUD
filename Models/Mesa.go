package Models

import (
	"first-api/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllMesas Fetch all Mesa data
func GetAllMesas(mesa *[]Mesa) (err error) {
	if err = Config.DB.Find(mesa).Error; err != nil {
		return err
	}
	return nil
}

// CreateMesa ... Insert New data
func CreateMesa(mesa *Mesa) (err error) {
	if err = Config.DB.Create(mesa).Error; err != nil {
		return err
	}
	return nil
}

// GetMesaByID ... Fetch only one Mesa by Id
func GetMesaByID(mesa *Mesa, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(mesa).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMesa ... Update Mesa
func UpdateMesa(mesa *Mesa, id string) (err error) {
	fmt.Println(mesa)
	Config.DB.Save(mesa)
	return nil
}

// DeleteMesa ... Delete Mesa
func DeleteMesa(mesa *Mesa, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(mesa)
	return nil
}
