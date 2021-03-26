package matriz

import "fmt"

type Node struct {
	value int
	up    *Node
	down  *Node
	right *Node
	left  *Node
}

//Matrix is a set of Nodes
type Matrix struct {
	head *Node
}

func (m *Matrix) Init() {
	m.head = &Node{value: 0}
}

func (m *Matrix) addRow(row int) {
	tmp := m.head
	if tmp.down == nil {
		newNode := &Node{value: row}
		tmp.down = newNode
		newNode.up = tmp
	} else {
		for tmp.down != nil && tmp.down.value < row {
			tmp = tmp.down
		}
		if tmp.down == nil && tmp.value != row {
			newNode := &Node{value: row}
			tmp.down = newNode
			newNode.up = tmp
		} else if tmp.down != nil && tmp.down.value != row {
			aux := tmp.down
			newNode := &Node{value: row}
			tmp.down = newNode
			newNode.up = tmp
			newNode.down = aux
			aux.up = newNode
		}
	}
}

func (m *Matrix) addCol(col int) {
	tmp := m.head
	if tmp.right == nil {
		newNode := &Node{value: col}
		tmp.right = newNode
		newNode.left = tmp
	} else {
		for tmp.right != nil && tmp.right.value < col {
			tmp = tmp.right
		}
		if tmp.right == nil && tmp.value != col {
			newNode := &Node{value: col}
			tmp.right = newNode
			newNode.left = tmp
		} else if tmp.right != nil && tmp.right.value != col {
			aux := tmp.right
			newNode := &Node{value: col}
			tmp.right = newNode
			newNode.left = tmp
			newNode.right = aux
			aux.left = newNode
		}
	}
}

// fix the error
func (m *Matrix) addNode(row, col, value int) {
	newNode := &Node{value: value}
	tmprow := m.head
	tmpcol := m.head
	for tmprow.down != nil {
		tmprow = tmprow.down
		if tmprow.value == row {
			tmprow.right = newNode
			newNode.left = tmprow
		}
	}
	for tmpcol.right != nil {
		tmpcol = tmpcol.right
		if tmpcol.value == col {
			tmpcol.down = newNode
			newNode.up = tmpcol
		}
	}
}

func (m *Matrix) Add(row, col, value int) {
	m.addRow(row)
	m.addCol(col)
	m.addNode(row, col, value)
}

func getCol(tmp *Node) int {
	var col int
	for tmp.up != nil {
		tmp = tmp.up
		col = tmp.value
	}
	return col
}
var ListTienda []int
var ListDia []int
var ListCodigo[]int
func (m *Matrix) Show() {
	tmprow := m.head
	for tmprow != nil {
		tmpcol := m.head

		tmpcol = tmprow.right
		if tmprow.value == 0 {
			for tmpcol != nil {

				tmpcol = tmpcol.right
			}
		} else {
			for tmpcol != nil {
				ListTienda=append(ListTienda,tmprow.value)
				ListDia=append(ListDia,getCol(tmpcol))
				ListCodigo=append(ListCodigo,tmpcol.value)
				//fmt.Println(tmprow.value, ",", getCol(tmpcol), "(", tmpcol.value, ") ")
				tmpcol = tmpcol.right
			}
		}
		tmprow = tmprow.down
		fmt.Println()
	}
}
