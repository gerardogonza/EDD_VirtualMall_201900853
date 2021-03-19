package archivos

import "fmt"

type nodo struct {
	indice   Producto
	altura   int
	izq, der *nodo
}
type Producto struct {
	Nombre string
	Codigo int
	Descripcion string
	Precio float64
	Cantidad int
	Imagen string
}

func newNodo(indice Producto) *nodo {
	return &nodo{indice, 0, nil, nil}
}

type AVL struct {
	raiz *nodo
}

func NewAVL() *AVL {
	return &AVL{nil}
}

func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func altura(temp *nodo) int {
	if temp != nil {
		return temp.altura
	}
	return -1
}

func rotacionIzquierda(temp **nodo) {
	aux := (*temp).izq
	(*temp).izq = aux.der
	aux.der = *temp
	(*temp).altura = max(altura((*temp).der), altura((*temp).izq)) + 1
	aux.altura = max(altura((*temp).izq), (*temp).altura) + 1
	*temp = aux
}

func rotacionDerecha(temp **nodo) {
	aux := (*temp).der
	(*temp).der = aux.izq
	aux.izq = *temp
	(*temp).altura = max(altura((*temp).der), altura((*temp).izq)) + 1
	aux.altura = max(altura((*temp).der), (*temp).altura) + 1
	*temp = aux
}

func rotacionDobleIzquierda(temp **nodo) {
	rotacionDerecha(&(*temp).izq)
	rotacionIzquierda(temp)
}

func rotacionDobleDerecha(temp **nodo) {
	rotacionIzquierda(&(*temp).der)
	rotacionDerecha(temp)
}

func insert(indice Producto, root **nodo) {
	if (*root) == nil {
		*root = newNodo(indice)
		return
	}
	if indice.Codigo < (*root).indice.Codigo {
		insert(indice, &(*root).izq)
		if (altura((*root).izq) - altura((*root).der)) == -2 {
			if indice.Codigo < (*root).izq.indice.Codigo{
				rotacionIzquierda(root)
			}else{
				rotacionDobleIzquierda(root)
			}
		}
	}else if indice.Codigo > (*root).indice.Codigo{
		insert(indice, &(*root).der)
		if (altura((*root).der) - altura((*root).izq)) == 2{
			if indice.Codigo > (*root).der.indice.Codigo {
				rotacionDerecha(root)
			}else{
				rotacionDobleDerecha(root)
			}
		}
	}else{
		fmt.Println("Ya se inserto el indice")
	}

	(*root).altura = max(altura((*root).izq), altura((*root).der))+1
}

func (avl *AVL) Insertar(indice Producto) {
	insert(indice, &avl.raiz)
}

func (avl *AVL) Print(){
	inOrden(avl.raiz)
}

func inOrden(temp *nodo)  {
	if temp != nil {
		inOrden(temp.izq)
		fmt.Println("Inorder: ", temp.indice)
		inOrden(temp.der)
	}
}
