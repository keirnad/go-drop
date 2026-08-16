package server

import (
	"godrop/internal/qrgen"
	"net/http"
)

func StartServer() {
	ip := GetOutboundIP()
	qrgen.QrGen("http://" + ip.String() + ":8080")
	http.ListenAndServe(ip.String()+":8080", http.FileServer(http.Dir("static")))
}
