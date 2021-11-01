package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/brenordv/go-count-db-stuff/internal/models"
	"github.com/brenordv/go-count-db-stuff/internal/utils"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func EnsureIsConnected(db *sql.DB) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	utils.ErrorHandler(err, fmt.Sprintf("Error pinging connection: %v", err))
}

func CreateConnectionPool(cfg *models.RuntimeConfig) *sql.DB {
	var err error
	var db *sql.DB

	log.Print("Connecting...")

	db, err = sql.Open("mssql", cfg.ConnectionString)
	utils.ErrorHandler(err, fmt.Sprintf("Error creating connection pool: %v", err))

	EnsureIsConnected(db)
	log.Println("Connected!")

	return db
}

func ExecuteCountQuery(sql string, db *sql.DB) (int, error) {
	EnsureIsConnected(db)

	stmt, err := db.Prepare(sql)
	utils.ErrorHandler(err, "Could not prepare query.")

	defer stmt.Close()

	row := stmt.QueryRow()

	var count int
	err = row.Scan(&count)
	if err != nil {
		return -1, fmt.Errorf("failed to read result for query. %v\n", err)
	}
	return count, nil
}
