package update

type Field struct {
	Name   string
	Parser func(u *Message, val string)
}
