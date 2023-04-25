package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"gopkg.in/yaml.v2"
	"strings"
)

type Searches struct {
	Searches []Search `yaml:"searches"`
}

type Search struct {
	ServiceName   string   `yaml:"serviceName"`
	AllowedMethods []string `yaml:"allowedMethods"`
	MaxLatencyMs  int      `yaml:"maxLatencyMs"`
	Bark bool			   `yaml:"bark"`
}

func main() {

	data, err := ioutil.ReadFile("searches.yaml")
	if err != nil {
		panic(err)
	}
	
	var searches Searches

	err = yaml.Unmarshal(data, &searches)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:3200"
	searchPath := "/api/search"

	for _, search := range searches.Searches {
		fmt.Printf("Service: %s, Allowed methods: %v, Max latency: %dms, Bark: %t\n", search.ServiceName, search.AllowedMethods, search.MaxLatencyMs, search.Bark)

		allowedMethodsString := strings.Join(search.AllowedMethods, "|")
		serviceName := search.ServiceName

		query := url.Values{}
		query.Set("q", fmt.Sprintf("{ .http.method =~ \"%s\" && .service.name = \"%s\"}", allowedMethodsString, serviceName))
		queryString := query.Encode()

		searchURL := fmt.Sprintf("%s%s?%s", baseURL, searchPath, queryString)

		req, err := http.NewRequest("GET", searchURL, nil)

		if err != nil {
			panic(err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		
		if err != nil {
			panic(err)
		}
		
		defer resp.Body.Close()


		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(body))
	}

}
