package routes

import (
	"codingTest/app/handler"
	likereview "codingTest/app/likeReview"
	"codingTest/app/members"
	"codingTest/app/products"
	reviewproduct "codingTest/app/reviewProduct"
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

	rProductRepository := reviewproduct.NewRepository(db)
	rProductService := reviewproduct.NewService(rProductRepository)

	lRev := likereview.NewRepository(db)
	Lservice := likereview.NewService(lRev)

	productRepository := products.NewRepository(db)
	productService := products.NewService(productRepository)

	memberHandler := handler.NewMemberHandler(memberService, rProductService)

	productHandler := handler.NewProductHandler(productService, rProductService)

	rProductHandler := handler.NewReviewProductHandler(rProductService, productService, memberService)

	likeRevHamdler := handler.NewLikeReviewHandler(Lservice, rProductService, memberService)

	api := router.Group("/api")

	member := api.Group("/members")
	member.GET("/", memberHandler.GetALlMembers)
	member.GET("/:id", memberHandler.GetMember)
	member.POST("/", memberHandler.Save)
	member.PUT("/:id", memberHandler.Update)
	member.DELETE("/:id", memberHandler.Delete)

	rProduct := api.Group("/review-products")
	rProduct.GET("/", rProductHandler.GetAllReviewProducts)
	rProduct.POST("/", rProductHandler.Save)

	api.POST("/like", likeRevHamdler.LikeRev)
	api.DELETE("/dislike", likeRevHamdler.Dislike)

	product := api.Group("/products")
	product.GET("/", productHandler.GetALlProduct)
	product.GET("/:id", productHandler.GetProduct)
	product.POST("/", productHandler.Save)
	product.PUT("/:id", productHandler.Update)
	product.DELETE("/:id", productHandler.Delete)

	router.Run(":3000")

	http.Serve(state.Listener, router)
}
