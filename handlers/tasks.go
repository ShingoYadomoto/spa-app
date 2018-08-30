package handlers

import (
	"net/http"

	"github.com/ShingoYadomoto/spa-app/models"
	"github.com/labstack/echo"
)

// GetTasks endpoint
func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetTasks())
}
