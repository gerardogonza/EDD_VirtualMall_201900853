package mingrafo

type edge struct {
	node   string
	weight int
}

type graph struct {
	nodes map[string][]edge
}

func NewGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) AddEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

var Hola []int
var Rutas []string
var RutaFinally []string
var S string

func (g *graph) GetPath(origin, destiny string) (int, []string) {
	h := newHeap()
	var rutstring string
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]
		if visited[node] {
			continue
		}

		if node == destiny {
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

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {

				h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
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
