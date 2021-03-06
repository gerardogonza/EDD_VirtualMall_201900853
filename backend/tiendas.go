package main

import (
	"./MerkleInventario"
	"./TableHash"
	"./arbolmerkle"
	"./archivos"
	"./markleUsuarios"
	"./matriz"
	"./merkleTiendas"
	"./mingrafo"
	"crypto/sha256"
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
	"time"
)

var path = "matriz.dot"
var file, err = os.OpenFile(path, os.O_RDWR, 0644)
var path1 = "graph.dot"
var file1, err1 = os.OpenFile(path1, os.O_RDWR, 0644)
var path2 = "tiendalinealizada.dot"
var file2, err2 = os.OpenFile(path2, os.O_RDWR, 0644)
var path3 = "graphrutes.dot"
var file3, err3 = os.OpenFile(path3, os.O_RDWR, 0644)
var path4 = "arbolB.dot"
var file4, err4 = os.OpenFile(path4, os.O_RDWR, 0644)
var tiendas_grafo, conexiones_grafo, datos_tiendas string
var nodo, valor, nodo1 int
var espacio []listD
var espacio1 []listE
var grafos grafo
var carrito1 carritoo
var pedidos pedido
var indices indice
var datosInventario arbol
var datosTiendas tiendas
var avl = archivos.NewAVL()

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
	Nombre         string
	Codigo         int
	Descripcion    string
	Precio         float64
	Cantidad       int
	Imagen         string
	Almacenamiento string
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
type carritoo struct {
	Nombre          string
	Descripcion     string
	Departamento    string
	Codigo          int
	Precio          int
	Cantidad        int
	Cantidad_pedida int
	Imagen          string
	Almacenamiento  string
}
type pedidosCreados struct {
	Id              int
	Nombre          string
	Descripcion     string
	Departamento    string
	Codigo          int
	Precio          int
	Cantidad        int
	Cantidad_pedida int
	Imagen          string
	Almacenamiento  string
}
type grafo struct {
	Nodos                []nodos
	PosicionInicialRobot string
	Entrega              string
}
type nodos struct {
	Nombre  string
	Enlaces []enlaces
}
type enlaces struct {
	Nombre    string
	Distancia int
}
type users struct {
	Usuarios []datos
}
type datos struct {
	Dpi      int
	Nombre   string
	Correo   string
	Password string
	Cuenta   string
}

var usuarios users
var ulogin login

type login struct {
	Dpi      int
	Nombre   string
	Correo   string
	Password string
	Cuenta   string
}

func main() {

	HashTable := TableHash.NewHashTable(7)
	HashTable.Insertar(5512340, "lena", "lena")
	HashTable.Insertar(2, "Marcos", "744899223")
	HashTable.Insertar(3, "Marcos", "744899223")
	HashTable.Insertar(4, "Marcos", "744899223")

	HashTable.Insertar(5, "Marcos", "744899223")
	HashTable.Insertar(74489223, "Marcos", "74489223")
	HashTable.Insertar(8665439, "scram", "socram")
	HashTable.Print()
	t := time.Now()
	fechaInicial = fmt.Sprintf("%02d",
		t.Minute())
	fmt.Println("minuto actual es =>", fechaInicial)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	router.HandleFunc("/mostrartiendas", mostrartiendas).Methods("GET")
	router.HandleFunc("/cargarinventario", cargarInventario).Methods("POST")
	router.HandleFunc("/mostrarinventario/{numero}", mostrarinventario).Methods("GET")
	router.HandleFunc("/crearcomentarios", crearcomentarios).Methods("POST")
	router.HandleFunc("/mostrarcomentarios", mostrarcomentarios).Methods("GET")
	router.HandleFunc("/crearcomentariosproductos", crearcomentariosproductos).Methods("POST")
	router.HandleFunc("/mostrarcomentariosproductos", mostrarcomentariosproductos).Methods("GET")
	router.HandleFunc("/cargarpedido", cargarpedido).Methods("POST")
	router.HandleFunc("/mostrarpedido/{numero}", mostrarpedido).Methods("GET")
	router.HandleFunc("/carrito", carrito).Methods("POST")
	router.HandleFunc("/mostrarcarrito", mostrarcarrito).Methods("GET")
	router.HandleFunc("/cargargrafo", cargargrafo).Methods("POST")
	router.HandleFunc("/mostrarlinealizacion", linealizacion).Methods("GET")
	router.HandleFunc("/rutamin", generarRuta).Methods("POST")
	router.HandleFunc("/cargarusuarios", cargarUsuarios).Methods("POST")
	router.HandleFunc("/mostrarusuarios", mostrarUsuarios).Methods("GET")
	router.HandleFunc("/registrarusuario", registrarUsuario).Methods("POST")
	router.HandleFunc("/creararbol", creararbol).Methods("GET")
	router.HandleFunc("/tiempo", tiempo).Methods("POST")
	router.HandleFunc("/comprobartiempo", combrobartiempo).Methods("POST")
	//log.Fatal(http.ListenAndServe(":3000", router))
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor Funcionando :D")
}

