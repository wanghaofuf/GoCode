/*************************************************************************
	> File Name: ../src/learnclosure/learnclosure.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Sat 21 Jan 2017 01:28:47 AM PST
 ************************************************************************/

package main

import "fmt"

func main() {

	var j int = 5
	a := func() (func()) {
		var i int = 10
			return func() {
				fmt.Printf("i,j: %d, %d\n",i,j)
			}
	}()

	a()
	j *= 2
	a()
	
}
