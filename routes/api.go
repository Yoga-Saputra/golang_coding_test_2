package routes

import (
	"codingTest/app/handler"
	"codingTest/app/members"
	"codingTest/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
)

func InitApi(state overseer.State) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())

	db := config.ConnectionDB()

	memberRepository := members.NewRepository(db)
	memberService := members.NewService(memberRepository)
	memberHandler := handler.NewMemberHandler(memberService)

	api := router.Group("/api")

	member := api.Group("/members")
	member.GET("/", memberHandler.GetALlMembers)
	member.GET("/:id", memberHandler.GetMember)
	member.POST("/", memberHandler.Save)
	member.PUT("/:id", memberHandler.Update)
	member.DELETE("/:id", memberHandler.Delete)

	router.Run(":3000")

	http.Serve(state.Listener, router)
}
