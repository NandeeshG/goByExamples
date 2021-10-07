package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadCSV(f *os.File) ([][]string, error) {
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		return lines, err
	}

	return lines, nil
}

func ValidateCohortCSV(lines [][]string) ([][]string, error) {

	yearsThreshold := 100
	past := (time.Now().AddDate(-yearsThreshold, 0, 0)).Unix()
	future := (time.Now().AddDate(yearsThreshold, 0, 0)).Unix()
	//fmt.Printf("past: %v\n", past)
	//fmt.Printf("future: %v\n", future)

	for idx := range lines {
		for idy := range lines[idx] {
			val := &lines[idx][idy]
			*val = strings.TrimSpace(*val)
			if idy > 0 { //if curr val is a timestamp

				//convert to float, round it off, convert to int
				fl, err := strconv.ParseFloat(*val, 64)
				if err != nil {
					return nil, err
				}
				ts := int64(math.Round(fl))
				//fmt.Printf("ts: %v and %v\n", ts, fl)

				////do we want to allow negative timestamps?
				//if fl < 0 {
				//	return nil, fmt.Errorf("Timestamp cannot be negative")
				//}

				//check time in range
				if ts < past {
					return nil, fmt.Errorf("Timestamp(%v) cannot be more than %v years old", ts, yearsThreshold)
				}
				if ts > future {
					return nil, fmt.Errorf("Timestamp(%v) should be within the next %v years", ts, yearsThreshold)
				}

				////(optional) Take time package's help to confirm timestamp integrity
				//timeStr := time.Unix(int64(ts), 0)
				//*val = strconv.FormatInt(timeStr.Unix(), 10)
				////fmt.Printf("timeStr: %v\n", timeStr)

				*val = strconv.FormatInt(ts, 10)
			}
		}
	}
	return lines, nil
}

func main() {
	//f, e := os.Open("readingCSV/patients.csv")
	//if e != nil {
	//	fmt.Println("Error: ", e)
	//} else {
	//	s, e := ReadCSV(f)
	//	if e != nil {
	//		fmt.Println("Read error: ", e)
	//	} else {
	//		s, e := ValidateCohortCSV(s[1:])
	//		if e != nil {
	//			fmt.Println("Validate error: ", e)
	//		} else {
	//			fmt.Printf("%T: %v\n", s, s)

	//		}
	//	}
	//}

	f, e := os.Open("readingCSV/patients.csv")
	if e != nil {
		fmt.Println("Error: ", e)
	}
	s, e := ReadCSV(f)
	if e != nil {
		fmt.Println("Read error: ", e)
	}
	v, e := ValidateCohortCSV(s[1:])
	if e != nil {
		fmt.Println("Validate error: ", e)
	}
	fmt.Printf("%T: %v\n", v, v)
}

// Only header but no data?
//
