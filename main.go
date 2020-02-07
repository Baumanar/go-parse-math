package main

import (
	"fmt"
	"github.com/Baumanar/go-parse-math/parser"
	"github.com/Baumanar/go-parse-math/utils"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

func funcHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	num, _ := strconv.ParseFloat(vars["num"], 64)
	fmt.Fprintf(w, "Function: %v\n evaluated on: %g is: %f", vars["id"], num, utils.Calc(vars["id"])(num) )
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}



func main() {
	// Setup the input
	r := mux.NewRouter()
	r.HandleFunc("/{id:.*}/{num:[0-9]+}", funcHandler)


	line := "ln(x)"
	is := antlr.NewInputStream(line)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}

	fmt.Printf("The answer is: %f\n", utils.Calc(line)(9))

	http.ListenAndServe(":8000", r)


}