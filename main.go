package main

import (
	"net/http"
	"tunaiku_tes/routers"
	"time"
	"fmt"
)

func main()  {
	mainServer := routers.RouterMain()
	serverMain := &http.Server{
		Addr:			":9008",
		Handler: 		mainServer,
		ReadTimeout:	10 * time.Second,
		WriteTimeout:	10 * time.Second,
		MaxHeaderBytes:	1 << 20,
	}
	fmt.Println("Running on port : http://localhost:9008/")

	serverMain.ListenAndServe() //running on port 9008
}