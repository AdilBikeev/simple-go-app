package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AdilBikeev/simple-go-mysql/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
)

// @title          Simple-GO-MySql
// @version        1.0
// @description    This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	routes.RegisterSwaggerRoutes(r)

	http.Handle("/", r)
	fmt.Println("Server started with port:8080!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
