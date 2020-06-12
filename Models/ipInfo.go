package Models

//IPinfo
//
// This represents the location information of an IP, given in a HTTP Service.
//
// swagger:model
type IPInfo struct {
	// the IP address of the endpoint.
	//
	// required: true
	// example: 172.217.0.46
	Query       string  `json:"query"`
	// the status of the endpoint.
	//
	// required: true
	// example: success
	Status      string  `json:"status"`
	// the country where is the endpoint server.
	//
	// required: true
	// example: United States
	Country     string  `json:"country"`
	// the code of the country where is the endpoint server.
	//
	// required: true
	// example: US
	CountryCode string  `json:"countryCode"`
	// the region where is the endpoint server.
	//
	// required: true
	// example: NY
	Region      string  `json:"region"`
	// the name of the region where is the endpoint server.
	//
	// required: true
	// example: New York
	RegionName  string  `json:"regionName"`
	// the city where is the endpoint server.
	//
	// required: true
	// example: New York
	City        string  `json:"city"`
	// the zip of the endpoint server.
	//
	// required: true
	// example: 10123
	Zip         string  `json:"zip"`
	// the lat of the endpoint server.
	//
	// required: true
	// example: 40.7128
	Lat         float64 `json:"lat"`
	// the lon of the endpoint server.
	//
	// required: true
	// example: 74.006
	Lon         float64 `json:"lon"`
	// the TimeZone where is the endpoint server.
	//
	// required: true
	// example: America/New_York
	Timezone    string  `json:"timezone"`
	// the isp of the endpoint server.
	//
	// required: true
	// example: Google LLC
	Isp         string  `json:"isp"`
	// the org of the endpoint server.
	//
	// required: true
	// example: Google LLC
	Org         string  `json:"org"`
	// the as of the endpoint server.
	//
	// required: true
	// example: AS15169 Google LLC
	As          string  `json:"as"`
}
