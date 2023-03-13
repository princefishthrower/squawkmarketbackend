package fear_and_greed

type FearAndGreedResponse struct {
	Fgi struct {
		Now struct {
			Value     int    `json:"value"`
			ValueText string `json:"valueText"`
		} `json:"now"`
	} `json:"fgi"`
}
