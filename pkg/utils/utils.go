package utils

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("error occurred: %s\n", err.Error())
	os.Exit(1)
}

func FatalPrint(s string) {
	fmt.Println(s)
	os.Exit(1)
}
