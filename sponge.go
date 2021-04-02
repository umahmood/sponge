package sponge

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

// Sponge instance
type Sponge struct {
	state []byte
	hash  hash.Hash
}

// New creates a new sponge instance. You can specify any hash function which
// implements the hash.Hash interface. Default hash is SHA256.
func New(h hash.Hash) *Sponge {
	if h == nil {
		h = sha256.New()
	}
	s := &Sponge{hash: h}
	s.transform()
	return s
}

// transform sponge internal state
func (s *Sponge) transform() {
	s.hash.Write(s.state)
	s.state = s.hash.Sum(nil)
}

// absorb bytes
func (s *Sponge) absorb(buffer []byte) {
	for _, b := range buffer {
		s.state[0] = b ^ s.state[0]
		s.transform()
	}
}

// AbsorbByte absorbs a single byte into the sponge
func (s *Sponge) AbsorbByte(b byte) {
	s.absorb([]byte{b})
}

// AbsorbBytes absorbs multiple bytes into the sponge
func (s *Sponge) AbsorbBytes(b []byte) {
	s.absorb(b)
}

// Squeeze N bytes out of the sponge
func (s *Sponge) Squeeze(n int) []byte {
	output := []byte{}
	for i := 0; i < n; i++ {
		b := s.state[0]
		s.transform()
		output = append(output, b)
	}
	return output
}

// String returns the sponge state hex encoded
func (s Sponge) String() string {
	return hex.EncodeToString(s.state)
}

// Hex returns the sponge state hex encoded
func (s *Sponge) Hex() string {
	return hex.EncodeToString(s.state)
}
