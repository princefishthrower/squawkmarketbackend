package db

import (
	"database/sql"
	"squawkmarketbackend/models"
	"squawkmarketbackend/utils"
)

func DoesSquawkExistAccordingToFeedCriterion(squawk string, symbols string, feedName string, insertThreshold float64) (bool, error) {
	squawks, err := GetSquawks()
	if err != nil {
		return false, err
	}

	// market-wide we only check the value of the squawk itself
	if feedName == "market-wide" || feedName == "crypto" {
		// get squawk strings from all Squawk objects
		var squawkStrings []string
		for _, squawk := range squawks {
			squawkStrings = append(squawkStrings, squawk.Squawk)
		}
		return utils.Contains(squawkStrings, squawk), nil
	}
	if feedName == "economic-prints" {
		// for economic prints, we return true only if we can't find the symbol.
		// in this case the 'symbols' is the name of the report with date, i.e. "fomcminutes20230201"

		// get symbol strings from all Squawk objects
		var symbolStrings []string
		for _, squawk := range squawks {
			symbolStrings = append(symbolStrings, squawk.Symbols)
		}
		return !utils.Contains(symbolStrings, symbols), nil
	}

	// TODO: for all finviz related ones, we check if the squawk fuzzy matches (using insertThreshhold) and if the symbols match

	return false, nil
}

func GetSquawks() ([]models.Squawk, error) {
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

	return squawks, nil
}

func InsertSquawk(link, symbols, feed, squawk string, mp3data []byte) error {
	// Open a database connection
	db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert a new squawk
	_, err = db.Exec("INSERT INTO squawks (link, symbols, feed, squawk, mp3data) VALUES (?, ?, ?, ?, ?)", link, symbols, feed, squawk, mp3data)
	if err != nil {
		return err
	}

	return nil
}

func GetLatestSquawkByFeed(feedName string) (models.Squawk, error) {
	// Open a database connection
	db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	if err != nil {
		return models.Squawk{}, err
	}
	defer db.Close()

	// Query the squawks table
	rows, err := db.Query("SELECT * FROM squawks WHERE feed = (?) ORDER BY created_at DESC LIMIT 1", feedName)
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
