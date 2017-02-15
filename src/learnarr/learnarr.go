/*************************************************************************
	> File Name: ../src/learnarr/learnarr.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Fri 20 Jan 2017 06:40:18 PM PST
 ************************************************************************/

package main

import "fmt"

func modify(array [10]int) {
	array[0] = 10
	fmt.Println("In modify(), array values:",array)
}

func main() {
	array := [10]int{1,2,3,4,5}
	modify(array)
	fmt.Println("In main(), array values:",array)
	for i,v := range array {
		fmt.Println("Array element[",i, "]=",v)
	}
}
