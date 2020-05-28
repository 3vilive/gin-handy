package main

import (
	"net/http"

	"github.com/3vilive/gin-handy/handy"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/query", func(c *gin.Context) {
		queryID, err := handy.GinCtxGetInt64(c, "id")
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"query_id": queryID,
		})
	})

	r.Run(":8080")
}
