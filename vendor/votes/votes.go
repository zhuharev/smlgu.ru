package votes

import (
	"bufio"
	//"fmt"
	"os"
	//"strings"
)

func GetScanner() (*bufio.Scanner, error) {
	f, e := os.OpenFile("vendor/votes/votes", os.O_RDONLY, 0777)
	if e != nil {
		return nil, e
	}
	scanner := bufio.NewScanner(f)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		start := 0
		for i := 0; i < len(data); i++ {
			start++
			if data[i] == '(' {
				break
			}
		}
		for i := start; i < len(data); i++ {
			if data[i] == ')' {
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
