package solutions

import "strconv"

/*

There are a number of people who will be attending ACM-ICPC World Finals. Each of them may be well versed in a number of topics. Given a list of topics known by each attendee, you must determine the maximum number of topics a 2-person team can know. Also find out how many ways a team can be formed to know that many topics. Lists will be in the form of bit strings, where each string represents an attendee and each position in that string represents a field of knowledge, 1 if its a known field or 0 if not.

For example, given three attendees' data as follows:

10101
11110
00010
These are all possible teams that can be formed:

Members Subjects
(1,2)   [1,2,3,4,5]
(1,3)   [1,3,4,5]
(2,3)   [1,2,3,4]
In this case, the first team will know all 5 subjects. They are the only team that can be created knowing that many subjects.

Function Description

Complete the acmTeam function in the editor below. It should return an integer array with two elements: the maximum number of topics any team can know and the number of teams that can be formed that know that maximum number of topics.

acmTeam has the following parameter(s):

topic: a string of binary digits

*/

// AcmTeam returns the number of teams that have knowledge in the most areas of the game
func AcmTeam(topic []string) []int32 {
	var knowledgeCount, topTeams, mostKnowledge int = 0, 0, 0

	for i := range topic[:len(topic)-1] {
		for j := range topic[i+1:] {
			knowledgeCount = 0
			for l := range topic[i] {
				thisI, _ := strconv.Atoi(string(topic[i][l]))
				thisJ, _ := strconv.Atoi(string(topic[i+1+j][l]))

				if thisI == 1 || thisJ == 1 {
					knowledgeCount++
					if knowledgeCount > mostKnowledge {
						mostKnowledge = knowledgeCount
						topTeams = 1
					} else if knowledgeCount == mostKnowledge {
						topTeams++
					}
				}
			}
		}
	}

	return []int32{int32(mostKnowledge), int32(topTeams)}
}