var lista_carrito []carritoo
var primerprodcuto = 1

type timer struct {
	Tiempo int
}

var timerr timer
var contador int
var fechaInicial string
var tiempoint int

func tiempo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &timerr)
	if err != nil {
		log.Fatal("Error")
	}
	str1 := fechaInicial
	/** converting the str1 variable into an int using Atoi method */
	tiempoint, err := strconv.Atoi(str1)
	if err == nil {
		fmt.Println(tiempoint)
	}

}
func combrobartiempo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &timerr)
	if err != nil {
		log.Fatal("Error")
	}
	t := time.Now()
	fechafinal := fmt.Sprintf("%02d",
		t.Minute())

	var tiempodeCierre = tiempoint + timerr.Tiempo
	str1 := fechafinal
	/** converting the str1 variable into an int using Atoi method */
	tiempofin, err := strconv.Atoi(str1)
	if err == nil {
		fmt.Println(tiempofin)
	}
	if tiempofin > tiempodeCierre {
		blockchain()
	}
}
func carrito(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &carrito1)
	if err != nil {
		log.Fatal("Error")
	}
	var nuevoProdcuto = 0
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	for i := 0; i < len(lista_carrito); i++ {
		if lista_carrito[i].Codigo == carrito1.Codigo {
			lista_carrito[i].Cantidad_pedida = lista_carrito[i].Cantidad_pedida + 1
			nuevoProdcuto = 1
			break
		}
	}
	if nuevoProdcuto == 0 {
		carrito1 = carritoo{
			Cantidad_pedida: 1,
			Nombre:          carrito1.Nombre,
			Descripcion:     carrito1.Descripcion,
			Codigo:          carrito1.Codigo,
			Precio:          carrito1.Precio,
			Cantidad:        carrito1.Cantidad,
			Departamento:    carrito1.Departamento,
			Imagen:          carrito1.Imagen,
			Almacenamiento:  carrito1.Almacenamiento,
		}
		lista_carrito = append(lista_carrito, carrito1)
	}
	json.NewEncoder(w).Encode(lista_carrito)
}
func cargargrafo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &grafos)
	if err != nil {
		log.Fatal("Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(grafos)
	fmt.Println("grafos cargados ")
	crearArchivo()
	generargrafo()
}

func generargrafo() {
	if existeError(err) {
		return
	}
	defer file1.Close()
	_, err = file1.WriteString("digraph grafica{ \n node [shape=box]\n")
	for i := 0; i < len(grafos.Nodos); i++ {
		for j := 0; j < len(grafos.Nodos[i].Enlaces); j++ {
			t := strconv.Itoa(grafos.Nodos[i].Enlaces[j].Distancia)
			_, err = file1.WriteString(grafos.Nodos[i].Nombre + "->" + grafos.Nodos[i].Enlaces[j].Nombre + "[label=\"" + t + "\"];\n")
		}
	}
	_, err = file1.WriteString(grafos.PosicionInicialRobot + "[fillcolor=blue, style=\"rounded,filled\"]\n")
	_, err = file1.WriteString(grafos.Entrega + "[fillcolor=green, style=\"rounded,filled\"]\n")
	_, err = file1.WriteString("}")
	s := "dot.exe -Tpng graph.dot -o frontend/src/assets/graph.png"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Ejecucicion Fallo", err)
	}
	fmt.Println("generarndo Imagen...", b)
}
func mostrarcarrito(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(lista_carrito)
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
	fmt.Println("tiendas cargadas")

}
func linealizacion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("generando linealizacion ")
	linealizar()
	crearArchivo()
	if existeError(err2) {
		return
	}
	defer file.Close()
	_, err2 = file2.WriteString("digraph grafica{ \nnode [shape=plaintext]\n")
	_, err2 = file2.WriteString(" vector [label=<<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n")
	_, err2 = file2.WriteString("<TR>\n")
	for i := 0; i < len(indices.Datos[0].Departamentos)*len(indices.Datos)*7; i++ { //letra
		t := strconv.Itoa(i)
		_, err2 = file2.WriteString("<TD PORT=\"" + t + "\">" + t + "</TD>\n")
	}
	_, err2 = file2.WriteString("</TR></TABLE>>];\n")
	_, err2 = file2.WriteString(tiendas_grafo)
	_, err2 = file2.WriteString(conexiones_grafo)
	_, err2 = file2.WriteString("}")
	s := "dot.exe -Tpng tiendalinealizada.dot -o frontend/src/assets/linealizacion.png"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Ejecucicion Fallo", err)
	}
	fmt.Println("generarndo Imagen...", b)
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

	var list []merkleTiendas.Hashable
	for i := 0; i < len(lista_tiendas1); i++ {
		list = append(list, merkleTiendas.Bloque(lista_tiendas1[i].Nombre+" "+lista_tiendas1[i].Contacto))
	}

	completadomerkle := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	for i := 0; i < len(completadomerkle); i++ {
		if len(list) < completadomerkle[i] {
			g := completadomerkle[i] - len(list)
			for j := 0; j < g; j++ {
				t := strconv.Itoa(j)
				list = append(list, merkleTiendas.Bloque("-1  "+t))
			}
			break
		}
	}
	data := "tiendas"
	merkleTiendas.RecorreTree(merkleTiendas.Tree(list)[0].(merkleTiendas.Node))

	generarMerkle(data)
}

