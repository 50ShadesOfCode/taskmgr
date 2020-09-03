package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kakty/taskmgr/middleware"

	"github.com/kakty/taskmgr/controllers"

	"github.com/go-redis/redis"

	"github.com/kakty/taskmgr/routes"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

var (
	db  *gorm.DB
	rdb *redis.Client
)

func init() {
	db = controllers.InitDb()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env values")
	}
}

func main() {
	PORT := os.Getenv("PORT")
	r := mux.NewRouter()
	routes.UseRoutes(r, db)
	r.Use(middleware.AuthMiddleware)
	srv := &http.Server{
		Handler: r,
		Addr:    ":" + PORT,
	}
	fmt.Println("Starting server on port: " + PORT)
	log.Fatal(srv.ListenAndServe())
}
