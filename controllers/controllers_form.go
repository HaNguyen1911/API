package controllers

import (
	"API/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getall
func GetALL_Form(c *gin.Context) {
	var form []models.ModelForm
	c.ShouldBindJSON(&form)
	err := models.GetAllHandle(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
}

// create
func Create_Form(c *gin.Context) {
	var form models.ModelForm
	c.ShouldBindJSON(&form)
	err := models.CreateFormHandle(&form)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}

}

// getformbyid
func GetById_Form(c *gin.Context) {
	var form models.ModelForm
	id := c.Params.ByName("id")
	c.ShouldBindJSON(&form)
	err := models.GetFormByIdHandle(&form, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
}

// update
func Update_Form(c *gin.Context) {
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
}

// delete
func Delete_Form(c *gin.Context) {
	var form models.ModelForm
	id := c.Params.ByName("id")
	c.ShouldBindJSON(&form)
	err := models.DeleteFormHandle(&form, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