type comentario struct {
	Cui        int
	Comentario string
}

var comentarios comentario
var lista_comentarios []comentario
var lista_comentariosproductos []comentario

func crearcomentarios(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &comentarios)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	coment := comentario{
		Cui:        comentarios.Cui,
		Comentario: comentarios.Comentario,
	}
	lista_comentarios = append(lista_comentarios, coment)
	json.NewEncoder(w).Encode(lista_comentarios)

}
func mostrarcomentarios(w http.ResponseWriter, r *http.Request) {

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(lista_comentarios)
}
func crearcomentariosproductos(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &comentarios)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	coment := comentario{
		Cui:        comentarios.Cui,
		Comentario: comentarios.Comentario,
	}
	lista_comentariosproductos = append(lista_comentariosproductos, coment)
	json.NewEncoder(w).Encode(lista_comentariosproductos)

}
func mostrarcomentariosproductos(w http.ResponseWriter, r *http.Request) {

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(lista_comentariosproductos)
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
						Tienda:         datosInventario.Inventarios[i].Tienda,
						Nombre:         datosInventario.Inventarios[i].Productos[j].Nombre,
						Codigo:         datosInventario.Inventarios[i].Productos[j].Codigo,
						Descripcion:    datosInventario.Inventarios[i].Productos[j].Descripcion,
						Precio:         datosInventario.Inventarios[i].Productos[j].Precio,
						Cantidad:       datosInventario.Inventarios[i].Productos[j].Cantidad,
						Departamento:   datosInventario.Inventarios[i].Departamento,
						Imagen:         datosInventario.Inventarios[i].Productos[j].Imagen,
						Almacenamiento: datosInventario.Inventarios[i].Productos[j].Almacenamiento,
					}
					lista_productos = append(lista_productos, datostiendasInventario)
				}
			}
		}
	}
	for nodoo := 0; nodoo < len(archivos.List); nodoo++ {
		if lista_tiendas[taskID].Nombre == lista_productos[nodoo].Tienda {
			productosTemporales = productotemporales{
				Nombre:         lista_productos[nodoo].Nombre,
				Codigo:         lista_productos[nodoo].Codigo,
				Descripcion:    lista_productos[nodoo].Descripcion,
				Precio:         lista_productos[nodoo].Precio,
				Cantidad:       lista_productos[nodoo].Cantidad,
				Departamento:   lista_productos[nodoo].Departamento,
				Imagen:         lista_productos[nodoo].Imagen,
				Almacenamiento: lista_productos[nodoo].Almacenamiento,
			}
			lista_temporal = append(lista_temporal, productosTemporales)
		}
	}
	json.NewEncoder(w).Encode(lista_temporal)
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	var list []MerkleInventario.Hashable
	for i := 0; i < len(lista_temporal); i++ {
		t := strconv.Itoa(lista_temporal[i].Codigo)
		list = append(list, MerkleInventario.Bloque(t+" "+lista_temporal[i].Nombre+" "+lista_temporal[i].Departamento+" "+fecha))
	}

	completadomerkle := []int{4, 8, 16, 32, 64, 128, 256, 512, 1024}

	for i := 0; i < len(completadomerkle); i++ {
		if len(list) < completadomerkle[i] {
			g := completadomerkle[i] - len(list)
			for j := 0; j < g; j++ {
				t := strconv.Itoa(j)
				list = append(list, MerkleInventario.Bloque("-1  "+t))
			}
			break
		}
	}
	datostree := "inventario"
	MerkleInventario.RecorreTree(MerkleInventario.Tree(list)[0].(MerkleInventario.Node))
	generarMerkle(datostree)
}

