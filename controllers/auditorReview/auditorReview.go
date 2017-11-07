package auditorReview

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"
	"github.com/liteByte/mascotas-backend/models/auditorReview"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(c *gin.Context) {
	auditorReviewStruct, err := structs.ParseBodyIntoAuditorReviewStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err = auditorReview.Create(auditorReviewStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `AuditorReview created successfully`)
}

func GetList(c *gin.Context) {
	auditorReviewStruct, err := auditorReview.GetList()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	auditorReviewJSON, _ := json.Marshal(auditorReviewStruct)
	c.JSON(200, json.RawMessage(string(auditorReviewJSON)))
}

func GetById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	auditorReviewStruct, err := auditorReview.GetById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	auditorReviewJSON, _ := json.Marshal(auditorReviewStruct)
	c.JSON(200, json.RawMessage(string(auditorReviewJSON)))
}

func UpdateById(c *gin.Context) {
	auditorReviewStruct, err := structs.ParseBodyIntoAuditorReviewStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	auditorReviewStruct.Id = frango.StringToInt(c.Params.ByName(`id`))

	if err = auditorReview.UpdateById(auditorReviewStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `AuditorReview updated successfully`)
}

func DeleteById(c *gin.Context) {
	id := frango.StringToInt(c.Params.ByName(`id`))

	if err := auditorReview.DeleteById(id); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `AuditorReview deleted successfully`)
}
