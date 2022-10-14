package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

	"clean2/infrastructure/database"
	"clean2/infrastructure/router"
	"clean2/registry"
)

func main() {

	db := database.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":7000")
	if err := e.Start(":7000"); err != nil {
		log.Fatalln(err)
	}
}
