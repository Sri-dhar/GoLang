package main

import(
  "log"
  "net/http"
  "github.com/gorilla/mux"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "/home/solomons/GoLang/go-sql/pkg/routes"
)

func main(){
  router := mux.NewRouter()
  routes.RegisterBookStoreRoutes(router)
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe("localhost:9010", r))  
}
