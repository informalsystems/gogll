
// Package lexer is generated by GoGLL. Do not edit.
package lexer

import (
	// "fmt"
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/goccmack/goutil/md"

	"github.com/goccmack/gogll/token"
)

type state int

const nullState state = -1


// Lexer contains both the input slice of runes and the slice of tokens
// parsed from the input
type Lexer struct {
	// I is the input slice of runes
	I      []rune

	// Tokens is the slice of tokens constructed by the lexer from I
	Tokens []*token.Token
}

/*
NewFile constructs a Lexer created from the input file, fname. 

If the input file is a markdown file NewFile process treats all text outside
code blocks as whitespace. All text inside code blocks are treated as input text.

If the input file is a normal text file NewFile treats all text in the inputfile
as input text.
*/
func NewFile(fname string) *Lexer {
	if strings.HasSuffix(fname, ".md") {
		src, err := md.GetSource(fname)
		if err != nil {
			panic(err)
		}
		return New([]rune(src))
	}
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return New([]rune(string(buf)))
}

/*
New constructs a Lexer from a slice of runes. 

All contents of the input slice are treated as input text.
*/
func New(input []rune) *Lexer {
	lex := &Lexer{
		I:      input,
		Tokens: make([]*token.Token, 0, 2048),
	}
	lext := 0
	for lext < len(lex.I) {
		for lext < len(lex.I) && unicode.IsSpace(lex.I[lext]) {
			lext++
		}
		if lext < len(lex.I) {
			tok := lex.scan(lext)
			lext = tok.Rext()
			lex.addToken(tok)
		}
	}
	lex.add(token.EOF, len(input), len(input))
	return lex
}

func (l *Lexer) scan(i int) *token.Token {
	// fmt.Printf("lexer.scan\n")
	s, typ, rext := state(0), token.Error, i
	for s != nullState {
		if rext >= len(l.I) {
			typ = accept[s]
			s = nullState
		} else {
			typ = accept[s]
			s = nextState[s](l.I[rext])
			if s != nullState || typ == token.Error {
				rext++
			}
		}
	}
	return token.New(typ, i, rext,l.I)
}

func escape(r rune) string {
	switch r {
	case '"':
		return "\""
	case '\\':
		return "\\\\"
	case '\r':
		return "\\r"
	case '\n':
		return "\\n"
	case '\t':
		return "\\t"
	}
	return string(r)
}

// GetLineColumn returns the line and column of rune[i] in the input
func (l *Lexer) GetLineColumn(i int) (line, col int) {
	line, col = 1, 1
	for j := 0; j < i; j++ {
		switch l.I[j] {
		case '\n':
			line++
			col = 1
		case '\t':
			col += 4
		default:
			col++
		}
	}
	return
}

// GetLineColumnOfToken returns the line and column of token[i] in the imput
func (l *Lexer) GetLineColumnOfToken(i int) (line, col int) {
	return l.GetLineColumn(l.Tokens[i].Lext())
}

// GetString returns the input string from the left extent of Token[lext] to
// the right extent of Token[rext]
func (l *Lexer) GetString(lext, rext int) string {
	return string(l.I[l.Tokens[lext].Lext():l.Tokens[rext].Rext()])
}

func (l *Lexer) add(t token.Type, lext, rext int) {
	l.addToken(token.New(t, lext, rext, l.I[lext:rext]))
}

func (l *Lexer) addToken(tok *token.Token) {
	l.Tokens = append(l.Tokens, tok)
}

func any(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return true
		}
	}
	return false
}

func not(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return false
		}
	}
	return true
}

var accept = []token.Type{ 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Type0, 
	token.Type1, 
	token.Type2, 
	token.Type3, 
	token.Type4, 
	token.Type5, 
	token.Type6, 
	token.Type7, 
	token.Type8, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Type21, 
	token.Type22, 
	token.Type23, 
	token.Error, 
	token.Error, 
	token.Type18, 
	token.Error, 
	token.Error, 
	token.Error, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type15, 
	token.Type10, 
	token.Type9, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type14, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type11, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type19, 
	token.Type12, 
	token.Type19, 
	token.Type16, 
	token.Type19, 
	token.Type20, 
	token.Type13, 
	token.Type17, 
}

