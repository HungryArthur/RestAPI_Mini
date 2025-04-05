package handlers

import (
	"RestAPI/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var db *gorm.DB

func GetHandler(c echo.Context) error {
	var messages []models.Message

	if err := db.Find(&messages).Error; err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Status:  "Error",
			Message: "Could not find the messages",
		})
	}

	return c.JSON(http.StatusOK, &messages)
}

func PostHandler(c echo.Context) error {
	var message models.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}

	if err := db.Create(&message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Could not create the message",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Message was successfully created",
	})
}

func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	var updatedMessage models.Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	if err := db.Model(&models.Message{}).Where("id = ?", id).Update("text", updatedMessage.Text).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Message was updated",
	})
}

func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	if err := db.Delete(&models.Message{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Could not delete the message",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Message was deleted",
	})
}
