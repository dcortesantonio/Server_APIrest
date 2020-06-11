package Controllers

import (
	"GoProject/Models"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
)

func getInfoSsl(domain string) (Models.Host, error) {
	//Request to http to get sll and servers information
	//Parameter: domain
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	if err != nil {
		log.Fatal("[Error] : Request HTTP SSL failed ", err)
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
		log.Fatal("[Error] : Request HTTP who is failed ", err)
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
func getLogo(name string) (string, string , error) {

	response, err := http.Get("http://" +name)
	if err != nil {
		log.Fatal("[Error] : Request HTTP failed ",err)
		return "","",err
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println("[Error] Response body, NewDocument ", err)
		return "","",err
	}

	title := document.Find("title").Text()
	logo := ""
	document.Find("link").Each(func(index int, element *goquery.Selection) {
		rel, exists := element.Attr("rel")
		if exists && (rel == "shortcut icon" || rel == "icon") {
			href, exists1 := element.Attr("href")
			if exists1 {
				logo = href
			}
		}
	})
	return title, logo, nil
}