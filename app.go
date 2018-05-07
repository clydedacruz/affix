package main

import "github.com/gin-gonic/gin"
import p "affix/lib"

func main() {
	r := gin.Default()

	// initialize persistence
	kvs := p.BadgerKV{}
	err := kvs.Init()
	if err != nil {
		panic(err.Error())
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
