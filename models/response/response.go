package response

type BitstampResponse struct {
	Data    BitstampData `json:"data"`
	Channel string       `json:"channel"`
	Event   string       `json:"event"`
}

type BitstampData struct {
	Timestamp      string          `json:"timestamp"`
	Microtimestamp string          `json:"microtimestamp"`
	Bids           [][]interface{} `json:"bids"`
	Asks           [][]interface{} `json:"asks"`
}
