package main

import (
	"fmt"
	"go_bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is running on port 9010...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
