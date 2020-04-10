package solutions

/*

Your local library needs your help! Given the expected and actual return dates for a library book, create a program that calculates the fine (if any). The fee structure is as follows:

If the book is returned on or before the expected return date, no fine will be charged (i.e.: .
If the book is returned after the expected return day but still within the same calendar month and year as the expected return date, .
If the book is returned after the expected return month but still within the same calendar year as the expected return date, the .
If the book is returned after the calendar year in which it was expected, there is a fixed fine of .
Charges are based only on the least precise measure of lateness. For example, whether a book is due January 1, 2017 or December 31, 2017, if it is returned January 1, 2018, that is a year late and the fine would be .

Function Description

Complete the libraryFine function in the editor below. It must return an integer representing the fine due.

libraryFine has the following parameter(s):

d1, m1, y1: returned date day, month and year
d2, m2, y2: due date day, month and year

*/

// LibraryFine returns the total fines due to lateness
func LibraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	if y1 < y2 {
		return int32(0)
	}
	if y1 > y2 {
		return int32(10000)
	}
	if m1 < m2 {
		return int32(0)
	}
	if m1 > m2 {
		return int32(500 * (m1 - m2))
	}
	if d1 > d2 {
		return int32(15 * (d1 - d2))
	}
	return int32(0)
}
