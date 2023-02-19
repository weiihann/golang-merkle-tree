package main

import (
	"github.com/ethereum/go-ethereum/crypto"
)
type BranchNode struct {
	Branches [16] Node
	Value []byte
}

func NewBranchNode() *BranchNode {
	return &BranchNode{
		Branches: [16]Node{},
	}
}

func (b BranchNode) Raw() []interface{} {
	hashes := make([]interface{}, 17)
	for i := 0; i < 16; i++ {
		if b.Branches[i] == nil {
			hashes[i] = EmptyNodeRaw
		} else {
			node := b.Branches[i]
			if len(Serialize(node)) >= 32 {
				hashes[i] = node.Hash()
			} else {
				// if node can be serialized to less than 32 bits, then
				// use Serialized directly.
				// it has to be ">=", rather than ">",
				// so that when deserialized, the content can be distinguished
				// by length
				hashes[i] = node.Raw()
			}
		}
	}

	hashes[16] = b.Value
	return hashes
}
func (b BranchNode) Serialize() []byte {
	return Serialize(b)
}

func (b BranchNode) Hash() []byte{
	return crypto.Keccak256(b.Serialize())
}

func (b BranchNode) AddBranch(nibbles Nibble, node Node) {
	b.Branches[int(nibbles)] = node
}

func (b BranchNode) RemoveBranch(nibbles Nibble){
	b.Branches[int(nibbles)] = nil
}

func (b BranchNode) SetValue(val []byte){
	b.Value = val 
}

func (b BranchNode) RemoveValue(){
	b.Value = nil
}