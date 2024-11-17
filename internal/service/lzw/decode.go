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
	previousString := dict.Reversed()[previousCode]
	result.WriteString(previousString)

	zap.S().Info(dict.Reversed())

	for _, currentCode := range encoded[1:] {
		var entry string
		if str, ok := dict.Reversed()[currentCode]; ok {
			entry = str
		} else {
			// Handle the special case where currentCode is not in the dictionary
			// Use runes to get the first character
			runes := []rune(previousString)
			if len(runes) == 0 {
				// Handle empty previousString to avoid panic
				zap.S().Error("previousString is empty when handling special case")
				return nil
			}
			entry = previousString + string(runes[0])
		}

		result.WriteString(entry)

		// Add new entry to the dictionary
		entryRunes := []rune(entry)
		if len(entryRunes) == 0 {
			// Handle empty entry to avoid panic
			zap.S().Error("entry is empty when adding to dictionary")
			return nil
		}
		dict.Add(previousString + string(entryRunes[0]))

		previousString = entry
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
