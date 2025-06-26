package main

import (
	"log"
	"net/http"
	"project-app-portfolio-golang-rahmadhany/database"
	"project-app-portfolio-golang-rahmadhany/handler"
	"project-app-portfolio-golang-rahmadhany/repository"
	"project-app-portfolio-golang-rahmadhany/router"
	"project-app-portfolio-golang-rahmadhany/service"
)

func main() {
	conn, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}

	repo := repository.NewRepository(conn)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	r := router.NewRouter(h)

	log.Println("server starting on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
