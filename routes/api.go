package routes

import (
	"codingTest/app/handler"
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
	memberHandler := handler.NewMemberHandler(memberService)

	productRepository := products.NewRepository(db)
	productService := products.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	rProductRepository := reviewproduct.NewRepository(db)
	rProductService := reviewproduct.NewService(rProductRepository)
	rProductHandler := handler.NewReviewProductHandler(rProductService)

	api := router.Group("/api")

	member := api.Group("/members")
	member.GET("/", memberHandler.GetALlMembers)
	member.GET("/:id", memberHandler.GetMember)
	member.POST("/", memberHandler.Save)
	member.PUT("/:id", memberHandler.Update)
	member.DELETE("/:id", memberHandler.Delete)

	rProduct := api.Group("/review-products")
	rProduct.GET("/", rProductHandler.GetAllReviewProducts)

	product := api.Group("/products")
	product.GET("/", productHandler.GetALlProduct)
	product.GET("/:id", productHandler.GetProduct)
	product.POST("/", productHandler.Save)
	product.PUT("/:id", productHandler.Update)
	product.DELETE("/:id", productHandler.Delete)

	router.Run(":3000")

	http.Serve(state.Listener, router)
}
