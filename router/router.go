package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"example.com/mod/controllers"
	"example.com/mod/database"
	"fmt"
	"strconv"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func routes() {
	router := setupRouter()
	database.ConnectDatabase("DB_NAME")
}