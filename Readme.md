# Proyecto de Práctica: Go Final REST API

Este proyecto de práctica utiliza las siguientes tecnologías:

- Golang: un lenguaje de programación de alto rendimiento y eficiencia.
- SQLite: una base de datos ligera y fácil de usar.
- Air: una herramienta de desarrollo en tiempo real para Go.
- Gin: un framework web rápido y minimalista para Go.

## Descripción

Este proyecto tiene como objetivo desarrollar una API REST utilizando Golang, SQLite, Air y Gin. La API permitirá realizar operaciones CRUD (Crear, Leer, Actualizar, Eliminar) en una base de datos SQLite.

## Requisitos

Antes de comenzar, asegúrate de tener instalados los siguientes componentes:

- Golang: [Descargar e instalar Golang](https://golang.org/dl/)
- SQLite: [Descargar e instalar SQLite](https://www.sqlite.org/download.html)
- Air: [Instalar Air](https://github.com/cosmtrek/air#installation)
- Gin: [Instalar Gin](https://github.com/gin-gonic/gin#installation)

## Configuración

1. Clona este repositorio en tu máquina local.
2. Abre una terminal y navega hasta el directorio del proyecto.
3. Ejecuta el siguiente comando para instalar las dependencias:

```shell
go mod download
```

4. Configura la conexión a la base de datos en el archivo `config.go`.
5. Ejecuta el siguiente comando para iniciar el servidor:

```shell
air
```

## Uso

Una vez que el servidor esté en funcionamiento, puedes realizar las siguientes operaciones:

- Crear un nuevo registro: `POST /api/registro`
- Obtener todos los registros: `GET /api/registro`
- Obtener un registro por ID: `GET /api/registro/{id}`
- Actualizar un registro existente: `PUT /api/registro/{id}`
- Eliminar un registro: `DELETE /api/registro/{id}`

## Contribución

Si deseas contribuir a este proyecto, sigue los siguientes pasos:

1. Haz un fork de este repositorio.
2. Crea una rama con tu nueva funcionalidad: `git checkout -b nueva-funcionalidad`
3. Realiza los cambios necesarios y realiza un commit: `git commit -m "Agrega nueva funcionalidad"`
4. Haz push a la rama: `git push origin nueva-funcionalidad`
5. Abre un pull request en este repositorio.

## Licencia

Este proyecto se encuentra bajo la Licencia MIT. Para más información, consulta el archivo `LICENSE`.
