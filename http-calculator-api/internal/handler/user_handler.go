package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// bodyBytes, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		//
		data := Todo{}
		//
		// if err := json.Unmarshal(bodyBytes, &data); err != nil {
		// 	log.Fatal(err)
		// }

		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&data); err != nil {
			log.Fatal("Error decoding :", err)
		}
		log.Printf("Data from api :%+v", data)
	}

}