var productosTemporales productotemporales

type productotemporales struct {
	Nombre         string
	Codigo         int
	Descripcion    string
	Precio         float64
	Cantidad       int
	Imagen         string
	Departamento   string
	Almacenamiento string
}

var datostiendasInventario productosInventario

type productosInventario struct {
	Tienda         string
	Nombre         string
	Codigo         int
	Descripcion    string
	Precio         float64
	Cantidad       int
	Imagen         string
	Departamento   string
	Almacenamiento string
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
	fmt.Println("pedidocargado")
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
	var _, err1 = os.Create(path1)
	var _, err2 = os.Create(path2)
	var _, err3 = os.Create(path3)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}

	if os.IsNotExist(err1) {
		var file1, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file1.Close()
	}
	if os.IsNotExist(err2) {
		var file, err2 = os.Create(path2)
		if existeError(err2) {
			return
		}
		defer file.Close()
	}
	if os.IsNotExist(err3) {
		var file, err3 = os.Create(path3)
		if existeError(err3) {
			return
		}
		defer file.Close()
	}
	if os.IsNotExist(err4) {
		var file, err4 = os.Create(path4)
		if existeError(err4) {
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
	s := "dot.exe -Tpng matriz.dot -o frontend/src/assets/matix.png"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Ejecucicion Fallo", err)
	}
	fmt.Println("generarndo Imagen...", b)
}
func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func (elist_e *listE) ShowData() {
	auxiliar := elist_e.first
	var temp, temp1 int
	for auxiliar != nil {
		if temp == nodo1 {
			temp1 = temp1 + 1
			if temp1 > 1 {
				t := strconv.Itoa(nodo1)
				g := strconv.Itoa(temp1)
				h := strconv.Itoa(temp1 - 1)
				tiendas_grafo = tiendas_grafo + "nodo" + t + "o" + g + "[shape=box label=\"" + auxiliar.Nombre + "\"];\n"
				conexiones_grafo = conexiones_grafo + "nodo" + t + "o" + h + " -> nodo" + t + "o" + g + "\n"
				conexiones_grafo = conexiones_grafo + "nodo" + t + "o" + g + " -> nodo" + t + "o" + h + "\n"
			}
			if temp1 == 1 {
				t := strconv.Itoa(nodo1)
				g := strconv.Itoa(temp1)
				tiendas_grafo = tiendas_grafo + "nodo" + t + "o" + g + "[shape=box label=\"" + auxiliar.Nombre + "\"];\n"
				conexiones_grafo = conexiones_grafo + "nodo" + t + " -> nodo" + t + "o" + g + "\n"
				conexiones_grafo = conexiones_grafo + "nodo" + t + "o" + g + " -> nodo" + t + "\n"
			}

		} else {
			t := strconv.Itoa(nodo1)
			tiendas_grafo = tiendas_grafo + "nodo" + t + "[shape=box label=\"" + auxiliar.Nombre + "\"];\n"
			conexiones_grafo = conexiones_grafo + "vector:" + t + " -> nodo" + t + "\n"

			temp = nodo1

		}
		auxiliar = auxiliar.next
	}
}

