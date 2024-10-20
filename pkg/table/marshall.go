package table

import (
	"bytes"
	"fmt"
)

func Marshall(m map[string]int) ([]byte, error) {
	var b bytes.Buffer

	for k, v := range m {
		b.WriteString(writePair(k, v))
	}

	return b.Bytes(), nil
}

func writePair(k, v any) string {
	return fmt.Sprintf("%v\t%v\n", k, v)
}
