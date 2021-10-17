package main

import (
	"github.com/tonyliang/coding-interview-note/array"
	"github.com/tonyliang/coding-interview-note/binarysearch"
	"github.com/tonyliang/coding-interview-note/dp"
	h "github.com/tonyliang/coding-interview-note/helper"
	"github.com/tonyliang/coding-interview-note/leetcode"
	"github.com/tonyliang/coding-interview-note/linkedlist"
	"github.com/tonyliang/coding-interview-note/matrix"
	"github.com/tonyliang/coding-interview-note/recursion"
	"github.com/tonyliang/coding-interview-note/stack"
	"github.com/tonyliang/coding-interview-note/strings"
	"github.com/tonyliang/coding-interview-note/tree"
)

func main() {
	runners := []h.Runnable{
		&binarysearch.P154{},
		&binarysearch.P475{},
		&binarysearch.P875{},
		&binarysearch.Offer34{},
		&binarysearch.Offer54{},
		&binarysearch.Offer33{},
		&tree.Offer26{},
		&tree.Offer34{},
		&tree.Offer07{},
		&dp.Offer47{},
		&dp.Offer46{},
		&dp.Offer10{},
		&dp.Offer42{},
		&dp.P343{},
		&strings.Offer05{},
		&strings.Offer48{},
		&strings.Offer58{},
		&strings.Offer67{},
		&linkedlist.P206{},
		&linkedlist.Offer06{},
		&linkedlist.Offer18{},
		&linkedlist.Offer22{},
		&linkedlist.Offer25{},
		&linkedlist.P426{},
		&recursion.Offer12{},
		&recursion.Offer13{},
		&recursion.Offer16{},
		&array.Offer53{},
		&array.X1{},
		&array.Offer40{},
		&array.Offer45{},
		&array.Offer57{},
		&array.Offer61{},
		&array.Offer66{},
		&matrix.Offer29{},
		&stack.P155{},
		&stack.Offer31{},
		&leetcode.Question416{},
		&leetcode.Question417{},
		&leetcode.Question494{},
		&leetcode.Question496{},
		&leetcode.Question500{},
		&leetcode.Question503{},
		&leetcode.Question507{},
		&leetcode.Question508{},
		&leetcode.Question513{},
		&leetcode.Question515{},
		&leetcode.Question518{},
		&leetcode.Question523{},
	}
	for _, r := range runners {
		r.Run()
	}
}
