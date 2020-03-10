package controllers

import (
	"net/http"
	"../structs"
	"../auth"
	
	"github.com/gin-gonic/gin"
)

func (idb *InDB) DeleteData(c *gin.Context){
	var (
		kasi structs.Kasi
		result gin.H
	)
		id := c.Param("id")
		err := idb.DB.Where("id = ?", id).First(&kasi).Error
		
		access := auth.Auth(c)
		if access == false {
			result = gin.H{
				"pesan": "not authorized",
				"status": "error",
			}
		} else {
			if err != nil {
				result = gin.H{
					"pesan": "data tidak ditemukan",
					"status": "not found",
				}
			} else {
				err = idb.DB.Delete(&kasi).Error
				if err != nil {
					result = gin.H{
						"pesan": "delete failed",
						"status": "error",
					}
				} else {
					result = gin.H{
						"pesan": "Data deleted successfully",
						"status": "success",
					}
				}
			}
		}	
	c.JSON(http.StatusOK, result)
}