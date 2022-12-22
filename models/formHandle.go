package models

import (
	"API/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// getall
func GetAllHandle(form *[]ModelForm) (err error) {
	if err = database.DB.Find(form).Error; err != nil {
		return err
	}
	return nil
}

// create
func CreateFormHandle(form *ModelForm) (err error) {
	err = database.DB.Create(form).Error
	if err != nil {
		return err
	}
	return nil
}

// getbyid
func GetFormByIdHandle(form *ModelForm, id string) (err error) {
	err = database.DB.Where("id = ?", id).First(form).Error
	if err != nil {
		return database.DB.Error
	}
	return nil
}

// update
func UpdateFormHandle(form *ModelForm, id string) (err error) {
	fmt.Println(form)
	database.DB.Save(form)
	return nil
}

// delete
func DeleteFormHandle(form *ModelForm, id string) (err error) {
	database.DB.Where("id = ?", id).Delete(form)
	return nil
}