type nodeE struct {
	next          *nodeE
	previous      *nodeE
	Indice        string
	Departamentos string
	Nombre        string
	Descripcion   string
	Contacto      string
	Calificacion  int
}

type listE struct {
	first *nodeE
	last  *nodeE
}

func NewListE() *listE {
	return &listE{nil, nil}
}

func (elist_e *listE) Insert(Nodo *nodeE) {

	if elist_e.first == nil {
		elist_e.last = Nodo
		elist_e.first = elist_e.last
	} else {
		Nodo.previous = elist_e.last
		elist_e.last.next = Nodo
		elist_e.last = Nodo
	}

}
func linealizar() {
	espacio1 = make([]listE, len(indices.Datos[0].Departamentos)*len(indices.Datos)*10)
	for i := 0; i < len(indices.Datos); i++ { //letra
		for j := 0; j < len(indices.Datos[i].Departamentos); j++ { //departamento
			for k := 0; k < len(indices.Datos[i].Departamentos[j].Tiendas); k++ { //tienda
				Calificacion := indices.Datos[i].Departamentos[j].Tiendas[k].Calificacion - 1
				colocacion := Calificacion + 5*(j+len(indices.Datos[i].Departamentos)*i)
				newNode := nodeE{Indice: indices.Datos[i].Indice, Departamentos: indices.Datos[i].Departamentos[j].Nombre, Nombre: indices.Datos[i].Departamentos[j].Tiendas[k].Nombre, Descripcion: indices.Datos[i].Departamentos[j].Tiendas[k].Descripcion, Contacto: indices.Datos[i].Departamentos[j].Tiendas[k].Contacto, Calificacion: indices.Datos[i].Departamentos[j].Tiendas[k].Calificacion}
				espacio1[colocacion].Insert(&newNode)
			}
		}
	}

	for nodo1 = 0; nodo1 < len(indices.Datos[0].Departamentos)*len(indices.Datos)*5; nodo1++ {
		espacio1[nodo1].ShowData()
	}

}

var idpedido = 0

func generarRuta(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		log.Fatal("Error")
	}
	if existeError(err3) {
		return
	}
	defer file3.Close()
	_, err = file3.WriteString("graph grafica{ \n node [shape=\"record\"]\n")
	for i := 0; i < len(grafos.Nodos); i++ {
		for j := 0; j < len(grafos.Nodos[i].Enlaces); j++ {
			t := strconv.Itoa(grafos.Nodos[i].Enlaces[j].Distancia)
			_, err = file3.WriteString("\"" + grafos.Nodos[i].Nombre + "\"--" + "\"" + grafos.Nodos[i].Enlaces[j].Nombre + "\"" + "[label=\"" + t + "\"];\n")
		}
	}
	graph := mingrafo.NewGraph()
	for i := 0; i < len(grafos.Nodos); i++ {
		for j := 0; j < len(grafos.Nodos[i].Enlaces); j++ {
			graph.AgregarNodo(grafos.Nodos[i].Nombre, grafos.Nodos[i].Enlaces[j].Nombre, grafos.Nodos[i].Enlaces[j].Distancia)

		}
	}
	for i := 0; i < len(lista_carrito); i++ {
		if i == 0 {
			graph.Obtenerruta(grafos.PosicionInicialRobot, lista_carrito[0].Almacenamiento)
			_, err = file3.WriteString(mingrafo.S)
		} else {
			graph.Obtenerruta(lista_carrito[i-1].Almacenamiento, lista_carrito[i].Almacenamiento)
			_, err = file3.WriteString(mingrafo.S)
		}
	}
	graph.Obtenerruta(lista_carrito[len(lista_carrito)-1].Almacenamiento, grafos.Entrega)
	_, err = file3.WriteString("tabla[shape=plaintext,fontsize=10, label=<\n<TABLE BORDER=\"3\">\n")
	for i := 0; i < len(mingrafo.RutaFinally); i++ {
		_, err = file3.WriteString("<TR><TD>" + mingrafo.RutaFinally[i] + "</TD></TR>\n")
	}
	_, err = file3.WriteString("</TABLE>>];")
	_, err = file3.WriteString("}")
	fmt.Println("Ruta Generada")
	s := "dot.exe -Tpng graphrutes.dot -o frontend/src/assets/rutmin.png"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Ejecucicion Fallo", err)
	}
	fmt.Println("generando Imagen...", b)

	var list []arbolmerkle.Hashable
	for i := 0; i < len(lista_carrito); i++ {
		t := strconv.Itoa(lista_carrito[i].Codigo)
		h := strconv.Itoa(lista_carrito[i].Cantidad_pedida)
		list = append(list, arbolmerkle.Bloque(t+" "+lista_carrito[i].Nombre+" "+h))
	}

	completadomerkle := []int{4, 8, 16, 32, 64, 128, 256, 512, 1024}

	for i := 0; i < len(completadomerkle); i++ {
		if len(list) < completadomerkle[i] {
			g := completadomerkle[i] - len(list)
			for j := 0; j < g; j++ {
				t := strconv.Itoa(j)
				list = append(list, arbolmerkle.Bloque("-1  "+t))
			}
			break
		}
	}
	datostree := "pedidos"
	arbolmerkle.RecorreTree(arbolmerkle.Tree(list)[0].(arbolmerkle.Node))

	generarMerkle(datostree)
}

