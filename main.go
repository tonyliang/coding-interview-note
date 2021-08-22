package main

import (
	"github.com/tonyliang/coding-interview-note/binarysearch"
	"github.com/tonyliang/coding-interview-note/dp"
	h "github.com/tonyliang/coding-interview-note/helper"
	"github.com/tonyliang/coding-interview-note/tree"
)

// "github.com/tonyliang/coding-interview-note/array/offer53"
// "github.com/tonyliang/coding-interview-note/array/x1"
// "github.com/tonyliang/coding-interview-note/strings/offer58"
// "github.com/tonyliang/coding-interview-note/linkedlist/p206"
// "github.com/tonyliang/coding-interview-note/linkedlist/offer06"
// "github.com/tonyliang/coding-interview-note/stack/p155"
// "github.com/tonyliang/coding-interview-note/permutation"
// "github.com/tonyliang/coding-interview-note/sliding-window/offer57_2"
// "github.com/tonyliang/coding-interview-note/sliding-window/p438"
// "github.com/tonyliang/coding-interview-note/sliding-window/p3"
// "github.com/tonyliang/coding-interview-note/sliding-window/p239"
// "github.com/tonyliang/coding-interview-note/binary-search/p475"
// "github.com/tonyliang/coding-interview-note/kmp"
// "github.com/tonyliang/coding-interview-note/binary-search/p875"

func main() {
	runners := []h.Runnable{
		&binarysearch.P154{},
		&binarysearch.P475{},
		&binarysearch.P875{},
		&tree.Offer26{},
		&dp.Offer10{},
		&dp.Offer42{},
	}
	for _, runner := range runners {
		runner.Run()
	}

	// offer53.Run()
	// x1.Run()
	// offer58.Run()
	// offer05.Run()
	// p206.Run()
	// offer06.Run()
	// p155.Run()
	// permutation.Run()
	// offer57_2.Run()
	// p438.Run()
	// p3.Run()
	// p239.Run()
	// kmp.Run()
	// p475.Run()
	// p875.Run()
}
