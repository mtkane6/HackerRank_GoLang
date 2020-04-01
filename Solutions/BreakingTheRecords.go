package solutions

/*

Maria plays college basketball and wants to go pro. Each season she maintains a record of her play.
She tabulates the number of times she breaks her season record for most points and least points in a game.
Points scored in the first game establish her record for the season, and she begins counting from there.

For example, assume her scores for the season are represented in the array . Scores are in the same order
as the games played. She would tabulate her results as follows:

                                 Count
Game  Score  Minimum  Maximum   Min Max
 0      12     12       12       0   0
 1      24     12       24       0   1
 2      10     10       24       1   1
 3      24     10       24       1   1
Given Maria's scores for a season, find and print the number of times she breaks her records for most and
least points scored during the season.

Function Description

Complete the breakingRecords function in the editor below. It must return an integer array containing the
numbers of times she broke her records. Index  is for breaking most points records, and index  is for
breaking least points records.

*/

// BreakingRecords returns the number of times her high, and low, scores are changed.
func BreakingRecords(scores []int32) []int32 {
	var high, low, max, min int32 = scores[0], scores[0], 0, 0
	for _, v := range scores {
		if v > high {
			high = v
			max++
		}
		if v < low {
			low = v
			min++
		}
	}

	return []int32{max, min}
}
