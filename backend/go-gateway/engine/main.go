package main

import (
	"go-gateway/engine/db/psqldb"
	"go-gateway/engine/services"

	"github.com/gin-gonic/gin"
)

func main() {
	psqldb.InitDB()
	defer psqldb.CloseDB()

	r := gin.Default()

	services.AddServices(r)

	r.Run(":8080")
}
