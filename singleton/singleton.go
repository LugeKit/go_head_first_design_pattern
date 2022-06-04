package singleton

import (
	"fmt"
	"sync"
)

type UniqueObject struct{}

var unique *UniqueObject
var once sync.Once

func GetUnique() *UniqueObject {
	once.Do(func() {
		unique = &UniqueObject{}
		fmt.Println("unique is created at the first time!")
	})
	return unique
}
