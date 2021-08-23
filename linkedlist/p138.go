package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	seen := make(map[*Node]*Node)
	seen[head] = &Node{
		Val: head.Val,
	}
	cur := head
	newHead := seen[head]
	newCur := newHead
	for cur != nil {
		if cur.Random != nil {
			if _, ok := seen[cur.Random]; !ok {
				seen[cur.Random] = &Node{
					Val: cur.Random.Val,
				}
			}
			newCur.Random = seen[cur.Random]
		}
		if cur.Next != nil {
			if _, ok := seen[cur.Next]; !ok {
				seen[cur.Next] = &Node{
					Val: cur.Next.Val,
				}
			}
			newCur.Next = seen[cur.Next]
		}
		cur = cur.Next
		newCur = newCur.Next
	}
	return newHead
}
