// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"

	"github.com/mikan/gopl/ch07/ex16/eval"
)

var ui = template.Must(template.New("calc").Parse(`
<html><head><title>Calc</title><style>
body { background-color: whitesmoke; }
#result, #error { font-size: 200%; }
#error { font-color: red; }
</style></head><body>
<form method="get" action="/">
<label for="expr">Expr: </label>
<input type="text" name="expr" id="expr" size="50" />
<input type="submit" value=" = ">
<p id="result">{{.Result}}</p>
<p id="error">{{.Error}}</p>
</form>
</body></html>
`))

type CalcResult struct {
	Result string
	Error  string
}

func main() {
	http.HandleFunc("/", calc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calc(w http.ResponseWriter, req *http.Request) {
	input := req.URL.Query().Get("expr")
	if input == "" {
		var calcResult CalcResult
		if err := ui.Execute(w, calcResult); err != nil {
			log.Fatal(err)
		}
		return
	}
	expr, err := eval.Parse(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		var calcResult CalcResult
		calcResult.Error = fmt.Sprintf("%v", err) // parse error
		if err := ui.Execute(w, calcResult); err != nil {
			log.Fatal(err)
		}
		return
	}
	env := eval.Env{"pi": math.Pi} // limitation: env customization is not supported
	got := fmt.Sprintf("%.6g", expr.Eval(env))
	var calcResult CalcResult
	calcResult.Result = got
	if err := ui.Execute(w, calcResult); err != nil {
		log.Fatal(err)
	}
}
