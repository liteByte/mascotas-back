package role

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
	"github.com/liteByte/mascotas-backend/models/role"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(c *gin.Context) {
	roleStruct, err := structs.ParseBodyIntoRoleStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err = role.Create(roleStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `Role created successfully`)
}

func GetList(c *gin.Context) {
	roleStruct, err := role.GetList()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	roleJSON, _ := json.Marshal(roleStruct)
	c.JSON(200, json.RawMessage(string(roleJSON)))
}

func GetById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	roleStruct, err := role.GetById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	roleJSON, _ := json.Marshal(roleStruct)
	c.JSON(200, json.RawMessage(string(roleJSON)))
}

func UpdateById(c *gin.Context) {
	roleStruct, err := structs.ParseBodyIntoRoleStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	roleStruct.Id = frango.StringToInt(c.Params.ByName(`id`))

	if err = role.UpdateById(roleStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `Role updated successfully`)
}

func DeleteById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	if err := role.DeleteById(id); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `Role deleted successfully`)
}
