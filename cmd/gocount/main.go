package main

import (
	"go-count-db-stuff/internal/constants"
	"go-count-db-stuff/internal/database"
	"go-count-db-stuff/internal/utils"
	"log"
	"time"
)

func main() {
	start := time.Now()
	log.Printf("** GO Count Stuff! (%s) **\n", constants.Version)
	cfgFile := utils.ParseCommandLine()
	cfg := utils.ReadConfigFile(cfgFile)

	log.Printf("Found '%d' queries to run. \n", len(cfg.Queries))
	db := database.CreateConnectionPool(cfg)

	for _, countQuery := range cfg.Queries {
		database.RunCount(countQuery, db)
	}

	log.Printf("All done! (elapsed time: %s)\n", time.Since(start))
}