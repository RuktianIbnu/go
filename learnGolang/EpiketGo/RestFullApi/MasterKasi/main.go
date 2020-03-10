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

	router.GET("/MasterKasi/GetAllData", inDB.GetAllData)
	router.POST("/MasterKasi/GetDataById/:id", inDB.GetDataById)	
	router.POST("/MasterKasi/GetDataByNama/:nama", inDB.GetDataByName)
	router.POST("/MasterKasi/CreateData", inDB.InsertData)
	router.PUT("/MasterKasi/UpdateData", inDB.UpdateData)
	router.DELETE("/MasterKasi/DeleteData/:id", inDB.DeleteData)
	router.Run(":3500")
}