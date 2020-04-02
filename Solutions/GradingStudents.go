package solutions

/*

HackerLand University has the following grading policy:

Every student receives a  in the inclusive range from  to .
Any  less than  is a failing grade.
Sam is a professor at the university and likes to round each student's  according to these rules:

If the difference between the  and the next multiple of  is less than , round  up to the next multiple of .
If the value of  is less than , no rounding occurs as the result will still be a failing grade.
For example,  will be rounded to  but  will not be rounded because the rounding would result in a number that is less than .

Given the initial value of  for each of Sam's  students, write code to automate the rounding process.

Function Description

Complete the function gradingStudents in the editor below. It should return an integer array consisting of rounded grades.

gradingStudents has the following parameter(s):

grades: an array of integers representing grades before rounding

*/

// GradingStudents rounds students grades if they deserve it
func GradingStudents(grades []int32) []int32 {
	for i, v := range grades {
		if v < 38 {
			continue
		}
		if v%5 >= 3 {
			grades[i] += (5 - (v % 5))
		}
	}
	return grades
}
