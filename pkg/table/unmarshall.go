package table

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func Unmarshall(buf []byte) (map[string]int, error) {
	m := make(map[string]int)

	r := bytes.NewReader(buf)

	scan := bufio.NewScanner(r)

	for scan.Scan() {
		k, v := mapMap(scan.Text())

		intVal, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal value %q", v)
		}

		m[k] = intVal
	}

	return m, nil
}

func mapMap(s string) (key, value string) {
	res := strings.Split(s, "\t")
	return res[0], res[1]
}
