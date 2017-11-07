package refugeAbility

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
	"github.com/liteByte/mascotas-backend/models/refugeAbility"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(c *gin.Context) {
	refugeAbilityStruct, err := structs.ParseBodyIntoRefugeAbilityStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err = refugeAbility.Create(refugeAbilityStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `RefugeAbility created successfully`)
}

func GetList(c *gin.Context) {
	refugeAbilityStruct, err := refugeAbility.GetList()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	refugeAbilityJSON, _ := json.Marshal(refugeAbilityStruct)
	c.JSON(200, json.RawMessage(string(refugeAbilityJSON)))
}

func GetById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	refugeAbilityStruct, err := refugeAbility.GetById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	refugeAbilityJSON, _ := json.Marshal(refugeAbilityStruct)
	c.JSON(200, json.RawMessage(string(refugeAbilityJSON)))
}

func UpdateById(c *gin.Context) {
	refugeAbilityStruct, err := structs.ParseBodyIntoRefugeAbilityStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	refugeAbilityStruct.Id = frango.StringToInt(c.Params.ByName(`id`))

	if err = refugeAbility.UpdateById(refugeAbilityStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `RefugeAbility updated successfully`)
}

func DeleteById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	if err := refugeAbility.DeleteById(id); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `RefugeAbility deleted successfully`)
}
