package controllers

import (
	"berita/configs"
	"berita/models/base"
	"berita/models/news"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeleteNewsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse {
			Status: false,
			Message: "id tidak di temukan",
			Data: nil,
		})
	}
	result := configs.DB.Delete(&news.News{}, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusNoContent, base.BaseResponse {
		Status: true,
		Message: "Success delete",
		Data: nil,
	})
}

func AddNewsController(c echo.Context) error {
	var newsRequestAdd news.NewsRequestAdd
	c.Bind(&newsRequestAdd)
	
	if newsRequestAdd.Title == ""{
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "title kosong",
			Data: nil,
		})
	}
	var newsDB news.News
	newsDB = newsDB.MapFromRequestAdd(newsRequestAdd)
	configs.DB.Create(&newsDB)
	result :=  configs.DB.Create(&newsDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status: false,
			Message: "title kosong",
			Data: nil,
	})

	
}
var newsResponse news.NewsResponse
newsResponse.MapFromRequestAdd(newsDB)
return c.JSON(http.StatusOK, base.BaseResponse{
	Status: true,
	Message: "suses",
	Data: newsResponse,
}) 
}

func GetNewsControllers(c echo.Context) error {
	var newsDB []news.News
	result := configs.DB.Find(&newsDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)

	}
	var newsResponse []news.NewsResponse
	newsResponse = news.MapFromDatabaseList(newsDB)

	return c.JSON(http.StatusOK, base.BaseResponse {
		Status: true,
		Message: "Success get News",
		Data: newsResponse,
	})
}