package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"access_token":  "demo-token",
			"refresh_token": "demo-refresh-token",
			"expires_in":    86400,
			"user": map[string]interface{}{
				"id":    "1",
				"email": "demo@example.com",
				"name":  "Demo User",
			},
		},
	})
}

func Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func RefreshToken(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"access_token": "new-token",
			"expires_in":   86400,
		},
	})
}

func GetCurrentUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"id":         "1",
			"email":      "demo@example.com",
			"name":       "Demo User",
			"avatar":     "",
			"created_at": "2024-01-01T00:00:00Z",
		},
	})
}

func UpdateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ListTeams(c echo.Context) error {
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

func CreateTeam(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func GetTeam(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UpdateTeam(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func DeleteTeam(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func AddTeamMember(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func RemoveTeamMember(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ListTeamMembers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items": []interface{}{},
		},
	})
}
