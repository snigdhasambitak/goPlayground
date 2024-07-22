package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
	return &MyError{
		when: time.Now(),
		what: "Something went wrong",
	}
}

func main() {
	v := MyError{time.Now(), "this is not an error"}
	fmt.Println(v.when, v.what)

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
