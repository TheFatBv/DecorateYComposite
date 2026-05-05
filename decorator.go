package main

import "fmt"

type Bebida interface {
	getDescripcion() string
	getPrecio() float64
}

type Cafe struct{}

func (c Cafe) getDescripcion() string {
	return "Café"
}

func (c Cafe) getPrecio() float64 {
	return 20.0
}

type DecoradorBebida struct {
	bebida Bebida
}

func (d DecoradorBebida) getDescripcion() string {
	return d.bebida.getDescripcion()
}

func (d DecoradorBebida) getPrecio() float64 {
	return d.bebida.getPrecio()
}

type Leche struct {
	DecoradorBebida
}

func (l Leche) getDescripcion() string {
	return l.bebida.getDescripcion() + " con leche"
}

func (l Leche) getPrecio() float64 {
	return l.bebida.getPrecio() + 5.0
}

type Azucar struct {
	DecoradorBebida
}

func (a Azucar) getDescripcion() string {
	return a.bebida.getDescripcion() + " con azúcar"
}

func (a Azucar) getPrecio() float64 {
	return a.bebida.getPrecio() + 2.0
}

type Chocolate struct {
	DecoradorBebida
}

func (c Chocolate) getDescripcion() string {
	return c.bebida.getDescripcion() + " con chocolate"
}

func (c Chocolate) getPrecio() float64 {
	return c.bebida.getPrecio() + 10.0
}

func main() {
	var bebida Bebida = Cafe{}
	bebida = Leche{DecoradorBebida{bebida}}
	bebida = Azucar{DecoradorBebida{bebida}}

	fmt.Println(bebida.getDescripcion())
	fmt.Println(bebida.getPrecio())

	bebida = Chocolate{DecoradorBebida{bebida}}
	fmt.Println(bebida.getDescripcion())
	fmt.Println(bebida.getPrecio())

	var bebida2 Bebida = Cafe{}
	bebida2 = Chocolate{DecoradorBebida{bebida2}}
	fmt.Println(bebida2.getDescripcion())
	fmt.Println(bebida2.getPrecio())
}
