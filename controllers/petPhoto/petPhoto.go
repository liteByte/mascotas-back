package petPhoto

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
	"github.com/liteByte/mascotas-backend/models/petPhoto"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(c *gin.Context) {
	petPhotoStruct, err := structs.ParseBodyIntoPetPhotoStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err = petPhoto.Create(petPhotoStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `PetPhoto created successfully`)
}

func GetList(c *gin.Context) {
	petPhotoStruct, err := petPhoto.GetList()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	petPhotoJSON, _ := json.Marshal(petPhotoStruct)
	c.JSON(200, json.RawMessage(string(petPhotoJSON)))
}

func GetById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	petPhotoStruct, err := petPhoto.GetById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	petPhotoJSON, _ := json.Marshal(petPhotoStruct)
	c.JSON(200, json.RawMessage(string(petPhotoJSON)))
}

func UpdateById(c *gin.Context) {
	petPhotoStruct, err := structs.ParseBodyIntoPetPhotoStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	petPhotoStruct.Id = frango.StringToInt(c.Params.ByName(`id`))

	if err = petPhoto.UpdateById(petPhotoStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `PetPhoto updated successfully`)
}

func DeleteById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	if err := petPhoto.DeleteById(id); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `PetPhoto deleted successfully`)
}
