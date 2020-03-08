package loader

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type stateFn func(*Lexer) stateFn
type itemType int

const (
	ItemDataDefinition itemType = iota
	ItemLineData
	ItemComment
	ItemEOF
	ItemError
)

const (
	eof rune = iota
)

type LexItem struct {
	Type  itemType
	Value string
}

func (i LexItem) String() string {
	switch i.Type {
	case ItemEOF:
		return "EOF"
	case ItemError:
		return i.Value
	}

	if len(i.Value) > 10 {
		return fmt.Sprintf("%d: %.10q...", i.Type, i.Value)
	}
	return fmt.Sprintf("%d: %q", i.Type, i.Value)
}

type Lexer struct {
	input string
	start int
	pos   int
	width int
	items chan LexItem
}

func Lex(input string) *Lexer {
	l := &Lexer{
		input: input,
		items: make(chan LexItem),
	}

	go l.run()
	return l
}

func (l *Lexer) run() {
	for state := lexLine; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func (l *Lexer) emit(t itemType) {
	l.items <- LexItem{
		Type:  t,
		Value: l.input[l.start:l.pos],
	}
	l.start = l.pos
}

func (l *Lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

func (l *Lexer) backup() {
	l.pos -= l.width
}

func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- LexItem{
		Type:  ItemError,
		Value: fmt.Sprintf(format, args),
	}

	return nil
}

func (l *Lexer) nextItem() LexItem {
	return <-l.items
}

func lexLine(l *Lexer) stateFn {
L:
	for {
		switch r := l.next(); {
		case isEndOfLine(r):
			l.ignore()
			return lexLine
		case r == eof:
			break L
		case r == '#':
			return lexComment
		case unicode.IsLetter(r):
			return lexDataDefinition
		}
	}

	l.emit(ItemEOF)
	return nil
}

func lexComment(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof || isEndOfLine(r):
			l.backup()
			l.emit(ItemComment)
			return lexLine
		}
	}
}

func lexDataDefinition(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
		case !unicode.IsLetter(r):
			l.backup()
			l.emit(ItemDataDefinition)
			return lexLineData
		}
	}
}

func lexLineData(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
		case unicode.IsSpace(r) && !isEndOfLine(r):
			l.ignore()
		case r == eof || isEndOfLine(r):
			l.ignore()
			return lexLine
		default:
			l.backup()
			return lexLineDataItem
		}
	}
}

func lexLineDataItem(l *Lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof || unicode.IsSpace(r):
			l.backup()
			l.emit(ItemLineData)
			return lexLineData
		}
	}
}

func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}
