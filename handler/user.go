package handler

import (
	"goauth/auth"
	"goauth/helper"
	"goauth/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authUser    auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.UserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponse("failed", http.StatusUnprocessableEntity, "failed to registered user", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("failed", http.StatusUnprocessableEntity, "failed to registered user", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// jika berhasil
	dataOutput := helper.OutputDataUser{
		Username: data.Username,
		Name:     data.Name,
		Email:    data.Email,
	}

	response := helper.APIResponse("success", http.StatusOK, "success registered new user", dataOutput)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponse("failed", http.StatusUnprocessableEntity, "failed to login", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	dataLogin, err := h.userService.Login(input.Email, input.Password)

	if err != nil {
		response := helper.APIResponse("failed", http.StatusUnprocessableEntity, "failed to login", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// jika berhasil login
	// generate token
	tokenLogin, err := h.authUser.GenerateToken(dataLogin.ID)

	if err != nil {
		response := helper.APIResponse("failed", http.StatusUnprocessableEntity, "failed to login", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	dataOutputLogin := helper.OutputDataUserLogin{
		Username: dataLogin.Username,
		Name:     dataLogin.Name,
		Email:    dataLogin.Email,
		Token:    tokenLogin, // nanti diisi sama JWT
	}

	response := helper.APIResponse("success", http.StatusOK, "success login user", dataOutputLogin)
	c.JSON(http.StatusOK, response)
}
