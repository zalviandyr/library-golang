package middleware

import (
	"library-be/app/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
}

func ErrorMiddleware(c *gin.Context, err any) {
	if error, ok := err.(pkg.InternalServerError); ok {
		c.JSON(pkg.NewResponse(http.StatusInternalServerError, error.Message).Build())
		return
	}

	if error, ok := err.(pkg.UnprocessableEntityError); ok {
		c.JSON(pkg.NewResponse(http.StatusUnprocessableEntity, error.Message).Body(error.Field).Build())
		return
	}

	if error, ok := err.(pkg.BadRequestError); ok {
		c.JSON(pkg.NewResponse(http.StatusBadRequest, error.Message).Build())
		return
	}

	if error, ok := err.(pkg.NotFoundError); ok {
		c.JSON(pkg.NewResponse(http.StatusNotFound, error.Message).Build())
		return
	}

	c.JSON(pkg.NewResponse(http.StatusInternalServerError, "Server Error").Body(err).Build())
}
