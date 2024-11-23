package table

import (
	"bytes"
	"fmt"
	"strings"
)

func Marshall(m map[string]int) ([]byte, error) {
	var b bytes.Buffer

	for k, v := range m {
		escapedKey := strings.ReplaceAll(k, "\n", "\\n")
		b.WriteString(writePair(escapedKey, v))
	}

	return b.Bytes(), nil
}

func writePair(k, v any) string {
	return fmt.Sprintf("%v\t%v\n", k, v)
}
