package utils

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type MonogramStats struct {
	logScores map[byte]float64
	total     int
}

func NewMonogramStats() *MonogramStats {
	file, err := os.Open("../data/english_monograms.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	counts := make(map[byte]int)
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsed := strings.Fields(scanner.Text())
		char := parsed[0][0]
		count, err := strconv.Atoi(parsed[1])
		if err != nil {
			log.Fatal(err)
		}
		counts[char] = count
		total += count
	}

	logScores := make(map[byte]float64)
	for char, count := range counts {
		logScores[char] = math.Log10(float64(count) / float64(total))
	}

	return &MonogramStats{
		logScores: logScores,
		total:     total,
	}
}

func (m *MonogramStats) Score(text string) float64 {
	upper := strings.ToUpper(text)
	score := 0.0
	for i := range upper {
		if logScore, ok := m.logScores[upper[i]]; ok {
			score += logScore
		} else {
			score += math.Log10(0.01 / float64(m.total))
		}
	}
	return score
}
