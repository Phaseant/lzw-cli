package lzw

import "github.com/phaseant/lzw-cli/internal/dictionary"

type Service interface {
	Encode(text []byte, dict *dictionary.Dictionary) []uint64
	Decode(text []uint64, dict *dictionary.Dictionary) []byte
}

type service struct {
}

func New() Service {
	return &service{}
}
