package controller

import (
	"library-be/app/domain/dto"
	"library-be/app/pkg"
	"library-be/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(
	service *service.AuthService,
) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (cont *AuthController) Login(c *gin.Context) {
	var login dto.LoginDto
	if err := pkg.ExtractValidateData(c, &login); err != nil {
		panic(*err)
	}

	if err := cont.service.Login(login); err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Success login").Build())
}

func (cont *AuthController) Register(c *gin.Context) {
	var register dto.RegisterDto
	if err := pkg.ExtractValidateData(c, &register); err != nil {
		panic(*err)
	}

	if err := cont.service.Register(register); err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Success register").Build())
}
