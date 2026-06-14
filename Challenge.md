# Interseguro - Coding Challenge
## División TI
**Fecha:** Junio 2024

---

## Desafío Técnico - Descripción

### Consideraciones técnicas
* [cite_start]Utilizar el lenguaje de programación Go (Golang) para una API y Node.js para la otra API[cite: 7].
* [cite_start]Implementar la solución utilizando los frameworks Fiber para la API en Go y Express.js para la API en Node.js[cite: 10].
* [cite_start]Documentar el código de manera clara y concisa, siguiendo las mejores prácticas de codificación[cite: 11].
* [cite_start]Utilizar Docker para contenerizar las aplicaciones y facilitar su despliegue en diferentes entornos[cite: 12].
* [cite_start]Implementar la comunicación entre las dos API utilizando un mecanismo como HTTP[cite: 13].
* [cite_start]Utilizar servicios en la nube para la implementación y el despliegue de las aplicaciones[cite: 14].

### Arquitectura de la solución
* [cite_start]**API en Go:** Esta API recibirá la matriz original como entrada, realizará la rotación de la matriz y luego enviará los datos resultantes a la segunda API en Node.js[cite: 17].
* [cite_start]**API en Node.js:** Esta API recibirá los datos de la matriz rotada de la API en Go, calculará estadísticas sobre los datos y devolverá estas estadísticas como resultado[cite: 18].

---

### Funcionalidad requerida
* **Crear dos API RESTful:**
    * [cite_start]Una API en Go que reciba como entrada un array de arrays de números que represente una matriz rectangular y devuelva la factorización QR de dicha matriz[cite: 22].
    * [cite_start]Otra API en Node.js que reciba el resultado de las matrices devueltas por la primera API y realice una operación adicional sobre los datos (Detalle en la sección operaciones adicionales)[cite: 23, 24].
* [cite_start]Implementar la lógica para realizar la rotación de la matriz y la operación adicional de manera eficiente y correcta en cada API[cite: 25].
* [cite_start]Aplicar un nivel de seguridad utilizando JWT para proteger las consultas a las APIs[cite: 28].
* [cite_start]Implementar pruebas unitarias y de integración para garantizar la calidad del código en ambas API[cite: 29].

### Funcionalidad opcional
* [cite_start]Implementar un frontend que consuma ambas APIs y muestre los resultados de la rotación de la matriz y las estadísticas adicionales[cite: 27].

---

### Operación adicional
[cite_start]La segunda API (Node.js) calculará lo siguiente sobre los datos de las matrices devueltas[cite: 32]:
* [cite_start]**Valor máximo:** El valor máximo encontrado en las matrices[cite: 33].
* [cite_start]**Valor mínimo:** El valor mínimo encontrado en las matrices[cite: 33].
* [cite_start]**Promedio:** El promedio de todos los valores de las matrices[cite: 34].
* [cite_start]**Suma total:** El suma total de todos los valores de las matrices[cite: 35].
* [cite_start]**Matriz diagonal:** Verificar si alguna matriz es diagonal[cite: 35].

### Consideraciones finales
* [cite_start]No hay un estándar específico para los nombres de los objetos creados, pero se espera coherencia en su estructura y documentación[cite: 38].
* [cite_start]En caso de dudas en el enunciado, se espera que el candidato tome decisiones informadas y las sustente durante la entrevista[cite: 39].
* [cite_start]Se valorará la eficiencia y la elegancia de la solución implementada, así como la capacidad del candidato para comunicar y defender sus decisiones técnicas[cite: 40].

---
[cite_start]Muchas gracias!! [cite: 42]