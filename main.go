package main

import (
	"./parser"
	"./utils"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gorilla/mux"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgsvg"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
)

func funcHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	expr, _ := vars["id"]
	function :=  utils.Calc(vars["id"])
	drawImg(function, expr)




	img, err := os.Open("currImg.jpg")
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, img)
}



func drawImg(function func(float64) float64, expr string ) {


	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = expr
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	gonumFctn := plotter.NewFunction(function)
	fmt.Println("function evaluated in: ", 3, function(3))
	gonumFctn.Color = color.RGBA{B: 255, A: 255}
	p.Add(gonumFctn)
	p.X.Min = -100
	p.X.Max = 100
	p.Y.Min = -100
	p.Y.Max = 100
	const dpi = 96
	
	c := vgsvg.New(3*vg.Inch, 3*vg.Inch)
	p.Draw(draw.New(c))
	// Save the image.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "currImg.jpg"); err != nil {
		panic(err)
	}
}






func main() {
	// Setup the input
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./static/")))
	r.HandleFunc("/{id:.*}", funcHandler)


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