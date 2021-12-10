package day_10

import (
	"fmt"
	"strings"

	parsec "github.com/prataprc/goparsec"
)

// * And, to combine a sequence of terminals and non-terminal parsers.
// * OrdChoice, to choose between specified list of parsers.
// * Kleene, to repeat the parser zero or more times.
// * Many, to repeat the parser one or more times.
// * ManyUntil, to repeat the parser until a specified end matcher.
// * Maybe, to apply the parser once or none.
func defineParser(ast *parsec.AST) parsec.Parser {
	var content parsec.Parser

	// terminal parsers.
	tagbro := parsec.Atom("(", "OT")
	tagbrc := parsec.Atom(")", "OT")
	tagsqo := parsec.Atom("[", "CT")
	tagsqc := parsec.Atom("]", "CT")
	tagcuro := parsec.Atom("{", "CT")
	tagcurcl := parsec.Atom("}", "CT")
	tagsmt := parsec.Atom("<", "CT")
	taggrt := parsec.Atom(">", "CT")

	brackets := ast.And(
		"brackets", nil, tagbro, &content, tagbrc,
	)
	// sqbrackets := ast.And(
	// 	"sqbrackets", nil, tagsqo, &content, tagsqc,
	// )
	// cubrackets := ast.And(
	// 	"curbrackets", nil, tagcuro, &content, tagcurcl,
	// )
	// arbrackets := ast.And(
	// 	"arbrackets", nil, tagsmt, &content, taggrt,
	// )
	// content = ast.OrdChoice("content", nil, brackets, sqbrackets, cubrackets, arbrackets)
	content = ast.OrdChoice("content", nil, tagsqo, tagsqc, tagcuro, tagcurcl, tagsmt, taggrt, brackets)
	contents := ast.Maybe(
		"maybecontents", nil, ast.Kleene("contents", nil, content, nil),
	)
	return ast.Kleene("c", nil, contents, nil)
}

func parseLine(line string) {
	ast := parsec.NewAST("line", 100)
	y := defineParser(ast)
	s := parsec.NewScanner([]byte(line))
	node, _ := ast.Parsewith(y, s)
	fmt.Println(node.GetValue())
	ast.Prettyprint()
	graph := ast.Dotstring("simplehtml")
	fmt.Println(graph)
	ch := make(chan parsec.Queryable, 100)
	ast.Query(".term", ch)
	for node := range ch {
		fmt.Printf("%s", node.GetValue())
	}
	fmt.Println()

}

func dumbWay(lines []string) {
	for _, line := range lines {
		if len(line) > 0 {
			findCorrect(line)
		}
	}
}

func findCorrect(line string) {
	elements := strings.Split(line, "")
	for _, e := range elements {
		switch e {
		case "(":
		case "<":
		case "{":
		case "[":

		}
	}
}
