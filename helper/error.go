package helper

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		panic(err)
	}
}
