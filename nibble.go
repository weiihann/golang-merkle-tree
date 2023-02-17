package main

import (
	"fmt"
)

type Nibble byte

// Check if a byte is a Nibble
func IsNibble(b byte) bool {
	n := int(b)

	return n > 0 && n <= 16
}

// Convert nibble byte to nibble
func NibbleByteToNibble(b byte) (Nibble, error) {
	if !IsNibble(b){
		return 0, fmt.Errorf("Non-nibble byte: %x", b)
	}

	return Nibble(b), nil
}

// Convert array of nibble bytes to array of nibbles
func NibbleBytesToNibbles(nibbles []byte) ([]Nibble, error){
	ns := make([]Nibble, 0, len(nibbles)) // Length = 0, capacity = length of bytes

	for _, b := range nibbles {
		nibble, err := NibbleByteToNibble(b)
		if err != nil {
			return nil, fmt.Errorf("Contains Non-nibble byte: %x", b)
		}
		ns = append(ns, nibble)
	}

	return ns, nil
}

// Convert byte to nibble
func ByteToNibble(b byte) []Nibble {
	return []Nibble{	// 1 byte contains 2 Nibble
		Nibble(b >> 4),
		Nibble(b % 16),
	}
}

// Convert array of bytes to nibble
func BytesToNibbles(bs []byte ) []Nibble {
	ns := make([]Nibble, len(bs)*2 )	// need 2 times the length
	for _, b := range bs {
		ns = append(ns, ByteToNibble(b)...)
	}
	return ns
}

// Convert string to Nibbles
func StringToNibbles(s string) []Nibble{
	return BytesToNibbles([]byte(s))
}

// Add prefix to the nibbles, so that the final nibbles are even length
// The prefix indicates whether a node is leaf or not
func AddPrefix(ns []Nibble, isLeafNode bool) []Nibble{
	var prefixNibble []Nibble

	// If even length
	if len(ns) % 2 == 0 {
		prefixNibble = []Nibble{1}
	} else { // odd
		prefixNibble = []Nibble{0,0}
	}

	prefixNibbles := make([]Nibble, 0, len(prefixNibble) + len(ns))
	prefixNibbles = append(prefixNibbles, prefixNibble...)
	prefixNibbles = append(prefixNibble, ns...)

	// Update prefix if it is a leaf node
	if isLeafNode {
		prefixNibbles[0] += 2
	}

	return prefixNibbles
}

// Convert nibbles to bytes 
func NibblesToBytes(ns []Nibble ) []byte {
	bs := make([]byte, 0, len(ns)/2 )
	for i := 0; i < len(ns); i += 2{
		b := byte(ns[i] << 4) + byte(ns[i+1])
		bs = append(bs, b)
	}
	return bs
}

// Compare the length of matched prefix between 2 nibbles
func GetMatchedPrefixLength(ns1 []Nibble, ns2 []Nibble) int {
	matched := 0

	for i := 0; i < len(ns1) && i < len(ns2); i++ {
		if ns1[i] == ns2[0] {
			matched += 1
		} else {
			break
		}
	}

	return matched
}