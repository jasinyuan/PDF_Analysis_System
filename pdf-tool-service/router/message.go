package routers

import (
	massagecor "main/controllers/massageCor"
	"main/model"

	"github.com/gin-gonic/gin"
)

func Message(r *gin.Engine) { 
	rdb := model.GetRedisClient()
	MassageControllers := massagecor.NewMassageControllers(rdb)
	r.GET("/",MassageControllers.MessageGet)
}