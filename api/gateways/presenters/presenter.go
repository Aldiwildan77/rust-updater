package presenters

type Response struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type GameResponse struct {
	UpdaterResponse
}

type UpdaterResponse struct {
	Client  int `json:"client"`
	Devblog int `json:"devblog"`
	Oxide   int `json:"oxide"`
	Server  int `json:"server"`
	Staging int `json:"staging"`
}
