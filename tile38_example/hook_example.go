package tile38_example

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// json struct
//{
//	"command": "set",
//	"detect": "enter",
//	"hook": "warehouse",
//	"time": "2016-03-20T09:37:49.567854293-07:00",
//	"key": "fleet",
//	"id": "truck02",
//	"object": {
//		"type": "Point",
//		"coordinates": [-112.2695, 33.4626]
//	}
//}

type Message struct {
	Command string `json:"command"`
	Detect  string `json:"detect"`
	Hook    string `json:"hook"`
	Time    string `json:"time"`
	Key     string `json:"key"`
	Id      string `json:"id"`
	Object  Point  `json:"object"`
}

type Point struct {
	Type        string    `json:"type"`
	Coordinates []float32 `json:"coordinates"`
}

func Listen(w http.ResponseWriter, r *http.Request) {
	result, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var msg Message
	json.Unmarshal(result, &msg)
	log.Println(msg)
}

//func main() {
//	http.HandleFunc("/", Listen)
//	err := http.ListenAndServe(":9000", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
