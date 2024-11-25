package main

import (
	"flag"
	"math/rand"
	"os"
	"time"
)

// argument -p specified percent of failure
// argument -f specifies input text file

// this program will print out text file supplied by -f, or print out "no input file received" if not supplied
// If -p is supplied, the program will fail with a probability 0-100. if less than 0, will be treated as 0, if more than 100, will be treated as 100
// the program will print if it is passing or failing, then pass or fail

func randomFailure(percentageSuccess int) bool {
	percentageFailure := 100 - percentageSuccess
	// if percentageFailure is less than 0, treat it as 0
	if percentageFailure < 0 {
		percentageFailure = 0
	}

	// if percentageFailure is more than 100, treat it as 100
	if percentageFailure > 100 {
		percentageFailure = 100
	}

	// if percentageFailure is 0, return true
	if percentageFailure == 0 {
		return true
	}

	// if percentageFailure is 100, return false
	if percentageFailure == 100 {
		return false
	}

	// if percentageFailure is between 0 and 100, generate a random number between 0 and 100
	// if the random number is less than percentageFailure, return false, else return true
	// this will give the program a percentageFailure chance of failing
	if rand.Intn(100) < percentageFailure {
		return false
	}
	return true
}

func createFailureFile() {
	file, err := os.Create("go-example-fail.txt")
	if err != nil {
		println("error creating file")
		os.Exit(1)
	}
	defer file.Close()
	dateTime := time.Now().Format(time.RFC3339)
	file.WriteString("failed on " + dateTime)
}

func main() {
	var percentageFailure int
	var inputFile string

	// setup flags with flags package
	flag.IntVar(&percentageFailure, "p", 100, "percentage of success chance")
	flag.StringVar(&inputFile, "f", "", "input file")
	flag.Parse()

	shouldPass := randomFailure(percentageFailure)

	if inputFile == "" {
		println("no input file received")
	} else {
		file, err := os.Open(inputFile)
		if err != nil {
			println("error opening file")
			os.Exit(1)
		}
		defer file.Close()
		// read in all file contents and print it out
		buf := make([]byte, 1024)
		for {
			n, _ := file.Read(buf)
			if n == 0 {
				break
			}
			os.Stdout.Write(buf[:n])
		}
	}

	if !shouldPass {
		println("fail")
		createFailureFile()
		os.Exit(1)
	}

	println("pass")
	os.Exit(0)
}
