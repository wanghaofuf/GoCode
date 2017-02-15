/*************************************************************************
	> File Name: ../src/learnerror/learnerror.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Sat 21 Jan 2017 07:25:19 PM PST
 ************************************************************************/

package main

import "fmt"
import "os"


func main() {
	_, err := os.Stat("a.txt")

	if err != nil {
		fmt.Println(err)
		if e, ok := err.(*os.PathError); ok && e.Err != nil {
			fmt.Println(e.Err)
		}
	}
}
