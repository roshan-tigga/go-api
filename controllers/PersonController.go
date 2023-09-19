package controllers

import (
	"crud/connection"
	"crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersons(c *gin.Context) {
	var persons []models.Person
	connection.DB.Find(&persons)
	c.JSON(http.StatusOK, persons)
}

func GetPerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := connection.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func CreatePerson(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	connection.DB.Create(&person)
	c.JSON(http.StatusCreated, person)
}

func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := connection.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	connection.DB.Save(&person)
	c.JSON(http.StatusOK, person)
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	if err := connection.DB.Delete(&models.Person{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}
