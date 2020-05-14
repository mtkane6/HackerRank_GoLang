package solutions

/*
You’re given the pointer to the head node of a linked list and the position of a node to delete. Delete the node at the given position and return the head node. A position of 0 indicates head, a position of 1 indicates one node away from the head and so on. The list may become empty after you delete the node.

Input Format

You have to complete the deleteNode(SinglyLinkedListNode* llist, int position) method which takes two arguments - the head of the linked list and the position of the node to delete. You should NOT read any input from stdin/console. position will always be at least 0 and less than the number of the elements in the list.

The first line of input contains an integer , denoting the number of elements in the linked list.
The next  lines contain an integer each in a new line, denoting the elements of the linked list in the order.
The last line contains an integer  denoting the position of the node that has to be deleted form the linked list.
*/

// DeleteNode removes a link node at a specified position
func DeleteNode(head *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	temp := head
	if position == 0 {
		temp = head.next
		return temp
	}

	for position != 1 && head != nil {
		head = head.next
		position--
	}

	if head != nil {
		head.next = head.next.next
	}

	return temp
}