func cargarUsuarios(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &usuarios)
	if err != nil {
		log.Fatal("Error")
	}
	for i := 0; i < len(usuarios.Usuarios); i++ {
		user := login{
			Dpi:      usuarios.Usuarios[i].Dpi,
			Password: usuarios.Usuarios[i].Password,
			Correo:   usuarios.Usuarios[i].Correo,
			Cuenta:   usuarios.Usuarios[i].Cuenta,
			Nombre:   usuarios.Usuarios[i].Nombre,
		}
		lista_usuarios = append(lista_usuarios, user)
	}
	fmt.Println("Usuarios Cargados")
	json.NewEncoder(w).Encode(usuarios)

	var list []markleUsuarios.Hashable
	for i := 0; i < len(lista_usuarios); i++ {
		t := strconv.Itoa(lista_usuarios[i].Dpi)
		list = append(list, markleUsuarios.Bloque(lista_usuarios[i].Nombre+" "+t))
	}

	completadomerkle := []int{4, 8, 16, 32, 64, 128, 256, 512, 1024}

	for i := 0; i < len(completadomerkle); i++ {
		if len(list) < completadomerkle[i] {
			g := completadomerkle[i] - len(list)
			for j := 0; j < g; j++ {
				t := strconv.Itoa(j)
				list = append(list, markleUsuarios.Bloque("-1  "+t))
			}
			break
		}
	}
	datostree := "usuarios"
	markleUsuarios.RecorreTree(markleUsuarios.Tree(list)[0].(markleUsuarios.Node))

	generarMerkle(datostree)
}

var lista_usuarios []login

func mostrarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(lista_usuarios)

}
func registrarUsuario(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &ulogin)
	if err != nil {
		log.Fatal("Error")
	}
	fmt.Println(ulogin)
	user := login{
		Dpi:      ulogin.Dpi,
		Password: ulogin.Password,
		Correo:   ulogin.Correo,
		Cuenta:   ulogin.Cuenta,
		Nombre:   ulogin.Nombre,
	}
	lista_usuarios = append(lista_usuarios, user)
	json.NewEncoder(w).Encode(ulogin)
}
func creararbol(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		log.Fatal("Error")
	}
	if existeError(err3) {
		return
	}

	_, err = file4.WriteString("graph grafica{ \n node [shape=\"record\"]\n")
	_, err = file4.WriteString("tabla[shape=plaintext,fontsize=10, label=<\n<TABLE BORDER=\"3\">\n")
	_, err = file4.WriteString("<TR><TD>Arbol Sin Cifrar</TD></TR>\n")
	_, err = file4.WriteString("<TR><TD>DPI</TD><TD>Nombre</TD><TD>Password</TD><TD>Correo</TD></TR>\n")
	for i := 0; i < len(lista_usuarios); i++ {
		t := strconv.Itoa(lista_usuarios[i].Dpi)
		_, err = file4.WriteString("<TR><TD>" + t + "</TD><TD>" + lista_usuarios[i].Nombre + "</TD><TD>" + lista_usuarios[i].Password + "</TD><TD>" + lista_usuarios[i].Correo + "</TD></TR>\n")
	}
	_, err = file4.WriteString("</TABLE>>];")
	_, err = file4.WriteString("tabla1[shape=plaintext,fontsize=10, label=<\n<TABLE BORDER=\"3\">\n")
	_, err = file4.WriteString("<TR><TD>Arbol Cifrado</TD></TR>\n")
	_, err = file4.WriteString("<TR><TD>DPI</TD><TD>Nombre</TD><TD>Password</TD><TD>Correo</TD></TR>\n")
	for i := 0; i < len(lista_usuarios); i++ {
		t := strconv.Itoa(lista_usuarios[i].Dpi)
		data := []byte(t)
		dpi := fmt.Sprintf("x", sha256.Sum256(data))
		data1 := []byte(lista_usuarios[i].Nombre)
		name := fmt.Sprintf("x", sha256.Sum256(data1))
		data2 := []byte(lista_usuarios[i].Password)
		password := fmt.Sprintf("x", sha256.Sum256(data2))
		data3 := []byte(lista_usuarios[i].Correo)
		correo := fmt.Sprintf("x", sha256.Sum256(data3))
		_, err = file4.WriteString("<TR><TD>" + dpi[:] + "</TD><TD>" + name[:] + "</TD><TD>" + password[:] + "</TD><TD>" + correo[:] + "</TD></TR>\n")
	}
	_, err = file4.WriteString("</TABLE>>];")
	_, err = file4.WriteString("tabla2[shape=plaintext,fontsize=10, label=<\n<TABLE BORDER=\"3\">\n")
	_, err = file4.WriteString("<TR><TD>Arbol Cifrado Sensible</TD></TR>\n")
	_, err = file4.WriteString("<TR><TD>DPI</TD><TD>Nombre</TD><TD>Password</TD><TD>Correo</TD></TR>\n")
	for i := 0; i < len(lista_usuarios); i++ {
		t := strconv.Itoa(lista_usuarios[i].Dpi)
		data := []byte(t)
		dpi := fmt.Sprintf("x", sha256.Sum256(data))
		data2 := []byte(lista_usuarios[i].Password)
		password := fmt.Sprintf("x", sha256.Sum256(data2))
		data3 := []byte(lista_usuarios[i].Correo)
		correo := fmt.Sprintf("x", sha256.Sum256(data3))
		_, err = file4.WriteString("<TR><TD>" + dpi[:] + "</TD><TD>" + lista_usuarios[i].Nombre + "</TD><TD>" + password[:] + "</TD><TD>" + correo[:] + "</TD></TR>\n")
	}
	_, err = file4.WriteString("</TABLE>>];")
	_, err = file4.WriteString("}")
	crearArchivo()
	s := "dot.exe -Tpng arbolB.dot -o frontend/src/assets/arbolB.png"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Ejecucicion Fallo", err)
	}
	fmt.Println("generando arbol...", b)
}

//func TreeMerkle()  {
//	var list []arbolmerkle.Hashable
//	list=append(list,arbolmerkle.Bloque("a"))
//	completadomerkle := []int{4,8,16,32,64,128,256,512,1024}
//
//	for i := 0; i < len(completadomerkle); i++ {
//		if len(list)<completadomerkle[i]{
//			g:=completadomerkle[i]-len(list)
//			for j := 0; j < g; j++ {
//				t := strconv.Itoa(j)
//				list=append(list,arbolmerkle.Bloque("-1  "+t))
//			}
//			break
//		}
//	}
//	arbolmerkle.RecorreTree(arbolmerkle.Tree(list)[0].(arbolmerkle.Node))
//}
func generarMerkle(data string) {
	s := "dot.exe -Tpng " + data + "merkletree.dot -o frontend/src/assets/" + data + "merkle.png"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Ejecucicion Fallo", err)
	}
	fmt.Println("generarndo Imagen...", data, b)
}

type blockchainn struct {
	Indice      int
	Fecha       string
	Data        string
	Nonce       int
	PreviusHash string
	Hash        string
}

var block blockchainn
var numerohash int

func blockchain() {
	fmt.Print("HOLA")

	numerohash = numerohash + 1

}
