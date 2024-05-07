package main

import (
	"database/sql"

	"GoDB/handlers"

	"GoDB/sqlc"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := sqlc.New(db)
	userHandler := handlers.NewUserHandler(query)

	r := gin.Default()
	r.POST("/users", userHandler.RegisterUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
	r.GET("/users/:id", userHandler.GetUser)
	r.PUT("/users/:id", userHandler.UpdateUser)

	r.Run(":8080")
}
