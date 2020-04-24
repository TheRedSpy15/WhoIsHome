package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Car struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Present bool   `json:"present"`
}
