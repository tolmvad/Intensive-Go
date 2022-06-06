package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func mean(intArr []int) float64 {
	sum := 0.0
	for _, v := range intArr {
		sum = sum + float64(v)
	}
	return sum / float64(len(intArr))
}

func median(intArr []int) float64 {
	if len(intArr)%2 == 0 {
		return float64(intArr[len(intArr)/2-1]+intArr[len(intArr)/2]) / 2.0
	}
	return float64(intArr[len(intArr)/2])
}

func mode(intArr []int) int {
	max, curV := 0, 0
	m := make(map[int]int)
	for _, v := range intArr {
		mV, ok := m[v]
		if ok {
			m[v] = mV + 1
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		if curV < v || (curV == v && k < max) {
			max = k
			curV = v
		}
	}
	return max
}

func sd(intArr []int) float64 {
	mid := mean(intArr)
	sum := 0.0
	for _, v := range intArr {
		sum = sum + math.Pow(float64(v)-mid, 2)
	}
	sum = math.Sqrt(sum) / float64(len(intArr))
	return sum
}

func checkLim(intArr []int) {
	if intArr[0] < -100000 {
		log.Fatal("some value < -100000")
	}
	if intArr[len(intArr)-1] > 100000 {
		log.Fatal("some value > 100000")
	}
}

func convSlice(str []string) []int {
	intArr := make([]int, 0)
	for _, v := range str {
		x, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		intArr = append(intArr, x)
	}
	sort.Ints(intArr)
	return intArr
}

func pars() []string {
	lineArr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			lineArr = append(lineArr, text)
		} else {
			break
		}
	}
	if len(lineArr) == 0 {
		fmt.Println("Error: values not given")
		os.Exit(1)
	}
	return lineArr
}

func main() {
	intArr := convSlice(pars())
	checkLim(intArr)
	meanFlag := flag.Bool("mean", false, "This flag to mean")
	medianFlag := flag.Bool("median", false, "This flag to median")
	modeFlag := flag.Bool("mode", false, "This flag to mode")
	sdFlag := flag.Bool("sd", false, "This flag to sd")
	flag.Parse()
	if *meanFlag || flag.NFlag() == 0 {
		fmt.Printf("Mean: %.2f\n", mean(intArr))
	}
	if *medianFlag || flag.NFlag() == 0 {
		fmt.Printf("Median: %.2f\n", median(intArr))
	}
	if *modeFlag || flag.NFlag() == 0 {
		fmt.Printf("Mode: %.2f\n", float64(mode(intArr)))
	}
	if *sdFlag || flag.NFlag() == 0 {
		fmt.Printf("SD: %.2f\n", sd(intArr))
	}
}
