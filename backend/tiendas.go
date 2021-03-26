package main

import (
	"./archivos"
	"./matriz"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
	Logo         string
	inventario   archivos.AVL
}

type arbol struct {
	Inventarios []Inventarios //todo si no jala inventario cambiarlo a como den el nuevo json
}
type Inventarios struct {
	Tienda       string
	Departamento string
	Calificacion int
	Productos    []productos
}
type productos struct {
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float64
	Cantidad    int
	Imagen      string
}

//esta estructura es la mando al front
type tiendas struct {
	Id           int
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
	Logo         string
}

type pedido struct {
	Pedidos []contenido
}
type contenido struct {
	Fecha        string
	Tienda       string
	Departamento string
	Calificacion int
	Productos    []producto
}
type producto struct {
	Codigo int
}

var pedidos pedido
var indices indice
var datosInventario arbol
var datosTiendas tiendas
var avl = archivos.NewAVL()

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	router.HandleFunc("/mostrartiendas", mostrartiendas).Methods("GET")
	router.HandleFunc("/cargarinventario", cargarInventario).Methods("POST")
	router.HandleFunc("/mostrarinventario/{numero}", mostrarinventario).Methods("GET")
	router.HandleFunc("/cargarpedido", cargarpedido).Methods("POST")
	router.HandleFunc("/mostrarpedido/{numero}", mostrarpedido).Methods("GET")
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

func cargarInventario(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &datosInventario)
	if err != nil {
		log.Fatal("Error")
	}
	fmt.Print("inventarioCargado")
	recorreInventario()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(datosInventario)
}

func recorreInventario() {
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

var lista_tiendas []tiendas

func mostrartiendas(w http.ResponseWriter, r *http.Request) {
	var id = 0
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var lista_tiendas1 []tiendas
	for i := 0; i < len(indices.Datos); i++ { //letra
		for j := 0; j < len(indices.Datos[i].Departamentos); j++ { //departamento
			for k := 0; k < len(indices.Datos[i].Departamentos[j].Tiendas); k++ {
				var tiendas_json = indices.Datos[i].Departamentos[j].Tiendas[k]

				datosTiendas = tiendas{
					Id:           id,
					Nombre:       tiendas_json.Nombre,
					Descripcion:  tiendas_json.Descripcion,
					Contacto:     tiendas_json.Contacto,
					Calificacion: tiendas_json.Calificacion,
					Logo:         tiendas_json.Logo,
				}
				id = id + 1
				lista_tiendas = append(lista_tiendas, datosTiendas)
				lista_tiendas1 = append(lista_tiendas1, datosTiendas)
			}
		}
	}

	json.NewEncoder(w).Encode(lista_tiendas1)
}

func mostrarinventario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["numero"])
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var lista_productos []productosInventario
	var lista_temporal []productotemporales
	for nodoo := 0; nodoo < len(archivos.List); nodoo++ {
		for i := 0; i < len(datosInventario.Inventarios); i++ {
			for j := 0; j < len(datosInventario.Inventarios[i].Productos); j++ {
				if archivos.List[nodoo] == datosInventario.Inventarios[i].Productos[j].Nombre {
					datostiendasInventario = productosInventario{
						Tienda:      datosInventario.Inventarios[i].Tienda,
						Nombre:      datosInventario.Inventarios[i].Productos[j].Nombre,
						Codigo:      datosInventario.Inventarios[i].Productos[j].Codigo,
						Descripcion: datosInventario.Inventarios[i].Productos[j].Descripcion,
						Precio:      datosInventario.Inventarios[i].Productos[j].Precio,
						Cantidad:    datosInventario.Inventarios[i].Productos[j].Cantidad,
						Imagen:      datosInventario.Inventarios[i].Productos[j].Imagen,
					}
					lista_productos = append(lista_productos, datostiendasInventario)
				}
			}
		}
	}
	for nodoo := 0; nodoo < len(archivos.List); nodoo++ {
		if lista_tiendas[taskID].Nombre == lista_productos[nodoo].Tienda {
			productosTemporales = productotemporales{
				Nombre:      lista_productos[nodoo].Nombre,
				Codigo:      lista_productos[nodoo].Codigo,
				Descripcion: lista_productos[nodoo].Descripcion,
				Precio:      lista_productos[nodoo].Precio,
				Cantidad:    lista_productos[nodoo].Cantidad,
				Imagen:      lista_productos[nodoo].Imagen,
			}
			lista_temporal = append(lista_temporal, productosTemporales)
		}
	}
	json.NewEncoder(w).Encode(lista_temporal)
}

var productosTemporales productotemporales

type productotemporales struct {
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float64
	Cantidad    int
	Imagen      string
}

var datostiendasInventario productosInventario

