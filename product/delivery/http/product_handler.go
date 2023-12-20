package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ddoffy/shopping/middleware"
	"github.com/ddoffy/shopping/model"
	productUcase "github.com/ddoffy/shopping/product"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpProductHandler struct {
	AUsecase productUcase.ProductUsecase
}

func NewProductHttpHandler(r *gin.Engine, us productUcase.ProductUsecase) {
	handler := &HttpProductHandler{
		AUsecase: us,
	}

	// API's with JWT authentication
	// auth := r.Group("/auth")
	// // the jwt middleware
	// authMiddleware := middleware.InitMiddleware().AuthMiddleware()
	// auth.Use(authMiddleware.MiddlewareFunc())
	// {
	// 	r.GET("/products", handler.Fetch)
	// 	r.POST("/products", handler.Store)
	// 	r.DELETE("/products", handler.Delete)
	// }
	auth := r.Group("/auth")
	authMiddleware := middleware.InitMiddleware().AuthMiddleware()
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/products", handler.Fetch)
		auth.POST("/products", handler.Store)
		auth.DELETE("/products", handler.Delete)

	}

}

// API for get product
func (a *HttpProductHandler) Fetch(c *gin.Context) {

	ctx := c
	fmt.Println("In Fetch")
	listP, err := a.AUsecase.Fetch(ctx)
	fmt.Println(listP)
	fmt.Println("In Fetch")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("In Fetch")
	c.JSON(http.StatusOK, gin.H{"Products": listP})
}

// Validate request
func isRequestValid(m *model.Product) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// API for add product
func (a *HttpProductHandler) Store(c *gin.Context) {
	var product model.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if ok, err := isRequestValid(&product); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c
	fmt.Println("Product:", product)
	pr, err := a.AUsecase.Store(ctx, &product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": pr})
}

// API for delete product
func (a *HttpProductHandler) Delete(c *gin.Context) {
	idP, err := strconv.Atoi(c.Query("id"))
	id := int(idP)
	ctx := c

	_, err = a.AUsecase.Delete(ctx, id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
