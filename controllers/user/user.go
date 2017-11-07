package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
	"github.com/liteByte/mascotas-backend/models/user"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(c *gin.Context) {
	userStruct, err := structs.ParseBodyIntoUserStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err = user.Create(userStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `User created successfully`)
}

func GetList(c *gin.Context) {
	userStruct, err := user.GetList()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	userJSON, _ := json.Marshal(userStruct)
	c.JSON(200, json.RawMessage(string(userJSON)))
}

func GetById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	userStruct, err := user.GetById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	userJSON, _ := json.Marshal(userStruct)
	c.JSON(200, json.RawMessage(string(userJSON)))
}

func GetByUsername(c *gin.Context) {
	username := c.Params.ByName(`username`)

	userStruct, err := user.GetByUsername(username)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	userJSON, _ := json.Marshal(userStruct)
	c.JSON(200, json.RawMessage(string(userJSON)))
}

func GetByEmail(c *gin.Context) {
	email := c.Params.ByName(`email`)

	userStruct, err := user.GetByEmail(email)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	userJSON, _ := json.Marshal(userStruct)
	c.JSON(200, json.RawMessage(string(userJSON)))
}

func UpdateById(c *gin.Context) {
	userStruct, err := structs.ParseBodyIntoUserStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	userStruct.Id = frango.StringToInt(c.Params.ByName(`id`))

	if err = user.UpdateById(userStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `User updated successfully`)
}

func UpdateByUsername(c *gin.Context) {
	userStruct, err := structs.ParseBodyIntoUserStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	userStruct.Username = c.Params.ByName(`username`)

	if err = user.UpdateByUsername(userStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `User updated successfully`)
}

func UpdateByEmail(c *gin.Context) {
	userStruct, err := structs.ParseBodyIntoUserStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	userStruct.Email = c.Params.ByName(`email`)

	if err = user.UpdateByEmail(userStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `User updated successfully`)
}

func DeleteById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	if err := user.DeleteById(id); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `User deleted successfully`)
}

func DeleteByUsername(c *gin.Context) {
	username := c.Params.ByName(`username`)

	if err := user.DeleteByUsername(username); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `User deleted successfully`)
}

func DeleteByEmail(c *gin.Context) {
	email := c.Params.ByName(`email`)

	if err := user.DeleteByEmail(email); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `User deleted successfully`)
}
