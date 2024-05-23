package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const COOKIE_LIFE_TIME_HOUERS = 1

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

type Product struct {
	gorm.Model
	ProductID   uint
	Name        string
	PriceRappen int
}

type ProductPosition struct {
	gorm.Model
	ProductPositionID uint
	Product           Product `gorm:"foreignKey:ProductID"`
	PurchasedQuantity uint
}

type User struct {
	gorm.Model
	UserID    uint
	FirstName string
	LastName  string
	City      string
	ZipCode   string
	Street    string
	Email     string
}

type Order struct {
	gorm.Model
	OrderID          uint
	Quantity         uint
	User             User              `gorm:"foreignKey:UserID"`
	ProductPositions []ProductPosition `gorm:"foreignKey:ProductPositionID"`
	PriceRappen      uint
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&ProductPosition{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Order{})

	return db
}

func Add(a, b int) int {
	return a + b
}

func serialize_cart(cart map[string]uint) string {
	serialized, err := json.Marshal(cart)
	if err != nil {
		log.Fatal("failed to serialize cart")
		os.Exit(1)
	}
	return string(serialized)
}

func deserialize_cart(cart string) map[string]uint {
	deserialized := make(map[string]uint)
	err := json.Unmarshal([]byte(cart), &deserialized)
	if err != nil {
		return make(map[string]uint)
	}
	return deserialized
}

func add_to_cart(c echo.Context, product string, quantity uint) error {
	cookie, err := c.Cookie("cart")
	if err != nil {
		cookie = new(http.Cookie)
		cookie.Name = "cart"
		emptyCart := make(map[string]uint)
		cookie.Value = url.QueryEscape(serialize_cart(emptyCart))
	}

	unescaped_value, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		log.Fatal("failed to unescape cookie value")
		os.Exit(1)
	}
	cart := deserialize_cart(unescaped_value)
	cart[product] += quantity

	cookie.Value = url.QueryEscape(serialize_cart(cart))
	cookie.Expires = time.Now().Add(COOKIE_LIFE_TIME_HOUERS * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "Product added to cart")
}

func main() {
	db := initDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get database connection")
		os.Exit(1)
	}
	defer sqlDB.Close()

	e := echo.New()
	//e.Use(middleware.Logger())

	e.Renderer = newTemplates()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/contact", func(c echo.Context) error {
		return c.Render(http.StatusOK, "contact", nil)
	})

	e.GET("/shop", func(c echo.Context) error {
		return c.Render(http.StatusOK, "shop", nil)
	})

	e.GET("/cart", func(c echo.Context) error {
		return c.Render(http.StatusOK, "cart", nil)
	})

	e.GET("/cart/addtocart/:id", func(c echo.Context) error {
		product_id := c.Param("id")

		add_to_cart(c, product_id, 1)
		return c.String(http.StatusOK, "Product added to cart")
	})

	e.GET("/cart/addtocart/:id/:quantity", func(c echo.Context) error {
		product_id := c.Param("id")
		quantity, err := strconv.ParseUint(c.Param("quantity"), 10, 32)
		if err != nil {
			return err
		}

		add_to_cart(c, product_id, uint(quantity))
		return c.String(http.StatusOK, "Product added to cart")
	})

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":1234"))
}
