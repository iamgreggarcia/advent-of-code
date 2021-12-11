package main

import "fmt"

func main() {
	values := newReport()

	slidingCount := slidingWindowCount(values)
	count := numberOfIncreases(values)
	fmt.Printf("There are %d measurements that are larger than the previous measurement\n", count)
	fmt.Printf("There ")
}

func numberOfIncreases(values report) int {
	var count int
	for i, val := range values {
		if i != 0 {
			if values[i - 1] < val {
				fmt.Printf("%d (increased)\n", val)
				count++
			} else {
				fmt.Printf("%d (decreased)\n", val)
			}
		} else {
			fmt.Printf("%d (N/A - no previous measurement\n", val)
		}
	}
	return count
}

func slidingWindowCount(values report) int {
	var count int
	for i, val := range values {
		if (i < len(values) - 2) {
			if i > 0 {
				curSum := val + values[i + 1] + values[i + 2]
				prevSum := values[i - 1] + values[i] + values[i + 1]
				fmt.Printf("Current sum: %d\n", curSum)
				fmt.Printf("Previous sum: %d\n", prevSum)
				if curSum > prevSum {
					fmt.Printf("%v: %d (inceased)\n\n", IntToLetters(i + 1), curSum)
					count++
				} else if (curSum == prevSum) {
					fmt.Printf("%v: %d (no change)\n\n", IntToLetters(i + 1), curSum)
				} else {
					fmt.Printf("%v: %d (decreased)\n\n", IntToLetters(i + 1), curSum)
				}
			} else {
				curSum := val + values[i + 1] + values[i + 2]
				nextSum := values[i + 1] + values[i + 2] + values[i + 3]
				fmt.Printf("First sum: %d\n", curSum)
				fmt.Printf("Next sum: %d\n", nextSum)
				fmt.Printf("%v: %d (N/A - no previous measurement)\n\n",IntToLetters(i + 1), curSum)
			}
		}
	}
	return count
}

func IntToLetters(number int) (letters string){
    number--
    if firstLetter := number/26; firstLetter >0{
        letters += IntToLetters(firstLetter)
        letters += string('A' + number%26)
    } else {
        letters += string('A' + number)
    }
        
    return 
}
