package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaiquecaires/rinha-backend-go/models"
)

type FindHandler struct {
	DbPool *pgxpool.Pool
}

func (fh FindHandler) Find(c *gin.Context) {
	t, exists := c.GetQuery("t")

	if !exists || t == "" {
		fmt.Print("não existe")
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	rows, err := fh.DbPool.Query(
		context.Background(),
		"SELECT id, apelido, nome, nascimento, stack FROM pessoas WHERE busca LIKE $1",
		"%"+t+"%",
	)

	for err != nil {
		fmt.Print(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	defer rows.Close()

	var pessoas []models.Pessoa

	for rows.Next() {
		var pessoa models.Pessoa
		rows.Scan(
			&pessoa.Id,
			&pessoa.Apelido,
			&pessoa.Nome,
			&pessoa.Nascimento,
			&pessoa.Stack,
		)
		pessoas = append(pessoas, pessoa)
	}

	if pessoas == nil {
		pessoas = []models.Pessoa{}
	}

	c.JSON(http.StatusOK, pessoas)
}
