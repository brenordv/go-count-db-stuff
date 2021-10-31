package database

import (
	"context"
	mssql "database/sql"
	"fmt"
	"go-count-db-stuff/internal/models"
	"go-count-db-stuff/internal/utils"
	"log"
)

func EnsureIsConnected(db *mssql.DB)  {
	ctx := context.Background()
	err := db.PingContext(ctx)
	utils.ErrorHandler(err, fmt.Sprintf("Error pinging connection: %v", err.Error()))
}

func CreateConnectionPool(cfg *models.RuntimeConfig) *mssql.DB {
	var err error
	var db *mssql.DB

	log.Print("Connecting...")
	db, err = mssql.Open("sqlserver", cfg.ConnectionString)
	utils.ErrorHandler(err, fmt.Sprintf("Error creating connection pool: %v", err.Error()))

	EnsureIsConnected(db)
	log.Println("Connected!")

	return db
}

func ExecuteCountQuery(sql string, db *mssql.DB) (int, error) {
	ctx := context.Background()
	EnsureIsConnected(db)

	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		return -1, fmt.Errorf("failed to run query. %v\n", err)
	}

	defer func(rows *mssql.Rows) {
		err := rows.Close()
		utils.ErrorHandler(err, "Failed to close rows cursor for query.")
	}(rows)

	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		if err != nil {
			return -1, fmt.Errorf("failed to read result for query. %v\n", err)
		}
		return count, nil
	}

	return -1, fmt.Errorf("No rows returned for query. Check the SQL statements.\n")
}
