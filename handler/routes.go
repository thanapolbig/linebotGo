package handler

import (
	"net/http"

	"api-wecode-supplychain/service/boibot"
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	//Validation  gin.HandlerFunc
}

type Routes struct {
	transaction []route
}

func (r Routes) InitTransactionRoute() http.Handler {

	boibot := boibot.NewEndpoint()

	r.transaction = []route{

		{
			Name:        "CallBack : POST ",
			Description: "CallBack",
			Method:      http.MethodPost,
			Pattern:     "/callback",
			Endpoint:    boibot.CallbackHandler,
		},
	}

	ro := gin.New()

	store := ro.Group("/app")
	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}
