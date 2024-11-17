package lzw

import (
	"github.com/phaseant/lzw-cli/internal/dictionary"
	"github.com/phaseant/lzw-cli/pkg/queue"
)

//func (s *service) Encode(text []byte, dict *dictionary.Dictionary) []int {
//	if len(text) == 0 {
//		return nil
//	}
//
//	var result []int
//	q := queue.New()
//	for _, c := range text {
//		// adding char to queue
//		q.Enqueue(c)
//
//		// if this symbol is not in dict
//		if _, ok := dict.Dictionary()[q.String()]; !ok {
//			dict.Add(q.String())
//		}
//
//		// adding prev to result
//		q.Dequeue()
//		result = append(result, dict.Dictionary()[q.String()])
//		//zap.S().Infof("res: %v, sym: %v", result, q.String())
//
//		// starting from cur symbol
//		q.Clean()
//		q.Enqueue(c)
//	}
//
//	return result
//}

func (s *service) Encode(text []byte, dict *dictionary.Dictionary) []int {
	if len(text) == 0 {
		return nil
	}

	var result []int
	q := queue.New() // Assume the queue can handle strings

	// Convert the text to a string to handle UTF-8 characters
	textStr := string(text)

	for _, r := range textStr {
		// `r` is a rune representing a character
		charStr := string(r)
		if charStr == "\n" {
			charStr = "\\n"
		}

		//adding char to queue
		q.Enqueue(charStr)

		// if this symbol is not in dict
		if _, ok := dict.Dictionary()[q.String()]; !ok {
			dict.Add(q.String())
		}

		// adding prev to result
		q.Dequeue()
		result = append(result, dict.Dictionary()[q.String()])
		//zap.S().Infof("res: %v, sym: %v", result, q.String())

		// starting from cur symbol
		q.Clean()
		q.Enqueue(charStr)
	}

	return result
}
