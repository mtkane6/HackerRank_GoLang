package solutions

/*

Lisa just got a new math workbook. A workbook contains exercise problems, grouped into chapters. Lisa believes a problem to be special if its index (within a chapter) is the same as the page number where it's located. The format of Lisa's book is as follows:

There are  chapters in Lisa's workbook, numbered from  to .
The  chapter has  problems, numbered from  to .
Each page can hold up to  problems. Only a chapter's last page of exercises may contain fewer than  problems.
Each new chapter starts on a new page, so a page will never contain problems from more than one chapter.
The page number indexing starts at .
Given the details for Lisa's workbook, can you count its number of special problems?

For example, Lisa's workbook contains  problems for chapter , and  problems for chapter . Each page can hold  problems. The first page will hold  problems for chapter . Problem  is on page , so it is special. Page  contains only Chapter , Problem , so no special problem is on page . Chapter  problems start on page  and there are  problems. Since there is no problem  on page , there is no special problem on that page either. There is  special problem in her workbook.

Note: See the diagram in the Explanation section for more details.

Function Description

Complete the workbook function in the editor below. It should return an integer that represents the number of special problems in the workbook.

workbook has the following parameter(s):

n: an integer that denotes the number of chapters
k: an integer that denotes the maximum number of problems per page
arr: an array of integers that denote the number of problems in each chapter

*/

// Workbook returns the number of times the problem # == page #
func Workbook(n int32, k int32, arr []int32) int32 {
	var prob, currentPage, problems, special int32
	currentPage = 1

	for _, v := range arr {
		problems = v
		for prob = 1; prob <= problems; prob++ {
			if prob != 1 && prob%k == 1 {
				currentPage++
			}
			if prob == currentPage {
				special++
				if prob == problems {
					currentPage++
				}
			} else {
				if prob == problems {
					currentPage++
				}
			}
		}
	}
	return special
}
