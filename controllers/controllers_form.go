package controllers

import (
	"API/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getall
func GetALL_Form(c *gin.Context) {
	log.Print("Start")
	var form []models.ModelForm
	c.ShouldBindJSON(&form)
	err := models.GetAllHandle(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End")
}

// create
func Create_Form(c *gin.Context) {
	log.Print("Start")
	var form models.ModelForm
	// c.ShouldBindJSON(&form)
	err := models.CreateFormHandle(&form)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End")
}

// getformbyid
func GetById_Form(c *gin.Context) {
	log.Print("Start")
	var form models.ModelForm
	id := c.Params.ByName("id")
	// c.ShouldBindJSON(&form)
	err := models.GetFormByIdHandle(&form, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End")
}

// update
func Update_Form(c *gin.Context) {
	log.Print("Start")
	var form models.ModelForm
	id := c.Params.ByName("id")
	err := models.GetFormByIdHandle(&form, id)
	if err != nil {
		c.JSON(http.StatusNotFound, form)
	}
	c.ShouldBindJSON(&form)
	err = models.UpdateFormHandle(&form, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End")
}

// delete
func Delete_Form(c *gin.Context) {
	log.Print("Start")
	var form models.ModelForm
	id := c.Params.ByName("id")
	// c.ShouldBindJSON(&form)
	err := models.DeleteFormHandle(&form, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}

	log.Print("End")
}
