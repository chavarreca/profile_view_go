package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var tpl *template.Template

type Cuenta struct {
	Usuario string `json:"Usuario"`
	Email   string `json:"Email"`
	Titulo  string `json:"Titulo"`
}

func main() {
	fmt.Println("Iniciar Server")

	tpl, _ = template.ParseFiles("./src/index.html")
	ruta := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", ruta))

	http.HandleFunc("/", infoHandler)
	http.ListenAndServe(":3030", nil)
}

func allowCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	allowCORS(&w)

	var cuenta Cuenta = getInfo()
	fmt.Println("Usuario: ", cuenta.Usuario, "\nEmail: ", cuenta.Email, "\nTitulo: ", cuenta.Titulo)

	err := tpl.Execute(w, cuenta)
	if err != nil {
		fmt.Println("Ha ocurrido un error al enviar la información xC : ", err)
	}
}

func getInfo() Cuenta {
	contenido, err := os.ReadFile("cuenta.json")
	if err != nil {
		fmt.Println("Error al leer el archivo :c", err)
	}

	var cuenta Cuenta

	err = json.Unmarshal(contenido, &cuenta)
	if err != nil {
		fmt.Println("Error al guardar los valores del JSON", err)
	}

	fmt.Println("Se habrio el archivo de forma exitosa!!! :3")

	return cuenta
}
