package main
import (
    "go-goof/handlers"
    "github.com/square/go-jose/cipher"
    "github.com/gin-gonic/gin"
)
func main() {
        r := gin.Default()
        r.GET("/ping", handlers.Ping)
        r.Run() // listen and serve on 0.0.0.0:8080
}
