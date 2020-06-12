//Package main of the API
//
// Documentation for Product API
// Schemes: HTTP
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// -application/json
//
// Produce:
// -application/json
// swager: meta

package main

import (
	"GoProject/Controllers"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)
func Default(ctx *fasthttp.RequestCtx)  {
	fmt.Fprintf(ctx,"Welcome honey...")
}
func main() {

	router := fasthttprouter.New()
	fmt.Print("INIT")
	router.GET("/", Default)
	router.GET("/history", Controllers.GetListServers)
	router.GET("/search/:domain", Controllers.GetInfoDomain )
	log.Fatal(fasthttp.ListenAndServe(":3000", router.Handler))


}

