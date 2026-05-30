# Estructuras-Lineales
Implementacion de estructras de datos lineales en go, usando colas, listas y pilas.

¿Para qué sirve?
Las tres estructuras son formas de organizar datos en memoria. La diferencia entre ellas es el orden en que puedes meter y sacar cosas.

Pila (Stack)
Funciona como una pila de platos: el último que pusiste es el primero que sacas, a esto se le llama LIFO (Last In, First Out). Tiene tres operaciones: Push para meter un dato arriba, Pop para sacar el de arriba, y Peek para ver cuál está arriba sin sacarlo. Se usa por ejemplo cuando el navegador guarda el historial de páginas para el botón "atrás".

Cola (Queue)
Funciona como una fila del banco: el primero que llegó es el primero que se atiende, a esto se le llama FIFO (First In, First Out). Tiene tres operaciones: Enqueue para meter un dato al final de la fila, Dequeue para sacar al que lleva más tiempo esperando, y Front para ver quién es el siguiente sin sacarlo. Se usa para procesar tareas en orden, como una impresora que atiende los documentos uno por uno.

Lista Enlazada (LinkedList)
Es como una cadena donde cada elemento sabe cuál es el siguiente. A diferencia de un arreglo normal, puedes insertar o borrar en cualquier posición sin mover todo lo demás, lo que la hace más eficiente en esos casos. Tiene tres operaciones: InsertAt para meter un dato en la posición que quieras, RemoveAt para borrar el dato en esa posición, y GetAt para obtener el dato en esa posición.

Para ejecutar codigo:
go run main.go
