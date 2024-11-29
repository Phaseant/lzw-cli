package table

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func Marshall(m map[string]int) ([]byte, error) {
	var b bytes.Buffer

	sorted := SortedByKeys(m)
	for _, item := range sorted {

		item.Val = strings.ReplaceAll(item.Val, "\n", "\\n")
		b.WriteString(item.String())
	}

	return b.Bytes(), nil
}

type Item struct {
	Key int
	Val string
}

func (i Item) String() string {
	return fmt.Sprintf("%v\t%v\n", i.Val, i.Key)
}

type SortedList []Item

func (s SortedList) String() string {
	var sb strings.Builder
	sb.WriteString("\n")
	for _, i := range s {
		sb.WriteString(i.String() + "\n")
	}
	return sb.String()
}

func SortedByKeys(m map[string]int) SortedList {
	sorted := make([]Item, 0, len(m))
	for k, v := range m {
		sorted = append(sorted, Item{Key: v, Val: k})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Key < sorted[j].Key
	})

	return sorted
}
