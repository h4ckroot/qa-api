package main

import (
	"github.com/rs/zerolog/log"
	"shamilabs.com/qa-api/config"
	"shamilabs.com/qa-api/models"
	"shamilabs.com/qa-api/routes"
)

func main() {

	// init config file
	c, _ := config.LoadConfig("config/config.yaml")
	log.Info().Msgf("Loaded configs: %s", c)

	// init DB
	dbConfig := c.DB

	log.Info().Msgf("db connection details: ", config.DbURL(&dbConfig))
	db := config.InitDBClient(config.DbURL(&dbConfig))

	qaDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer qaDB.Close()

	// Run DB Migraitons
	db.AutoMigrate(&models.Question{})
	db.AutoMigrate(&models.Answer{})
	db.AutoMigrate(&models.Tag{})

	r := routes.SetupRouter()
	r.Run()
}
