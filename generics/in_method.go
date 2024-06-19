package generics

type stringer interface {
	String() string
}

type MyString string

func (s MyString) String() string {
	return string(s)
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, v := range vals {
		result += v.String()
	}
	return result
}
