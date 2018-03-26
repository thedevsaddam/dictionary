package data

import (
	"archive/zip"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	// used for dictionary sqlite3 database
	_ "github.com/mattn/go-sqlite3"
)

const (
	dictionaryDB    = "./dictionary.db"
	dictionaryDBZip = "./dictionary.db.zip"
	dbZipURL        = "https://raw.githubusercontent.com/thedevsaddam/dictionary/master/dictionary.db.zip"
)

// Entry describes an word entry
type Entry struct {
	Word       string
	Wordtype   string
	Definition string
}

// Data represents a dinctionary to find and search words
type Data interface {
	Load(limit int) []Entry
	Find(word string) []Entry
	Fuzzy(word string) []Entry
	Close() error
}

// DictionarySQL implements the Dictionarier interface to find words from sql
type DictionarySQL struct {
	DB *sql.DB
}

// New ...
func New() Data {
	db, err := sql.Open("sqlite3", dictionaryDB)
	if err != nil {
		log.Fatal(err)
	}
	return &DictionarySQL{
		DB: db,
	}
}

// Load laod the dictionary with default entry
func (d *DictionarySQL) Load(limit int) []Entry {
	entries := []Entry{}
	rows, err := d.DB.Query(`SELECT * FROM entries LIMIT $1`, limit)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var word, wordtype, definition string
		err = rows.Scan(&word, &wordtype, &definition)
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, Entry{
			Word:       word,
			Wordtype:   wordtype,
			Definition: definition,
		})
	}

	return entries
}

// Find find a exact match of word
func (d *DictionarySQL) Find(word string) []Entry {
	entries := []Entry{}
	rows, err := d.DB.Query(`SELECT * FROM entries WHERE word = $1 ORDER BY word ASC`, word)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var word, wordtype, definition string
		err = rows.Scan(&word, &wordtype, &definition)
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, Entry{
			Word:       word,
			Wordtype:   wordtype,
			Definition: definition,
		})
	}

	return entries
}

// Fuzzy find fuzzy match of words
func (d *DictionarySQL) Fuzzy(word string) []Entry {
	entries := []Entry{}
	rows, err := d.DB.Query(`SELECT * FROM entries WHERE word LIKE '%' || $1 || '%' ORDER BY word ASC`, word)
	checkErr(err)

	defer rows.Close()
	for rows.Next() {
		var word, wordtype, definition string
		err = rows.Scan(&word, &wordtype, &definition)
		checkErr(err)

		entries = append(entries, Entry{
			Word:       word,
			Wordtype:   wordtype,
			Definition: definition,
		})
	}

	return entries
}

// Close close the db connection
func (d *DictionarySQL) Close() error {
	return d.DB.Close()
}

// Setup ...
func Setup() {
	fmt.Println("Setting up dictionary...")
	if _, err := os.Stat(dictionaryDBZip); os.IsNotExist(err) {
		fmt.Println("Downloading database...")
		// download the zip database
		out, err := os.Create(dictionaryDBZip)
		checkErr(err)
		defer out.Close()

		resp, err := http.Get(dbZipURL)
		checkErr(err)
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)
		checkErr(err)
	}

	fmt.Println("Unzipping database...")
	r, err := zip.OpenReader(dictionaryDBZip)
	checkErr(err)

	rc := r.File[0]

	df, err := os.OpenFile(dictionaryDB, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, rc.Mode())
	checkErr(err)

	sf, err := rc.Open()
	checkErr(err)

	_, err = io.Copy(df, sf)
	checkErr(err)
	fmt.Println("Dictionary setup completed successfully!")
}

// checkErr ....
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
