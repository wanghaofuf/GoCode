/*************************************************************************
	> File Name: ../src/learnmap/learnmap.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Fri 20 Jan 2017 10:47:11 PM PST
 ************************************************************************/

package main

import "fmt"

type PersonInfo struct {
	ID string
	Name string
	Address string
}


func main(){
	var personDB map[string] PersonInfo
	personDB = make(map[string] PersonInfo)

	personDB["12345"] = PersonInfo{"12345","Tom","Room 203,..."}
	personDB["1"] = PersonInfo{"1","Jack","Room 101,..."}

	person, ok := personDB["1"]
	if ok {
		fmt.Println("Found person",person.Name,"with ID ",person.ID,".")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}

}
