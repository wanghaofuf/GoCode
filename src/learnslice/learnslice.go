/*************************************************************************
	> File Name: ../src/learnslice/learnslice.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Fri 20 Jan 2017 07:25:37 PM PST
 ************************************************************************/

package main

import "fmt"

func main() {
	fmt.Println(".....................................")
	var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,0}
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray: ")
	for _,v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice")
	for _,v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println(".....................................")
	mySlice1 := make([]int ,5)
	fmt.Println("\nElements of mySlice1")
	for _,v := range mySlice1 {
		fmt.Print(v, " ")
	}
	mySlice2 := make([]int ,5,10)
	fmt.Println("\nElements of mySlice2")
	for _,v := range mySlice2 {
		fmt.Print(v, " ")
	}
	mySlice3 := []int{1,2,3,4,5}
	fmt.Println("\nElements of mySlice3")
	for _,v := range mySlice3 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println(".....................................")

	fmt.Println("len(mySlice2):",len(mySlice2))
	fmt.Println("cap(mySlice2):",cap(mySlice2))
	fmt.Println()
	fmt.Println(".....................................")
	mySlice2 = append(mySlice2,1,2,3)
	fmt.Println("\nElements of mySlice2")
	for _,v := range mySlice2 {
		fmt.Print(v, " ")
	}
	mySlice2 = append(mySlice2,mySlice3...)
	fmt.Println("\nElements of mySlice2")
	for _,v := range mySlice2 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println(".....................................")
	fmt.Println("len(mySlice2):",len(mySlice2))
	fmt.Println("cap(mySlice2):",cap(mySlice2))
	fmt.Println()
	fmt.Println(".....................................")
	mySlice4 := mySlice2[:3]
	fmt.Println("\nElements of mySlice4")
	for _,v := range mySlice4 {
		fmt.Print(v, " ")
	}
	mySlice5 := mySlice2[:15]
	fmt.Println("\nElements of mySlice5")
	for _,v := range mySlice5 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println("cap(mySlice2):",cap(mySlice2))
	fmt.Println(".....................................")
	copy(mySlice3,mySlice5);
	fmt.Println("\nElements of mySlice3")
	for _,v := range mySlice3 {
		fmt.Print(v, " ")
	}
	copy(mySlice5,mySlice);
	fmt.Println("\nElements of mySlice5")
	for _,v := range mySlice5 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println("cap(mySlice5):",cap(mySlice5))
	fmt.Println(".....................................")


}
