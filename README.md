# Patrones de Diseño: Composite y Decorator en Go

Este proyecto presenta una implementación práctica y comparativa de los patrones de diseño **Composite** y **Decorator** utilizando el lenguaje de programación Go.

---

## Patrones implementados

### Composite — Sistema de archivos

**Problema que resuelve:** Representar jerarquías de objetos donde los clientes deben tratar de forma uniforme tanto a objetos individuales (hojas) como a grupos de objetos (contenedores). En este ejemplo, archivos y carpetas se gestionan con la misma interfaz `Componente`, permitiendo calcular tamaños y mostrar el árbol completo de forma recursiva sin distinguir si se trata de un archivo simple o una carpeta con contenido anidado.

**Estructura:**
```
composite/
├── composite.go   # Interfaz Componente, structs Archivo (hoja) y Carpeta (contenedor)
└── go.mod
```

---

### Decorator — Personalización de bebidas

**Problema que resuelve:** Añadir responsabilidades adicionales a un objeto de forma dinámica, sin modificar su clase ni crear subclases para cada combinación posible. En este ejemplo, un `Cafe` base puede decorarse con `Leche`, `Azucar` y `Chocolate` en cualquier orden y combinación, acumulando descripción y precio en cada capa sin alterar el objeto original.

**Estructura:**
```
decorator/
├── decorator.go   # Interfaz Bebida, struct Cafe (componente), DecoradorBebida + decoradores concretos
└── go.mod
```

---
## 3. Ejecución del Proyecto

### Requisitos previos

-   Go: Versión 1.21 o superior.
-   No se utilizan dependencias externas. Ambos módulos usan únicamente la biblioteca estándar de Go.

---

### Comandos para ejecutar los ejemplos

**Clonar el repositorio:**

``` bash
git clone https://github.com/TheFatBv/DecorateYComposite.git
```

## Cómo ejecutar

Cada patrón es un módulo independiente. Ejecutar desde la raíz del repositorio:

### Composite

```bash
cd composite
go run composite.go
```

**Salida esperada:**
```
=== Arbol completo ===
📁 Inicio/ (9597.0 KB)
   📁 Documentos/ (9585.0 KB)
      📄 tesis.pdf (1200.0 KB)
      📄 cv.docx (85.0 KB)
      📁 Musica/ (8300.0 KB)
         📄 cancion1.mp3 (4500.0 KB)
         📄 cancion2.mp3 (3800.0 KB)
   📄 notas.txt (12.0 KB)
💾 Tamaño total: 9597.0 KB

=== Solo la carpeta Musica ===
📁 Musica/ (8300.0 KB)
   📄 cancion1.mp3 (4500.0 KB)
   📄 cancion2.mp3 (3800.0 KB)
💾 Tamaño total: 8300.0 KB
```

---

### Decorator

```bash
cd decorator
go run decorator.go
```

**Salida esperada:**
```
Café con leche con azúcar
27
Café con leche con azúcar con chocolate
37
Café con chocolate
30
```

---

## Estructura del repositorio

```
/
├── composite/
│   ├── composite.go
│   └── go.mod
├── decorator/
│   ├── decorator.go
│   └── go.mod
├── .gitignore
└── README.md
```

---

## Relación entre los patrones

Composite y Decorator comparten una estructura de composición recursiva similar, pero tienen propósitos distintos: Composite agrupa múltiples hijos para tratarlos como una unidad, mientras que Decorator envuelve un único objeto para extender su comportamiento. Puedes colaborar: un Decorator puede aplicarse sobre un nodo específico de un árbol Composite.

------------------------------------------------------------------------

## 4. Herramientas y Versiones

-   **Lenguaje:** Go (v1.21+)
-   **IDE recomendado:** VS Code con extensión de Go o GoLand

------------------------------------------------------------------------
