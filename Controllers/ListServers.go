package Controllers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func getListServers(ctx *fasthttp.RequestCtx) {
	enc := json.NewEncoder(ctx)
	//items :=
	//Call to DATA BASE.
	//Query: FROM DOMAIN
	//		ORDER BY MAX(consulted_time)
	//		LIMIT 100;
	//record = FindServersRecords()
	//enc.Encode(&record)
	ctx.SetStatusCode(fasthttp.StatusOK)
	//record := ModelsAPI.ServersConsulted{}

}
