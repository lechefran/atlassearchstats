package main

import (
	"atlassearchstatus/model"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var collScanStats []model.FileStats
	var idxScanStats []model.FileStats
	var atlasSearchStats []model.FileStats
	var avgs []float64
	var mins []float64
	var maxs []float64

	// collection scan results
	collScanStats = append(collScanStats, getStats("./log/col-scan-restaurant-id-results.txt"))
	collScanStats = append(collScanStats, getStats("./log/col-scan-owner-name-results.txt"))
	collScanStats = append(collScanStats, getStats("./log/col-scan-city-results.txt"))
	collScanStats = append(collScanStats, getStats("./log/col-scan-state-results.txt"))
	collScanStats = append(collScanStats, getStats("./log/col-scan-country-results.txt"))

	// index scan results
	idxScanStats = append(idxScanStats, getStats("./log/idx-scan-restaurant-id-results.txt"))
	idxScanStats = append(idxScanStats, getStats("./log/idx-scan-owner-name-results.txt"))
	idxScanStats = append(idxScanStats, getStats("./log/idx-scan-city-results.txt"))
	idxScanStats = append(idxScanStats, getStats("./log/idx-scan-state-results.txt"))
	idxScanStats = append(idxScanStats, getStats("./log/idx-scan-country-results.txt"))

	// atlas search results
	atlasSearchStats = append(atlasSearchStats, getStats("./log/atlas-search-restaurant-id-results.txt"))
	atlasSearchStats = append(atlasSearchStats, getStats("./log/atlas-search-owner-name-results.txt"))
	atlasSearchStats = append(atlasSearchStats, getStats("./log/atlas-search-city-results.txt"))
	atlasSearchStats = append(atlasSearchStats, getStats("./log/atlas-search-state-results.txt"))
	atlasSearchStats = append(atlasSearchStats, getStats("./log/atlas-search-country-results.txt"))

	for _, s := range collScanStats {
		printStats(s)
		avgs = append(avgs, s.Average)
		mins = append(mins, s.Min)
		maxs = append(maxs, s.Max)
	}

	fmt.Printf("Average average time for collection scan: %.3fs\n", avgVal(avgs))
	fmt.Printf("Average minimum time for collection scan: %.3fs\n", avgVal(mins))
	fmt.Printf("Average maximum time for collection scan: %.3fs\n", avgVal(maxs))

	avgs = []float64{}
	mins = []float64{}
	maxs = []float64{}

	for _, s := range idxScanStats {
		printStats(s)
		avgs = append(avgs, s.Average)
		mins = append(mins, s.Min)
		maxs = append(maxs, s.Max)
	}

	fmt.Printf("Average average time for index scan: %.3fs\n", avgVal(avgs))
	fmt.Printf("Average minimum time for index scan: %.3fs\n", avgVal(mins))
	fmt.Printf("Average maximum time for index scan: %.3fs\n", avgVal(maxs))

	avgs = []float64{}
	mins = []float64{}
	maxs = []float64{}

	for _, s := range atlasSearchStats {
		printStats(s)
		avgs = append(avgs, s.Average)
		mins = append(mins, s.Min)
		maxs = append(maxs, s.Max)
	}

	fmt.Printf("Average average time for atlas search: %.3fs\n", avgVal(avgs))
	fmt.Printf("Average minimum time for atlas search: %.3fs\n", avgVal(mins))
	fmt.Printf("Average maximum time for atlas search: %.3fs\n", avgVal(maxs))
}

func getStats(f string) model.FileStats {
	floats := parseTextFile(f)
	return model.FileStats{
		FileName: f,
		Average:  avgVal(floats),
		Min:      minVal(floats),
		Max:      maxVal(floats),
	}
}

func printStats(stats model.FileStats) {
	fmt.Printf("File Name: %s\n", stats.FileName)
	fmt.Printf("Average time: %.3fs\n", stats.Average)
	fmt.Printf("Min time: %.3fs\n", stats.Min)
	fmt.Printf("Max time: %.3fs\n\n", stats.Max)
}

func parseTextFile(f string) []float64 {
	var strings []string
	var res []float64
	rgx := regexp.MustCompile(`\d+\.\d+`)

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, s := range strings {
		tmp := s
		idx := rgx.FindStringIndex(tmp)
		if idx != nil {
			val, _ := strconv.ParseFloat(tmp[idx[0]:idx[1]], 64)
			res = append(res, val)
		}
	}
	return res
}

func avgVal(floats []float64) float64 {
	var res float64
	for _, f := range floats {
		res += f
	}
	return res / float64(len(floats))
}

func minVal(floats []float64) float64 {
	res := math.MaxFloat64
	for _, f := range floats {
		if f < res {
			res = f
		}
	}
	return res
}

func maxVal(floats []float64) float64 {
	res := math.SmallestNonzeroFloat64
	for _, f := range floats {
		if f > res {
			res = f
		}
	}
	return res
}
