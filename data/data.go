package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDb() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

// CreateTable creates a new table case doesn't exist
func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS quick_notes (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"note" TEXT,		
		"category" TEXT
	  );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("The quick_notes table created")
}

// Add add a new note on the quick notes table
func Add(note string, category string) {
	insertQuery := `INSERT INTO quick_notes (note, category) VALUES (?,?)`

	statement, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(note, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Note added successfully")
}

// GetAll retrieves all notes
func GetAll() {
	row, err := db.Query("SELECT * FROM quick_notes order by id")
	if err != nil {
		log.Fatalln(err)
	}
	defer row.Close()
	
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)

	fmt.Fprintln(w, "Note\t\tCategory\t\tID")

	for row.Next() {
		var id int
		var note string
		var category string

		row.Scan(&id, &note, &category)
		
		fmt.Fprintf(w, "%s\t\t%s\t\t%d\n", note, category, id)
	}	
	w.Flush()
}

// DeleteById removes the note by ID
func DeleteById(id int) {
	deleteQuery := `DELETE FROM quick_notes WHERE id = (?)`

	statement, err := db.Prepare(deleteQuery)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(id)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Note deleted successfully")
}
