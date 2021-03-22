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
var espacio []listD

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
	inventario archivos.AVL

}
type arbol struct {
Invetarios[] Inventarios //todo si no jala inventario cambiarlo a como den el nuevo json
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
//esta estructura es la mando al front
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
	router.HandleFunc("/cargarinventario", cargarInventario).Methods("POST")
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

func cargarInventario(w http.ResponseWriter, r *http.Request)  {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &datosInventario)
	if err != nil {
		log.Fatal("Error")
	}

	recorreInventario()

	avl.Print()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func recorreInventario()  {
	espacio = make([]listD, len(indices.Datos[0].Departamentos)*len(indices.Datos)*5)
	for i := 0; i < len(indices.Datos); i++ { //letra
		for j := 0; j < len(indices.Datos[i].Departamentos); j++ { //departamento
			for k := 0; k < len(indices.Datos[i].Departamentos[j].Tiendas); k++ { //tienda
				Calificacion := indices.Datos[i].Departamentos[j].Tiendas[k].Calificacion - 1
				colocacion := Calificacion + 5*(j+len(indices.Datos[i].Departamentos)*i)
				newNode := nodeD{Indice: indices.Datos[i].Indice, Departamentos: indices.Datos[i].Departamentos[j].Nombre, Nombre: indices.Datos[i].Departamentos[j].Tiendas[k].Nombre, Descripcion: indices.Datos[i].Departamentos[j].Tiendas[k].Descripcion, Contacto: indices.Datos[i].Departamentos[j].Tiendas[k].Contacto, Calificacion: indices.Datos[i].Departamentos[j].Tiendas[k].Calificacion}
				espacio[colocacion].Insert(&newNode)
			}
		}
	}
	for nodo = 0; nodo < len(espacio); nodo++ {
		espacio[nodo].Showtienda()
	}


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

type nodeD struct {
	next          *nodeD
	previous      *nodeD
	Indice        string
	Departamentos string
	Nombre        string
	Descripcion   string
	Contacto      string
	Calificacion  int
	inventario archivos.AVL
}

type listD struct {
	first *nodeD
	last  *nodeD
}

func NewList() *listD {
	return &listD{nil, nil}
}

func (elist_d *listD) Insert(Nodo *nodeD) {

	if elist_d.first == nil {
		elist_d.last = Nodo
		elist_d.first = elist_d.last
	} else {
		Nodo.previous = elist_d.last
		elist_d.last.next = Nodo
		elist_d.last = Nodo
	}

}
func (elist_d *listD)Showtienda() {
	auxiliar := elist_d.first
	for auxiliar != nil {
		for i := 0; i < len(datosInventario.Invetarios); i++ {
			if (datosInventario.Invetarios[i].Tienda == auxiliar.Nombre && datosInventario.Invetarios[i].Departamento == auxiliar.Departamentos && datosInventario.Invetarios[i].Calificacion == auxiliar.Calificacion) {
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
						auxiliar.inventario.Insertar(productos)

					}
				}
			}
			fmt.Println(auxiliar.Nombre)
			auxiliar.inventario.Print()

			auxiliar = auxiliar.next

		}
	}
