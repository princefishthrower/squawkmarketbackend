package db

import (
	"database/sql"
	"squawkmarketbackend/models"
	"squawkmarketbackend/utils"
)

func DoesSquawkExistAccordingToFeedCriterion(squawk string, symbols string, feedName string, insertThreshold float64) (bool, error) {
	existingSquawks, err := GetSquawks()
	if err != nil {
		return false, err
	}

	switch feedName {
	// market-wide / crypto we only check the value of the squawk itself
	case "market-wide":
	case "crypto":
		// get squawk strings from all Squawk objects
		var existingSquawkStrings []string
		for _, existingSquawk := range existingSquawks {
			existingSquawkStrings = append(existingSquawkStrings, existingSquawk.Squawk)
		}
		return utils.Contains(existingSquawkStrings, squawk), nil

	case "economic-prints":
		// for economic prints, we return true only if we can't find the symbol.
		// in this case the 'symbols' is the name of the report with date, i.e. "fomcminutes20230201"

		// get symbol strings from all Squawk objects
		var symbolStrings []string
		for _, squawk := range existingSquawks {
			symbolStrings = append(symbolStrings, squawk.Symbols)
		}
		return !utils.Contains(symbolStrings, symbols), nil

	case "unusual-trading-volume":
	case "most-volatile":
	case "most-active":
	case "new-highs":
	case "new-lows":
	case "overbought":
	case "oversold":
	case "top-gainers":
	case "top-losers":
		//for all finviz related ones, we check if the squawk fuzzy matches (using insertThreshold) and if the symbols match
		// get all existing squawks today that match the feedName
		existingSquawks, err := GetSquawksByFeedName(feedName)
		if err != nil {
			return false, err
		}

		// get squawk strings from all Squawk objects
		var squawkStrings []string
		for _, squawk := range existingSquawks {
			squawkStrings = append(squawkStrings, squawk.Squawk)
		}

		// check if the squawk fuzzy matches (using utils.IsDisimilarEnough and insertThreshold) and if the symbols match
		for i, existingSquawk := range existingSquawks {
			if utils.IsDisimilarEnough(squawk, existingSquawk.Squawk, insertThreshold) && existingSquawk.Symbols == symbols {
				return true, nil
			}
			// if we've reached the end of the list and haven't found a match, return false
			if i == len(existingSquawks)-1 {
				return false, nil
			}
		}
	default:
		return false, nil
	}

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
