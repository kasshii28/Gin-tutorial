package models

import (
	//"crypto/sha1"
	"backend/config"
	"database/sql"
	"fmt"
	"log"

	//"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln("failed to open database",err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME
	)`, tableNameUser)

	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalln("failed to execute user_table",err)
	}

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		content_tag TEXT,
		done BOOLEAN,
		user_id INTEGER,
		created_at DATETIME,
		updated_at DATETIME,
		deadline DATETIME)`, tableNameTodo)

	_, err = Db.Exec(cmdT)
	if err != nil {
		log.Fatalln("failed to execute todo_table",err)
	}
}

// uuid作成関数
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// パスワードのハッシュ化関数
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}