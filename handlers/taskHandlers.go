package handlers

import (
	"fmt"
	"net/http"

	"github.com/AvinFajarF/initializers"
	"github.com/AvinFajarF/model"
	"github.com/gin-gonic/gin"
)



func CreateTask(c *gin.Context){

	id, _ := c.Get("id")
	

    // Mengkonversi id menjadi uint
    // var userID uint
    // if idInt, ok := id.(int64); ok {
    //     userID = uint(idInt)
    // }


	
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
