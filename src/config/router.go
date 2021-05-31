package config

import (
	"github.com/gin-gonic/gin"

	"artics-api/src/registry"
)

func Router(reg *registry.Registry) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// User
		api.POST("/users", reg.V1User.Create)
		api.GET("/users/:id", reg.V1User.Show)
		api.PATCH("/users", reg.V1User.Update)
		api.DELETE("/users/suspend", reg.V1User.Suspend)
		api.GET("/users/:id/followings", reg.V1User.Followings)

		// Follow
		api.POST("/follows", reg.V1Follow.Create)
		api.DELETE("/follows", reg.V1Follow.Delete)

		// Content
		// api.POST("/contents", reg.ContentHandler.CreateContent)
		api.GET("/contents/favorites", reg.V1Content.Favorites)
		// api.GET("/contents/:id", reg.ContentHandler.GetContent)
		// api.PATCH("/contents/:id", reg.ContentHandler.UpdateContent)
		// api.DELETE("/contents/:id", reg.ContentHandler.DestroyContent)

		// // お気に入り
		// api.POST("/favorites", reg.FavoriteHandler.CreateFavorite)
		// api.GET("/favorites", reg.FavoriteHandler.GetFavorites)
		// api.DELETE("/favorites", reg.FavoriteHandler.DestroyFavorite)

		// // カテゴリ
		// api.GET("/categories", reg.CategoryHandler.GetCategories)
	}

	return r
}
