package pkg2

import "github.com/ldsute/go-practice/oo/shared"

type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

func (pNode *PowerNode) SetValue(v int) error {
	if pNode == nil {
		return shared.ErrInvalidNode
	}
	pNode.value = v * 10
	return nil
}

func (pNode *PowerNode) GetValue() int {
	return pNode.value
}

func NewPowerNode() *PowerNode {
	return &PowerNode{PNodeMessage: "This is the PowerNode default message"}
}
