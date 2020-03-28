package solutions

/*
Consider a staircase of size :

   #
  ##
 ###
####
Observe that its base and height are both equal to , and the image is drawn using # symbols and spaces. The last line is not preceded by any spaces.

Write a program that prints a staircase of size .

Function Description

Complete the staircase function in the editor below. It should print a staircase as described above.

staircase has the following parameter(s):

n: an integer

*/

import (
	"fmt"
	"strings"
)

// Staircase prints as the descrioption above describes.
func Staircase(n int32) {
	var newN int = int(n)

	for i := 0; i < newN; i++ {
		str := strings.Repeat(" ", (newN-1)-i) + strings.Repeat("#", i+1)
		fmt.Println(str)
	}
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

// 	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
// 	checkError(err)
// 	n := int32(nTemp)

// 	staircase(n)
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
