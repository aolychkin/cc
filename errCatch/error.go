package errCatch

import (
	"fmt"
	"log"
	"os"
)

func ErrDefaultDetect(err error, name string) {
	if err != nil {
		log.Fatal("ListenAndServe "+name+" : ", err)
	}
}

func Out(err error, text string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", text, err)
	}
}