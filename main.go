package main

import (
	"GoProject/Controllers"
	"github.com/buaazp/fasthttprouter"
	_ "github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	//router.GET("/{domain}", Controllers.ListServers )
	//router.GET("/history", Controllers.getListServers )

}
