package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo_app/config"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// テーブルを作成するファイル
var Db *sql.DB

var err error

/* PostgreSQL対応に伴いコメントアウト
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)
*/

func init() {
	url := os.Getenv("DATABASE_URL") //HerokuのDATABASEのURLを取得
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}

	/* PostgreSQL対応に伴いコメントアウト
	// データベースとの接続
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)

	if err != nil {
		log.Fatalln(err)
	}

	// Create User Table
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		eMAIL STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

	// Create Todo Table
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	Db.Exec(cmdT)

	// Create Session Table
	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	Db.Exec(cmdS)
	*/
}

// Create UUID
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// Encrypt Password
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

//
