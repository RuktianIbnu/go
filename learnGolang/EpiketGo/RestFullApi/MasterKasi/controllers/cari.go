package controllers

import (
	"net/http"
	"../structs"
	"../auth"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetAllData(c *gin.Context){
	var (
		result gin.H 
		kasi []structs.Kasi
	)

	idb.DB.Find(&kasi)
	access := auth.Auth(c)

	if access == false {
		result = gin.H{
			"pesan": "not authorized",
			"status": "error",
		}
	} else {
		
			if len(kasi) <= 0 {
				result = gin.H {
					"result": nil,
					"count": 0,
					"status": "Data Kosong",
				}
			} else {
				result = gin.H {
					"count": len(kasi),
					"status": "success",
					"data_kasi": kasi,
				}
			}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetDataByName(c *gin.Context){
	var (
		result gin.H 
		kasi []structs.Kasi
	)

	nama := c.Param("nama")
	idb.DB.Where("nama_kasi = ?", nama).Find(&kasi)
	access := auth.Auth(c)

	if access == false {
		result = gin.H{
			"pesan": "not authorized",
			"status": "error",
		}
	} else {
		if len(kasi) <= 0 {
			result = gin.H {
				"result": nil,
				"count": 0,
				"status": "Data Kosong",
			}
		} else {
			result = gin.H {
				"count": len(kasi),
				"status": "success",
				"data_kasi": kasi,
			}
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetDataById(c *gin.Context){
	var (
		result gin.H 
		kasi []structs.Kasi
	)

	id := c.Param("id")
	idb.DB.Where("id = ?", id).Find(&kasi)
	access := auth.Auth(c)

	if access == false {
		result = gin.H{
			"pesan": "not authorized",
			"status": "error",
		}
	} else {
		if len(kasi) <= 0 {
			result = gin.H {
				"result": nil,
				"count": 0,
				"status": "Data Kosong",
			}
		} else {
			result = gin.H {
				"count": len(kasi),
				"status": "success",
				"data_kasi": kasi,
			}
		}
	}
	c.JSON(http.StatusOK, result)
}