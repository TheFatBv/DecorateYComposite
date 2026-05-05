# Patrones de Diseño: Composite y Decorator en Go

Este proyecto presenta una implementación práctica y comparativa de los patrones de diseño **Composite** y **Decorator** utilizando el lenguaje de programación Go.

## 1. Patrón Composite (Estructura de Árbol)
El archivo `composite.go` implementa un sistema de archivos simplificado donde existen **Archivos** (Hojas) y **Carpetas** (Contenedores/Composite).

- **Propósito:** Permite tratar objetos individuales y composiciones de objetos de manera uniforme.
- **Uso:** Ideal para estructuras jerárquicas tipo "parte-todo".

## 2. Patrón Decorator (Envoltorios Dinámicos)
El archivo `decorator.go` implementa un sistema de pedidos de café donde se añaden ingredientes extra (Leche, Azúcar, Chocolate).

- **Propósito:** Añade responsabilidades adicionales a un objeto de forma dinámica sin recurrir a la herencia.
- **Uso:** Alternativa flexible a la creación de múltiples subclases para cada combinación posible.


## 3. Ejecución del Proyecto

### Requisitos previos

-   Go: Versión 1.21 o superior.

### Comandos para ejecutar los ejemplos

**Clonar el repositorio:**

``` bash
git clone https://github.com/TheFatBv/DecorateYComposite.git
```

**Ejecutar ejemplo de Composite:**

``` bash
go run composite.go
```

**Ejecutar ejemplo de Decorator:**

``` bash
go run decorator.go
```

------------------------------------------------------------------------

## 4. Herramientas y Versiones

-   **Lenguaje:** Go (v1.21+)
-   **IDE recomendado:** VS Code con extensión de Go o GoLand
-   **Gestor de dependencias:** Go Modules

------------------------------------------------------------------------

## 5. Referencias Bibliográficas

-   Gamma, E., Helm, R., Johnson, R., & Vlissides, J. (1994). *Design
    Patterns: Elements of Reusable Object-Oriented Software*.
    Addison-Wesley.\
    (Págs. 163--184 para Composite, 175--188 para Decorator).

-   Shvets, A. (2019). *Dive Into Design Patterns*.

-   Refactoring.Guru. *Design Patterns - Composite & Decorator*.\
    https://refactoring.guru/design-patterns
