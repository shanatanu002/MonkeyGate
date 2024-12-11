package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"interpreter/Monkey/evaluator"
	"interpreter/Monkey/lexer"
	"interpreter/Monkey/object"
	"interpreter/Monkey/parser"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHTML)
	r.HandleFunc("/run", runREPL).Methods("POST")

	port := getenv("PORT", "8080")
	fmt.Printf("Server is listening on port %s...\n", port)

	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}

func getenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func runREPL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REPL request received")
	var req struct {
		Code string `json:"code"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		fmt.Println("Invalid request:", err)
		return
	}

	env := object.NewEnvironment()
	output, errors := evaluate(req.Code, env)
	if len(errors) != 0 {
		printParserErrors(w, errors)
		return
	}

	io.WriteString(w, output)
}

func evaluate(input string, env *object.Environment) (string, []string) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return "", p.Errors()
	}

	var result string
	for _, stmt := range program.Statements {
		evaluated := evaluator.Eval(stmt, env)
		if evaluated != nil {
			if str, ok := evaluated.(*object.String); ok {
				result += str.Value
			} else if evaluated.Inspect() != "null" {
				result += evaluated.Inspect() + "\n"
			}
		}
	}

	return result, nil
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
