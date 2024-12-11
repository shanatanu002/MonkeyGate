package repl

import (
	"Monkey/evaluator"
	"Monkey/lexer"
	"Monkey/object"
	"Monkey/parser"
	"bufio"
	"fmt"
	"io"
	"strings"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	var inputLines []string

	for {
		if len(inputLines) == 0 {
			fmt.Printf("%s", PROMPT)
		} else {
			fmt.Printf("%s", ".. ")
		}

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if strings.HasSuffix(line, "\\") {
			// remove the trailing backslash and store the line
			inputLines = append(inputLines, strings.TrimSuffix(line, "\\"))
			continue
		}

		// add the final line and join all lines together
		inputLines = append(inputLines, line)
		fullInput := strings.Join(inputLines, "\n")
		inputLines = inputLines[:0] // clear the input lines

		l := lexer.New(fullInput)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil && evaluated.Inspect() != "null" {
			io.WriteString(out, "\n")
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MONKEY_FACE = `
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some trouble here!\n\n")
	io.WriteString(out, "👾👾👾👾👾👾👾👾👾👾👾\n\n")
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
