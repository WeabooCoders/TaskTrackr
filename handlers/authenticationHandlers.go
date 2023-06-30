package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AvinFajarF/initializers"
	"github.com/AvinFajarF/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func SignUp(c *gin.Context) {
	var user model.User

	ids :=  uuid.New().String()

	fmt.Println(ids)

	user.Uuid = ids

	// mengambil semua data user dari request postman dan juga pengecekan error
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"error":   err.Error(),
			"massage": "silahkan di cek kembali",
		})
		return
	}

	// hash password user
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"massage": "silahkan di cek kembali",
		})
	}

	// menyimpan data user
	userModel := model.User{Username: user.Username, Password: string(hash), Email: user.Email}
	result := initializers.DB.Create(&userModel)

	// pengecekan error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"massage": "silahkan di cek kembali",
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": "success",
			"data":   userModel,
		})
	}
}

func SignIn(c *gin.Context) {
	var userRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// mengambil semua data user dari request postman dan juga pengecekan error
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"error":   err.Error(),
			"massage": "silahkan di cek kembali",
		})
		return
	}

	var user model.User
	initializers.DB.First(&user, "email = ?", userRequest.Email)


	// hash cek

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"massage": "silahkan masukan email dan username yang benar",
		})
		return
	}

	// membuat token

	key := []byte(os.Getenv("SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"massage": "error membuat token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.Header("Authorization", tokenString)

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"token":   tokenString,
	})

}

func Coba (c *gin.Context) {
	user, _ := c.Get("id")

    // Mendapatkan ID dari user
    c.JSON(http.StatusOK, gin.H{
        "status": "success",
        "data":   user,
    })
}
