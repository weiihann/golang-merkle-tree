package main

import (
	"github.com/ethereum/go-ethereum/crypto"
)

type ExtensionNode struct {
	Path []Nibble
	Next Node
}

func NewExtensionNode(nibbles []Nibble, next Node) *ExtensionNode{
	return &ExtensionNode{
		Path: nibbles,
		Next: next,
	}
}

// Serialize node into bytes
func (e ExtensionNode) Serialize() []byte {
	return Serialize(e)
}

// Implicitly implement the Node interface 
func (e ExtensionNode) Hash() []byte {
	return crypto.Keccak256(e.Serialize())
}

// Implicitly implmenent the Node interface
func (e ExtensionNode) Raw() []interface{} {
	hashes := make([]interface{}, 2)
	hashes[0] = NibblesToBytes(AddPrefix(e.Path, false))
	if len(Serialize(e.Next)) >= 32 {
		hashes[1] = e.Next.Hash()
	} else {
		hashes[1] = e.Next.Raw()
	}
	return hashes
}