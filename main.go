package main

import (
	"github.com/anousoneFS/go-fiber-postgres-workshop/internal/district"
	"github.com/anousoneFS/go-fiber-postgres-workshop/internal/province"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProvinceResponse struct {
	Name string `json:"name"`
}

var DB *gorm.DB

func main() {
	dsn := "postgres://ajlrhvob:hR5QQMvvokaydTHQ5ygiSW-OkBoeFNuX@tiny.db.elephantsql.com/ajlrhvob"
	dial := postgres.Open(dsn)
	var err error
	DB, err = gorm.Open(dial)
	if err != nil {
		panic(err)
	}
	if err = DB.AutoMigrate(province.Province{}, district.District{}); err != nil {
		panic(err)
	}
	app := fiber.New()

	// Province
	provinceRepo := province.NewRepository(DB)
	provinceUsecase := province.NewUsecase(provinceRepo)
	province.NewHandler(app, provinceUsecase)

	// District
	districtRepo := district.NewRepository(DB)
	districUsecase := district.NewUsecase(districtRepo)
	district.NewHandler(app, districUsecase)

	app.Listen(":3000")
}
