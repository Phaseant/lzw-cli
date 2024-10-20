package lzw

import (
	"bytes"
	"github.com/phaseant/lzw-cli/internal/dictionary"
	"go.uber.org/zap"
)

func (s *service) Decode(encoded []int, dict *dictionary.Dictionary) []byte {
	if len(encoded) == 0 {
		return nil
	}

	zap.S().Info(encoded)

	var result bytes.Buffer
	// Initialize the previous code with the first code from the encoded data
	previousCode := encoded[0]
	result.WriteString(dict.Reversed()[previousCode])

	currentString := dict.Reversed()[previousCode]

	for _, currentCode := range encoded[1:] {

		var entry string
		if str, ok := dict.Reversed()[currentCode]; ok {
			entry = str
		} else {
			entry = currentString + string(currentString[0])
		}
		result.WriteString(entry)

		zap.S().Info(result.String())

		dict.Add(currentString + string(entry[0]))

		currentString = entry
	}

	return result.Bytes()
}
