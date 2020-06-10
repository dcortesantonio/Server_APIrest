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
	router.GET("/information/:domain", Controllers.GetInfoDomain )
	log.Fatal(fasthttp.ListenAndServe(":3000", router.Handler))




}

