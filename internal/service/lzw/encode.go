package lzw

import (
	"github.com/phaseant/lzw-cli/internal/dictionary"
	"github.com/phaseant/lzw-cli/pkg/queue"
)

func (s *service) Encode(text []byte, dict *dictionary.Dictionary) []int {
	if len(text) == 0 {
		return nil
	}

	var result []int
	q := queue.New()
	for _, c := range text {
		// adding char to queue
		q.Enqueue(c)

		// if this symbol is not in dict
		if _, ok := dict.Dictionary()[q.String()]; !ok {
			dict.Add(q.String())

			// adding prev to result
			q.Dequeue()
			result = append(result, dict.Dictionary()[q.String()])

			// starting from cur symbol
			q.Clean()
			q.Enqueue(c)
		}
	}

	// adding last
	if q.Len() != 0 {
		result = append(result, dict.Dictionary()[q.String()])
	}

	return result
}
