package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "mrt-schedules/modules/station"
    "time"
)

func main() {
    initiateRouter()
}

func initiateRouter() {
    router := gin.Default()

    // Middleware CORS dengan konfigurasi lebih fleksibel
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // Jika ingin spesifik, ubah ke "http://127.0.0.1:5500"
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    api := router.Group("/v1/api")
    station.Initiate(api)

    router.Run(":8080")
}

