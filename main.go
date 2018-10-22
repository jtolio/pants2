package main

import (
	"fmt"
	"io"
	"os"
)

func handleErr(err error) {
	if err == nil {
		return
	}
	if IsHandledError(err) {
		_, err = fmt.Println(err)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func main() {
	p := NewParser("<input>")
	ls := NewReaderLineSource(os.Stdin, func() error {
		_, err := fmt.Printf("> ")
		return err
	})
	for {
		stmt, err := p.ParseNext(ls)
		if err != nil {
			if err == io.EOF {
				break
			}
			handleErr(err)
		}
		handleErr(Run(stmt))
	}
}

func Run(stmt Stmt) error {
	panic("TODO")
}