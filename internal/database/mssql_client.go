package database

import (
	"context"
	"database/sql"
	"fmt"
	"go-count-db-stuff/internal/models"
	"go-count-db-stuff/internal/utils"
	"log"
)

func EnsureIsConnected(db *sql.DB)  {
	ctx := context.Background()
	err := db.PingContext(ctx)
	utils.ErrorHandler(err, fmt.Sprintf("Error pinging connection: %v", err.Error()))
}

func CreateConnectionPool(cfg *models.RuntimeConfig) *sql.DB {
	var err error
	var db *sql.DB

	log.Print("Connecting...")
	db, err = sql.Open("sqlserver", cfg.ConnectionString)
	utils.ErrorHandler(err, fmt.Sprintf("Error creating connection pool: %v", err.Error()))

	EnsureIsConnected(db)
	log.Println("Connected!")

	return db
}

func RunCount(countQuery models.CountQuery, db *sql.DB) bool {
	ctx := context.Background()
	EnsureIsConnected(db)

	rows, err := db.QueryContext(ctx, countQuery.Sql)
	utils.ErrorHandler(err, fmt.Sprintf("Failed to run query '%s'", countQuery.Name))
	defer func(rows *sql.Rows) {
		err := rows.Close()
		utils.ErrorHandler(err, fmt.Sprintf("Failed to close rows cursor for query '%s'\n", countQuery.Name))
	}(rows)

	successText := "SUCCESS"
	failText := "FAILED"

	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		utils.ErrorHandler(err, fmt.Sprintf("Failed to read result for query '%s'", countQuery.Name))

		var resultText string
		success :=  count == countQuery.ExpectedCount
		if success {
			resultText = successText
		} else {
			resultText = failText
		}

		log.Printf("%s: %s\n", countQuery.Name, resultText)
		return success
	}

	log.Printf("No rows returned for query '%s'. Check the SQL statements.\n", countQuery.Name)
	return false
}