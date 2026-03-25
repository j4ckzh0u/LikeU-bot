package router

import (
	"net/http"

	"ai-coding-assistant/internal/api/handlers"
	"ai-coding-assistant/internal/api/middleware"

	"github.com/labstack/echo/v4"
)

type Handler = middleware.Handler

func SetupRoutes(e *echo.Echo, h *Handler) {
	api := e.Group("/api/v1")

	api.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    0,
			"message": "success",
		})
	})

	auth := api.Group("/auth")
	auth.POST("/login", handlers.Login)
	auth.POST("/logout", handlers.Logout)
	auth.POST("/refresh", handlers.RefreshToken)

	user := api.Group("/user")
	user.GET("/me", handlers.GetCurrentUser)
	user.PUT("/me", handlers.UpdateUser)

	teams := api.Group("/teams")
	teams.GET("", handlers.ListTeams)
	teams.POST("", handlers.CreateTeam)
	teams.GET("/:id", handlers.GetTeam)
	teams.PUT("/:id", handlers.UpdateTeam)
	teams.DELETE("/:id", handlers.DeleteTeam)
	teams.POST("/:id/members", handlers.AddTeamMember)
	teams.DELETE("/:id/members/:user_id", handlers.RemoveTeamMember)
	teams.GET("/:id/members", handlers.ListTeamMembers)

	projects := api.Group("/projects")
	projects.GET("", handlers.ListProjects)
	projects.POST("", handlers.CreateProject)
	projects.GET("/:id", handlers.GetProject)
	projects.PUT("/:id", handlers.UpdateProject)
	projects.DELETE("/:id", handlers.DeleteProject)

	tasks := api.Group("/tasks")
	tasks.GET("", handlers.ListTasks)
	tasks.POST("", handlers.CreateTask)
	tasks.GET("/:id", handlers.GetTask)
	tasks.PUT("/:id", handlers.UpdateTask)
	tasks.DELETE("/:id", handlers.DeleteTask)
	tasks.PUT("/:id/status", handlers.UpdateTaskStatus)
	tasks.POST("/:id/messages", handlers.SendTaskMessage)
	tasks.GET("/:id/messages", handlers.GetTaskMessages)

	assistant := api.Group("/assistant")
	agents := assistant.Group("/agents")
	agents.GET("", handlers.ListAgents)
	agents.POST("", handlers.CreateAgent)
	agents.GET("/:id", handlers.GetAgent)
	agents.PUT("/:id", handlers.UpdateAgent)
	agents.DELETE("/:id", handlers.DeleteAgent)

	channels := assistant.Group("/channels")
	channels.GET("", handlers.ListChannels)
	channels.POST("", handlers.CreateChannel)
	channels.PUT("/:id", handlers.UpdateChannel)
	channels.PATCH("/:id/status", handlers.UpdateChannelStatus)

	sessions := assistant.Group("/sessions")
	sessions.GET("", handlers.ListSessions)
	sessions.POST("", handlers.CreateSession)
	sessions.GET("/:id", handlers.GetSession)
	sessions.POST("/:id/messages", handlers.SendSessionMessage)
	sessions.GET("/:id/messages", handlers.GetSessionMessages)
	sessions.DELETE("/:id/messages", handlers.ClearSessionMessages)

	skills := assistant.Group("/skills")
	skills.GET("", handlers.ListSkills)
	skills.POST("/:id/install", handlers.InstallSkill)
	skills.DELETE("/:id/install", handlers.UninstallSkill)
	skills.PUT("/:id/config", handlers.UpdateSkillConfig)

	canvas := assistant.Group("/canvas")
	canvas.POST("", handlers.CreateCanvas)
	canvas.GET("/:id", handlers.GetCanvas)
	canvas.POST("/:id/share", handlers.ShareCanvas)
	canvas.PUT("/:id/content", handlers.UpdateCanvasContent)

	voice := assistant.Group("/voice")
	voice.POST("/call", handlers.InitiateVoiceCall)
	voice.GET("/calls/:id", handlers.GetVoiceCallStatus)
	voice.POST("/calls/:id/continue", handlers.ContinueVoiceCall)
	voice.DELETE("/calls/:id", handlers.EndVoiceCall)
	voice.PUT("/config", handlers.UpdateVoiceConfig)

	ai := api.Group("/ai")
	ai.GET("/models", handlers.ListModels)
	ai.POST("/keys", handlers.SetAPIKey)
	ai.DELETE("/keys/:provider", handlers.DeleteAPIKey)

	ide := api.Group("/ide")
	ide.POST("/environments", handlers.CreateIDEEnvironment)
	ide.GET("/environments/:id", handlers.GetIDEEnvironment)
	ide.DELETE("/environments/:id", handlers.CloseIDEEnvironment)

	reviews := api.Group("/reviews")
	reviews.POST("", handlers.CreateReview)
	reviews.GET("", handlers.ListReviews)
	reviews.GET("/:id", handlers.GetReview)
	reviews.POST("/webhook", handlers.SetupWebhook)

	resources := api.Group("/resources")
	hosts := resources.Group("/hosts")
	hosts.GET("", handlers.ListHosts)
	hosts.POST("", handlers.CreateHost)
	hosts.DELETE("/:id", handlers.DeleteHost)

	images := resources.Group("/images")
	images.GET("", handlers.ListImages)
	images.POST("/build", handlers.BuildImage)
}
