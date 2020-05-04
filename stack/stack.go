package stack

type Stack []interface{}

func Stack() *Stack {
	s := Stack(make([]interface{}, 0))
	return &s
}
func (s *Stack) Push(i interface{}) {
	*s = append(*s, i)
}
func (s *Stack) Pop() interface{} {
	lastIndex := len(*s) - 1
	if lastIndex == -1 {
		return nil
	}
	i := (*s)[lastIndex]
	*s = (*s)[0:lastIndex]
	return i
}
func (s *Stack) Top() interface{} {
	lastIndex := len(*s) - 1
	if lastIndex == -1 {
		return nil
	}
	return (*s)[lastIndex]
}
