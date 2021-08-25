package main

import (
	"github.com/tonyliang/coding-interview-note/binarysearch"
	"github.com/tonyliang/coding-interview-note/dp"
	h "github.com/tonyliang/coding-interview-note/helper"
	"github.com/tonyliang/coding-interview-note/linkedlist"
	"github.com/tonyliang/coding-interview-note/recursion"
	"github.com/tonyliang/coding-interview-note/strings"
	"github.com/tonyliang/coding-interview-note/tree"
)

func main() {
	runners := []h.Runnable{
		&binarysearch.P154{},
		&binarysearch.P475{},
		&binarysearch.P875{},
		&tree.Offer26{},
		&dp.Offer47{},
		&dp.Offer46{},
		&dp.Offer10{},
		&dp.Offer42{},
		&strings.Offer05{},
		&strings.Offer48{},
		&strings.Offer58{},
		&linkedlist.P206{},
		&linkedlist.Offer06{},
		&linkedlist.Offer18{},
		&linkedlist.Offer22{},
		&linkedlist.Offer25{},
		&recursion.Offer12{},
		&recursion.Offer13{},
	}
	for _, runner := range runners {
		runner.Run()
	}
}
