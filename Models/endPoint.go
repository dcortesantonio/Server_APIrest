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
