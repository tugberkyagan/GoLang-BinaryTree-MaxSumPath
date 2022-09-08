package services

import (
	"fmt"
	"go-wrk/pkg/models"
	"math"
)

type BinaryTreeServiceInterface interface {
	Main(data models.JsonTree) float64
}

type BinaryTreeServiceStruct struct{}

var response float64

func (s BinaryTreeServiceStruct) Main(data models.JsonTree) float64 {

	nodeMap := make(map[string]models.JsonNode)
	for _, node := range data.Tree.Nodes {
		nodeMap[node.Id] = models.JsonNode{
			Left:  node.Left,
			Right: node.Right,
			Value: node.Value,
		}
	}

	nodes := s.generateNodes(data.Tree.Root, nodeMap)

	s.findMaxUtil(nodes)

	fmt.Println(response)

	return response
}

func (s BinaryTreeServiceStruct) generateNodes(root string, nodeMap map[string]models.JsonNode) *models.Node {
	if root == "" {
		return nil
	}

	var val int

	val = nodeMap[root].Value

	return &models.Node{
		Left:  s.generateNodes(nodeMap[root].Right, nodeMap),
		Right: s.generateNodes(nodeMap[root].Left, nodeMap),
		Value: float64(val),
	}
}

func (s BinaryTreeServiceStruct) findMaxUtil(root *models.Node) float64 {
	if root == nil {
		return 0
	}

	l := s.findMaxUtil(root.Left)
	r := s.findMaxUtil(root.Right)

	maxSingle := math.Max(math.Max(l, r)+root.Value, root.Value)

	maxTop := math.Max(maxSingle, l+r+root.Value)

	response = math.Max(response, maxTop)

	return maxSingle

}
