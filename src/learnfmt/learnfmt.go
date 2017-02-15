/*************************************************************************
	> File Name: learnfmt.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Fri 20 Jan 2017 06:23:11 PM PST
 ************************************************************************/

package main

import "fmt"

func main() {
	fval := 110.123456789
	ival := 200
	sval := "This is a string. "
	fmt.Println("The value of fval is", fval)
	fmt.Printf("fval=%f, ival=%d, sval=%s\n",fval, ival, sval)
	fmt.Printf("fval=%v, ival=%v, sval=%v\n",fval, ival, sval)
}
