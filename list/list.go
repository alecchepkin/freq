package list

type List interface {
	Insert(s string) Item
	Frequency(n uint) []Item
}

type Item struct {
	Name  string
	Count int
}
