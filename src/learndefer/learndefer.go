/*************************************************************************
	> File Name: ../src/learndefer/learndefer.go
	> Author: haofu.wang
	> Mail: wanghaofu@hotmail.com 
	> Created Time: Sat 21 Jan 2017 10:46:59 PM PST
 ************************************************************************/

package main

import "fmt"
import "os"
import "io"


func CopyFile (dst, src string) (w int64, err error) {

	srcFile, err := os.Open(src)
	if err != nil {
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile,srcFile)
}

func main(){
	CopyFile("b.txt","a.txt")
	fmt.Println()
}

