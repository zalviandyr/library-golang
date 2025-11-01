package controller

import (
	"library-be/app/domain/dto"
	"library-be/app/pkg"
	"library-be/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookController struct {
	service *service.BookService
}

func NewBookController(
	service *service.BookService,
) *BookController {
	return &BookController{
		service: service,
	}
}

func (cont *BookController) Index(c *gin.Context) {
	publishers := cont.service.FindAll()

	c.JSON(pkg.NewResponse(http.StatusOK, "All book").Body(publishers).Build())
}

func (cont *BookController) Show(c *gin.Context) {
	id := c.Param("id")
	book, err := cont.service.Find(id)
	if err != nil {
		panic(pkg.NotFoundError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Show book").Body(book).Build())

}

func (cont *BookController) Create(c *gin.Context) {
	var book dto.BookDto
	if err := pkg.ExtractValidateData(c, &book); err != nil {
		panic(*err)
	}

	result, err := cont.service.Save(book)
	if err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusCreated, "Created book").Body(result).Build())
}

func (cont *BookController) Update(c *gin.Context) {
	id := c.Param("id")
	uuid := uuid.MustParse(id)
	book := dto.BookDto{ID: &uuid}

	if err := pkg.ExtractValidateData(c, &book); err != nil {
		panic(*err)
	}

	result, err := cont.service.Save(book)
	if err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Updated book").Body(result).Build())
}

func (cont *BookController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := cont.service.Delete(id); err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Deleted book").Build())
}
