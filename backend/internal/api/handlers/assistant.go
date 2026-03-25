package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListAgents(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items": []map[string]interface{}{
				{
					"id":      "dev",
					"name":    "开发助手",
					"type":    "dev",
					"builtin": true,
					"config":  map[string]interface{}{"model": "gpt-4o"},
				},
				{
					"id":      "review",
					"name":    "审查助手",
					"type":    "review",
					"builtin": true,
					"config":  map[string]interface{}{"model": "gpt-4o"},
				},
			},
		},
	})
}

func CreateAgent(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func GetAgent(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UpdateAgent(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func DeleteAgent(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ListChannels(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items": []map[string]interface{}{
				{
					"id":      "web",
					"type":    "web",
					"name":    "Web 聊天",
					"enabled": true,
				},
			},
		},
	})
}

func CreateChannel(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UpdateChannel(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UpdateChannelStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ListSessions(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items": []interface{}{},
		},
	})
}

func CreateSession(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func GetSession(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func SendSessionMessage(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func GetSessionMessages(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items": []interface{}{},
		},
	})
}

func ClearSessionMessages(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ListSkills(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"items": []map[string]interface{}{
				{"id": "browser", "name": "Browser", "version": "1.0.0", "builtin": true, "enabled": true},
				{"id": "git", "name": "Git", "version": "1.0.0", "builtin": true, "enabled": true},
				{"id": "terminal", "name": "Terminal", "version": "1.0.0", "builtin": true, "enabled": true},
				{"id": "file", "name": "File", "version": "1.0.0", "builtin": true, "enabled": true},
				{"id": "canvas", "name": "Canvas", "version": "1.0.0", "builtin": true, "enabled": true},
				{"id": "voice", "name": "Voice", "version": "1.0.0", "builtin": true, "enabled": true},
			},
		},
	})
}

func InstallSkill(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UninstallSkill(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UpdateSkillConfig(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func CreateCanvas(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"id":  "canvas-1",
			"url": "http://localhost:18793/canvas/abc123",
		},
	})
}

func GetCanvas(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ShareCanvas(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UpdateCanvasContent(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func InitiateVoiceCall(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func GetVoiceCallStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func ContinueVoiceCall(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func EndVoiceCall(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

func UpdateVoiceConfig(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}
