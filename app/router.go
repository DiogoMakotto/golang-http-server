package app		//nome do pacote

import ( 		//imports necessários
	"net/http"
	"log"
	"fmt"
)

func Routes(){		//função para declarar endpoints
	http.HandleFunc("/", returnString)
	log.Fatal(http.ListenAndServe(":8089", nil))
}

func returnString(w http.ResponseWriter, r *http.Request){		//função que recebe os retornos de parametos
	fmt.Fprintf(w, "testing some data as response =P")
}