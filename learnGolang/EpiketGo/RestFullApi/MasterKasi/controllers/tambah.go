package controllers

import (
	"net/http"
	"../structs"
	"../auth"
	
	"github.com/gin-gonic/gin"
)

func (idb *InDB) InsertData(c *gin.Context){
	var (
		result gin.H 
		kasi structs.Kasi
	)

	kode_kasi := c.PostForm("Kode_kasi")
	nama_kasi := c.PostForm("Nama_kasi")

	kasi.Kode_Kasi = kode_kasi
	kasi.Nama_Kasi = nama_kasi

	access := auth.Auth(c)

	if access == false {
		result = gin.H{
			"pesan": "not authorized",
			"status": "error",
		}
	} else {
		if kode_kasi == "" || nama_kasi == "" {
			result = gin.H{
				"status": "warning!",
				"pesan": "Silahkan lengkapi lembar isian",
			}
		} else {
			status := idb.DB.Create(&kasi).Error
			if status != nil {
				result = gin.H {
					"status": "error",
					"pesan": status,
				}
			} else {
				x := kasi.ID
				result = gin.H {
					"status": "Success",
					"pesan": "Data berhasil ditambahkan",
					"id": x,
				}
			}
		}
	}
	c.JSON(http.StatusOK, result)
}
