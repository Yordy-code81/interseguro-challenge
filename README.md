# Interseguro Matrix Challenge 🚀

Este proyecto es una solución integral para el procesamiento matemático de matrices, construido bajo una arquitectura de microservicios. Consta de dos APIs conectadas y orquestadas mediante Docker.

## 🏗️ Arquitectura del Sistema

La solución se divide en dos microservicios altamente acoplados por red y protegidos mediante JSON Web Tokens (JWT):

## 🌐 Demo en Vivo

Puedes probar la API completamente funcional y documentada que he desplegado en el siguiente enlace:
👉 **[http://20.97.194.76:3000/swagger/](http://20.97.194.76:3000/swagger/)**

1. **Go Matrix API (Golang / Fiber) - Puerto `3000`**
   - Actúa como la puerta de enlace (API Gateway) para el cliente.
   - Es responsable de generar los Tokens JWT de acceso.
   - Aplica operaciones matemáticas complejas: Rotación de matrices 90° a la derecha y Factorización QR (Gram-Schmidt Modificado).
   - Delega la analítica estadística a Node.js y unifica la respuesta final.

2. **Node Analytics API (Node.js / Express) - Puerto `4000`**
   - Servicio interno (no expuesto directamente al cliente final en un entorno real).
   - Recibe las matrices $Q$ y $R$ factorizadas.
   - Calcula estadísticas en O(N): Valor máximo, mínimo, promedio, suma total y verifica si las matrices son diagonales.

---

## 🛠️ Requisitos Previos

Solo necesitas tener instalado en tu sistema:
- [Docker](https://www.docker.com/products/docker-desktop)
- Docker Compose

*No es necesario tener Go ni Node.js instalados localmente, ya que Docker se encargará de construir los contenedores y descargar las dependencias.*

## ⚙️ Configuración de Variables de Entorno (.env)

Por seguridad, los archivos `.env` reales no se suben al repositorio. Al clonar este proyecto, debes crear tus propios archivos basándote en los ejemplos provistos.

En la terminal, ejecuta los siguientes comandos para crear una copia de los archivos de ejemplo:

```bash
# Para la API de Go
cp go-matrix-api/.env.example go-matrix-api/.env

# Para la API de Node.js
cp node-analytics-api/.env.example node-analytics-api/.env
```

*Nota: Aunque Docker Compose inyecta algunas de estas variables automáticamente, es una excelente práctica tenerlas creadas localmente para ejecutar pruebas o correr los servidores fuera de Docker.*

---

## 🚀 Levantando el Proyecto

Para levantar todo el ecosistema (y ejecutar automáticamente los pipelines de pruebas unitarias), sitúate en la raíz del proyecto y ejecuta:

```bash
docker-compose up --build -d
```

Este comando:
1. Construirá la imagen de Node.js (ejecutando tests de Jest en el proceso).
2. Construirá la imagen de Go (ejecutando tests de Go en el proceso y auto-generando Swagger).
3. Desplegará ambos contenedores conectados por una red virtual privada.

---

## 📖 Uso y Pruebas con Swagger UI (Recomendado)

La forma más amigable de probar la API es utilizando la interfaz gráfica de Swagger que hemos implementado en la API de Go.

1. Abre tu navegador e ingresa a: `http://localhost:3000/swagger/`
2. **Generar Token:**
   - Despliega el endpoint `GET /api/auth/token`.
   - Haz clic en **Try it out** y luego en **Execute**.
   - Copia el string del token que aparece en el cuerpo de respuesta.
3. **Autorizar Swagger:**
   - Sube al inicio de la página y haz clic en el botón verde **Authorize**.
   - En el cuadro de texto, escribe la palabra `Bearer ` seguida de un espacio y pega tu token (ej. `Bearer eyJhbG...`).
   - Clic en **Authorize** y luego cierra la ventana.
4. **Procesar Matriz:**
   - Despliega el endpoint `POST /api/matrix`.
   - Haz clic en **Try it out**.
   - Edita la matriz de ejemplo si lo deseas (prueba con matrices rectangulares como `[[2,4,6,8], [1,3,5,7]]`).
   - Haz clic en **Execute** y disfruta de los resultados.

---

## 💻 Uso con cURL o Postman

Si prefieres la consola, el flujo es exactamente el mismo:

**1. Obtener Token:**
```bash
curl -X GET http://localhost:3000/api/auth/token
```

**2. Enviar Matriz (Sustituye `<TU_TOKEN>`):**
```bash
curl -X POST http://localhost:3000/api/matrix \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TU_TOKEN>" \
  -d '{"matrix": [[1,2], [3,4], [5,6]]}'
```

---

## ✅ Pruebas Automatizadas (CI Pipeline)

El proyecto cuenta con robustas pruebas unitarias y de integración. Estas pruebas **se ejecutan automáticamente al momento de construir los contenedores (`docker-compose build`)**. Si alguna prueba falla, el despliegue se cancela para proteger la integridad del ecosistema.

- **Go (Golang Testing):** Valida la rotación correcta de matrices $M \times N$ y la exactitud matemática de la factorización QR mediante la multiplicación $Q \times R = M$.
- **Node.js (Jest + Supertest):** Valida la precisión de los cálculos estadísticos (min, max, avg, sum) incluyendo manejo de números negativos y ceros, y evalúa el ciclo HTTP simulando peticiones con y sin Tokens JWT.
