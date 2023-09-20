package main

import (
	"bufio"  // Sisend-väljund puhverdamiseks
	"fmt"    // Vormindatud I/O
	"math"   // Matemaatilised funktsioonid
	"os"     // Operatsioonisüsteemi interaktsioon
	"sort"   // Sorteerimise algoritmid
	"strconv"// Stringide teisendamine numbriteks
)

func main() {
	// Kontrolli, kas failitee on antud
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as an argument.")
		return
	}
	filePath := os.Args[1]

	// Ava fail
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() 

	var numbers []float64 // Massiiv numbrite jaoks
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // Loeb numbreid failist
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		numbers = append(numbers, num)
	}

	// Kontrolli, kas lugemisel ilmnes viga
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Kontrolli, kas numbrid leiti
	if len(numbers) == 0 {
		fmt.Println("No numbers found in the file.")
		return
	}

	// Arvuta keskmine
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	average := sum / float64(len(numbers))

	// Arvuta mediaan
	sort.Float64s(numbers)
	var median float64
	n := len(numbers)
	if n%2 == 0 {
		median = (numbers[n/2-1] + numbers[n/2]) / 2
	} else {
		median = numbers[n/2]
	}

	// Arvuta variatsioon
	var variance float64
	for _, num := range numbers {
		diff := num - average
		variance += diff * diff
	}
	variance = variance / float64(len(numbers))

	// Arvuta standardhälve
	stdDev := math.Sqrt(variance)

	// Väljasta tulemused
	fmt.Printf("Average: %.0f\n", math.Round(average))
	fmt.Printf("Median: %.0f\n", math.Round(median))
	fmt.Printf("Variance: %.0f\n", math.Round(variance))
	fmt.Printf("Standard Deviation: %.0f\n", math.Round(stdDev))
}
