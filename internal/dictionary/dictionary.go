package dictionary

import table2 "github.com/phaseant/lzw-cli/pkg/table"

type Dictionary struct {
	m map[string]int
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

func (d *Dictionary) Dictionary() map[string]int {
	return d.m
}

func (d *Dictionary) Len() int {
	return len(d.m)
}

func (d *Dictionary) Add(key string) {
	d.m[key] = len(d.m)
}

func (d *Dictionary) Reversed() map[int]string {
	reversed := make(map[int]string)

	for k, v := range d.m {
		reversed[v] = k
	}

	return reversed
}

func (d *Dictionary) Marshall() ([]byte, error) {
	return table2.Marshall(d.m)
}

func populateMap(text []byte) map[string]int {
	m := make(map[string]int)

	for _, char := range text {
		if _, ok := m[string(char)]; !ok {
			m[string(char)] = len(m)
		}
	}

	return m
}
