package controllers

import (
	"net/http"
	"../structs"
	"strconv"
	"io/ioutil"
	"fmt"
	
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

func (idb *InDB) GetToken(c *gin.Context) {
	var result gin.H
	response, err := http.Get("http://localhost:3000/api/token")
	if err != nil {
		result = gin.H {
			"status": "failed communication",
		}
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		result = gin.H {
			"data": data,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get data by id
func (idb *InDB) GetSubditById(c *gin.Context) {
	var (
		subdit []structs.Subdit
		result gin.H
		id string
	)

	id = c.Param("id")
	idInt,_ := strconv.Atoi(id)
	err := idb.DB.Find(&subdit, idInt)

		if err != nil {
			result = gin.H {
				"result": err,
				"count": 0,
				"status": "error",
			}
		} else {
			result = gin.H {
				"data_subdit": err,
				"status": "success",
				//"id": id,
			}
		}
	c.JSON(http.StatusOK, result)
}

// get all data
func (idb *InDB) GetSubdit(c *gin.Context) {
	var (
		subdits []structs.Subdit
		result gin.H
	)

	idb.DB.Find(&subdits)
	access := Auth(c)
	
	if access == false {
		result = gin.H{
			"pesan": "not authorized",
			"status": "error",
		}
	} else {
		if len(subdits) <= 0 {
			result = gin.H {
				"result": nil,
				"count": 0,
				"status": "Data Kosong",
			}
		} else {
			result = gin.H {
				"count": len(subdits),
				"status": "success",
				"data_subdit": subdits,
			}
		}
	}
	c.JSON(http.StatusOK, result)
}

// Search data by nama
func (idb *InDB) GetSubditByNama(c *gin.Context) {
	var (
		subdit []structs.Subdit
		result gin.H
	)
	nama_subdit := c.Param("nama_subdit")
	query := "%" + nama_subdit + "%"
	err := idb.DB.Where("nama_subdirektorat LIKE ?", query).Find(&subdit).Error
	if err != nil {
		result = gin.H {
			"data": "kosong",
			"count": 0,
			"status": "error",
			"nama_subdit": nama_subdit,
		}
	} else {
		result = gin.H {
			"data_subdit": subdit,
			"status": "Success",
			"count": len(subdit),
			"nama_subdit": nama_subdit,
			"query": query,
		}
	}
	c.JSON(http.StatusOK, result)
}

// create new data
func (idb *InDB) CreateSubdit(c *gin.Context) {
	var (
		subdit structs.Subdit
		result gin.H
	)
	kode_subdit := c.PostForm("Kode_Subdirektorat")
	nama_subdit := c.PostForm("Nama_Subdirektorat")

	subdit.Kode_Subdirektorat = kode_subdit
	subdit.Nama_Subdirektorat = nama_subdit

	if subdit.Kode_Subdirektorat == "" || subdit.Nama_Subdirektorat == "" {
		result = gin.H{
			"status": "error",
			"pesan": "Silahkan lengkapi lembar isian",
		}
	} else {
		status := idb.DB.Create(&subdit).Error
		if status != nil {
			result = gin.H {
				"status": "error",
				"pesan": status,
			}
		} else {
			result = gin.H {
				"status": "Success",
				"pesan": "Data berhasil ditambahkan",
			}
		}
	}
	c.JSON(http.StatusOK, result)
}

// update data with {id} as query
func (idb *InDB) UpdateSubdit(c *gin.Context) {
	id := c.PostForm("id")
	kode_subdit := c.PostForm("Kode_Subdirektorat")
	nama_subdit := c.PostForm("Nama_Subdirektorat")
	var (
		subdit    structs.Subdit
		newSubdit structs.Subdit
		result    gin.H
	)

	CekDataSubdit := idb.DB.Where("id = ?", id).First(&subdit).Error 
	if CekDataSubdit != nil {
		result = gin.H{
			"pesan": "data tidak ditemukan",
			"status": "not_found",
			"CekDataSubdit": CekDataSubdit,
		}
	} else if CekDataSubdit == nil {
		newSubdit.Kode_Subdirektorat = kode_subdit
		newSubdit.Nama_Subdirektorat = nama_subdit
		err := idb.DB.Model(&subdit).Where("id = ?", id).Updates(newSubdit).Error
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
	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (idb *InDB) DeleteSubdit(c *gin.Context) {
	var (
		subdit structs.Subdit
		result gin.H
	)
		id := c.Param("id")
		err := idb.DB.Where("id = ?", id).First(&subdit).Error
		
		access := Auth(c)
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
				err = idb.DB.Delete(&subdit).Error
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

func Auth(c *gin.Context) bool{
	var status bool
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{},error){
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		status = true
	} else {
		status = false
	}
	return status
}