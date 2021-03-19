package main

import (
	"./archivos"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gorilla/handlers"
)

var path = "grafo.dot"
var file, err = os.OpenFile(path, os.O_RDWR, 0644)
var path1 = "documento.json"
var file1, err1 = os.OpenFile(path1, os.O_RDWR, 0644)
var tiendas_grafo, conexiones_grafo, datos_tiendas string
var nodo, valor int

type indice struct {
	Datos []indicess
}
type indicess struct {
	Indice        string
	Departamentos []departamentos
}
type departamentos struct {
	Nombre  string
	Tiendas []tienda
}
type tienda struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
	Logo 		string
}
type arbol struct {
Invetarios[] Inventarios
}
type Inventarios struct {
	Tienda string
	Departamento string
	Calificacion int
	Productos [] productos
	}
type productos struct {
	Nombre string
	Codigo int
	Descripcion string
	Precio float64
	Cantidad int
	Imagen string
}
//esta estructura es la mando al fron
type tiendas struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
	Logo 		string
}

var indices indice
var datosInventario arbol
var datosTiendas tiendas
var avl = archivos.NewAVL()
func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	router.HandleFunc("/mostrartiendas", mostrartiendas).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", TiendaEspecifica).Methods("POST")
	router.HandleFunc("/id/{numero}", busquedaporPosicion).Methods("GET")
	router.HandleFunc("/arbol", cargararbol).Methods("POST")
	//log.Fatal(http.ListenAndServe(":3000", router))
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor Funcionando :D")
}

func cargartienda(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &indices)
	if err != nil {
		log.Fatal("Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(indices)
}


func busquedaporPosicion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["numero"])
	if err != nil {
		return
	}

	valor = taskID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datos_tiendas)

}

func getArreglo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("<img src=\"imagengrafo.jpg\">")

}
func TiendaEspecifica(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("<img src=\"dinosaur.jpg\">")

}

func cargararbol(w http.ResponseWriter, r *http.Request)  {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &datosInventario)
	if err != nil {
		log.Fatal("Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(datosInventario)
	recorreInventario()
}

func recorreInventario()  {
	for i:=0; i<len(datosInventario.Invetarios);i ++{
		for j:=0; j<len(datosInventario.Invetarios[i].Productos);j ++{
			var productosInventario=datosInventario.Invetarios[i].Productos[j]
			productos:=archivos.Producto{
				Nombre: productosInventario.Nombre,
				Cantidad: productosInventario.Cantidad,
				Descripcion: productosInventario.Descripcion,
				Precio: productosInventario.Precio,
				Codigo: productosInventario.Codigo,
				Imagen: productosInventario.Imagen,
			}

			avl.Insertar(productos)

		}
	}
	fmt.Println("Valores de AVL")
	avl.Print()
}

func mostrartiendas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var lista_tiendas []tiendas
	for i := 0; i < len(indices.Datos); i++ { //letra
		for j := 0; j < len(indices.Datos[i].Departamentos); j++ { //departamento
			for k := 0; k < len(indices.Datos[i].Departamentos[j].Tiendas); k++ {
				var tiendas_json=indices.Datos[i].Departamentos[j].Tiendas[k]
				datosTiendas=tiendas{
					Nombre: tiendas_json.Nombre,
					Descripcion: tiendas_json.Descripcion,
					Contacto: tiendas_json.Contacto,
					Calificacion: tiendas_json.Calificacion,
					Logo: tiendas_json.Logo,
				}
				lista_tiendas=append(lista_tiendas,datosTiendas)

			}
		}
	}
	json.NewEncoder(w).Encode(lista_tiendas)


}


