package ModelsAPI

type Item struct {
	Domain string `json:"domain"`
	Info   Server `json:"infoserver"`
}