var nextState = []func(r rune) state{ 
	// Set0
	func(r rune) state {
		switch { 
		case r == '"':
			return 1 
		case r == '\'':
			return 2 
		case r == '(':
			return 3 
		case r == ')':
			return 4 
		case r == '.':
			return 5 
		case r == ':':
			return 6 
		case r == ';':
			return 7 
		case r == '<':
			return 8 
		case r == '>':
			return 9 
		case r == '[':
			return 10 
		case r == ']':
			return 11 
		case r == 'a':
			return 12 
		case r == 'e':
			return 13 
		case r == 'l':
			return 14 
		case r == 'n':
			return 15 
		case r == 'p':
			return 16 
		case r == 'u':
			return 17 
		case r == '{':
			return 18 
		case r == '|':
			return 19 
		case r == '}':
			return 20 
		case unicode.IsLower(r):
			return 21 
		case unicode.IsUpper(r):
			return 22 
		}
		return nullState
	}, 
	// Set1
	func(r rune) state {
		switch { 
		case r == '"':
			return 23 
		case r == '\\':
			return 24 
		case not(r, []rune{'"','\\'}):
			return 1 
		}
		return nullState
	}, 
	// Set2
	func(r rune) state {
		switch { 
		case r == '\\':
			return 25 
		case not(r, []rune{'\'','\\'}):
			return 26 
		}
		return nullState
	}, 
	// Set3
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set4
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set5
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set6
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set7
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set8
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set9
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set10
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set11
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set12
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'n':
			return 28 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set13
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'm':
			return 29 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set14
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'e':
			return 30 
		case r == 'o':
			return 31 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set15
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'o':
			return 32 
		case r == 'u':
			return 33 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set16
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'a':
			return 34 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set17
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'p':
			return 35 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set18
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set19
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set20
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set21
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set22
	func(r rune) state {
		switch { 
		case r == '_':
			return 36 
		case unicode.IsLetter(r):
			return 36 
		case unicode.IsNumber(r):
			return 36 
		}
		return nullState
	}, 
	// Set23
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set24
	func(r rune) state {
		switch { 
		case any(r, []rune{'"','\\','n','r','t'}):
			return 1 
		}
		return nullState
	}, 
	// Set25
	func(r rune) state {
		switch { 
		case any(r, []rune{'\'','\\','n','r','t'}):
			return 26 
		}
		return nullState
	}, 
	// Set26
	func(r rune) state {
		switch { 
		case r == '\'':
			return 37 
		}
		return nullState
	}, 
	// Set27
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set28
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'y':
			return 38 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set29
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'p':
			return 39 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set30
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 't':
			return 40 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set31
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'w':
			return 41 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set32
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 't':
			return 42 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set33
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'm':
			return 43 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set34
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'c':
			return 44 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set35
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'c':
			return 45 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set36
	func(r rune) state {
		switch { 
		case r == '_':
			return 36 
		case unicode.IsLetter(r):
			return 36 
		case unicode.IsNumber(r):
			return 36 
		}
		return nullState
	}, 
	// Set37
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set38
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set39
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 't':
			return 46 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set40
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 't':
			return 47 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set41
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'c':
			return 48 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set42
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set43
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'b':
			return 49 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set44
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'k':
			return 50 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set45
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'a':
			return 51 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set46
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'y':
			return 52 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set47
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'e':
			return 53 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set48
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'a':
			return 54 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set49
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'e':
			return 55 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set50
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'a':
			return 56 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set51
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 's':
			return 57 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set52
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set53
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'r':
			return 58 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set54
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 's':
			return 59 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set55
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'r':
			return 60 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set56
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'g':
			return 61 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set57
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'e':
			return 62 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set58
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set59
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'e':
			return 63 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set60
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set61
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case r == 'e':
			return 64 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set62
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set63
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
	// Set64
	func(r rune) state {
		switch { 
		case r == '_':
			return 27 
		case unicode.IsLetter(r):
			return 27 
		case unicode.IsNumber(r):
			return 27 
		}
		return nullState
	}, 
}
