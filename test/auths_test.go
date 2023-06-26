package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AvinFajarF/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T)  {
	
	server := gin.Default()


	requestUser := struct{
		Username string  `json:"username"`
		Password string  `json:"password"`
		Email    string  `json:"email"`
	}{
		Username: "test",
		Password: "test",
		Email: "email@gmail.com",
	}

	bodyReqs, err := json.Marshal(requestUser)

	if err != nil {
		return
	}


	server.POST("/v1/signup", handlers.SignUp)

	req, _ := http.NewRequest(http.MethodPost, "/v1/signup", bytes.NewBuffer(bodyReqs))

	fmt.Println(req)

	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	t.Log(resp.Result().StatusCode)
	t.Log(resp.Body)

	assert.Equal(t, 200, resp.Result().StatusCode)
}