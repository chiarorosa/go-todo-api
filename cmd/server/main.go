package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Cria uma instância padrão do servidor Gin
    r := gin.Default()

    // Define uma rota GET simples
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Servidor Go com Gin está rodando!",
        })
    })

    // Inicia o servidor na porta 8080
    r.Run(":8080")
}