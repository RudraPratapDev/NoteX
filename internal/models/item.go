package models

// Title_ has the trailing underscore so it doesn't shadow the Title() method.
type Item struct {
	Title_, Desc string
}

func (i Item) Title() string       { return i.Title_ }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Title_ }

func NewItem(title, desc string) Item {
	return Item{Title_: title, Desc: desc}
}
