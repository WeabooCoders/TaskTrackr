package handlers

import (
	"fmt"
	"net/http"

	"github.com/AvinFajarF/initializers"
	"github.com/AvinFajarF/model"
	"github.com/gin-gonic/gin"
)


func GetTaskAllById(c *gin.Context) {
	id, _ := c.Get("id")

	var tasks []model.Task

	// Menggunakan metode Find untuk mengambil data task berdasarkan user_id
	result := initializers.DB.Where("user_id = ?", id).Find(&tasks)
	
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "gagal mendapatkan data task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   tasks,
	})
}

func CreateTask(c *gin.Context){

	// mendapatkan id dari user yang login
	id, _ := c.Get("id")
	
	var task model.Task

	task.User.ID = id.(uint)
	fmt.Println("coba : ", task)


	err := c.ShouldBindJSON(&task)

	// mengecek error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "error ketika mengambil data dari request",
		})
	}

	// mengirim data dari request user ke database
	result := initializers.DB.Create(&task)

		// pengecekan error
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"massage": "silahkan di cek kembali",
			})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"status": "success",
				"massage": "data berhasil di buat",
			})
		}

}
