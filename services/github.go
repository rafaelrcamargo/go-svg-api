package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ApiData []struct {
	Total int   `json:"total"`
	Week  int64 `json:"week"`
	Days  []int `json:"days"`
}

func GetData() ApiData {
	resp, reqErr := http.Get("https://api.github.com/repos/RafaelRCamargo/RafaelRCamargo/stats/commit_activity")
	if reqErr != nil {
		panic(reqErr)
	}

	// We Read the response body on the line below.
	body, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		panic(ioErr)
	}

	// Convert the body to type string
	/* sb := string(body)
	println(sb) */

	// We create a empty ApiData struct
	apiJson := make(ApiData, 0)

	// Convert the JSON to the ApiData struct
	jsonErr := json.Unmarshal(body, &apiJson)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return apiJson
}
