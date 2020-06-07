package ModelsAPI

type Server struct {
	Servers_Changed    bool         `json:"servers_changed"`
	Min_SSL_Grade      string       `json:"ssl_grade"`
	Previous_SSL_Grade string       `json:"previous_ssl_grade"`
	Logo               string       `json:"logo"`
	Title              string       `json:"title"`
	Is_Down            bool         `json:"is_down"`
	Servers            []ServerInfo `json:"servers"`
}
