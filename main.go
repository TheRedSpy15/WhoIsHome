package main

import (
	"crypto/subtle"
	"database/sql"

	"github.com/TheRedSpy15/WhoIsHome/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sql.DB
var csv [4]string

func main() {
	e := echo.New()
	csv = utils.ReadCsvFile("OpenCV/database.csv")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(csv[2])) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(csv[3])) == 1 {
			return true, nil
		}
		return false, nil
	}))
	e.Use(middleware.BodyLimit("2M"))
	e.Debug = true

	// Routes
	e.GET("/", index)

	e.Logger.Fatal(e.Start(":1323"))
}
