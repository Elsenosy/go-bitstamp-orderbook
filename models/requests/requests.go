package requests

type Message struct {
	Event string  `json:"event"`
	Data  Channel `json:"data"`
}

type Channel struct {
	Channel string `json:"channel"`
}
