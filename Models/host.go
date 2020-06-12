package Models

//Endpoint
//
// This represents an endpoint of a domain searched, given in a HTTP Service.
//
// swagger:model
type Endpoint struct {
	// the IP address of the endpoint.
	//
	// required: true
	// example: 172.217.0.46
	IpAddress         string `json:"ipAddress"`
	// the name of the server searched.
	//
	// required: true
	// example: lga15s43-in-f46.1e100.net
	Server_Name       string `json:"serverName"`
	// the status of the endpoint searched.
	//
	// required: true
	// example: READY
	Status_Message    string `json:"statusMessage"`
	// the SSL grade qualification of this endpoint.
	//
	// required: true
	// example: B
	SSL_Grade         string `json:"grade"`
	// the SSL grade trust of this endpoint.
	//
	// required: true
	// example: B
	GradeTrustIgnored string `json:"gradeTrustIgnored"`
	// if this endpoint have warnings.
	//
	// required: true
	// example: false
	Has_Warnings      bool   `json:"hasWarnings"`
	// if the server is exceptional.
	//
	// required: true
	// example: false
	Is_Exceptional    bool   `json:"isExceptional"`
	// the progress of this endpoint.
	//
	// required: true
	// example: 100
	Progress          int64  `json:"progress"`
	// the duration of this endpoint.
	//
	// required: true
	// example: 100569
	Duration          int64  `json:"duration"`
	// the delegation of this endpoint.
	//
	// required: true
	// example: 1
	Delegation        int8   `json:"delegation"`
}

//Host
//
// This represents the SSL and servers information of a domain searched, given in a HTTP Service.
//
// swagger:model
type Host struct {
	// the domain searched
	//
	// required: true
	// example: google.com
	Host             string     `json:"host"`
	// the port of the domain searched
	//
	// required: true
	// example: 443
	Port             int16      `json:"port"`
	// the protocol that use the domain searched
	//
	// required: true
	// example: http
	Protocol         string     `json:"protocol"`
	// if the the domain searched is public
	//
	// required: true
	// example: false
	Is_Public        bool       `json:"isPublic"`
	// the port of the domain searched
	//
	// required: true
	// example: 443
	Status           string     `json:"status"`
	// the status of the domain searched
	//
	// required: true
	// example: 1591917234433
	StartTime        int64      `json:"startTime"`
	// the test time of the domain searched
	//
	// required: true
	// example: 1591917435246
	TestTime         int64      `json:"testTime"`
	// the engine version of the domain searched
	//
	// required: true
	// example: 2.1.5
	Engine_Version   string     `json:"engineVersion"`
	// the criteria version of the domain searched
	//
	// required: true
	// example: 2009q
	Criteria_Version string     `json:"criteriaVersion"`
	// the endpoints version of the domain searched
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
	Endpoints        []Endpoint `json:"endpoints"`
}
