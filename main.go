package main

import (
	"context"
	docs "kiennguyen94/go_calendar/docs"
	routes "kiennguyen94/go_calendar/routes"

	"database/sql"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

const db_path string = "/Volumes/SSD-PGU3/Documents/programming_proj/go_bun/calendar.db"

var ctx context.Context
var db *bun.DB

func main() {
	ctx = context.Background()
	sqldb, err := sql.Open(sqliteshim.ShimName, db_path)
	if err != nil {
		panic(err)
	}
	db = bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/doctor", routes.GetDoctor)
		v1.POST("/doctor", routes.PostDoctor)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
