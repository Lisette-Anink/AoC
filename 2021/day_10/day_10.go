package day_10

import (
	"fmt"
	"strings"

	"lisette.anink/aoc/utils"
)

// * And, to combine a sequence of terminals and non-terminal parsers.
// * OrdChoice, to choose between specified list of parsers.
// * Kleene, to repeat the parser zero or more times.
// * Many, to repeat the parser one or more times.
// * ManyUntil, to repeat the parser until a specified end matcher.
// * Maybe, to apply the parser once or none.
// func defineParser(ast *parsec.AST) parsec.Parser {
// 	var content parsec.Parser

// 	// terminal parsers.
// 	tagbro := parsec.Atom("(", "OT")
// 	tagbrc := parsec.Atom(")", "OT")
// 	tagsqo := parsec.Atom("[", "CT")
// 	tagsqc := parsec.Atom("]", "CT")
// 	tagcuro := parsec.Atom("{", "CT")
// 	tagcurcl := parsec.Atom("}", "CT")
// 	tagsmt := parsec.Atom("<", "CT")
// 	taggrt := parsec.Atom(">", "CT")

// 	brackets := ast.And(
// 		"brackets", nil, tagbro, &content, tagbrc,
// 	)
// 	// sqbrackets := ast.And(
// 	// 	"sqbrackets", nil, tagsqo, &content, tagsqc,
// 	// )
// 	// cubrackets := ast.And(
// 	// 	"curbrackets", nil, tagcuro, &content, tagcurcl,
// 	// )
// 	// arbrackets := ast.And(
// 	// 	"arbrackets", nil, tagsmt, &content, taggrt,
// 	// )
// 	// content = ast.OrdChoice("content", nil, brackets, sqbrackets, cubrackets, arbrackets)
// 	content = ast.OrdChoice("content", nil, tagsqo, tagsqc, tagcuro, tagcurcl, tagsmt, taggrt, brackets)
// 	contents := ast.Maybe(
// 		"maybecontents", nil, ast.Kleene("contents", nil, content, nil),
// 	)
// 	return ast.Kleene("c", nil, contents, nil)
// }

// func parseLine(line string) {
// 	ast := parsec.NewAST("line", 100)
// 	y := defineParser(ast)
// 	s := parsec.NewScanner([]byte(line))
// 	node, _ := ast.Parsewith(y, s)
// 	fmt.Println(node.GetValue())
// 	ast.Prettyprint()
// 	graph := ast.Dotstring("simplehtml")
// 	fmt.Println(graph)
// 	ch := make(chan parsec.Queryable, 100)
// 	ast.Query(".term", ch)
// 	for node := range ch {
// 		fmt.Printf("%s", node.GetValue())
// 	}
// 	fmt.Println()
// }
func mapFromString(in string) map[int]string {
	elements := strings.Split(in, "")
	mapElements := map[int]string{}
	for i, e := range elements {
		mapElements[i] = e
	}
	return mapElements
}
func dumbWay(lines []string) []string {
	list := []string{}
	for _, line := range lines {
		if len(line) > 0 {
			elements := mapFromString(line)
			if ok, wrong := findCorrect(elements, 0, 1); !ok {
				fmt.Println("wrong", wrong, line)
				list = append(list, wrong)
			} else {
				fmt.Println("right", line)

			}

		}
	}
	return list
}

// if open
// look for closed
// if first is right closed !top
// als einde van map niet is bereikt open volgnede
// elsif first is open dan
// look for closed
// elsif first is wrong
// return false
func findCorrect(elements map[int]string, current, close int) (bool, string) {
	fmt.Println("start", current, close, elements[current])
	if isOpen(elements[current]) {
		// fmt.Println("isOpen", current, close)
		if isClose(elements[current], elements[close]) {
			fmt.Println("isClose", current, close)
			return true, elements[current]
		} else if isOpen(elements[close]) {
			fmt.Println("isNewOpen", close, elements[close])
			return findCorrect(elements, close, close+1)
		} else {
			fmt.Println("isWrong", close, elements[close])
			return false, elements[current]
		}
	}
	// fmt.Println("!isOpen", current, close)
	return false, elements[current]
}

var allOpen = []string{"(", "<", "{", "["}

func isOpen(current string) bool {
	return utils.IncludesAll(allOpen, []string{current})
}

func isClose(open, close string) bool {
	switch open {
	case "(":
		return close == ")"
	case "<":
		return close == ">"
	case "{":
		return close == "}"
	case "[":
		return close == "]"
	default:
		return false
	}
}
