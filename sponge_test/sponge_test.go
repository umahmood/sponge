package sponge_test

import (
	"testing"

	"github.com/umahmood/sponge"
)

func TestAbsorbByte(t *testing.T) {
	s := sponge.New(nil)
	s.AbsorbByte(42)
	wantHash := "2893d6fe0e6c2e98c68020d29acce67ef76bf1d35bb1294979572805c69221be"
	gotHash := s.Hex()
	if gotHash != wantHash {
		t.Errorf("fail got hash %s want hash %s", gotHash, wantHash)
	}
}

func TestAbsorbBytes(t *testing.T) {
	s := sponge.New(nil)
	s.AbsorbBytes([]byte("Though this be madness, yet there is method in't"))
	wantHash := "2c22bfeae7c0fdbd8985eee4f62be39750b6533b42869eb804de81ccee573af8"
	gotHash := s.Hex()
	if gotHash != wantHash {
		t.Errorf("fail got hash %s want hash %s", gotHash, wantHash)
	}
}

func TestSqueeze(t *testing.T) {
	s := sponge.New(nil)
	s.AbsorbBytes([]byte("But break, my heart, for I must hold my tongue"))
	gotBytes := s.Squeeze(10)
	wantBytes := []byte{12, 127, 203, 111, 181, 59, 8, 225, 120, 148}
	for idx, b := range gotBytes {
		if wantBytes[idx] != b {
			t.Errorf("fail got bytes %s want bytes %s", gotBytes, wantBytes)
		}
	}
}
