package stk

import "sync"

// Stack the classic `FILO` structrued collection
type Stack struct {
	lock sync.Mutex
	ts   bool
	item []interface{}
}

// NewStack instantiate a new stack, the stack is **threadsafe**
// when and **only** when the argument `ts` is true
func NewStack(ts bool) *Stack {
	return &Stack{ts: ts, item: []interface{}{}}
}

// Push push an item on top of this stack
func (stk *Stack) Push(v interface{}) {
	if stk.ts {
		stk.lock.Lock()
		defer stk.lock.Unlock()
	}

	stk.item = append(stk.item, v)
}

// Pop pop the item on the top, or an nil if this stack is empty
func (stk *Stack) Pop() interface{} {
	if stk.ts {
		stk.lock.Lock()
		defer stk.lock.Unlock()
	}

	l := stk.Size()
	if l == 0 {
		return nil
	}

	top := stk.item[l-1]
	stk.item = stk.item[:l-1]
	return top
}

// Peek peek the item on the top
func (stk *Stack) Peek() interface{} {
	if stk.ts {
		stk.lock.Lock()
		defer stk.lock.Unlock()
	}

	l := stk.Size()
	if l == 0 {
		return nil
	}

	return stk.item[l-1]
}

// Size return the size of this stack
func (stk *Stack) Size() int {
	return len(stk.item)
}
