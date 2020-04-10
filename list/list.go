package list

type node struct {
	next *node
	data interface{}
}

type List struct {
	// Attributes
	fist *node
	last *node
}

func (l *List) IsEmpty() bool {
	return l.fist == nil
}
