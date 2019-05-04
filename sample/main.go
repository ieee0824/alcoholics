package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ieee0824/alcoholics"
)

func main() {
	r := gin.Default()
	r.Use(alcoholics.New([]alcoholics.Option{
		{
			Probability: 5,
			StatusCode:  200,
		},
		{
			Probability:  1,
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "なんか無いよ",
		},
		{
			Probability:  1,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "なんかおかしい",
		},
	}).Drunk())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "うごいた")
	})

	r.Run()
}
