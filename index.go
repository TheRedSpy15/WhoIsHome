package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Car struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Present bool   `json:"present"`
}

func index(c echo.Context) error {
	db, err := sql.Open("mysql", "hunter:HDrumm_17@tcp(127.0.0.1:3306)/testdatabase")
	if err != nil {
		panic(err.Error())
	} else {
		c.Logger().Print("MySQL: Connection Established")
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // need proper error handling instead of panic
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM cars")
	defer results.Close()

	var viewData string
	for results.Next() {
		var car Car
		err = results.Scan(&car.ID, &car.Name, &car.Present)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		viewData += `<p style="width:20px;height:20px;">` + car.Name + `</p>`
	}

	return c.HTML(http.StatusOK, viewData)
}
