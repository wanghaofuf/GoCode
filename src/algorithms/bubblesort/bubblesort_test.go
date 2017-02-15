/*************************************************************************
	> File Name: ../src/sorter/algorithms/bubblesort/bubblesort_test.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Mon 06 Feb 2017 11:22:57 PM PST
 ************************************************************************/

package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{5,4,3,2,1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("BubbleSort() failed. Got",values, "Expected 1 2 3 4 5")
	}
}

func TestBubbleSort2 (t *testing.T) {
	values := []int{5,5,3,2,1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("BubbleSort() failed. Got",values, "Expected 1 2 3 5 5")
	}
}

func TestBubbleSort3 (t *testing.T) {
	values := []int{5}
	BubbleSort(values)
	if values[0] != 5 {
		t.Error("BubbleSort() failed. Got",values, "Expected 5")
	}
}


