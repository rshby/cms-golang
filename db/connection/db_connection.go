package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// create function to connect with database
func ConnectDatabase() *sql.DB {
	// koneksi dari docker container -> host.docker.internal

	db, err := sql.Open("mysql", fmt.Sprintf("root:@tcp(host.docker.internal:3306)/%v?parseTime=true", os.Getenv("DATABASE_NAME")))

	// root:@tcp(host.docker.internal:3306)/BRI_CORESYSTEM?parseTime=True?loc=Local

	if err != nil {
		log.Fatalf("error connection %v\n", err.Error())
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	db.SetConnMaxIdleTime(30 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
