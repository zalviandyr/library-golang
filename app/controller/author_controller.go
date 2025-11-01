package controller

import (
	"library-be/app/domain/dto"
	"library-be/app/pkg"
	"library-be/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthorController struct {
	service *service.AuthorService
}

func NewAuthorController(
	service *service.AuthorService,
) *AuthorController {
	return &AuthorController{
		service: service,
	}
}

func (cont *AuthorController) Index(c *gin.Context) {
	authors := cont.service.FindAll()

	c.JSON(pkg.NewResponse(http.StatusOK, "All author").Body(authors).Build())
}

func (cont *AuthorController) Show(c *gin.Context) {
	id := c.Param("id")
	author, err := cont.service.Find(id)
	if err != nil {
		panic(pkg.NotFoundError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Show author").Body(author).Build())

}

func (cont *AuthorController) Create(c *gin.Context) {
	var author dto.AuthorDto
	if err := pkg.ExtractValidateData(c, &author); err != nil {
		panic(*err)
	}

	result, err := cont.service.Save(author)
	if err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusCreated, "Created author").Body(result).Build())
}

func (cont *AuthorController) Update(c *gin.Context) {
	id := c.Param("id")
	uuid := uuid.MustParse(id)
	author := dto.AuthorDto{ID: &uuid}

	if err := pkg.ExtractValidateData(c, &author); err != nil {
		panic(*err)
	}

	result, err := cont.service.Save(author)
	if err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Updated author").Body(result).Build())
}

func (cont *AuthorController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := cont.service.Delete(id); err != nil {
		panic(pkg.BadRequestError{Message: err.Error()})
	}

	c.JSON(pkg.NewResponse(http.StatusOK, "Deleted author").Build())
}
