package server

import (
	"github.com/gin-gonic/gin"
	"./httpHandler"
	"log"
	"time"
)

type GinRouter struct {
	// required parameters for data source VariFlight API validation
	Appid            string
	RegistrationCode string

	// time control parameters for avoiding unnecessary extra API calling
	MinDurAgainstExtraRequest time.Duration
	MaxDurAgainstExtraRequest time.Duration

	// arguments for connecting a database
	DriverName string
	DataSourceName string

	variFlightHttpHandler *httpHandler.VariFlightHttpHandler
}

func New(appid, registrationCode string, minDurAgainstExtraRequest, maxDurAgainstExtraRequest time.Duration, driverName, dataSourceName string) *GinRouter {
	r := GinRouter{
		Appid: appid,
		RegistrationCode: registrationCode,

		MinDurAgainstExtraRequest: minDurAgainstExtraRequest,
		MaxDurAgainstExtraRequest: maxDurAgainstExtraRequest,

		DriverName: driverName,
		DataSourceName: dataSourceName,
	}

	variFlightHttpHandler := httpHandler.New(appid, registrationCode, minDurAgainstExtraRequest, maxDurAgainstExtraRequest, driverName, dataSourceName)
	r.variFlightHttpHandler = variFlightHttpHandler
	
	return &r
}

func (r *GinRouter) Run() {
	router := gin.Default()
	gin.WrapH(r.variFlightHttpHandler)

	router.GET("/variflight", r.variFlightHttpHandler)
	if err := router.Run("localhost:8000"); err != nil {
		log.Fatal(err)
	}
}

