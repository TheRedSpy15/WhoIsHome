package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	// Database connection
	db, err := sql.Open("mysql", "hunter:HDrumm_17@tcp(127.0.0.1:3306)/testdatabase") // TODO: read csv
	if err != nil {
		panic("MySQL: Failed to open database")
	} else {
		fmt.Println("MySQL: Connection Established")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("MySQL: Database ping failed: \n" + err.Error())
	}

	results, err := db.Query("SELECT * FROM cars")
	defer results.Close()

	// Displaying data
	var viewData string
	for results.Next() {
		var car Car
		err = results.Scan(&car.ID, &car.Name, &car.Present)
		if err != nil {
			fmt.Println("Failed to read data")
		}
		viewData += `<p style="width:20px;height:20px;">` + car.Name + `</p>` // TODO: read from template file
	}

	return c.HTML(http.StatusOK, viewData)
}
