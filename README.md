# PETS microservice

## Description
Este microservicio es responsable de almacenar y gestionar la información de las mascotas de los usuarios. Proporciona una API RESTful para realizar operaciones CRUD sobre los registros de mascotas almacenados en una base de datos.

## Tech Stack
* **Lenguaje:** Desarrollando en GO v1.23.2.
* **Api:** OpenApi 2.0 mediante go-swagger
* **Base de datos:** PostgreSQL

## Software architecture
La estructura de proyecto se ha modelado siguiendo el patrón de arquitectura hexagonal al cual se le ha añadido vertical slicing y screaming architecture. El código sigue los principios propios de las arquitecturas limpias buscando facilitar aspectos como la mantenibilidad, escalabilidad y testeabilidad.

## Estructura de carpetas
>> /internal/commonlib --> todo lo que cuelga de aquí debería estar en una librería separada para ser compartida entre los diferentes proyectos.

## Base de datos
Port: 5432 | admin : admin | iskaypet : iskaypet

## API Reference

#### Get all items

```http
  GET /api/items
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get item

```http
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### add(num1, num2)

Takes two numbers and returns the sum.


## Información útil para despliegue y desarrollo

### Readme.md
README.md online editor (https://readme.so/editor)

### Swagger & GO-Swagger
- [Project repo&documentation](https://github.com/go-swagger/go-swagger)
- [Installing from source](https://goswagger.io/go-swagger/install/install-source/)
- [Generate a server from a swagger spec](https://goswagger.io/go-swagger/generate/server/)
- [Server generation and customization](https://goswagger.io/go-swagger/faq/faq_server/)
- [Serve a documentation site (SwaggerUI)](https://goswagger.io/go-swagger/usage/serve_ui/)
- [OpenApi 2.0 definition](https://swagger.io/docs/specification/v2_0/basic-structure/)
- [Swagger UI](http://localhost:5555/docs)
- [??](??)

https://posener.github.io/openapi-intro/
https://posener.github.io/go-swagger/
https://friendlyuser.github.io/posts/tech/2023/Using_Go-Swagger_in_Golang_A_Comprehensive_Guide/

### Validate API definition
```bash
swagger validate ./swagger.yaml
```

### Generate server code
```bash
swagger generate server --exclude-main -A backend
```

### Publish Swagger UI
```bash
swagger serve --flavor=swagger --port 5555 swagger.yaml
swagger serve --flavor=swagger --port 5555 --no-open swagger.yaml
```

### GO modules
Update all dependencies
```bash
go get -u ./...
```
Refresh go.mod
```bash
go mod tidy
```

### Build and run
```bash
  cd iskaypetchallenge
  go build ./cmd/backend-server ./cmd/backend-server/
  go run ./cmd/backend-server ./cmd/backend-server/.
```

### Health
```bash
http://localhost:8081/health
```

### Swagger UI
```bash
http://localhost:5555/docs
```