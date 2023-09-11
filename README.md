# plantillas-crud
✔️ Check: CRUD para el nuevo módulo de Plantillas

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)


### Variables de Entorno
```shell
# parametros de api
PLANTILLAS_CRUD_USER=[nombre del usuario]
PLANTILLAS_CRUD_PASS=[contraseña del usuario]
PLANTILLAS_CRUD_HOST=[direccion de host]
PLANTILLAS_CRUD_PORT=[puerto de ejecucion] BD
PLANTILLAS_CRUD_DB=[nombre de la base de datos]


### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/plantillas_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/plantillas_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
PLANTILLAS_CRUD_PORT=27017 PLANTILLAS_CRUD_HOST=127.0.0.1 PLANTILLAS_CRUD_SOME_VARIABLE=some_value bee run

