package merkleTiendas

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

var Path = "tiendasmerkletree.dot"
var File, Err = os.OpenFile(Path, os.O_RDWR, 0644)

type Hash [32]byte
type Bloque string

func Tree(parts []Hashable) []Hashable {
	var nodes []Hashable
	var i int
	for i = 0; i < len(parts); i += 2 {
		if i+1 < len(parts) {
			nodes = append(nodes, Node{left: parts[i], right: parts[i+1]})
		} else {
			nodes = append(nodes, Node{left: parts[i], right: EmptyBlock{}})
		}
	}
	if len(nodes) == 1 {
		return nodes
	} else if len(nodes) > 1 {
		return Tree(nodes)
	} else {
		panic("f")
	}
}

type Hashable interface {
	hash() Hash
}

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

func (b Bloque) hash() Hash {
	return hash([]byte(b)[:])
}

type EmptyBlock struct {
}

func (_ EmptyBlock) hash() Hash {
	return [32]byte{}
}
func text(data []byte) string {
	return string(data[:])
}

type Node struct {
	left  Hashable
	right Hashable
}

func (n Node) hash() Hash {
	var l, r [sha256.Size]byte
	l = n.left.hash()
	r = n.right.hash()
	return hash(append(l[:], r[:]...))
}

func hash(data []byte) Hash {
	return sha256.Sum256(data)
}

func RecorreTree(node Node) {
	crearArchivo()

	if existeError(Err) {
		return
	}
	_, Err = File.WriteString("graph g {\n node [shape=\"record\"];\ngraph [rankdir=\"BT\"];\n")
	PrintNode(node, 0)
	_, Err = File.WriteString("}")
	fmt.Println("Tree Merkle")

}

var nodenum = 0

func PrintNode(node Node, nivel int) {

	//fmt.Printf("(%d) %s %s\n", level, strings.Repeat(" ", level), node.hash())
	if l, ok := node.left.(Node); ok {
		//fmt.Println(nivel, node,node.left,node.right,"nodo")
		//t := strconv.Itoa(nodenum)
		//o := strconv.Itoa(level+1)
		if nodenum == 0 {
			_, Err = File.WriteString("\"" + node.hash().String() + "\" [label=\"" + node.hash().String() + " \\n " + node.left.hash().String() + "\\n " + node.right.hash().String() + "\"];\n")
			_, Err = File.WriteString("\"" + node.left.hash().String() + "\"--\"" + node.hash().String() + "\"\n")
			_, Err = File.WriteString("\"" + node.right.hash().String() + "\"--\"" + node.hash().String() + "\"\n")
		}

		//nodenum=nodenum+1
		//h := strconv.Itoa(nodenum)
		_, Err = File.WriteString("\"" + l.hash().String() + "\" [label=\"" + l.hash().String() + " \\n " + l.left.hash().String() + "\\n " + l.right.hash().String() + "\"];\n")
		//_, err = file.WriteString("\""+node.left.hash().String()+"\"--\""+node.hash().String()+"\"\n")
		//_, err = file.WriteString("\""+node.right.hash().String()+"\"--\""+node.hash().String()+"\"\n")
		_, Err = File.WriteString("\"" + l.left.hash().String() + "\"--\"" + l.hash().String() + "\"\n")
		_, Err = File.WriteString("\"" + l.right.hash().String() + "\"--\"" + l.hash().String() + "\"\n")
		nodenum = nodenum + 1
		//fmt.Println(nivel+1, l,"10")
		PrintNode(l, nivel+1)
	} else if l, ok := node.left.(Bloque); ok {

		nodenum = nodenum + 1
		//t := strconv.Itoa(nodenum)
		//o := strconv.Itoa(level+1)

		_, Err = File.WriteString("\"" + l.hash().String() + "\" [label=\"" + l.hash().String() + " \\n " + text([]byte(l)) + "\"];\n")

		//fmt.Println(nivel+1, node.left,node.right,"11")
		//fmt.Println(nivel+1, l,"fdf",t)

		//.Printf("(%d) %s %s (data: %s)\n", level + 1, strings.Repeat(" ", level + 1), l.hash(), l)
	}
	if r, ok := node.right.(Node); ok {
		//nodenum=nodenum+1
		//t := strconv.Itoa(nodenum)
		//o := strconv.Itoa(level+1)

		_, Err = File.WriteString("\"" + r.hash().String() + "\" [label=\"" + r.hash().String() + " \\n " + r.left.hash().String() + "\\n " + r.right.hash().String() + "\"];\n")
		_, Err = File.WriteString("\"" + r.left.hash().String() + "\"--\"" + r.hash().String() + "\"\n")
		_, Err = File.WriteString("\"" + r.right.hash().String() + "\"--\"" + r.hash().String() + "\"\n")
		//fmt.Println(nivel+1, r,"32")
		PrintNode(r, nivel+1)

	} else if r, ok := node.right.(Bloque); ok {
		//data1 := []byte(r)
		nodenum = nodenum + 1

		//t := strconv.Itoa(nodenum)
		//o := strconv.Itoa(level+1)

		//name := fmt.Sprintf( ".",string(data1))
		_, Err = File.WriteString("\"" + r.hash().String() + "\" [label=\"" + r.hash().String() + " \\n " + text([]byte(r)) + "\"];\n")
		//fmt.Println(nivel+1, r,"Fdf")

		//fmt.Printf("(%d) %s %s (data: %s)\n", level + 1, strings.Repeat(" ", level + 1), r.hash(), r)
	}

}

func crearArchivo() {
	var _, err = os.Create(Path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(Path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}

}
func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
