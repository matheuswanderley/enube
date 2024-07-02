package routes

import (
    "teste-enube/controllers"
    "github.com/gin-gonic/gin"
)

func SetupDataRoutes(router *gin.RouterGroup) {

        router.GET("/data/:field", controllers.GetDataByField)
    }