type productosInventario struct {
	Tienda      string
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float64
	Cantidad    int
	Imagen      string
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
	inventario    archivos.AVL
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
func (elist_d *listD) Showtienda() {
	auxiliar := elist_d.first
	for auxiliar != nil {
		for i := 0; i < len(datosInventario.Inventarios); i++ {
			if datosInventario.Inventarios[i].Tienda == auxiliar.Nombre && datosInventario.Inventarios[i].Departamento == auxiliar.Departamentos && datosInventario.Inventarios[i].Calificacion == auxiliar.Calificacion {
				for j := 0; j < len(datosInventario.Inventarios[i].Productos); j++ {
					var productosInventario = datosInventario.Inventarios[i].Productos[j]
					productos := archivos.Producto{
						Nombre:      productosInventario.Nombre,
						Cantidad:    productosInventario.Cantidad,
						Descripcion: productosInventario.Descripcion,
						Precio:      productosInventario.Precio,
						Codigo:      productosInventario.Codigo,
						Imagen:      productosInventario.Imagen,
					}
					auxiliar.inventario.Insertar(productos)

				}
			}
		}
		//fmt.Println(auxiliar.Nombre)
		auxiliar.inventario.Print()
		auxiliar = auxiliar.next
	}
}

var mat = matriz.Matrix{}

func cargarpedido(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &pedidos)
	if err != nil {
		log.Fatal("Error")
	}

}
func mostrarpedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["numero"])
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var strdate string
	var mes int
	mat.Init()
	for i := 0; i < len(pedidos.Pedidos); i++ {
		strdate = pedidos.Pedidos[i].Fecha
		day, _ := strconv.Atoi(strdate[:2])
		mes, _ = strconv.Atoi(strdate[3:5])
		if mes == taskID {
			for j := 0; j < len(pedidos.Pedidos[i].Productos); j++ {
				if len(pedidos.Pedidos[i].Productos) == 1 {
					mat.Add(i+1, day, pedidos.Pedidos[i].Productos[j].Codigo)
				}
				if len(pedidos.Pedidos[i].Productos) > 1 {
					mat.Add(i+1, day, pedidos.Pedidos[i].Productos[j].Codigo)
				}

			}
		} else {
			fmt.Println("No hay pedidos en este mes")
		}
	}

	if mes == taskID {
		mat.Show()
		crearArchivo()
		escribeArchivo()
	}
}

func crearArchivo() {
	var _, err = os.Create(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
}
func escribeArchivo() {
	if existeError(err) {
		return
	}
	defer file.Close()
	_, err = file.WriteString("digraph grafica{ \n node [shape=box]\n")
	_, err = file.WriteString(" Mt[ label = \"Matriz\" group = 1 ];\n")
	_, err = file.WriteString("e0[ shape = point, width = 0 ];\n e1[ shape = point, width = 0 ]; \n ")
	for i := 0; i < len(matriz.ListTienda); i++ {
		t := strconv.Itoa(matriz.ListTienda[i] - 1)
		_, err = file.WriteString("nodo" + t + "[label = \"" + pedidos.Pedidos[matriz.ListTienda[i]-1].Departamento + "\"    group = 1 ];\n")
	}
	for i := 0; i < len(matriz.ListTienda); i++ {
		t := strconv.Itoa(matriz.ListDia[i])
		_, err = file.WriteString("D" + t + "[label = \"" + t + "\"    group = " + t + " ];\n")
	}
	for i := 0; i < len(matriz.ListTienda); i++ {
		t := strconv.Itoa(matriz.ListCodigo[i])
		h := strconv.Itoa(matriz.ListDia[i])
		_, err = file.WriteString("C" + t + "[label = \"" + t + "\"    group = " + h + " ];\n")
	}
	for i := 0; i < len(matriz.ListTienda); i++ {
		h := strconv.Itoa(matriz.ListDia[i])
		t := strconv.Itoa(matriz.ListTienda[i] - 1)
		g := strconv.Itoa(matriz.ListCodigo[i])
		_, err = file.WriteString("nodo" + t + " ->C" + g + ";\n")
		_, err = file.WriteString("C" + g + " ->nodo" + t + ";\n")
		_, err = file.WriteString("D" + h + " ->C" + g + ";\n")
		_, err = file.WriteString("C" + g + " ->D" + h + ";\n")

	}
	_, err = file.WriteString("{ rank = same; Mt;")
	for i := 0; i < len(matriz.ListTienda); i++ {
		h := strconv.Itoa(matriz.ListDia[i])
		_, err = file.WriteString("D" + h + ";")
	}
	_, err = file.WriteString("}\n")

	for i := 0; i < len(matriz.ListTienda); i++ {
		t := strconv.Itoa(matriz.ListTienda[i] - 1)
		f := strconv.Itoa(matriz.ListTienda[i])
		_, err = file.WriteString("nodo" + t + " ->nodo" + f + ";\n")
	}

	for i := 0; i < len(matriz.ListTienda); i++ {
		g := strconv.Itoa(matriz.ListCodigo[i])
		t := strconv.Itoa(matriz.ListTienda[i] - 1)
		_, err = file.WriteString("{ rank = same; ")
		_, err = file.WriteString("C" + g + ";")
		_, err = file.WriteString("nodo" + t + ";")
		_, err = file.WriteString("}\n")
	}

	for i := 0; i < 1; i++ {
		h := strconv.Itoa(matriz.ListDia[i])
		t := strconv.Itoa(matriz.ListTienda[i] - 1)
		_, err = file.WriteString("Mt -> D" + h + ";\n  Mt -> nodo" + t + ";\n")
	}

	//_, err = file.WriteString(tiendas_grafo)
	//_, err = file.WriteString(conexiones_grafo)
	_, err = file.WriteString("}")
	generarImagen()
	fmt.Println("Grafo editado correctamente.")
}
func generarImagen() {
	s := "dot.exe -Tpng grafo.dot -o frontend/src/assets/imagengrafo.png"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Ejucicion Fallo", err)
	}
	fmt.Println("%s\n", b)
}
func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
