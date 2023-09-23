package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaiquecaires/rinha-backend-go/models"
	"github.com/kaiquecaires/rinha-backend-go/utils"
)

type CreateHandler struct {
	DbPool *pgxpool.Pool
}

func (ch CreateHandler) Create(c *gin.Context) {
	var data models.Pessoa

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to parse JSON data"})
		return
	}

	if !utils.IsDateValid(data.Nascimento) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	var InsertedId string

	err := ch.DbPool.QueryRow(
		context.Background(),
		"INSERT INTO pessoas (apelido, nome, nascimento, stack) VALUES ($1, $2, $3, $4) RETURNING id",
		data.Apelido,
		data.Nome,
		data.Nascimento,
		data.Stack,
	).Scan(&InsertedId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.Header("Location", "/pessoas/"+InsertedId)
	c.JSON(http.StatusCreated, gin.H{})
}
