package repl

import (
	"Monkey/evaluator"
	"Monkey/lexer"
	"Monkey/object"
	"Monkey/parser"
	"fmt"
	"io"
	"os"
	"strings"
)

const PROMPT = ">> "

func Start(filePath string, out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(out, "Could not open file: %s\n", err)
		return
	}
	defer file.Close()

	// Read the entire file content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(out, "Error reading file: %s\n", err)
		return
	}

	env := object.NewEnvironment()

	// Split the content into statements
	statements := strings.Split(string(content), "\n\n")

	for _, statement := range statements {
		if strings.TrimSpace(statement) == "" {
			continue
		}

		l := lexer.New(statement)
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
