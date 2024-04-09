package gotextenc

type CharLike interface {
	byte | uint16 | rune
}
