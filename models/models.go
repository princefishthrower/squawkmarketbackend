package models

type Squawk struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Link      string `json:"link"`
	Symbols   string `json:"symbols"`
	Feed      string `json:"feed"`
	Squawk    string `json:"squawk"`
	Mp3Data   []byte `json:"mp3data"`
}
