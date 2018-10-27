package backtracking_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jszroberto/go-algorithms/backtracking"
)

var _ = Describe("Ratinmaze", func() {
	It("solves the maze", func() {
		maze := [][]int{
			[]int{1, 0, 0, 0},
			[]int{1, 1, 0, 1},
			[]int{0, 1, 0, 0},
			[]int{1, 1, 1, 1},
		}
		want := [][]int{
			[]int{1, 0, 0, 0},
			[]int{1, 1, 0, 0},
			[]int{0, 1, 0, 0},
			[]int{0, 1, 1, 1},
		}
		got, isResolved := ResolveMaze(0, 0, 3, 3, maze)
		Expect(isResolved).To(BeTrue())
		Expect(got).To(Equal(want))
	})
})
