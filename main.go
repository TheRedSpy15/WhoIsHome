package main

import (
	"crypto/subtle"
	"os"

	"github.com/TheRedSpy15/WhoIsHome/controllers"
	"github.com/TheRedSpy15/WhoIsHome/utils"
	"github.com/akamensky/argparse"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	parser := argparse.NewParser("WhoIsHome - Web Server", "Connects to MySQL database and securely serves the data")

	// Create flags
	d := parser.Flag("d", "Debug", &argparse.Options{Required: false, Help: "Enable debug mode"})

	err := parser.Parse(os.Args) // parse arguments
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()
	csv := utils.ReadCsvFile("OpenCV/database.csv")

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
	e.Debug = *d

	// Routes
	e.GET("/", controllers.Index)

	e.Logger.Fatal(e.Start(":1323"))
}
