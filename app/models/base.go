package models

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	"app/config"
)

var (
	Db  *sql.DB
	err error
)

const (
	userTable = "users"
)

func init() {
	Db, err = sql.Open(config.Config.Db.Driver, config.Config.Db.Name)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			name STRING,
			email STRING,
			password STRING,
			created_at DATETIME
		)`,
		userTable,
	)
	Db.Exec(cmd)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha256.Sum256([]byte(plaintext)))
	return cryptext
}
