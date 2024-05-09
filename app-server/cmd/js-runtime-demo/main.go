package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justyanna/toy-box/src/runtime"
)

func main() {
	router := gin.Default()

	runtime.RegisterRoutes(router)

	router.Run()
}
