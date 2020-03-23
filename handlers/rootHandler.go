package handlers

import "net/http"

//RootHandler handles the request
func RootHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Asset not found\n"))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Running API v1\n"))
}
