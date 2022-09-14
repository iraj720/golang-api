package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"snapp/internal/service"

	"github.com/gin-gonic/gin"
)

type productController struct {
	ps service.IProductService
}

// swagger:model vote
type Product struct {
	// example: 1
	// required: false
	Id uint `json:"id"`
	// The name of product
	// example: TV
	// required: false
	Name string `json:"name"`
	// The value for a vote point
	// example: 6.5
	// required: true
	Point float32 `json:"point"`
}

type products []Product

func NewProductController(ps service.IProductService) productController {
	return productController{ps: ps}
}

// vote godoc
// @Summary This is for voting
// @Description you can vote with id or name (every porduct that matches that conditions will consider your point)
// @ID get-all
// @Tags get-all
// @Router /get-all [get]
func (pc *productController) GetAllProducts(c *gin.Context) {
	products := pc.ps.GetAllProducts(c)
	c.JSON(http.StatusOK, products)
}

// swagger:parameters vote, filter
type _ struct {
	// The body for voting
	// in:body
	// required: true
	Body Product
}

// vote godoc
// @Summary This is for filtering
// @Description you can get Product with id or name
// @ID filter
// @Tags filter
// @Param Body body Product true "filter products"
// @Router /filter [post]
func (pc *productController) GetFilteredProducts(c *gin.Context) {
	var p Product
	c.BindJSON(&p)
	products := pc.ps.GetSpecialProducts(c, p.Id, p.Name, p.Point)
	c.JSON(http.StatusOK, products)
}

// vote godoc
// @Summary This is for voting
// @Description you can vote with id or name (every porduct that matches that conditions will consider your point)
// @ID vote
// @Tags vote
// @Param Body body Product true "The body to create a vote"
// @Router /vote [post]
func (pc *productController) VoteToProduct(c *gin.Context) {
	var p Product
	c.BindJSON(&p)
	if p.Point < 0 || p.Point > 10 {
		c.JSON(http.StatusNotAcceptable, "your point should be between 0 and 10")
	}
	pc.ps.VoteToProduct(c, p.Id, p.Name, p.Point)
	c.JSON(http.StatusOK, nil)
}

// swagger:parameters bulkyVote
type _ struct {
	// The body for voting
	// in:body
	// required: true
	Body []*Product
}

// bulkyVote godoc
// @Summary This is for voting
// @Description you can vote with id or name (every porduct that matches that conditions will consider your point)
// @ID bulkyVote
// @Tags bulkyVote
// @Param Body body products true "The body to create a vote"
// @Router /bulky-vote [post]
func (pc *productController) BulkyVoteToProduct(c *gin.Context) {
	var products products
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = json.Unmarshal(body, &products)
	for _, p := range products {
		if p.Point < 0 || p.Point > 10 {
			c.JSON(http.StatusNotAcceptable, "your point should be between 0 and 10")
		}
		pc.ps.VoteToProduct(c, p.Id, p.Name, p.Point)
	}
	c.JSON(http.StatusOK, nil)
}
