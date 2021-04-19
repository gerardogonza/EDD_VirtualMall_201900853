package arbolb

type User struct {
	Dpi      int
	Nombre   string
	Correo   string
	Password string
	Cuenta   string
}

type ArbolB struct {
	Root *Nodo
}

type Nodo struct {
	hoja         bool
	numero_users int
	Child        [6]*Nodo
	Users        [5]*User
	Parent       *Nodo
}

// constructor nodo para poner con un padre
func Nodo_(Parent *Nodo) *Nodo {
	return &Nodo{Parent: Parent, hoja: true, numero_users: 0}
}

// constructor para el arbol
func Btree() *ArbolB {
	return &ArbolB{Root: Nodo_(nil)}
}

// inserta los datos ordenado de una vez
func (nodo *Nodo) Insert(user *User) {
	nodo.numero_users++
	for i := 0; i < nodo.numero_users; i++ {
		if nodo.Users[i] == nil {
			nodo.Users[i] = user
			break
		} else if nodo.Users[i].Dpi > user.Dpi {
			for j := nodo.numero_users; j > i; j-- {
				nodo.Users[j] = nodo.Users[j-1]
			}
			nodo.Users[i] = user
		}
	}
}

// buscar el dpi en cada nodo
func (nodo *Nodo) Find(dpi int) *User {
	for i := 0; i < len(nodo.Users); i++ {
		if nodo.Users[i].Dpi == dpi {
			return nodo.Users[i]
		}
	}
	return nil
}

// insertar en el arbol
func (arbol *ArbolB) Insertar(user *User, tmp *Nodo) {
	if tmp.hoja {
		tmp.Insert(user)
	} else {
		encontrado := false
		for i := 0; i < tmp.numero_users-1; i++ {
			if user.Dpi < tmp.Users[i].Dpi {
				encontrado = true
				arbol.Insertar(user, tmp.Child[i])
				break
			}
		}
		if !encontrado {
			arbol.Insertar(user, tmp.Child[tmp.numero_users])
		}
	}
	if tmp.numero_users == 5 {
		if tmp.Parent != nil {
			middleKey := tmp.Users[2]
			tmp.Parent.Insert(middleKey)
			index := 0
			for index = 0; index < tmp.Parent.numero_users; index++ {
				if tmp.Parent.Users[index] == middleKey {
					break
				}
			}
			for i := tmp.Parent.numero_users; i > index+1; i-- {
				if tmp.Parent.numero_users < 5 {
					tmp.Parent.Child[i] = tmp.Parent.Child[i-1]
				} else {
					arbol.Insertar(user, tmp.Parent)
				}
			}
			tmp.Parent.Child[index+1] = Nodo_(tmp.Parent)
			tmp.Parent.Child[index+1].Child[0] = tmp.Child[3]
			if tmp.Parent.Child[index+1].Child[0] != nil {
				tmp.Parent.Child[index+1].Child[0].Parent = tmp.Parent.Child[index+1]
				tmp.Parent.Child[index+1].hoja = false
			}
			for i := 3; i < 5; i++ {
				tmp.Parent.Child[index+1].Insert(tmp.Users[i])
				tmp.Parent.Child[index+1].Child[i-2] = tmp.Child[i+1]
				if tmp.Parent.Child[index+1].Child[i-2] != nil {
					tmp.Parent.Child[index+1].Child[i-2].Parent = tmp.Parent.Child[index+1]
					tmp.Parent.Child[index+1].hoja = false
				}
			}
		} else {

		}

	}

}
