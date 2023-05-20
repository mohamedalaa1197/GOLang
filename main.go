package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// const ==> means its value can't be overwritten again.
// its scope is on the package itself
const problemsFilename = "problems.csv"

func main() {

	// as the function returns two variables, then we can define two variables
	f, err := os.Open(problemsFilename)
	// in go we should handle the error in each functio, as there is no throw try/catch checkes in Go, if we didn't handle this then a panic will happen
	// and may cos the program to stop.
	if err != nil {
		fmt.Printf("failed to open the file %s", err)
		return
	}
	// to make sure after we finish dealing with the file we close the file, and release it back to the IO system, also useful when error happen as it will also run if an error happen
	defer f.Close()
	r = csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("failed to read the file %v", err)
		return
	}
	// fmt.Println(r)

	// to explain.
	// myR := myReader{}

	// r := csv.NewReader(myR)

	// fmt.Println(r)
}

// define a class in Go
type myReader struct {
	myInt    int
	myString string
	mybool   bool
}

// define a function in go
func Read(p []byte) (n int, err error) {
	return 0, nil
}

// to say this function related to this class in Go,
// adding a reciver in the function
func (r myReader) Write(p []byte) (n string, err error) {
	return "Hi there", nil
}

// type Reader interface {
// 	Read(p []byte)(n int, err error)
// };

func (r myReader) Read(p []byte) (n int, err error) {

	return 0, nil
}
