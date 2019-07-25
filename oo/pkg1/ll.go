package pkg1

import "github.com/ldsute/go-practice/oo/shared"

type SLLNode struct {
	next         *SLLNode
	value        int
	SNodeMessage string
}

func (sNode *SLLNode) SetValue(v int) error {
	if sNode == nil {
		return shared.ErrInvalidNode
	}
	sNode.value = v
	return nil
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return &SLLNode{SNodeMessage: "This is the SLLNode default message"}
}
