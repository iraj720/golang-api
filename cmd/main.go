package main

import (
	api "snapp/api/productController"
	"snapp/confs"
	_ "snapp/docs"
	"snapp/fixture"
	"snapp/internal/repository"
	"snapp/internal/service"
	"snapp/migrations"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := confs.ConnectToDb()
	migrations.MigrateUp(db)
	fixture.ProductTable(db)
	re := repository.NewProductRepository(db)
	ps := service.NewProductService(re)
	pc := api.NewProductController(ps)
	r := gin.Default()
	pc.RegisterRoutes(&r.RouterGroup)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	confs.RunGin(r)
}
