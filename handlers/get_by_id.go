package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaiquecaires/rinha-backend-go/models"
)

type GetByIdHanlder struct {
	DbPool *pgxpool.Pool
}

func (g GetByIdHanlder) GetById(c *gin.Context) {
	id := c.Param("id")

	var pessoa models.Pessoa

	err := g.DbPool.QueryRow(
		context.Background(),
		"SELECT p.id, p.nome, p.apelido, p.nascimento, p.stack FROM pessoas p WHERE p.id = $1",
		id,
	).Scan(&pessoa.Id, &pessoa.Nome, &pessoa.Apelido, &pessoa.Nascimento, &pessoa.Stack)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, pessoa)
}
