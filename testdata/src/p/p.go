package p

import "errors"
import "fmt"


func funcWithReversedErrHandling() {
	err := errors.New("error")

	if err != nil {
		//"err" usage in Right place
		fmt.Println(err)
		fmt.Println("hello world")
	} else {
		//"err" not used in else, as it should
		fmt.Println("hello world")
	}

	i := 1
	if i != 2 {
		fmt.Println("skipped block")
	}

  var f *int
	if f != nil {
		fmt.Println("skipped block")
	}
		core := "test"
	if core != "" {
		fmt.Println("skipped block")
	}


	if err != nil {
		//"err" not used
		fmt.Println("hello world")
	} else {
		//"err" usage in wrong place
		fmt.Println(err)
	}
}
