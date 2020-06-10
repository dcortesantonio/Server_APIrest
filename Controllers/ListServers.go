package Controllers

import (
	"GoProject/ModelsAPI"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

func GetListServers(ctx *fasthttp.RequestCtx) {
	fmt.Print(ctx,"1")
	listDomains := ModelsAPI.ServersConsulted{}
	fmt.Print(ctx,"2")
	//DB consul
	listDomains = getListServersDB()
	res, err := json.Marshal(listDomains)
	if err != nil {
		log.Println("[Error]: Marshaling struct ServerConsuled. ", err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetStatusCode(200)
	ctx.Response.SetBody(res)
	ctx.Write(res)
}
