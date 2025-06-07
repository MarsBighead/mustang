package linked

import (
	"fmt"
	"testing"
)

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
