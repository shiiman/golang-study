package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = NewValidator()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.POST("/path_param/:path", pathParam)
	e.POST("/query_param", queryParam)
	e.POST("/form_param", formParam)
	e.GET("/bind_param/:id", bindParam)

	e.Logger.Fatal(e.Start(":1323"))
}

func saveUser(c echo.Context) error {
	fmt.Println("POST")
	return c.String(http.StatusOK, "POST")
}

func getUser(c echo.Context) error {
	fmt.Println("GET")
	return c.String(http.StatusOK, "GET")
}

func updateUser(c echo.Context) error {
	fmt.Println("PUT")
	return c.String(http.StatusOK, "PUT")
}

func deleteUser(c echo.Context) error {
	fmt.Println("DELETE")
	return c.String(http.StatusOK, "DELETE")
}

// =============================================================================

// Path Parameters.
func pathParam(c echo.Context) error {
	fmt.Println("pathParam")
	path := c.Param("path")
	return c.String(http.StatusOK, path)
}

// Query Parameters.
func queryParam(c echo.Context) error {
	fmt.Println("queryParam")
	query := c.QueryParam("query")
	return c.String(http.StatusOK, query)
}

// Form Parameters.
func formParam(c echo.Context) error {
	fmt.Println("formParam")
	form := c.FormValue("form")
	return c.String(http.StatusOK, form)
}

// =============================================================================

type User struct {
	ID   string `param:"id"`
	Name string `query:"name" validate:"required"`
	// Mail   string `form:"mail"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Token  string `header:"token"`
}

// Bind Parameters.
func bindParam(c echo.Context) error {
	fmt.Println("bindParam")

	var user User
	c.Bind(&user)
	(&echo.DefaultBinder{}).BindHeaders(c, &user)

	if err := c.Validate(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}

	return c.JSON(http.StatusOK, user)
}

// =============================================================================

// CustomValidator
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator
func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

// Validate validate
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
