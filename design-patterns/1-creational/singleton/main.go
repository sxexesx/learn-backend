package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	fmt.Scanln()
}

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			singleInstance = &single{}
			fmt.Println("Single instance created now!")
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return singleInstance
}
