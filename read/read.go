package read

import (
	"bufio"
	"os"
)

func Lines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var text []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}
