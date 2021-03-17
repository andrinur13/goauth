package main

import (
	"fmt"
	"goauth/auth"
	"goauth/handler"
	"goauth/helper"
	"goauth/product"
	"goauth/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_auth?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")

	authUser := auth.NewService() // jwt

	// user
	repoUser := user.NewRepository(db)
	servUser := user.NewService(repoUser)
	userHandler := handler.NewUserHandler(servUser, authUser)

	// product
	repoProduct := product.NewRepository(db)
	servProduct := product.NewService(repoProduct)
	productHandler := handler.NewProductHandler(servProduct)

	router := gin.Default()
	api := router.Group("/api/v1/")

	// user
	api.POST("/user", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	// product
	api.POST("/product", authMiddleware, productHandler.InsertProduct)
	api.GET("/products", authMiddleware, productHandler.GetAllProducts)

	router.Run()
}

func authMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	// jika tidak ada header auth
	if !strings.Contains(authHeader, "Bearer") {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// jika ada header auth maka validasi token
	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	statusToken, err := auth.NewService().ValidateToken(tokenString)

	if err != nil {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if !statusToken.Valid {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

}
