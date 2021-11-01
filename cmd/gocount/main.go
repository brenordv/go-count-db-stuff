package main

import (
	"database/sql"
	"fmt"
	"github.com/brenordv/go-count-db-stuff/internal/constants"
	"github.com/brenordv/go-count-db-stuff/internal/database"
	"github.com/brenordv/go-count-db-stuff/internal/models"
	"github.com/brenordv/go-count-db-stuff/internal/utils"
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
		DoQueryCount(countQuery, db)
	}

	log.Printf("All done! (elapsed time: %s)\n", time.Since(start))
}

func DoQueryCount(query models.CountQuery, db *sql.DB) {
	actualCount, err := database.ExecuteCountQuery(query.Sql, db)
	utils.ErrorHandler(err, fmt.Sprintf("Query '%s' failed to run.", query.Name))

	var resultText string
	success := actualCount >= 0 && query.ExpectedCount >= 0 && actualCount == query.ExpectedCount

	if success {
		resultText = fmt.Sprintf("%s: %s%s%s", query.Name, constants.ColorGreen, "SUCCESS", constants.ColorReset)
	} else {
		resultText = fmt.Sprintf("%s: %s%s%s", query.Name, constants.ColorRed, "FAILED", constants.ColorReset)
	}

	log.Println(resultText)

}
