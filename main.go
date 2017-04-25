// Package classification Petstore API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<aa@bb.cc> http://john.doe.com
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     - api_key:
//          type: apiKey
//          name: KEY
//          in: header
//
// swagger:meta
//go:generate swagger generate spec -o swagger.json
package main

import (
	"fmt"
	"time"

	"github.com/houjunchen/go-swagger-test/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	fmt.Println("vim-go")

	router := gin.Default()

	router.GET("/health", GetHealth)
	router.GET("/products", GetProducts)
	router.GET("/pets", GetPets)

	router.Run()
}

// GetHealth swagger:route GET /health getHeahth
//
// Lists pets filtered by some parameters.
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
// Responses:
//   default: genericError
//       200: ResHealth
//       422: validationError
func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "healthy",
	})
}

// GetPoints swagger:route GET /products getProducts
//
// Lists all products
//
// Responses:
//   default: genericError
//       200: resProduct
//       422: validationError
func GetProducts(c *gin.Context) {
	c.JSON(200, models.ResProduct{})
}

// GetPets swagger:route GET /pets pets listPets
//
// Lists the pets known to the store.
//
// By default it will only lists pets that are available for sale.
// This can be changed with the status flag.
//
// Responses:
// 		default: genericError
// 		    200: []pet
func GetPets(c *gin.Context) {
	// some actual stuff should happen in here
}

// A Pet is the main product in the store.
// It is used to describe the animals available in the store.
//
// swagger:model ResHealth
type ResHealth struct {
	// The health message
	//
	// required: true
	Message string `json:"message"`
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32 `json:"code"`
		Message error `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// A Pet is the main product in the store.
// It is used to describe the animals available in the store.
//
// swagger:model pet
type Pet struct {
	// The id of the pet.
	//
	// required: true
	ID int64 `json:"id"`

	// The name of the pet.
	//
	// required: true
	// pattern: \w[\w-]+
	// minimum length: 3
	// maximum length: 50
	Name string `json:"name"`

	// The photo urls for the pet.
	// This only accepts jpeg or png images.
	//
	// items pattern: \.(jpe?g|png)$
	PhotoURLs []string `json:"photoUrls,omitempty"`

	// The current status of the pet in the store.
	Status string `json:"status,omitempty"`

	// The pet's birthday
	//
	// swagger:strfmt date
	Birthday time.Time `json:"birthday"`
}
