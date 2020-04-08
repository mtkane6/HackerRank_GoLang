package solutions

import "strings"

/*

When you select a contiguous block of text in a PDF viewer, the selection is highlighted with a blue rectangle. In this PDF viewer, each word is highlighted independently. For example:

PDF-highighting.png

In this challenge, you will be given a list of letter heights in the alphabet and a string. Using the letter heights given, determine the area of the rectangle highlight in  assuming all letters are  wide.

For example, the highlighted . Assume the heights of the letters are  and . The tallest letter is  high and there are  letters. The hightlighted area will be  so the answer is .

Function Description

Complete the designerPdfViewer function in the editor below. It should return an integer representing the size of the highlighted area.

designerPdfViewer has the following parameter(s):

h: an array of integers representing the heights of each letter
word: a string

*/

// DesignerPdrViewer returns the length of the word * height of the tallest letter
func DesignerPdfViewer(h []int32, word string) int32 {
	var maxHeight int32 = 0
	letterMap := make(map[string]int)

	i := 1
	for i <= 26 {
		letterMap[string(toChar(i))] = i - 1
		i++
	}

	for _, v := range word {
		if h[int32(letterMap[strings.ToUpper(string(v))])] > maxHeight {
			maxHeight = h[int32(letterMap[strings.ToUpper(string(v))])]
		}
	}
	return maxHeight * int32(len(word))
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}
