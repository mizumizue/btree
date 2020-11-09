package btree

type BTree struct {
	root *Node
}

func NewBTree() *BTree {
	return &BTree{}
}

func (t *BTree) Add(i int) {
	if t.root == nil {
		t.root = &Node{
			v:      i,
			parent: nil,
			l:      nil,
			r:      nil,
		}
		return
	}

	currentNode := t.root
	for {
		var next *Node
		gt := false

		// 比較して大きければ左、そうでなければ右へ
		if i > currentNode.v {
			next = currentNode.l
			gt = true
		} else {
			next = currentNode.r
		}

		// 次の幹が存在しなければ追加
		if next == nil {
			newItem := &Node{
				parent: currentNode,
				v:      i,
			}
			if gt {
				currentNode.l = newItem
			} else {
				currentNode.r = newItem
			}
			break // 追加完了したら抜ける
		}
		currentNode = next

	}
}

func (t *BTree) Max() int {
	cs := t.root
	count := 0
	for {
		count++
		if !cs.hasLeft() {
			break
		}
		cs = cs.l
	}
	return cs.v
}

func (t *BTree) Min() int {
	cs := t.root
	count := 0
	for {
		count++
		if !cs.hasRight() {
			break
		}
		cs = cs.r
	}
	return cs.v
}

func (t *BTree) Search(target int) bool {
	cs := t.root
	count := 0
	for {
		count++
		if cs.v == target {
			return true
		}
		if target > cs.v && cs.hasLeft() {
			cs = cs.l
			continue
		}
		if target < cs.v && cs.hasRight() {
			cs = cs.r
			continue
		}
		break
	}
	return false
}

// TODO Delete
