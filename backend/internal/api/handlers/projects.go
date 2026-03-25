package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListProjects(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items":     []interface{}{},
			"total":     0,
			"page":      1,
			"page_size": 20,
		},
	})
}

func CreateProject(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func GetProject(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UpdateProject(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func DeleteProject(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ListTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items":     []interface{}{},
			"total":     0,
			"page":      1,
			"page_size": 20,
		},
	})
}

func CreateTask(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func GetTask(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UpdateTask(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func DeleteTask(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UpdateTaskStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func SendTaskMessage(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func GetTaskMessages(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items": []interface{}{},
		},
	})
}
