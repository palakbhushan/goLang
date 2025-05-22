package main

import (
	"fmt"
	"net/http"
	"math/cmplx"
)

// tree shaking removes unneccessary code , and mark it as a dead code, and remove it in build process

func main(){
	fmt.Println("go import!")

	resp, err:= http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil{
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Http resp: ", resp.Status)
}