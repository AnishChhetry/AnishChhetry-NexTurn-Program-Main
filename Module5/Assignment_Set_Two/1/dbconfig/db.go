package dbconfig

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {
	var err error

	_, err = os.Stat("./blog.db")
	if os.IsNotExist(err) {
		log.Println("Database does not exist. Creating a new one...")

		db, err = sql.Open("sqlite3", "./blog.db")
		if err != nil {
			log.Fatal("Error opening database: ", err)
		}
		defer db.Close()

		createTableSQL := `
		CREATE TABLE IF NOT EXISTS Blog (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			content TEXT,
			author TEXT,
			timestamp DATETIME
		);
		`
		_, err = db.Exec(createTableSQL)
		if err != nil {
			log.Fatal("Error creating table: ", err)
		}

		insertDataSQL := `
		INSERT INTO Blog (title, content, author, timestamp)
		VALUES
			("First Blog Post", "This is the content of the first blog post.", "Author 1", datetime('now')),
			("Second Blog Post", "This is the content of the second blog post.", "Author 2", datetime('now'));
		`
		_, err = db.Exec(insertDataSQL)
		if err != nil {
			log.Fatal("Error inserting initial data: ", err)
		}

		log.Println("Database created and initial data inserted successfully.")

	} else if err != nil {
		log.Fatal("Error checking database file: ", err)
	} else {
		db, err = sql.Open("sqlite3", "./blog.db")
		if err != nil {
			log.Fatal("Error opening existing database: ", err)
		}
		log.Println("Successfully connected to the existing database.")
	}
}

func GetDB() *sql.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		log.Fatal("Error closing database: ", err)
	} else {
		log.Println("Database connection closed")
	}
}
