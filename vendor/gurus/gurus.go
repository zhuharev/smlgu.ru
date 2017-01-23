package gurus

import (
	"bufio"
	"encoding/csv"
	//"fmt"
	//"fmt"
	"os"
	"strings"
)

func GetScanner() (*bufio.Scanner, error) {
	f, e := os.OpenFile("vendor/gurus/gurus", os.O_RDONLY, 0777)
	if e != nil {
		return nil, e
	}
	scanner := bufio.NewScanner(f)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		start := 0
		quotesCount := 0
		for i := 0; i < len(data); i++ {
			start++
			if data[i] == '(' {
				break
			}
		}
		for i := start; i < len(data); i++ {
			if data[i] == '\'' {
				quotesCount++
			}
			if data[i] == ')' && quotesCount%2 == 0 {
				return i + 1, data[start:i], nil
			}
		}
		return start, nil, nil
	}

	scanner.Split(split)

	return scanner, nil

	//
	// Scan.
	/*i := 0
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ",")

	}
	fmt.Println(i)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}*/
}

func ScanValues(in string) (res []string) {
	in = strings.Replace(in, "'", "\"", -1)
	rdr := strings.NewReader(in)
	crdr := csv.NewReader(rdr)
	res, _ = crdr.Read()
	return
	/*scanner := bufio.NewScanner(rdr)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		isWords := false
		start := 0

		if len(data) > 1 && data[0] == '\'' && data[1] == '\'' {
			return 4, []byte("''"), nil
		}

		if data[start] == '\'' {
			start++
			isWords = true
			if len(data) > start+1 && data[start+1] == '\'' {
				return start + 1, nil, nil
			}
		}

		for i := start; i < len(data); i++ {
			if data[i] == ',' && !isWords || data[i] == '\'' {
				fmt.Println(string(data[start:i]))
				return i + 1, data[start:i], nil
			}
		}

		return start, nil, nil
	}
	scanner.Split(split)
	for scanner.Scan() {
		res = append(res, strings.TrimSpace(scanner.Text()))
	}
	return*/
}
