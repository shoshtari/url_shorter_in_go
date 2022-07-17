package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type url struct {
	Url string `json:"url"`
}

func Shorter(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data url
	decoder.Decode(&data)
	fmt.Fprintln(w, data)
}
