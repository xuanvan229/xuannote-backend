package helper

import (
	"encoding/json"
	"net/http"
)

func SimpleLog(err error) {
	if err != nil {
		panic(err)
	}
	return
}

func Log(err error, w http.ResponseWriter) http.ResponseWriter {
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusInternalServerError)
	return w
}

func SetResponse(w *http.ResponseWriter, message string, status int) {
		response := map[string]string{"message": message}
		js, _ := json.Marshal(response)
		(*w).Header().Set("Content-Type", "application/json")
		(*w).WriteHeader(status)
		(*w).Write(js)
	//(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func WriteResponse(w *http.ResponseWriter, data []byte) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusOK)
	(*w).Write(data)
}