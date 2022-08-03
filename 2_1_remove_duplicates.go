package main


func RemoveDuplicates(n *Node){
	// which numbers we've seen in the list so far
	seen := map[int]struct{}{}

	// as long as there's another item in the list, continue to remove duplicates
	for n.next != nil {
		// record that we saw this item
		seen[n.data] = struct{}{}

		// check if the next item has been seen yet
		if _, ok := seen[n.next.data]; ok {
			// if it has been seen, remove it from the list
			// by linking the current item to the seen item's next
			n.next = n.next.next
		} else {
			// otherwise proceed to the next item in the list
			n = n.next
		}
	}
}

func RemoveDuplicatesWithoutTemporaryBuffer(n *Node){
	// as long as there's another item in the list, continue to remove duplicates
	for n != nil {
		n.removeNumberFromRemainderOfList(n.data)
		n = n.next
	}
}

func (n *Node) removeNumberFromRemainderOfList(num int) {
	curr := n
	for curr.next != nil {
		// check if the next item should be removed
		if curr.next.data == num {
			// remove it from the list by linking the current item to the seen item's next
			curr.next = curr.next.next
		} else {
			// otherwise proceed to the next item in the list
			curr = curr.next
		}
	}
}
