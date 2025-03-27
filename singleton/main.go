package main

import (
	"fmt"
	"sync"
)

// Singleton struct
type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

// GetInstance ensures that only one instance is created
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating a new instance")
		instance = &Singleton{data: "Singleton Instance"}
	})
	return instance
}

func main() {
	// Multiple calls to GetInstance() return the same instance
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println("Instance 1:", s1)
	fmt.Println("Instance 2:", s2)

	// Check if both instances are the same
	if s1 == s2 {
		fmt.Println("Both instances are the same")
	} else {
		fmt.Println("Instances are different")
	}
}
