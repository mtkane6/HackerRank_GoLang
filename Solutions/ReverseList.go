package solutions

/*
You’re given the pointer to the head node of a linked list. Change the next pointers of the nodes so that their order is reversed. The head pointer given may be null meaning that the initial list is empty.

Input Format

You have to complete the SinglyLinkedListNode reverse(SinglyLinkedListNode head) method which takes one argument - the head of the linked list. You should NOT read any input from stdin/console.

The input is handled by the code in the editor and the format is as follows:

The first line contains an integer , denoting the number of test cases.
Each test case is of the following format:

The first line contains an integer , denoting the number of elements in the linked list.
The next  lines contain an integer each, denoting the elements of the linked list.
*/

// Reverse the connections of all links
func Reverse(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	var prev, next *SinglyLinkedListNode
	curr := head

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}
