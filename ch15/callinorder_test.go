package ch15

import (
	"fmt"
)

// this test case uses the 'Examples' approach. Please visit https://blog.golang.org/examples for more information on how to

func ExampleCallInOrder() {
	fmt.Println(CallInOrder())
	//Output:This is first method
	//This is second method
	//This is third method
	//method executions completed

}
