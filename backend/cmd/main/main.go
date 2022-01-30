package main

import (
	"fmt"

	"github.com/mitul/easy-lease/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	config.Connect()
	db = config.GetDB()
	fmt.Println(db)
}
