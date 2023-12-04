package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Datarows []Datarow

type Datarow interface{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		os.Exit(1)
	}
	data, _ := ReadInput(os.Args[1])

	rv, err := ProcessData(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("error:%+v\nresult:\n%+v\n", err, rv)

	os.Exit(0)
}

func ReadInput(filename string) (Datarows, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file:", filename)
		return nil, err
	}

	lines := bytes.Split(data, []byte("\n"))

	rv := make(Datarows, len(lines))

	idx := 0
	for i, line := range lines {
		if len(line) == 0 {
			idx++
			continue
		}

		datarow := transformInputLine(line)

		rv[i] = datarow

	}

	return rv, nil
}

func transformInputLine(line []byte) Datarow {
	var rv Datarow

	rv = string(line)

	return rv
}

func ProcessData(data Datarows) (interface{}, error) {

	var rv=int(0)

	for _, item := range data {
		rv += ProcessLine(item.(string) )
	}

	return rv, nil
}

func ProcessLine(line string) int {
	r1 := regexp.MustCompile(`^[^\d]*(\d{1}).*?`)
	r2 := regexp.MustCompile(`.*?(\d{1})[^\d]*$`)

	r1m := r1.FindStringSubmatch(line)
	r2m := r2.FindStringSubmatch(line)
	
	ld := r1m[1]
	rd := r2m[1]

	ldi,_ := strconv.Atoi(ld)
	rdi,_ := strconv.Atoi(rd)

	rv := ldi*10+rdi
	
	return rv
}