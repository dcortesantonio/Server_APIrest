package Controllers

import (
	"GoProject/Models"
	"GoProject/ModelsAPI"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"reflect"
	"regexp"
)

func GetInfoDomain(ctx *fasthttp.RequestCtx) {
	//ConectionDB()
	fmt.Print(ctx, "<<<------INIT...->>>> ")

	//Get URL parameter [domain]
	str := ctx.UserValue("domain")
	domainValue := fmt.Sprintf("%v", str)

	//Evalue if the domain is valid
	regex := regexp.MustCompile(`[a-zA-Z]+.[a-zA-Z]+`)
	if regex.MatchString(domainValue) {
		fmt.Print("Siiii pasa", domainValue)
	}

	//Get information of SSL and servers.
	infoSSL := Models.Host{}
	infoSSL, err := getInfoSsl(domainValue)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	fmt.Print("\n firts infoSSL:", infoSSL.Host, infoSSL.Protocol, len(infoSSL.Endpoints))
	//Information of the domain searched.
	ServersConsulted := ModelsAPI.Server{}

	//Get HTML information [Title, Logo]
	title := ""
	logo := ""
	title, logo, err = getLogo(domainValue)
	if err != nil {
		log.Println("Error get Logo and title", err)
	}
	ServersConsulted.Title = title
	ServersConsulted.Logo = logo

	ServersConsulted.Is_Down = true
	if infoSSL.Status != "ERROR" {
		ServersConsulted.Is_Down = false
	}

	//Get information of servers of the domain searched and the min value of SSL Grade of its.
	ServerItems, Min_SSLGrade, err := getListServersItems(infoSSL.Endpoints)
	ServersConsulted.Servers = ServerItems
	ServersConsulted.Min_SSL_Grade = Min_SSLGrade
	fmt.Print("\n infoSSL:::--> - " + ServersConsulted.Min_SSL_Grade + " - " + ServersConsulted.Logo + " - " + ServersConsulted.Title)

	//Review changes in previous and actual information: Grade and changes in servers.
	gradeChange, serversChange := compareInformation(domainValue, ServersConsulted.Servers)
	ServersConsulted.Servers_Changed = serversChange
	ServersConsulted.Previous_SSL_Grade = Min_SSLGrade
	if gradeChange != "" {
		ServersConsulted.Previous_SSL_Grade = gradeChange
	}

	//Insert in DB the new information of the domain searched.
	success := insertDomain(domainValue, ServersConsulted)
	if success {
		fmt.Print("Insert Successfully!")
	}

	//Generate JSON of information of servers.
	res, err := json.Marshal(ServersConsulted)
	if err != nil {
		log.Println("Error marshaling. ", err)
		return
	}
	fmt.Print("-----END", res)

	//Return JSON of information of servers.
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetStatusCode(200)
	ctx.Response.SetBody(res)
	ctx.Write(res)
}

func getListServersItems(endPointsList []Models.Endpoint) ([]ModelsAPI.ServerItem, string, error) {
	Min_SSLGrade := "A+"
	var err error
	ServerItems := []ModelsAPI.ServerItem{}
	for _, endpoint := range endPointsList {
		IP := endpoint.IpAddress
		IPInfo, err := whoIs(IP)
		if err != nil {
			log.Println("Error whois <IP>", err)
			return ServerItems, Min_SSLGrade, err
		}
		serverItem := ModelsAPI.ServerItem{}
		serverItem.Address = endpoint.IpAddress
		serverItem.SSL_Grade = endpoint.SSL_Grade
		if findMinGrade(endpoint.SSL_Grade, Min_SSLGrade) {
			Min_SSLGrade = endpoint.SSL_Grade
		}
		serverItem.Country = IPInfo.CountryCode
		serverItem.Owner = IPInfo.Org
		ServerItems = append(ServerItems, serverItem)
	}
	err = nil
	return ServerItems, Min_SSLGrade, err
}

func findMinGrade(SSL_Grade string, min_SSL_Grade string) bool {
	MapSSL_Grades := make(map[string]int)
	MapSSL_Grades["F-"] = 0
	MapSSL_Grades["F"] = 1
	MapSSL_Grades["F+"] = 2
	MapSSL_Grades["E-"] = 3
	MapSSL_Grades["E"] = 4
	MapSSL_Grades["E+"] = 5
	MapSSL_Grades["D-"] = 6
	MapSSL_Grades["D"] = 7
	MapSSL_Grades["D+"] = 8
	MapSSL_Grades["C-"] = 9
	MapSSL_Grades["C"] = 10
	MapSSL_Grades["C+"] = 11
	MapSSL_Grades["B-"] = 12
	MapSSL_Grades["B"] = 13
	MapSSL_Grades["B+"] = 14
	MapSSL_Grades["A-"] = 15
	MapSSL_Grades["A"] = 16
	MapSSL_Grades["A+"] = 17

	if SSL, ok := MapSSL_Grades[SSL_Grade]; ok {
		if minGrade, ok := MapSSL_Grades[min_SSL_Grade]; ok {
			if minGrade > SSL {
				return true
			}
			return false
		}
	}
	return false
}
func compareInformation(domain string, actualServers []ModelsAPI.ServerItem) (string, bool) {
	serversChange := true
	previousServers := getPreviousServersItems(domain)
	if reflect.DeepEqual(previousServers, actualServers) == true {
		serversChange = false
	}

	previous_SSLGrade := getPrevious_SSL_Grade(domain)
	return previous_SSLGrade, serversChange
}
