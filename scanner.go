package main

//Librerias que se utilizan en la implementacion
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Tipos de los tokens a leer
const (
	NUMBER  = "NUMBER"
	PLUS    = "PLUS"
	MINUS   = "MINUS"
	STAR    = "STAR"
	SLASH   = "SLASH"
	LPAREN  = "LPAREN"
	RPAREN  = "RPAREN"
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
)

// Estructura de para almacenar la respuesta
type Token struct {
	Type   string
	Lexeme string
}

// Estructura de datos para almacenar el Scanner
type Scanner struct {
	input []rune
	pos   int
}

// Contructor del Scanner
func NewScanner(input string) *Scanner {
	return &Scanner{input: []rune(strings.TrimSpace(input))}
}

/*
Metodo asociado al scanner que lee
el caracter en la posicion actual sin modificar la posicion
*/
func (s *Scanner) peek() rune {
	if s.pos >= len(s.input) {
		return 0
	}
	fmt.Println(s.pos, s.input[s.pos])
	return s.input[s.pos]
}

/*
Metodo asociado al scanner que modifica la posision en la
que se encuentra el scanner
*/
func (s *Scanner) advance() rune {
	ch := s.peek()
	s.pos++
	return ch
}

/*
Metodo asociado al scanner que discrimina los espacion en blanco
*/
func (s *Scanner) skipWhitespace() {
	for unicode.IsSpace(s.peek()) {
		s.advance()
	}
}

/*
Metodo asociado al scanner que
lee los tokens y asigna el tipo
*/
func (s *Scanner) NextToken() Token {
	s.skipWhitespace()
	ch := s.peek()

	if ch == 0 {
		return Token{Type: EOF, Lexeme: ""}
	}

	switch {
	case unicode.IsDigit(ch):
		return s.scanNumber()
	case ch == '+':
		s.advance()
		return Token{Type: PLUS, Lexeme: "+"}
	case ch == '-':
		s.advance()
		return Token{Type: MINUS, Lexeme: "-"}
	case ch == '*':
		s.advance()
		return Token{Type: STAR, Lexeme: "*"}
	case ch == '/':
		s.advance()
		return Token{Type: SLASH, Lexeme: "/"}
	case ch == '(':
		s.advance()
		return Token{Type: LPAREN, Lexeme: "("}
	case ch == ')':
		s.advance()
		return Token{Type: RPAREN, Lexeme: ")"}
	default:
		s.advance()
		return Token{Type: ILLEGAL, Lexeme: string(ch)}
	}
}

/*
Metodo asociado al scanner que lee numeros y considera multidigitos
*/
func (s *Scanner) scanNumber() Token {
	start := s.pos
	for unicode.IsDigit(s.peek()) {
		s.advance()
	}
	return Token{Type: NUMBER, Lexeme: string(s.input[start:s.pos])}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write a math expression:")
	expression, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Input:", expression)

	scanner := NewScanner(expression)

	for {
		token := scanner.NextToken()
		if token.Type == EOF {
			break
		}
		fmt.Printf("Token: %-7s Lexeme: %s\n", token.Type, token.Lexeme)
	}
}
