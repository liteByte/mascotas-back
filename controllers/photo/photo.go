package photo

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
	"github.com/liteByte/mascotas-backend/models/photo"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(c *gin.Context) {
	photoStruct, err := structs.ParseBodyIntoPhotoStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err = photo.Create(photoStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `Photo created successfully`)
}

func GetList(c *gin.Context) {
	photoStruct, err := photo.GetList()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	photoJSON, _ := json.Marshal(photoStruct)
	c.JSON(200, json.RawMessage(string(photoJSON)))
}

func GetById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	photoStruct, err := photo.GetById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	photoJSON, _ := json.Marshal(photoStruct)
	c.JSON(200, json.RawMessage(string(photoJSON)))
}

func UpdateById(c *gin.Context) {
	photoStruct, err := structs.ParseBodyIntoPhotoStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	photoStruct.Id = frango.StringToInt(c.Params.ByName(`id`))

	if err = photo.UpdateById(photoStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `Photo updated successfully`)
}

func DeleteById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	if err := photo.DeleteById(id); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `Photo deleted successfully`)
}
