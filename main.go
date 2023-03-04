package main

import (
	"enerbit/prueba/db"
	_ "enerbit/prueba/docs"
	"enerbit/prueba/handlers"
	"enerbit/prueba/models"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	baseRouter := router.Group("/api/v1")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	baseRouter.GET("/listemeter", handlers.ListEMeter)
	baseRouter.GET("/lastemeter/:brand/:serial", handlers.LastEMeter)
	baseRouter.GET("/listdisableemeter", handlers.ListDisableEMeter)
	baseRouter.POST("/newemeter", handlers.NewEMeter)
	baseRouter.POST("/deleteemeter", handlers.DeleteEMeter)
	baseRouter.POST("/updateemeterstatus", handlers.UpdateEMeterStatus)

	router.Run(":8080")
}

// @title		Prueba técnica Enerbit
// @version		1.0
// @description	Servicio de gestion contadores
// @host		localhost:8080
// @BasePath	/api/v1
func init() {
	// migracion

	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	dbConn, err := db.PostgresConnection()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	err = dbConn.AutoMigrate(&models.Datameds{})
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
