package handler

import (
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// resp := make(map[string]string)
	// resp["Message"] = "Hello World!"
	//
	// jsonResp, err := json.Marshal(resp)
	//
	// if err != nil {
	// 	log.Fatalf("Error handling JSON marshal. Err:%v", err)
	// }
	response := "Hello World!"
	_, _ = w.Write([]byte(response))

}
