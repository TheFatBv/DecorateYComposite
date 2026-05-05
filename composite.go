package main

import (
	"fmt"
)

// COMPONENTE
type Componente interface {
	Nombre() string
	TamanioKB() float64
	Mostrar(indent string)
}

// HOJA
type Archivo struct {
	nombre    string
	tamanioKB float64
}

func (a *Archivo) Nombre() string {
	return a.nombre
}

func (a *Archivo) TamanioKB() float64 {
	return a.tamanioKB
}

func (a *Archivo) Mostrar(indent string) {
	fmt.Printf("%s📄 %s (%.1f KB)\n", indent, a.nombre, a.tamanioKB)
}

// CONTENEDOR
type Carpeta struct {
	nombre string
	hijos  []Componente
}

func (c *Carpeta) Nombre() string {
	return c.nombre
}

func (c *Carpeta) TamanioKB() float64 {
	total := 0.0

	for i := 0; i < len(c.hijos); i++ {
		hijo := c.hijos[i]
		total = total + hijo.TamanioKB()
	}

	return total
}

func (c *Carpeta) Mostrar(indent string) {
	fmt.Printf("%s📁 %s/ (%.1f KB)\n", indent, c.nombre, c.TamanioKB())

	for i := 0; i < len(c.hijos); i++ {
		hijo := c.hijos[i]
		nuevoIndent := indent + "   "
		hijo.Mostrar(nuevoIndent)
	}
}

func (c *Carpeta) Agregar(comp Componente) {
	c.hijos = append(c.hijos, comp)
}

// CLIENTE
func imprimirInfo(comp Componente) {
	comp.Mostrar("")

	tamanio := comp.TamanioKB()
	fmt.Printf("💾 Tamaño total: %.1f KB\n\n", tamanio)
}

func main() {
	musica := &Carpeta{nombre: "Musica"}

	archivo1 := &Archivo{nombre: "cancion1.mp3", tamanioKB: 4500}
	archivo2 := &Archivo{nombre: "cancion2.mp3", tamanioKB: 3800}

	musica.Agregar(archivo1)
	musica.Agregar(archivo2)

	docs := &Carpeta{nombre: "Documentos"}

	doc1 := &Archivo{nombre: "tesis.pdf", tamanioKB: 1200}
	doc2 := &Archivo{nombre: "cv.docx", tamanioKB: 85}

	docs.Agregar(doc1)
	docs.Agregar(doc2)
	docs.Agregar(musica)

	raiz := &Carpeta{nombre: "Inicio"}

	nota := &Archivo{nombre: "notas.txt", tamanioKB: 12}

	raiz.Agregar(docs)
	raiz.Agregar(nota)

	fmt.Println("=== Arbol completo ===")
	imprimirInfo(raiz)

	fmt.Println("=== Solo la carpeta Musica ===")
	imprimirInfo(musica)
}
