package ModelsAPI

type ServerInfo struct {
	Address   string `json:"address"`
	SSL_Grade string `json:"ssl_grade"`
	Country   string `json:"country"`
	Owner     string `json:"owner"`
}
