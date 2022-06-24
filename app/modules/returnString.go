package modules

import(
	"fmt"
	"net/http"
)

//função que recebe os retornos de parametos
func ReturnString(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "testing some data as response :P")
}