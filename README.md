# cake-pieces (v1.0) [ES]

Cake Pieces es un proyecto que ofrece soluciones a diversos desafíos de algoritmos y estructuras de datos, centrándose en la encriptación de mensajes y la eliminación de subarreglos consecutivos que suman cero.

## Detalles de la API

-   Usa go version 1.22.8

## Inicia el proyecto

-   go run main.go


## Nuevas funcionalidades

### 1. Encripta tu mensaje
Has sido encargado de desarrollar una nueva forma de encriptar comunicaciones. 
Básicamente, cada vocal del mensaje de entrada deberá ser precedida por otra cadena, 
llamada la clave. La función recibirá dos parámetros de cadena: el primero será la clave y el 
segundo, el mensaje. La función devolverá una cadena

### 2. Suma a cero
Dado un arreglo de enteros, elimina todos los nodos consecutivos cuya suma sea cero y devuelve los nodos restantes. Un arreglo vacío también puede ser un resultado válido. Si se recibe un valor nulo, devuelve un arreglo vacío.
Ejemplo:
Entrada: [3, 4, -7, 5, -6, 2, 5, -1, 8]
Salida: [5, 8]
