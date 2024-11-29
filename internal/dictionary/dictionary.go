package dictionary

import (
	"fmt"
	table2 "github.com/phaseant/lzw-cli/pkg/table"
	"sort"
	"strings"
)

type Dictionary struct {
	m map[string]uint64
}

func New(input []byte) (*Dictionary, error) {
	m := populateMap(input)
	return &Dictionary{m: m}, nil
}

func Open(table []byte) (*Dictionary, error) {
	m, err := table2.Unmarshall(table)
	if err != nil {
		return nil, err
	}

	return &Dictionary{m: m}, nil
}

func (d *Dictionary) Dictionary() map[string]uint64 {
	return d.m
}

func (d *Dictionary) Len() int {
	return len(d.m)
}

func (d *Dictionary) Add(key string) {
	d.m[key] = uint64(len(d.m))
}

func (d *Dictionary) Reversed() map[uint64]string {
	reversed := make(map[uint64]string)

	for k, v := range d.m {
		reversed[v] = k
	}

	return reversed
}

func (d *Dictionary) Marshall() ([]byte, error) {
	return table2.Marshall(d.m)
}

func populateMap(text []byte) map[string]uint64 {
	m := make(map[string]uint64)

	// Convert the byte slice to a string to handle UTF-8 decoding
	textStr := string(text)

	for _, r := range textStr {
		ch := string(r)
		if _, ok := m[ch]; !ok {
			m[ch] = uint64(len(m))
		}
	}

	return m
}

type Item struct {
	Key uint64
	Val string
}

func (i Item) String() string {
	return fmt.Sprintf("%d: %s", i.Key, i.Val)
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

func (d *Dictionary) SortedByKeys() SortedList {
	sorted := make([]Item, 0, len(d.m))
	for k, v := range d.m {
		sorted = append(sorted, Item{Key: v, Val: k})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Key < sorted[j].Key
	})

	return sorted
}
