/*************************************************************************
	> File Name: ../src/sorter/algorithms/qsort/qsort.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Mon 06 Feb 2017 11:32:21 PM PST
 ************************************************************************/

package qsort

func quickSort (values []int , left,right int) {
	p := (left+right) / 2
	temp := values[p]
	i,j := left, right
	for  i< j {
		for values[i] < temp && i < p {
			i++
		}
		if i < p {
			values[p] = values[i]
			p = i
		}
		for j > p && values[j] >= temp {
			j--
		}
		if j > p {
			values[p] = values[j]
			p = j
		}
	}
	values[p] = temp
	if p - left > 1 {
		quickSort(values,left,p-1)
	}
	if right - p > 1 {
		quickSort(values,p+1,right)
	}
}
func QuickSort(values []int) {
	quickSort(values,0,len(values)-1)
}
