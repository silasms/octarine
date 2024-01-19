package handler

import (
	"net/http"
	"fmt"
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
)

type PostImg struct {
	Img string
}

func ImageGetAll(w http.ResponseWriter, r *http.Request) {
	entries, err := os.ReadDir("./../../internal/images")
	if err != nil {
		log.Fatal(err)
	}
	var list []string
	for _, e := range entries {
		list = append(list, string(e.Name()))
	}
	result := struct {
		Result []string
	}{ list }
	resJson, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%+v\n", string(resJson))
}

func SendImg(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// body, err := io.ReadAll(r.Body)
	var data PostImg
	json.NewDecoder(r.Body).Decode(&data)
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(data)
	r.Body.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(data.silas)
	// w.Write([]byte(data))
	fmt.Fprintf(w, data.Img)
	// arc, _ := os.Open("./../../internal/images/" + data.Img)
	arc, _ := ioutil.ReadFile("./../../internal/images/" + data.Img)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(arc)
}