package db

import (
	"database/sql"
	"squawkmarketbackend/models"
	"squawkmarketbackend/utils"
)

func DoesSquawkExist(squawk string) (bool, error) {
	squawks, err := GetSquawks()
	if err != nil {
		return false, err
	}

	return utils.Contains(squawks, squawk), nil
}

func GetSquawks() ([]string, error) {
	// Open a database connection
	db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query the squawks table
	rows, err := db.Query("SELECT squawk FROM squawks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through the results and print each squawk
	squawks := []models.Squawk{}
	for rows.Next() {
		var h models.Squawk
		err := rows.Scan(&h.Squawk)
		if err != nil {
			return nil, err
		}
		squawks = append(squawks, h)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// convert to strings
	var squawkStrings []string
	for _, squawk := range squawks {
		squawkStrings = append(squawkStrings, squawk.Squawk)
	}
	return squawkStrings, nil
}

func InsertSquawkIfNotExists(link, symbols, feed, squawk string, mp3data []byte) error {
	// Open a database connection
	db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Check if squawk already exists
	exists, err := DoesSquawkExist(squawk)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	// Insert a new squawk
	_, err = db.Exec("INSERT INTO squawks (link, symbols, feed, squawk, mp3data) VALUES (?, ?, ?, ?, ?)", link, symbols, feed, squawk, mp3data)
	if err != nil {
		return err
	}

	return nil
}

func GetLatestSquawk() (models.Squawk, error) {
	// Open a database connection
	db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	if err != nil {
		return models.Squawk{}, err
	}
	defer db.Close()

	// Query the squawks table
	rows, err := db.Query("SELECT * FROM squawks ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		return models.Squawk{}, err
	}
	defer rows.Close()

	// Loop through the results and print each squawk
	squawks := []models.Squawk{}
	for rows.Next() {
		var h models.Squawk
		err := rows.Scan(&h.ID, &h.CreatedAt, &h.Link, &h.Symbols, &h.Feed, &h.Squawk, &h.Mp3Data)
		if err != nil {
			return models.Squawk{}, err
		}
		squawks = append(squawks, h)
	}
	if err = rows.Err(); err != nil {
		return models.Squawk{}, err
	}

	return squawks[0], nil
}

func DeleteAllSquawks() error {
	// Open a database connection
	db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Delete all squawks
	_, err = db.Exec("DELETE FROM squawks")
	if err != nil {
		return err
	}

	return nil
}
