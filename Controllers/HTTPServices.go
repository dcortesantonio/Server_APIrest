package Controllers

import (
	"GoProject/Models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getInfoSsl(domain string) (Models.Host, error) {
	//Request to http to get sll and servers information
	//Parameter: domain
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	if err != nil {
		log.Printf(" Request HTTP failed : [Error %s\n]", err)
		return Models.Host{}, err
	}
	defer response.Body.Close()

	// Reading the data of th response
	dataResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Models.Host{}, err
	}

	//Inverse Encoding that Marshall use
	//Get data information to strut Host
	actual_host := Models.Host{}
	err = json.Unmarshal(dataResponse, &actual_host)
	if err != nil {
		return Models.Host{}, err
	}
	return actual_host, nil
}

func whoIs(domain string) (Models.IPInfo, error) {
	//Request to http to get country information
	//Parameters: IP
	response, err := http.Get("http://ip-api.com/json/" + domain)
	if err != nil {
		log.Println("The HTTP request failed with error %s\n", err)
		return Models.IPInfo{}, err
	}
	defer response.Body.Close()

	// Reading the data of th response
	dataResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Models.IPInfo{}, err
	}

	//Inverse Encoding that Marshall use
	//Get data information to strut IPInfo
	ip_information := Models.IPInfo{}
	err = json.Unmarshal(dataResponse, &ip_information)
	if err != nil {
		return Models.IPInfo{}, err
	}

	return ip_information, nil
}
