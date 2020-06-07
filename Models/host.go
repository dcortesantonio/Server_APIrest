package Models

type Endpoint struct {
	IpAddress         string `json:"ipAddress"`
	Server_Name       string `json:"serverName"`
	Status_Message    string `json:"statusMessage"`
	SSL_Grade         string `json:"grade"`
	GradeTrustIgnored string `json:"gradeTrustIgnored"`
	Has_Warnings      bool   `json:"hasWarnings"`
	Is_Exceptional    bool   `json:"isExceptional"`
	Progress          int64  `json:"progress"`
	Duration          int64  `json:"duration"`
	Delegation        int8   `json:"delegation"`
}

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
