package bsearch

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func TestSearch(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	idx := search(nums, target)
	fmt.Println("find index=", idx)
	g.Expect(idx).To(gomega.Equal(4))
	nums = []int{-1, 0, 5}
	target = 5
	idx = search(nums, target)
	fmt.Println("find index=", idx)
	//g.Expect(idx).To(gomega.Equal(4))

}
