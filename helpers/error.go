package helpers

import "fmt"

func CheckPanicError(err error)  {
	fmt.Println("inside check for panic error")
	if err != nil {
		panic(err)
	}
}