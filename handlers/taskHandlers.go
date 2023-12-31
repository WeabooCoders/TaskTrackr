package handlers

import (
	"net/http"

	"github.com/AvinFajarF/initializers"
	"github.com/AvinFajarF/model"
	"github.com/gin-gonic/gin"
)


func GetTaskAllById(c *gin.Context) {
	id, _ := c.Get("id")

	var tasks []model.Task

	// Menggunakan metode Find untuk mengambil data task berdasarkan user_id
	result := initializers.DB.Preload("User").Where("user_id = ?", id).Find(&tasks)
	
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

func FindTaskByTitle(c *gin.Context){

	var task model.Task

	// mendapatkan parameter task
	titleParam := c.Param("title")
	

	// mencari data title by title
	result := initializers.DB.Preload("User").Where("title = ?", titleParam).Find(&task)

	if result.Error!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "gagal mendapatkan data task",
			"error": result.Error,
		})
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "success",
		"data": task,
	})
}


func UpdateTaskByTitle(c *gin.Context){

	var task model.Task

	// mendapatkan request dari user
	err := c.ShouldBindJSON(&task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
            "message": "error ketika mengambil data dari request",
		})
	}

	// mendapatkan parameter task
	titleParam := c.Param("title")

	result := initializers.DB.Preload("User").Where("title = ?", titleParam).Updates(&task)

 // Menangani hasil pembaruan
 if result.Error != nil {
	// Tanggapan jika terjadi kesalahan pembaruan
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  "error",
		"message": "error ketika memperbarui tugas",
	})
	return
}

// Tanggapan jika pembaruan berhasil
c.JSON(http.StatusOK, gin.H{
	"status":  "success",
	"message": "tugas berhasil diperbarui",
})

}



