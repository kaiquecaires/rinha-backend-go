package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kaiquecaires/rinha-backend-go/db"
	"github.com/kaiquecaires/rinha-backend-go/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	pool, err := db.CreatePostgresPool()
	if err != nil {
		log.Fatalf("Error to create pool: %v", pool)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API IS ON FIRE",
		})
	})

	r.POST("/pessoas", func(ctx *gin.Context) {
		handler := handlers.CreateHandler{
			DbPool: pool,
		}

		handler.Create(ctx)
	})

	r.GET("/pessoas/:id", func(ctx *gin.Context) {
		handler := handlers.GetByIdHanlder{
			DbPool: pool,
		}

		handler.GetById(ctx)
	})

	r.GET("/pessoas", func(ctx *gin.Context) {
		handler := handlers.FindHandler{
			DbPool: pool,
		}

		handler.Find(ctx)
	})

	r.GET("/contagem-pessoas", func(ctx *gin.Context) {
		handler := handlers.CountHandler{
			DbPool: pool,
		}

		handler.Count(ctx)
	})

	r.Run(os.Getenv("APP_PORT"))
}
