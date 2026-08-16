package server

import (
	"godrop/internal/qrgen"
	"net/http"
	"path"
)

func StartServerDownload(filepath string) {
	ip := GetOutboundIP()
	filename := path.Base(filepath)

	qrgen.QrGen("http://" + ip.String() + ":8080")

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		http.ServeFile(w, r, filepath)
	})

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.ListenAndServe(ip.String()+":8080", nil)
}
