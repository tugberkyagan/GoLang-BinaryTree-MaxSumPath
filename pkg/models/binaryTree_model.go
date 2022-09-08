package models

type JsonTree struct {
	Tree struct {
		Nodes []JsonNode `json:"nodes"`
		Root  string     `json:"root"`
	} `json:"tree"`
}

type JsonNode struct {
	Id    string `json:"id"`
	Left  string `json:"left"`
	Right string `json:"right"`
	Value int    `json:"value"`
}

type Node struct {
	Left  *Node
	Right *Node
	Value float64
}
