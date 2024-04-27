package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
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


	e.Static("/static", "static")


	e.Logger.Fatal(e.Start(":1234"))
}
