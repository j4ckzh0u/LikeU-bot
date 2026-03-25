package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListModels(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": []map[string]interface{}{
			{"id": "gpt-4o", "name": "GPT-4o", "provider": "openai", "enabled": true},
			{"id": "gpt-4o-mini", "name": "GPT-4o Mini", "provider": "openai", "enabled": true},
			{"id": "claude-3-5-sonnet", "name": "Claude 3.5 Sonnet", "provider": "anthropic", "enabled": true},
			{"id": "claude-3-5-haiku", "name": "Claude 3.5 Haiku", "provider": "anthropic", "enabled": true},
			{"id": "qwen-max", "name": "Qwen Max", "provider": "aliyun", "enabled": true},
			{"id": "glm-4-plus", "name": "GLM-4-Plus", "provider": "zhipu", "enabled": true},
		},
	})
}

func SetAPIKey(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func DeleteAPIKey(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func CreateIDEEnvironment(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func GetIDEEnvironment(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func CloseIDEEnvironment(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func CreateReview(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ListReviews(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items": []interface{}{},
		},
	})
}

func GetReview(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func SetupWebhook(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ListHosts(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items": []interface{}{},
		},
	})
}

func CreateHost(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func DeleteHost(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ListImages(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items": []interface{}{},
		},
	})
}

func BuildImage(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}
