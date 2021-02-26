package main

type action struct {
	//Name     string `json:"name"`
	//Document string `json:"document"`
	//Input    string `json:"input"`
	//Output   string `json:"output"`
}

type metadata struct {
	ApiVersion       string `json:"apiVersion"`
	APIBrief         string `json:"api_brief"`
	ServiceNameCN    string `json:"serviceNameCN"`
	ServiceShortName string `json:"serviceShortName"`
}

type object struct {
	//	useless now
}

type model struct {
	Meta metadata `json:"metadata"`
	//Version string            `json:"version"`
	Actions map[string]action `json:"actions"`
	Objects map[string]object `json:"objects"`
}
