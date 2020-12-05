package requel

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fatih/color"
)

// TimeReq request timer
func TimeReq() float64 {
	start := time.Now()
	exectime := time.Since(start)
	return exectime.Seconds()
}

//LogReq  http request logger
func LogReq(next http.Handler) http.Handler {
	var resptime float64
	defer func() {
		resptime = TimeReq()
	}()
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		greeen := color.New(color.FgHiGreen).SprintfFunc()
		path := req.RequestURI
		method := req.Method
		ip := req.RemoteAddr

		fmt.Printf("%s"+" "+"|"+"%s"+" "+"|"+"%ss"+"   "+"%s"+" "+"%s\n", green(method), yellow(path), greeen(strconv.FormatFloat(resptime, 'f', 6, 64)), green("=>"), yellow(ip))
		next.ServeHTTP(resp, req)
	})
}
