package config

import (
	"bufio"
	"os"
	"regexp"
)

func readPodLabels(labelFileName string) map[string]string {
	labels := make(map[string]string)
	labelFile, err := os.Open(labelFileName)
	defer labelFile.Close()

	if err == nil {
		scanner := bufio.NewScanner(labelFile)
		for scanner.Scan() {
			line := scanner.Text()
			key, value := convertLine(line)
			labels[key] = value
		}
	}

	return labels
}

func convertLine(line string) (string, string) {
	r := regexp.MustCompile(`^(.*)="?(.*?)"?$`)
	m := r.FindStringSubmatch(line)
	return m[1], m[2]
}
