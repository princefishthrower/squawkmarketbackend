package db

import (
	"database/sql"
	"log"
	"squawkmarketbackend/models"
	"squawkmarketbackend/utils"
)

func DoesSquawkAlreadyExistAccordingToFeedCriterion(squawk string, symbols string, feedName string, insertThreshold float64) (bool, error) {
	existingSquawks, err := GetSquawks()
	if err != nil {
		return false, err
	}

	// log that we are in this function and the length of the existing squawks
	log.Println("DoesSquawkAlreadyExistAccordingToFeedCriterion, feedName: ", feedName, ", len(existingSquawks): ", len(existingSquawks))

	// get all squawk strings
	var existingSquawkStrings []string
	for _, existingSquawk := range existingSquawks {
		existingSquawkStrings = append(existingSquawkStrings, existingSquawk.Squawk)
	}

	switch feedName {
	// market-wide / crypto we only check the value of the squawk itself
	case "market-wide":
		fallthrough
	case "crypto":
		exists := utils.Contains(existingSquawkStrings, squawk)
		return exists, nil

	case "us-economic-prints":
		fallthrough
	case "eu-economic-prints":
		fallthrough
	case "cny-economic-prints":
		// for economic prints, the scraper itself is only forward looking so we can just use the squawk itself
		exists := utils.Contains(existingSquawkStrings, squawk)
		return exists, nil

	case "unusual-trading-volume":
		fallthrough
	case "most-volatile":
		fallthrough
	case "most-active":
		fallthrough
	case "new-highs":
		fallthrough
	case "new-lows":
		fallthrough
	case "overbought":
		fallthrough
	case "oversold":
		fallthrough
	case "top-gainers":
		fallthrough
	case "top-losers":
		// for all finviz related ones, we check if the squawk fuzzy matches (using insertThreshold) and if the symbols match
		// get all existing squawks today that match the feedName
		existingSquawks, err := GetSquawksByFeedName(feedName)
		if err != nil {
			return false, err
		}

		// check if the squawk matches (using utils.IsTooSimilar and insertThreshold) and if the symbols match
		for _, existingSquawk := range existingSquawks {
			// if the strings are too similar and the symbols match, return true
			if utils.AreStringsTooSimilar(squawk, existingSquawk.Squawk, insertThreshold) && existingSquawk.Symbols == symbols {
				return true, nil
			}
		}
		// if we get here, we didn't find a too similar squawk, so return false
		return false, nil
	default:
		return false, nil
	}
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

func GetSquawksByFeedName(feedName string) ([]models.Squawk, error) {
	// Open a database connection
	db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query the squawks table
	rows, err := db.Query("SELECT * FROM squawks WHERE feed = (?)", feedName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through the results and print each squawk
	squawks := []models.Squawk{}
	for rows.Next() {
		var h models.Squawk
		err := rows.Scan(&h.ID, &h.CreatedAt, &h.Link, &h.Symbols, &h.Feed, &h.Squawk, &h.Mp3Data)
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
