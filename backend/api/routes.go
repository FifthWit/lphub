package api

import (
	"lphub/handlers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		// Swagger
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		v1.GET("/", func(c *gin.Context) {
			c.File("docs/index.html")
		})
		// Tokens, login
		v1.GET("/token", handlers.GetCookie)
		v1.DELETE("/token", handlers.DeleteCookie)
		v1.GET("/login", handlers.Login)
		// Users, profiles
		v1.GET("/profile", IsAuthenticated, handlers.Profile)
		v1.PUT("/profile", IsAuthenticated, handlers.UpdateCountryCode)
		v1.POST("/profile", IsAuthenticated, handlers.UpdateUser)
		v1.GET("/users/:userid", IsAuthenticated, handlers.FetchUser)
		// Maps
		// - Summary
		v1.GET("/maps/:mapid/summary", handlers.FetchMapSummary)
		v1.POST("/maps/:mapid/summary", IsAuthenticated, handlers.CreateMapSummary)
		v1.PUT("/maps/:mapid/summary", IsAuthenticated, handlers.EditMapSummary)
		v1.DELETE("/maps/:mapid/summary", IsAuthenticated, handlers.DeleteMapSummary)
		v1.PUT("/maps/:mapid/image", IsAuthenticated, handlers.EditMapImage)
		// - Leaderboards
		v1.GET("/maps/:mapid/leaderboards", handlers.FetchMapLeaderboards)
		v1.POST("/maps/:mapid/record", IsAuthenticated, handlers.CreateRecordWithDemo)
		v1.DELETE("/maps/:mapid/record/:recordid", IsAuthenticated, handlers.DeleteRecord)
		v1.GET("/demos", handlers.DownloadDemoWithID)
		// - Discussions
		v1.GET("/maps/:mapid/discussions", handlers.FetchMapDiscussions)
		v1.GET("/maps/:mapid/discussions/:discussionid", handlers.FetchMapDiscussion)
		v1.POST("/maps/:mapid/discussions", IsAuthenticated, handlers.CreateMapDiscussion)
		v1.POST("/maps/:mapid/discussions/:discussionid", IsAuthenticated, handlers.CreateMapDiscussionComment)
		v1.PUT("/maps/:mapid/discussions/:discussionid", IsAuthenticated, handlers.EditMapDiscussion)
		v1.DELETE("/maps/:mapid/discussions/:discussionid", IsAuthenticated, handlers.DeleteMapDiscussion)
		// Rankings, search
		v1.GET("/rankings/lphub", handlers.RankingsLPHUB)
		v1.GET("/rankings/steam", handlers.RankingsSteam)
		v1.GET("/search", handlers.SearchWithQuery)
		// Games, chapters, maps
		v1.GET("/games", handlers.FetchGames)
		v1.GET("/games/:gameid", handlers.FetchChapters)
		v1.GET("/chapters/:chapterid", handlers.FetchChapterMaps)
		v1.GET("/games/:gameid/maps", handlers.FetchMaps)
		// Logs
		v1.GET("/logs/score", handlers.ScoreLogs)
		// v1.GET("/logs/mod", IsAuthenticated, handlers.ModLogs)
	}
}
