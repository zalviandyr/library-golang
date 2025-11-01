package controller

import (
	"library-be/app/domain/dto"
	"library-be/app/pkg"
	"library-be/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PublisherController struct {
	service *service.PublisherService
}

func NewPublisherController(
	service *service.PublisherService,
) *PublisherController {
	return &PublisherController{
		service: service,
	}
}

func (cont *PublisherController) Index(c *gin.Context) {
	publishers := cont.service.FindAll()

	c.JSON(pkg.NewResponse(http.StatusOK, "All publisher").Body(publishers).Build())
}

func (cont *PublisherController) Show(c *gin.Context) {
	id := c.Param("id")
	publisher, err := cont.service.Find(id)
	if err != nil {
		panic(pkg.NotFoundError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Show publisher").Body(publisher).Build())

}

func (cont *PublisherController) Create(c *gin.Context) {
	var publisher dto.PublisherDto
	if err := pkg.ExtractValidateData(c, &publisher); err != nil {
		panic(*err)
	}

	result, err := cont.service.Save(publisher)
	if err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusCreated, "Created publisher").Body(result).Build())
}

func (cont *PublisherController) Update(c *gin.Context) {
	id := c.Param("id")
	uuid := uuid.MustParse(id)
	publisher := dto.PublisherDto{ID: &uuid}

	if err := pkg.ExtractValidateData(c, &publisher); err != nil {
		panic(*err)
	}

	result, err := cont.service.Save(publisher)
	if err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Updated publisher").Body(result).Build())
}

func (cont *PublisherController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := cont.service.Delete(id); err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Deleted publisher").Build())
}
