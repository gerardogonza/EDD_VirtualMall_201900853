package matriz

import "fmt"

type Product struct {
	Codigo      int    //Se cambio a primer letra mayuscula para poder acceder
	Nombre      string //Se cambio a primer letra mayuscula para poder acceder
	Descripcion string //Se cambio a primer letra mayuscula para poder acceder
}

type nodo1 struct {
	//Estos atributos son especificos para la matriz
	x, y                              int      //Saber en que cabecera estoy
	producto                          *Product //tipo de objeto
	izquierda, derecha, arriba, abajo *nodo1   //nodos con los que nos desplazamos dentro de la matriz
	//Estos atributos son especificos para la lista
	header              int    //tipo interno de la cabecera
	siguiente, anterior *nodo1 // nodos con los que nos vamos a desplazar dentro de las listas
}

type lista struct {
	first, last *nodo1
}

type matriz struct {
	lst_h, lst_v *lista
}

func nodoMatriz(x int, y int, producto *Product) *nodo1 {
	return &nodo1{x, y, producto, nil, nil, nil, nil, 0, nil, nil}
}

func nodoLista(header int) *nodo1 {
	return &nodo1{0, 0, nil, nil, nil, nil, nil, header, nil, nil}
}

func newLista() *lista {
	return &lista{nil, nil}
}

//Se cambio a primer letra mayuscula para poder acceder
func NewMatriz() *matriz {
	return &matriz{newLista(), newLista()}
}

func (n *nodo1) headerX() int { return n.x }
func (n *nodo1) headerY() int { return n.y }
func (n *nodo1) toString() string {
	return "Nombre: " + n.producto.Nombre + "\nDescripcion: " + n.producto.Descripcion
}

func (l *lista) ordenar(nuevo *nodo1) {
	aux := l.first
	for aux != nil {
		if nuevo.header > aux.header {
			aux = aux.siguiente
		} else {
			if aux == l.first {
				nuevo.siguiente = aux
				aux.anterior = nuevo
				l.first = nuevo
			} else {
				nuevo.anterior = aux.anterior
				aux.anterior.siguiente = nuevo
				nuevo.siguiente = aux
				aux.anterior = nuevo
			}
			return
		}
	}
	l.last.siguiente = nuevo
	nuevo.anterior = l.last
	l.last = nuevo
}

func (l *lista) insert(header int) {
	nuevo := nodoLista(header)
	if l.first == nil {
		l.first = nuevo
		l.last = nuevo
	} else {
		l.ordenar(nuevo)
	}
}

func (l *lista) search(header int) *nodo1 {
	temp := l.first
	for temp != nil {
		if temp.header == header {
			return temp
		}
		temp = temp.siguiente
	}
	return nil
}

func (l *lista) print() {
	temp := l.first
	for temp != nil {
		fmt.Println("Cabecera:", temp.header)
		temp = temp.siguiente
	}
}

func (m *matriz) Insert(producto *Product, x int, y int) {
	h := m.lst_h.search(x)
	v := m.lst_v.search(y)

	if h == nil && v == nil {
		m.noExisten(producto, x, y)
	} else if h == nil && v != nil {
		m.existeVertical(producto, x, y)
	} else if h != nil && v == nil {
		m.existeHorizontal(producto, x, y)
	} else {
		m.existen(producto, x, y)
	}
}

func (m *matriz) noExisten(producto *Product, x int, y int) {
	m.lst_h.insert(x) //insertamos en la lista que emula la cabecera horizontal
	m.lst_v.insert(y) //insertamos en la lista que emula la cabecera vertical

	h := m.lst_h.search(x) //vamos a buscar el nodo que acabos de insertar para poder enlazarlo
	v := m.lst_v.search(y) //vamos a buscar el nodo que acabos de insertar para poder enlazarlo

	nuevo := nodoMatriz(x, y, producto) //creamos nuevo nodo tipo matriz

	h.abajo = nuevo  //enlazamos el nodo horizontal hacia abajo
	nuevo.arriba = h //enlazmos el nuevo nodo hacia arriba

	v.derecha = nuevo   //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierda = v //enlazamos el nuevo nodo hacia la izquierda
}

func (m *matriz) existeVertical(producto *Product, x int, y int) {

}

func (m *matriz) existeHorizontal(producto *Product, x int, y int) {

}

func (m *matriz) existen(producto *Product, x int, y int) {

}
