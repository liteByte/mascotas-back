package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
	"github.com/liteByte/mascotas-backend/authentication"
	"github.com/liteByte/mascotas-backend/models/user"
	"github.com/liteByte/mascotas-backend/structs"
)

func Login(c *gin.Context) {
	loginStruct, err := structs.ParseBodyIntoLoginStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	loginStruct.Password = frango.Hash(loginStruct.Username, loginStruct.Password)

	if err = user.CheckLogin(loginStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	token := authentication.CreateToken(loginStruct.Username)

	c.JSON(200, token)
}

func Signup(c *gin.Context) {
	userStruct, err := structs.ParseBodyIntoUserStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	userStruct.Password = frango.Hash(userStruct.Username, userStruct.Password)

	if err = user.Create(userStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "Signup successful")
}
