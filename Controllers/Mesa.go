package Controllers

import (
	"first-api/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMesas ... Get all Mesas
func GetMesas(c *gin.Context) {
	var mesa []Models.Mesa
	err := Models.GetAllMesas(&mesa)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, mesa)
	}
}

// CreateMesa ... Create Mesa
func CreateMesa(c *gin.Context) {
	var mesa Models.Mesa
	c.BindJSON(&mesa)
	err := Models.CreateMesa(&mesa)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, mesa)
	}
}

// GetMesaByID ... Get the Mesa by id
func GetMesaByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var mesa Models.Mesa
	err := Models.GetMesaByID(&mesa, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, mesa)
	}
}

// UpdateMesa ... Update the Mesa information
func UpdateMesa(c *gin.Context) {
	var mesa Models.Mesa
	id := c.Params.ByName("id")
	err := Models.GetMesaByID(&mesa, id)
	if err != nil {
		c.JSON(http.StatusNotFound, mesa)
	}
	c.BindJSON(&mesa)
	err = Models.UpdateMesa(&mesa, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, mesa)
	}
}

// DeleteMesa ... Delete the Mesa
func DeleteMesa(c *gin.Context) {
	var mesa Models.Mesa
	id := c.Params.ByName("id")
	err := Models.DeleteMesa(&mesa, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id " + id: "is deleted"})
	}
}
