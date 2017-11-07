package pet

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
	"github.com/liteByte/mascotas-backend/models/pet"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(c *gin.Context) {
	petStruct, err := structs.ParseBodyIntoPetStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err = pet.Create(petStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `Pet created successfully`)
}

func GetList(c *gin.Context) {
	petStruct, err := pet.GetList()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	petJSON, _ := json.Marshal(petStruct)
	c.JSON(200, json.RawMessage(string(petJSON)))
}

func GetById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	petStruct, err := pet.GetById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	petJSON, _ := json.Marshal(petStruct)
	c.JSON(200, json.RawMessage(string(petJSON)))
}

func UpdateById(c *gin.Context) {
	petStruct, err := structs.ParseBodyIntoPetStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	petStruct.Id = frango.StringToInt(c.Params.ByName(`id`))

	if err = pet.UpdateById(petStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `Pet updated successfully`)
}

func DeleteById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	if err := pet.DeleteById(id); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `Pet deleted successfully`)
}
