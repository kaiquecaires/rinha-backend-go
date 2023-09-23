package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CountHandler struct {
	DbPool *pgxpool.Pool
}

func (ch CountHandler) Count(c *gin.Context) {
	var count int

	err := ch.DbPool.QueryRow(
		context.Background(),
		"SELECT count(*) as count FROM pessoas",
	).Scan(&count)

	for err != nil {
		fmt.Print(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	c.JSON(http.StatusOK, count)
}
