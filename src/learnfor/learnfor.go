/*************************************************************************
	> File Name: ../src/learnfor/learnfor.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Fri 20 Jan 2017 11:11:43 PM PST
 ************************************************************************/

package main

import "fmt"

func main(){
	a:=[]int{1,2,3,4,5,6,7,8,9}
	for i,j:=0,len(a)-1;i<j;i,j = i+1,j-1 {
		a[i],a[j]= a[j],a[i]
	}
	for _,v := range a {
		fmt.Print(" ",v)
	}
	fmt.Println()
}
