package controllers

import (
	"net/http"
	"../structs"
	"../auth"
	
	"github.com/gin-gonic/gin"
)

func (idb *InDB) UpdateData(c *gin.Context) {
	id := c.PostForm("id")
	kode_kasi := c.PostForm("Kode_Kasi")
	nama_kasi := c.PostForm("Nama_Kasi")
	var (
		kasi    structs.Kasi
		newKasi structs.Kasi
		result    gin.H
	)

	access := auth.Auth(c)

	if access == false {
		result = gin.H{
			"pesan": "not authorized",
			"status": "error",
		}
	} else {
		CekDataKasi := idb.DB.Where("id = ?", id).First(&kasi).Error 
		if CekDataKasi != nil {
			result = gin.H{
				"pesan": "data tidak ditemukan",
				"status": "not_found",
				"CekDataKasi": CekDataKasi,
			}
		} else if CekDataKasi == nil {
			newKasi.Kode_Kasi = kode_kasi
			newKasi.Nama_Kasi = nama_kasi
			err := idb.DB.Model(&kasi).Where("id = ?", id).Updates(newKasi).Error
			if err != nil {
				result = gin.H{
					"pesan": "update failed",
					"status": "error",
				}
			} else {
				result = gin.H{
					"pesan": "successfully updated data",
					"status": "success",
				}
			}	
		}
	}
	c.JSON(http.StatusOK, result)
}