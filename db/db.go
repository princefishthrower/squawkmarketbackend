package db

import (
	"database/sql"
	headlinesTypes "squawkmarketbackend/headlines/types"
	"squawkmarketbackend/utils"
)

func DoesHeadlineExist(headline string) (bool, error) {
	headlines, err := GetHeadlines()
	if err != nil {
		return false, err
	}

	return utils.Contains(headlines, headline), nil
}

func GetHeadlines() ([]string, error) {
	// Open a database connection
	db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query the headlines table
	rows, err := db.Query("SELECT headline FROM headlines")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through the results and print each headline
	headlines := []headlinesTypes.Headline{}
	for rows.Next() {
		var h headlinesTypes.Headline
		err := rows.Scan(&h.Headline)
		if err != nil {
			return nil, err
		}
		headlines = append(headlines, h)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// convert to strings
	var headlineStrings []string
	for _, headline := range headlines {
		headlineStrings = append(headlineStrings, headline.Headline)
	}
	return headlineStrings, nil
}

func AddHeadline(headline string, mp3data []byte) error {
	// Open a database connection
	db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Check if headline already exists
	exists, err := DoesHeadlineExist(headline)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	// Insert a new headline
	_, err = db.Exec("INSERT INTO headlines (headline, mp3data) VALUES (?, ?)", headline, mp3data)
	if err != nil {
		return err
	}

	return nil
}

func GetLatestHeadline() (headlinesTypes.Headline, error) {
	// Open a database connection
	db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	if err != nil {
		return headlinesTypes.Headline{}, err
	}
	defer db.Close()

	// Query the headlines table
	rows, err := db.Query("SELECT * FROM headlines ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		return headlinesTypes.Headline{}, err
	}
	defer rows.Close()

	// Loop through the results and print each headline
	headlines := []headlinesTypes.Headline{}
	for rows.Next() {
		var h headlinesTypes.Headline
		err := rows.Scan(&h.ID, &h.CreatedAt, &h.Headline, &h.Mp3Data)
		if err != nil {
			return headlinesTypes.Headline{}, err
		}
		headlines = append(headlines, h)
	}
	if err = rows.Err(); err != nil {
		return headlinesTypes.Headline{}, err
	}

	return headlines[0], nil
}
