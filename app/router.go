package app		//nome do pacote

import ( 		//imports necessários
	"net/http"
	"log"
	"golang-http-server/app/modules"
)

func Routes(){		//função para declarar endpoints
	http.HandleFunc("/", modules.ReturnString)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
