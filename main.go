package main

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	cartHttpDeliver "github.com/ddoffy/shopping/cart/delivery/http"
	cartRepo "github.com/ddoffy/shopping/cart/repository"
	cartUcase "github.com/ddoffy/shopping/cart/usecase"
	"github.com/ddoffy/shopping/controller"
	"github.com/gin-gonic/gin"

	"github.com/ddoffy/shopping/middleware"
	"github.com/ddoffy/shopping/model"
	productHttpDeliver "github.com/ddoffy/shopping/product/delivery/http"
	productRepo "github.com/ddoffy/shopping/product/repository"
	productUcase "github.com/ddoffy/shopping/product/usecase"

	promotionHttpDeliver "github.com/ddoffy/shopping/promotion/delivery/http"
	promotionRepo "github.com/ddoffy/shopping/promotion/repository"
	promotionUcase "github.com/ddoffy/shopping/promotion/usecase"

	loginHttpDeliver "github.com/ddoffy/shopping/login/delivery/http"
	loginRepo "github.com/ddoffy/shopping/login/repository"
	loginUcase "github.com/ddoffy/shopping/login/usecase"
	// gin-swagger middleware
	// swagger embed files
)

func init() {
	// setupRouter()
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	r := gin.Default()
	// the jwt middleware
	authMiddleware := middleware.InitMiddleware().AuthMiddleware()
	// API's to get JWT toekn
	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// API's with JWT authentication
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{

		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		// auth.POST("/products", controller.HandlePostProducts)
		// auth.DELETE("/products", controller.HandleDelProducts)

		// auth.GET("/cart", controller.HandleGetCart)
		// auth.POST("/cart", controller.HandlePostCart)
		// auth.DELETE("/cart", controller.HandleDelCart)

		// auth.GET("/promotion", controller.HandleGetPromotion)
		// auth.POST("/promotion", controller.HandlePostPromotion)
		// auth.DELETE("/promotion", controller.HandleDelPromotion)

	}

	// r.Use(Cors())

	// Public API's without JWT authentication
	// HealthCheck
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/", controller.HandlerRoot)
	// r.GET("/products", controller.HandleGetProducts)

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r

}

func main() {
	r := setupRouter()
	db := model.GetDBInstance()
	timeoutContext := time.Duration(10) * time.Second

	prR := productRepo.NewProductRepository(db.SDB)
	pu := productUcase.NewProductUsecase(prR, timeoutContext)
	productHttpDeliver.NewProductHttpHandler(r, pu)

	caR := cartRepo.NewERepository(db.SDB)
	cu := cartUcase.NewEUsecase(caR, timeoutContext)
	cartHttpDeliver.NewCartHttpHandler(r, cu)

	poR := promotionRepo.NewERepository(db.SDB)
	pou := promotionUcase.NewEUsecase(poR, timeoutContext)
	promotionHttpDeliver.NewPromotionHttpHandler(r, pou)

	loR := loginRepo.NewERepository(db.SDB)
	loU := loginUcase.NewEUsecase(loR, timeoutContext)
	loginHttpDeliver.NewLoginHttpHandler(r, loU)

	r.Run(":8080")
}
