package ModelsAPI

// ServerItem
//
// This represents server information.
//
// swagger:model
type ServerItem struct {
	// the IP or host of this server
	//
	// required: true
	// example: 34.193.69.252
	Address   string `json:"address"`
	// the SSL grade qualification of this server
	//
	// required: true
	// example: B
	SSL_Grade string `json:"ssl_grade"`
	// the country of this server
	//
	// required: true
	// example: US
	Country   string `json:"country"`
	// the owner of this server
	//
	// required: true
	// example: Amazon.com, Inc
	Owner     string `json:"owner"`
}

// Server
//
// This represents all servers information.
//
// swagger:model
type Server struct {
	// if the server changed between an hour or more
	//
	// required: true
	// example: true
	Servers_Changed    bool         `json:"servers_changed"`
	// the minimum value of SSL grade qualifications of servers
	//
	// required: true
	// example: B-
	Min_SSL_Grade      string       `json:"ssl_grade"`
	// the value of SSL grade qualifications between an hour or more
	//
	// required: true
	// example: A
	Previous_SSL_Grade string       `json:"previous_ssl_grade"`
	// the logo of the domain searched
	//
	// required: true
	Logo               string       `json:"logo"`
	// the title of the page of the domain searched
	//
	// required: true
	Title              string       `json:"title"`
	// if the server is down
	//
	// required: true
	// example: false
	Is_Down            bool         `json:"is_down"`
	// the servers of the domain searched
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
	Servers            []ServerItem `json:"servers"`
}
