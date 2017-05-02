package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("starting...")

	numberStr := os.Getenv("NUMBER")
	if len(numberStr) == 0 {
		numberStr = "3000000"
	}
	fmt.Println("will run", numberStr, "loops...")

	number, err := strconv.Atoi(numberStr)
	check(err)

	var (
		data     map[int]int
		filePath = "/tmp/testfile"
		count    uint64
	)

	for {
		fmt.Println("loop", count)
		start := time.Now()
		genData := start
		data = make(map[int]int)
		for i := 0; i < number; i++ {
			data[i] = i
		}
		byt, err := json.Marshal(data)
		check(err)
		fmt.Println("generating data:", time.Now().Sub(genData))

		genData = time.Now()
		err = ioutil.WriteFile(filePath, byt, 0644)
		check(err)
		fmt.Println("writting file:", time.Now().Sub(genData))

		genData = time.Now()
		err = os.Remove(filePath)
		check(err)
		fmt.Println("deleting file:", time.Now().Sub(genData))

		fmt.Println("-------- total:", time.Now().Sub(start))
		time.Sleep(500 * time.Millisecond)
		count++
	}
}
