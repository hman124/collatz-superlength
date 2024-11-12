package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {

	var superLengthList [][]int = make([][]int, 0)
	var resultCache map[int]int = make(map[int]int)

	for baseNumber := 1; baseNumber <= 1_000_000; baseNumber++ {
		fmt.Println(baseNumber)
		var countList []int = []int{baseNumber}
		var functionLength int = 1
		var currentNumber int = baseNumber

		for {
			if baseNumber == 1_000_000 {
				fmt.Println(currentNumber)
			}
			if val, ok := resultCache[currentNumber]; ok {
				functionLength = val
			} else {
				startNumber := currentNumber
				for currentNumber != 1 {
					if currentNumber%2 == 0 {
						currentNumber = currentNumber / 2
					} else {
						currentNumber = (currentNumber * 3) + 1
					}

					functionLength++
				}

				resultCache[startNumber] = functionLength

			}

			// include the start
			if functionLength == 1 || slices.Contains(countList, functionLength) {
				break
			}

			countList = append(countList, functionLength)
			currentNumber = functionLength
			functionLength = 1

		}

		// fmt.Println(len(countList))
		superLengthList = append(superLengthList, []int{baseNumber, len(countList)})
	}

	fmt.Println("finished calculations")

	file, err := os.Create("./result.csv")
	if err != nil {
		return
	}
	defer file.Close()

	// Wrap the file in a buffered writer
	bufferedWriter := bufio.NewWriter(file)
	defer bufferedWriter.Flush() // Ensure any remaining data is written

	for _, i := range superLengthList {
		fmt.Fprintf(bufferedWriter, "%d,%d\n", i[0], i[1])
	}

	// f, err := os.Create("./result.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	defer f.Close()
	// }

	// var resultText string = "NUMBER,COUNT\n"
	// for i := 0; i < len(superLengthList); i++ {
	// 	v := superLengthList[i]
	// 	resultText += fmt.Sprintf("%s,%s\n", strconv.Itoa(v[0]), strconv.Itoa(v[1]))
	// }
	// f.Write([]byte(resultText))

	fmt.Println("algorithm finished")
}
