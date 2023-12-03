package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	readFile, _ := os.Open(os.Args[1])
	fileScanner := bufio.NewScanner(readFile)
	defer readFile.Close()
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// part 1
		for i := 0; i < 10000000; i++ {
			hash := md5.Sum([]byte(fmt.Sprintf("%s%d", line, i)))
			if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {
				fmt.Println(i)
				break
			}
		}

		// part 2
		for i := 0; i < 10000000; i++ {
			hash := md5.Sum([]byte(fmt.Sprintf("%s%d", line, i)))
			if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
				fmt.Println(i)
				break
			}
		}
	}

	duration := time.Since(start)
	fmt.Println(duration)
}
