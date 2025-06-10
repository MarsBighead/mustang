package linked

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	nums1 := []int{4, 1, 8, 4, 5}
	headA := ArrayToList(nums1)
	nums2 := []int{5, 6}
	headB := ArrayToList(nums2)
	right := headB
	for right != nil && right.Next != nil {
		right = right.Next
	}
	fmt.Printf("%#v\n", right)
	i := 0
	common := headA
	for common != nil && right.Next != nil {
		i++
		if i == 2 {
			right.Next = common
			break
		}
		common = common.Next
	}
	right.Next = common
	//fmt.Printf("%#v\n", common)

	node := getIntersectionNode(headA, headB)
	fmt.Printf("%#v\n", node.Val)
	right.Next = headA
	node = getIntersectionNode(headA, headB)
	fmt.Printf("%#v\n", node.Val)

}

func TestSearch(t *testing.T) {
	//g := gomega.NewGomegaWithT(t)
	//ops := []string{"MyLinkedList", "addAtHead", "deleteAtIndex"}
	ops := []string{"MyLinkedList", "addAtTail", "addAtTail"}
	var link MyLinkedList
	for _, op := range ops {
		switch op {
		case "MyLinkedList":
			link = Constructor()
			fmt.Printf("constructor: %#v\n", link)
		case "addAtHead":
			link.AddAtHead(1)
			fmt.Printf("addHead: %#v\n", link)
		case "addAtTail":
			link.AddAtTail(1)
			fmt.Printf("addAtTail: %#v\n", link)
		case "deleteAtIndex":
			link.DeleteAtIndex(0)
			fmt.Printf("deleteAtIndex: %#v\n", link)

		}
	}
}
