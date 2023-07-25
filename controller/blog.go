package controller

import (
	"apm/db"
	"apm/dto"
	"apm/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InsertBlog(c echo.Context) error {

	db := db.GetDb()

	db = db.WithContext(c.Request().Context())

	var blogDto dto.BlogDto

	if err := c.Bind(&blogDto); err != nil {
		return err
	}

	blogModel := model.Blog{
		Title:       blogDto.Title,
		Description: blogDto.Description,
	}

	db.Save(&blogModel)

	return c.JSON(http.StatusCreated, blogModel)

}
