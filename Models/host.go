package Models

type Host struct {
	Host             string     `json:"host"`
	Port             int16      `json:"port"`
	Protocol         string     `json:"protocol"`
	Is_Public        bool       `json:"isPublic"`
	Status           string     `json:"status"`
	StartTime        int64      `json:"startTime"`
	TestTime         int64      `json:"testTime"`
	Engine_Version   string     `json:"engineVersion"`
	Criteria_Version string     `json:"criteriaVersion"`
	Endpoints        []Endpoint `json:"endpoints"`
}
