package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Item struct {
	prodID uint
	prodName string
	price_rp int
}

type AmountItem struct {
	amount uint
	item Item
}

func (aItem AmountItem) totalPrice() int{
	return aItem.item.price_rp * int(aItem.amount)
}


type Cart struct {
	items []AmountItem
}

func (c Cart) totalPrice() int{
	acc := 0
	for _,aItem := range c.items{
		acc += aItem.totalPrice()
	}
	return acc
}

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("*.html")),
	}
}

func setUpUrlHandlers() *echo.Echo {
	e := echo.New()
	//e.Use(middleware.Logger())

	e.Renderer = newTemplates()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/HelloWorld",func(c echo.Context) error{
		return c.String(http.StatusOK,"Hello World")
	})
	
	e.Static("/static", "static")

	return e;
}

func main() {
	e := setUpUrlHandlers()

	e.Logger.Fatal(e.Start(":1234"))
}
