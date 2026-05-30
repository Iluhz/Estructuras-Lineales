package main

import (
	"errors"
	"fmt"
)

// Pila

type Stack struct {
	datos []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.datos = append(s.datos, value)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.datos) == 0 {
		return nil, errors.New("la pila esta vacia")
	}
	ultimo := s.datos[len(s.datos)-1]
	s.datos = s.datos[:len(s.datos)-1]
	return ultimo, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if len(s.datos) == 0 {
		return nil, errors.New("la pila esta vacia")
	}
	return s.datos[len(s.datos)-1], nil
}

func (s *Stack) IsEmpty() bool { return len(s.datos) == 0 }
func (s *Stack) Size() int     { return len(s.datos) }

// Cola

type Queue struct {
	datos []interface{}
}

func (q *Queue) Enqueue(value interface{}) {
	q.datos = append(q.datos, value)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.datos) == 0 {
		return nil, errors.New("la cola esta vacia")
	}
	primero := q.datos[0]
	q.datos = q.datos[1:]
	return primero, nil
}

func (q *Queue) Front() (interface{}, error) {
	if len(q.datos) == 0 {
		return nil, errors.New("la cola esta vacia")
	}
	return q.datos[0], nil
}

func (q *Queue) IsEmpty() bool { return len(q.datos) == 0 }
func (q *Queue) Size() int     { return len(q.datos) }

//Lista enlazada

type nodo struct {
	valor     interface{}
	siguiente *nodo
}

type LinkedList struct {
	cabeza *nodo
	tamano int
}

func (l *LinkedList) InsertAt(index int, value interface{}) error {
	if index < 0 || index > l.tamano {
		return errors.New("indice fuera de rango")
	}
	nuevo := &nodo{valor: value}
	if index == 0 {
		nuevo.siguiente = l.cabeza
		l.cabeza = nuevo
		l.tamano++
		return nil
	}
	actual := l.cabeza
	for i := 0; i < index-1; i++ {
		actual = actual.siguiente
	}
	nuevo.siguiente = actual.siguiente
	actual.siguiente = nuevo
	l.tamano++
	return nil
}

func (l *LinkedList) RemoveAt(index int) (interface{}, error) {
	if index < 0 || index >= l.tamano {
		return nil, errors.New("indice fuera de rango")
	}
	if index == 0 {
		valor := l.cabeza.valor
		l.cabeza = l.cabeza.siguiente
		l.tamano--
		return valor, nil
	}
	actual := l.cabeza
	for i := 0; i < index-1; i++ {
		actual = actual.siguiente
	}
	valor := actual.siguiente.valor
	actual.siguiente = actual.siguiente.siguiente
	l.tamano--
	return valor, nil
}

func (l *LinkedList) GetAt(index int) (interface{}, error) {
	if index < 0 || index >= l.tamano {
		return nil, errors.New("indice fuera de rango")
	}
	actual := l.cabeza
	for i := 0; i < index; i++ {
		actual = actual.siguiente
	}
	return actual.valor, nil
}

func (l *LinkedList) Size() int { return l.tamano }

func main() {
	// Pila
	fmt.Println("=== Pila ===")
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	v, _ := s.Pop()
	fmt.Println("Pop:", v) // 3  (ultimo en entrar, primero en salir)
	top, _ := s.Peek()
	fmt.Println("Peek:", top) // 2
	fmt.Println("Size:", s.Size())

	// Cola
	fmt.Println("\n=== COla ===")
	q := &Queue{}
	q.Enqueue("primero")
	q.Enqueue("segundo")
	q.Enqueue("tercero")
	d, _ := q.Dequeue()
	fmt.Println("Dequeue:", d) // primero (el que lleva mas tiempo esperando)
	f, _ := q.Front()
	fmt.Println("Front:", f) // segundo
	fmt.Println("Size:", q.Size())

	// Lista enlazada
	fmt.Println("\n=== Lista ===")
	l := &LinkedList{}
	l.InsertAt(0, "A")
	l.InsertAt(1, "B")
	l.InsertAt(2, "C")
	val, _ := l.GetAt(1)
	fmt.Println("GetAt(1):", val) // B
	l.RemoveAt(1)
	val, _ = l.GetAt(1)
	fmt.Println("Despues de RemoveAt(1), GetAt(1):", val) // C
	fmt.Println("Size:", l.Size())
}
