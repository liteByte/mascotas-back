package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/liteByte/mascotas-backend/config"
	"github.com/liteByte/mascotas-backend/controllers/auditorReview"
	"github.com/liteByte/mascotas-backend/controllers/authentication"
	"github.com/liteByte/mascotas-backend/controllers/pet"
	"github.com/liteByte/mascotas-backend/controllers/petPhoto"
	"github.com/liteByte/mascotas-backend/controllers/photo"
	"github.com/liteByte/mascotas-backend/controllers/refugeAbility"
	"github.com/liteByte/mascotas-backend/controllers/role"
	"github.com/liteByte/mascotas-backend/controllers/user"
	"github.com/liteByte/mascotas-backend/middleware"
	"os"
)

var router *gin.Engine

func ConfigureRouter() {
	if config.GetConfig().ENV != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func CreateRouter() {
	router = gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authentication", "Content-Type", "Authorization", "X-API-KEY"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	}))

	public := router.Group("/")
	{
		public.POST("/signup", authentication.Signup)
		public.POST("/login", authentication.Login)
	}

	api := router.Group("/", middleware.ValidateToken())
	{
		api.POST(`/user`, user.Create)
		api.GET(`/user/list`, user.GetList)
		api.GET(`/user/id/:id`, user.GetById)
		api.PUT(`/user/id/:id`, user.UpdateById)
		api.DELETE(`/user/id/:id`, user.DeleteById)
		api.GET(`/user/username/:username`, user.GetByUsername)
		api.PUT(`/user/username/:username`, user.UpdateByUsername)
		api.DELETE(`/user/username/:username`, user.DeleteByUsername)
		api.GET(`/user/email/:email`, user.GetByEmail)
		api.PUT(`/user/email/:email`, user.UpdateByEmail)
		api.DELETE(`/user/email/:email`, user.DeleteByEmail)
		api.POST(`/role`, role.Create)
		api.GET(`/role/list`, role.GetList)
		api.GET(`/role/id/:id`, role.GetById)
		api.PUT(`/role/id/:id`, role.UpdateById)
		api.DELETE(`/role/id/:id`, role.DeleteById)
		api.POST(`/photo`, photo.Create)
		api.GET(`/photo/list`, photo.GetList)
		api.GET(`/photo/id/:id`, photo.GetById)
		api.PUT(`/photo/id/:id`, photo.UpdateById)
		api.DELETE(`/photo/id/:id`, photo.DeleteById)
		api.POST(`/refugeAbility`, refugeAbility.Create)
		api.GET(`/refugeAbility/list`, refugeAbility.GetList)
		api.GET(`/refugeAbility/id/:id`, refugeAbility.GetById)
		api.PUT(`/refugeAbility/id/:id`, refugeAbility.UpdateById)
		api.DELETE(`/refugeAbility/id/:id`, refugeAbility.DeleteById)
		api.POST(`/auditorReview`, auditorReview.Create)
		api.GET(`/auditorReview/list`, auditorReview.GetList)
		api.GET(`/auditorReview/id/:id`, auditorReview.GetById)
		api.PUT(`/auditorReview/id/:id`, auditorReview.UpdateById)
		api.DELETE(`/auditorReview/id/:id`, auditorReview.DeleteById)
		api.POST(`/pet`, pet.Create)
		api.GET(`/pet/list`, pet.GetList)
		api.GET(`/pet/id/:id`, pet.GetById)
		api.PUT(`/pet/id/:id`, pet.UpdateById)
		api.DELETE(`/pet/id/:id`, pet.DeleteById)
		api.POST(`/petPhoto`, petPhoto.Create)
		api.GET(`/petPhoto/list`, petPhoto.GetList)
		api.GET(`/petPhoto/id/:id`, petPhoto.GetById)
		api.PUT(`/petPhoto/id/:id`, petPhoto.UpdateById)
		api.DELETE(`/petPhoto/id/:id`, petPhoto.DeleteById)

	}
}

func RunRouter() {
	herokuPort, _ := os.LookupEnv("PORT")

	if herokuPort == "" {
		herokuPort = config.GetConfig().PORT
	}

	router.Run(":" + herokuPort)
}
