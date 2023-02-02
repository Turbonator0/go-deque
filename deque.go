package deque

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
	d.content = append(d.content, v)
	return *d
}

func (d *deque) popleft() any {
	yield := d.content[0]
	d.content = d.content[1:]
	return yield
}

func (d *deque) pushleft(v any) deque {
	d.content = append([]any{v}, d.content...)
	return *d
}

func (d *deque) len() int {
	return len(d.content)
}

func (d *deque) count(v any) int {
	count := 0
	for _, c := range d.content { // super slow but oh well
		if c == v {
			count++
		}
	}
	return count
}
