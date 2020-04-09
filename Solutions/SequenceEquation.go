package solutions

/*

Given a sequence of  integers,  where each element is distinct and satisfies . For each  where , find any integer  such that  and print the value of  on a new line.

For example, assume the sequence . Each value of  between  and , the length of the sequence, is analyzed as follows:

, so
, so
, so
, so
, so
The values for  are .

Function Description

Complete the permutationEquation function in the editor below. It should return an array of integers that represent the values of .

permutationEquation has the following parameter(s):

p: an array of integers

*/

// PermutationEquation finds the index of the nth item of p, within p
func PermutationEquation(p []int32) []int32 {
	valueMap := make(map[int32]int32)
	var result []int32

	for i, v := range p {
		valueMap[v] = int32(i + 1)
	}

	var i int32 = 1
	for i <= int32(len(p)) {
		result = append(result, valueMap[valueMap[i]])
		i++
	}
	return result
}
