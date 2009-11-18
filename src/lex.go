package golambda

import(
	"utf8";
)

var line string;
var linep int;
var lookahead int;

func getrune() (result int) {
	var n int;

	if linep >= len(line) {
		return 0;
	}
	result,n = utf8.DecodeRuneInString(line[linep:len(line)]);
	linep += n;
	if result == '\n' {
		result = 0;
	}
	return result;
}

func Lex() int {
	c := lookahead;
	lookahead = getrune();
	switch true {
	case c == ' ' || c == '\t':
		return Lex();
	case c == 'f' && lookahead == 'n':
		lookahead = getrune();
		return FN;
	case isIdent(c):
		str := string(c);
		for isIdent(lookahead) {
			str += string(lookahead);
			lookahead = getrune();
		}
		yylval.ident = str;
		return IDENT;
	}
	return c;
}

func isIdent(c int) bool {
	return
		('a' <= c && c <= 'z') ||
		('A' <= c && c <= 'Z') ||
		('0' <= c && c <= '9') ||
		c == '_';
}
