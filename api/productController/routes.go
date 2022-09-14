package api

import "github.com/gin-gonic/gin"

func (pc *productController) RegisterRoutes(router *gin.RouterGroup) {

	router.GET("/get-all", pc.GetAllProducts)
	router.POST("/filter", pc.GetFilteredProducts)
	router.POST("/vote", pc.VoteToProduct)
	router.POST("/bulky-vote", pc.BulkyVoteToProduct)

	// router.GET("/", UserRetrieve)
	// router.PUT("/", UserUpdate)

	// router.GET("/:username", ProfileRetrieve)
	// router.POST("/:username/follow", ProfileFollow)
	// router.DELETE("/:username/follow", ProfileUnfollow)
}
