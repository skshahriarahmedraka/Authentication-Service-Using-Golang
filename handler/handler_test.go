package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/config"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/pkg/monodb"
	"github.com/stretchr/testify/assert"
)

func TestALLRoute(t *testing.T) {
	
	config.LoadEnvVars()
	config.AdminEmails = map[string]bool{
		"skshahra@gmail.com": true,
		"skraka@gmail.com" : true ,
		"raka@gmail.com" : true ,
	}
	r := gin.New()
	
	r.Use(gin.Logger())

	
	mongoConn := monodb.MongodbConnection()
	H := monodb.DatabaseInitialization(mongoConn)

	// Mock route for testing
	r.GET("/public", H.Home())
	r.POST("/login", H.Login())
	r.POST("/register", H.Register())
	r.GET("/logout", H.Logout())
	r.GET("/:id", H.UserData())
	r.GET("/alluser", H.AllUserData())



	// TEST /public ROUTE
	publicResponse := performPublicRequest(t, r)
	assert.Equal(t, http.StatusOK, publicResponse.Code)

	// TEST /register ROUTE
	registerResponse := performRegisterRequest(t, r)
	assert.Equal(t, http.StatusOK, registerResponse.Code)

	// TEST /login ROUTE
	loginResponse := performLoginRequest(t, r)
	assert.Equal(t, http.StatusOK, loginResponse.Code)

	var loginData map[string]interface{}
    err := json.Unmarshal(loginResponse.Body.Bytes(), &loginData)
	assert.NoError(t, err)
	id, ok := loginData["id"].(string)
	if !ok {
		t.Fatal("❌🔥 error in c.bindjson() ")
	}
	// TEST /:id ROUTE
	UserDataResponse := performUserdataRequest(t, r,id, loginResponse)
	assert.Equal(t, http.StatusOK, UserDataResponse.Code)

	// TEST /alluser ROUTE  (only for admin)
	alluserdataResponse := performAllUserdataRequest(t, r, loginResponse)
	assert.Equal(t, http.StatusOK, alluserdataResponse.Code)
}
func performPublicRequest(t *testing.T, r *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/public", nil)
	assert.NoError(t, err)
	r.ServeHTTP(w, req)
	return w
}

func performRegisterRequest(t *testing.T, r *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	registerData := []byte(`{
		"firstname": "sk",
		"lastname": "raka",
		"email": "skshahra@gmail.com",
		"password": "111187" 
	}`)
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(registerData))
	assert.NoError(t, err)
	r.ServeHTTP(w, req)
	return w
}

func performLoginRequest(t *testing.T, r *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	loginData := []byte(`{
		"email": "skshahra@gmail.com",
		"password": "111187" 
	}`)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(loginData))
	assert.NoError(t, err)
	r.ServeHTTP(w, req)
	return w
}

func performUserdataRequest(t *testing.T, r *gin.Engine,id string, loginResponse *httptest.ResponseRecorder) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/"+id, nil)
	assert.NoError(t, err)
	req.Header.Set("Cookie", loginResponse.Header().Get("Set-Cookie"))
	r.ServeHTTP(w, req)

	return w
}

func performAllUserdataRequest(t *testing.T, r *gin.Engine, loginResponse *httptest.ResponseRecorder) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/alluser", nil)
	assert.NoError(t, err)
	req.Header.Set("Cookie", loginResponse.Header().Get("Set-Cookie"))
	r.ServeHTTP(w, req)
	return w
}