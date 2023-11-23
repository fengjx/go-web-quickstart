package kit

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadLine(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer closeq(f)
	r := bufio.NewReader(f)
	var bytes []byte
	var lines []string
	for {
		line, isPrefix, err := r.ReadLine()
		if err != nil {
			break
		}
		bytes = append(bytes, line...)
		if !isPrefix {
			str := strings.TrimSpace(string(bytes))
			if len(str) > 0 {
				lines = append(lines, str)
				bytes = []byte{}
			}
		}
	}
	if len(bytes) > 0 {
		lines = append(lines, string(bytes))
	}
	return lines, nil
}

func closeq(v interface{}) {
	if c, ok := v.(io.Closer); ok {
		silently(c.Close())
	}
}

func silently(_ ...interface{}) {}

func IsFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
