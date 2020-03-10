package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/subdit/:id", inDB.GetSubditById)
	router.GET("/subdit", inDB.GetSubdit)
	router.POST("/subdit/GetSubditByNama/:nama_subdit", inDB.GetSubditByNama)
	router.POST("/subdit/CreateSubdit", inDB.CreateSubdit)
	router.PUT("/subdit/UpdateSubdit", inDB.UpdateSubdit)
	router.DELETE("/subdit/:id", inDB.DeleteSubdit)
	router.GET("/GetToken/FromAPI-Login", inDB.GetToken)
	router.Run(":3300")
}