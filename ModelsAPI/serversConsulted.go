package ModelsAPI

//Item
//
// This represents a consult of a domain.
//
// swagger:model
type Item struct {
	// the domain searched
	//
	// required: true
	// example: google.com
	Domain string `json:"domain"`
	// the information of the domain searched
	//
	// required: true
	// Extensions:
	// ---
	// x-property-value: value
	// x-property-array:
	//   - value1
	//   - value2
	// x-property-array-obj:
	//   - name: obj
	//     value: field
	// ---
	Info   Server `json:"infoserver"`
}

//ServersConsulted
//
// This represents the history of domains searched.
//
// swagger:model
type ServersConsulted struct {
	// the list of domains searched
	//
	// required: true
	// Extensions:
	// ---
	// x-property-value: value
	// x-property-array:
	//   - value1
	//   - value2
	// x-property-array-obj:
	//   - name: obj
	//     value: field
	// ---
	Items []Item `json:"items"`
}
