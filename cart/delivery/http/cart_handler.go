package http

import (
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	cartUcase "github.com/ddoffy/shopping/cart"
	"github.com/ddoffy/shopping/middleware"
	"github.com/ddoffy/shopping/model"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

// API struct with business logic
type HttpCartHandler struct {
	AUsecase cartUcase.EUsecase
}

// To create new API handler with business logic
func NewCartHttpHandler(r *gin.Engine, us cartUcase.EUsecase) {
	handler := &HttpCartHandler{
		AUsecase: us,
	}
	// Authentication using JWT
	auth := r.Group("/auth")
	authMiddleware := middleware.InitMiddleware().AuthMiddleware()
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/cart", handler.Fetch)
		auth.POST("/cart", handler.Store)
		auth.DELETE("/cart", handler.Delete)

	}

}

// Get cart item for user
func (a *HttpCartHandler) Fetch(c *gin.Context) {

	ctx := c
	fmt.Println("In Fetch")
	claims := jwt.ExtractClaims(c)
	// fmt.Println("user:", claims["id"].(string))
	user := claims["id"].(string)
	listC, err := a.AUsecase.Fetch(ctx, user)
	total := a.AUsecase.GetTotalCartValue(listC)
	fmt.Println(listC)
	fmt.Println("In Fetch")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("In Fetch")
	c.JSON(http.StatusOK, gin.H{"Cart": listC, "Total": total})
}

// Validate request
func isRequestValid(m *model.Cart) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Add a new cart item to user
func (a *HttpCartHandler) Store(c *gin.Context) {
	var cart model.Cart
	err := c.BindJSON(&cart)
	claims := jwt.ExtractClaims(c)
	// fmt.Println("user:", claims["id"].(string))
	user := claims["id"].(string)
	cart.Code = user
	// Binding error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// Validate request
	if ok, err := isRequestValid(&cart); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c
	fmt.Println("Cart:", cart)
	// Call business logic
	pr, err := a.AUsecase.Store(ctx, &cart)
	// Handle error from business logic
	if err != nil {
		fmt.Println("handler error:" + err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}
	c.JSON(http.StatusOK, gin.H{"success": pr})
}

// Delete a cart item
// TODO : verify whether user and id are matching with request value
func (a *HttpCartHandler) Delete(c *gin.Context) {
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
