package ModelsAPI

//Item: A consult of a domain.
type Item struct {
	Domain string `json:"domain"`
	Info   Server `json:"infoserver"`
}

//List of wanted domains.
type ServersConsulted struct {
	Items []Item `json:"items"`
}
