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

	zap.S().Info(dict.Reversed())

	currentString := dict.Reversed()[previousCode]

	for _, currentCode := range encoded[1:] {
		var entry string
		if str, ok := dict.Reversed()[currentCode]; ok {
			entry = str
		} else {
			entry = currentString + string(currentString[0])
		}
		result.WriteString(entry)

		dict.Add(currentString + string(entry[0]))

		currentString = entry
	}

	return result.Bytes()
}

//func (s *service) Decode(encoded []int, dict *dictionary.Dictionary) []byte {
//	if len(encoded) == 0 {
//		return nil
//	}
//
//	var result bytes.Buffer
//
//	for _, code := range encoded {
//		entry, ok := dict.Reversed()[code]
//		if !ok {
//			// Handle the error appropriately, e.g., return an error or skip
//			zap.S().Errorf("Code %d not found in dictionary", code)
//			return nil
//		}
//		result.WriteString(entry)
//	}
//
//	return result.Bytes()
//}
