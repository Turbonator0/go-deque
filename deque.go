package main

import "fmt"

type deque struct {
	content []any
}

// No items by default
func deque_init(v []any) deque {
	d := deque{content: v}
	return d
}

func (d *deque) pop() any { // Pop has to yield the popped value
	yield := d.content[len(d.content)-1]
	d.content = d.content[:len(d.content)-1]
	return yield
}

func (d *deque) push(v any) deque {
	return *d
}

func (d *deque) popleft() any {
	yield := d.content[0]
	d.content = d.content[1:]
	return yield
}

func (d *deque) pushleft(v any) deque {
	return *d
}

func (d *deque) len() deque {
	return *d
}

func (d *deque) count(v any) deque {
	return *d
}

func main() {
	testque := deque_init([]any{1, 2, 3, 4, 5})
	fmt.Println(testque.content...)
	fmt.Println(testque.pop())
	fmt.Println(testque.content...)
	fmt.Println(testque.popleft())
	fmt.Println(testque.content...)
}
