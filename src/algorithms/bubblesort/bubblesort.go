/*************************************************************************
	> File Name: ../src/sorter/algorithms/bubblesort/bubblesort.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Mon 06 Feb 2017 11:14:55 PM PST
 ************************************************************************/

package bubblesort

func BubbleSort (values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j:=0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j+1],values[j] = values[j],values[j+1]
			} //end if
		} //end for j = ...
	} //end for i = ... 
}
