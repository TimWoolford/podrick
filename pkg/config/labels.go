package config

import (
	"os"
	"strings"
	"bufio"
)

func readPodLabels(labelFile string) map[string]string {
	labels := make(map[string]string)
	labelData, err := os.Open(labelFile)

	if err == nil {
		scanner := bufio.NewScanner(labelData)
		for scanner.Scan() {
			line := scanner.Text()
			i := strings.IndexRune(line, '=')
			labels[line[:i]] = strings.TrimLeft(strings.TrimRight(line[i+1:],"\"" ), "\"")
		}
	}

	return labels
}