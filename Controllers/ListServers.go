package Controllers

import (
	"GoProject/ModelsAPI"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)
// swagger:route GET /history
// Returns the list of domains searched.
//     Produces:
//     - application/json
//     Responses:
//       200: ServersConsulted
//       400: Client Error
func GetListServers(ctx *fasthttp.RequestCtx) {
	//The list of the domains searched
	listDomains := ModelsAPI.ServersConsulted{}

	//Get the list of domains searched.
	listDomains = getListServersDB()
	//Generate JSON of domains searched.
	res, err := json.Marshal(listDomains)
	if err != nil {
		log.Println("[Error]: Marshaling struct ServerConsuled ", err)
		ctx.SetStatusCode(400)
		return
	}
	//Return JSON of domains searched.
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetStatusCode(200)
	ctx.Write(res)
}
