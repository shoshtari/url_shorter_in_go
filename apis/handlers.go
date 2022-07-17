package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shorter/utils"
)

type url struct {
	Url string `json:"url"`
}

var links = make(map[string]string)

func Shorter(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data url
	decoder.Decode(&data)
	shorted := utils.ShortLinkGen()
	links[shorted] = data.Url
	fmt.Fprintln(w, shorted, data.Url)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Path[len("/redirect/"):]
	if url, ok := links[token]; ok {
		//fmt.Fprintln(w, url)
		http.Redirect(w, r, url, 301)
		return
	}
	fmt.Fprintln(w, "Error")
}
