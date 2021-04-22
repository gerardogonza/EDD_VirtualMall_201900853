package mingrafo

type adya struct {
	nodo      string
	distancia int
}

type graph struct {
	nodos map[string][]adya
}

func NewGraph() *graph {
	return &graph{nodos: make(map[string][]adya)}
}

func (g *graph) AgregarNodo(inicio, final string, distancia int) {
	g.nodos[inicio] = append(g.nodos[inicio], adya{nodo: final, distancia: distancia})
	g.nodos[final] = append(g.nodos[final], adya{nodo: inicio, distancia: distancia})
}

func (g *graph) obtenerNodo(node string) []adya {
	return g.nodos[node]
}

var Hola []int
var Rutas []string
var RutaFinally []string
var S string

func (g *graph) Obtenerruta(inicio, fin string) (int, []string) {
	h := newHeap()
	var rutstring string
	h.push(path{value: 0, nodes: []string{inicio}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]
		if visited[node] {
			continue
		}

		if node == fin {
			//Hola=append(Hola,p.value)
			//Rutas=append(Rutas,p.nodes[len(p.nodes)-1])
			//ComparadordeRutas()
			for i := 0; i < len(p.nodes); i++ {
				S = S + p.nodes[i] + "[fillcolor=red, style=\"rounded,filled\"];\n"
				rutstring = rutstring + p.nodes[i] + "--"
			}
			RutaFinally = append(RutaFinally, rutstring)
			return p.value, p.nodes
		}

		for _, e := range g.obtenerNodo(node) {
			if !visited[e.nodo] {

				h.push(path{value: p.value + e.distancia, nodes: append([]string{}, append(p.nodes, e.nodo)...)})
			}
		}
		visited[node] = true
	}

	return 0, nil
}

//var numberMinor int
//func ComparadordeRutas()  {
//	fmt.Println(Hola)
//	if len(Hola)>1{
//		numberMinor = Hola[0]
//		for _, numero := range Hola {
//			if numero < numberMinor {
//				if numberMinor>0 {
//					numberMinor = numero
//				}
//
//			}
//		}
//	}
//
//}
//func Min()  {
//	for i := 0; i < len(Hola); i++ {
//		if numberMinor==Hola[i] {
//			if numberMinor>0{
//				RutaFinally=append(RutaFinally,Rutas[i])
//				Hola[i]=0
//				fmt.Println(RutaFinally)
//			}
//		}
//	}
//}
