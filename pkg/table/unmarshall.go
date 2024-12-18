package table

import (
	"bufio"
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

func Unmarshall(buf []byte) (map[string]uint64, error) {
	m := make(map[string]uint64)

	r := bytes.NewReader(buf)

	scan := bufio.NewScanner(r)

	for scan.Scan() {
		k, v := mapMap(scan.Text())

		intVal, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal value %q", v)
		}

		m[k] = intVal
	}

	return m, nil
}

func mapMap(s string) (key, value string) {
	if s == "" {
		zap.S().Info("empty string")
		return
	}
	res := strings.Split(s, "\t")
	return res[0], res[1]
}
