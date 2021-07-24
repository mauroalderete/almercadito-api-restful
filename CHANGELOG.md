# CHANGELOG

## 24/07/2021 Release v1.1.0

Libero una primera version que solo accede al maestro de clientes

## 24/07/2021 Preparo deploy por docker

Construyo un dockerfile que compila la aplicacion en una imagen y prepara su ejecucion en otra. El resultado es una imagen que permite ejecutar el microservicio.

Un dockercompose de ejemplo muestra la configuracion necesaria para levantar el servicio.

Se libera una instancia para pruebas de rendimiento.

## 24/07/2021 hash y mejoras en api de clientes

Decidi descartar el uso de un verbo reload para recargar la cache del repositorio.
Asi cada llamada consulta siempre la api de google y garantiza que los datos esten disponibles.

Tuve que crear una structura para manejar los ID hexadecimales que plantea appsheet. Todavia no es seguro que tipo de datos se utilizara para la clave principal en las tablas, como se operara al insertar un nuevo elemento y como se dara la verificacion en caso de claves repetidas.

Por esta razones, decidi crear un elemento id del tipo hash que mantiene el valor en dos variantes numericas Hex (expresado como string) e Int64.

Tambien se creo un metodo especifico para el parseo de filas de spreadsheets debido a que los tamaños de los arreglos devueltos no son siempre iguales, sino que va de acuerdo al contenido de la fila.

Se define una varaible de entorno para cada modulo. De esta forma es posible acceder a los diferentes libros y paginas que cada modulo requiere.

## 11/07/2021 montaje simple de webapi con gin

Implemento una arquitectura simple para separar:

- El control de flags
- La inicialización de un contexto de coneccion
- Las rutas y funciones del servicio de webapi

Diseñe dos rutas GET, una que retorna un estado simple y otra que toma el contexto de coneccion y devuelve un listado simple de clientes

## 10/07/2021 importo auth y login

Traslado el mecanismo de login a un proyecto diferente con su propio respositorio

https://gitlab.com/rayquen-google/golang/login

De esta forma se consigue que el cli de login funcione para otros proyectos que requieran acceso a la api de google por token.

Implemento un pequeño script que extrae el listado de nombres de clientes, para verificar que el mecanismo de login, y la importacion de la libreria funcione correctamente.

## 10/07/2021 Feature imrpot auth y login

Se construyo un script en /login que permite generar un token de acceso a un usuario de google. El script requiere un archivo de credencial válido y utiliza la libreria auth de rayquen-google para completar el login y guardar el token.

La implementación del script derivo en el hotfix v1.0.1 de la libreria auth

Por otra parte se dificultó el tratamiento de librerias en repositorios privados con subgrupos anidados de gitlab.

Al parecer go get es demasido basico, y no implementa un caso de uso como el requerido en esta situación

Para resolver el problema se siguieron varios pasos descriptos en:

- https://duythhuynh.medium.com/import-private-go-modules-from-gitlab-repositories-8933fcd79c79
- https://stackoverflow.com/questions/29707689/how-to-use-go-with-a-private-gitlab-repo

La mayoría coincide en

- configurar .gitconfig
- agregar la variable de entorno GOPRIVATE
- preparar go.mod para subgrupos

Los pasos que segui fueron:

```bash
git config --global user.name Rayquen
git config --global user.email rayquen.zero@gmail.com
```

```bash
go env -w GOPRIVATE="gitlab.com/rayquen-google/golang/"
```

```go
require ( 
    gitlab.com/rayquen-google/golang/auth v1.0.1
)

replace (
    gitlab.com/rayquen-google/golang/auth => gitlab.com/rayquen-google/golang/auth.git v1.0.1
)
```

> Existe un problema al usar GOPRIVATE y es que no se permiten mas de una unica url. Y esta variable de entorno se utiliza en todos los proyectos. Por lo que si existe un proyecto que utiliza librerias privadas de varios propietarios con subgrupos diferentes puede haber un grave problema.

> Por otra parte, GOPRIVATE al parecer requiere que se indiquen todos los subgrupos necesarios hasta alcanzar el repositorio. Utilizar subgrupos no parece natural para golan.

## 09/07/2021 Release v0.0.1

Version de inicialización

## 09/07/2021 Inicialización

Se crean los documentos principales y el arbol de directorios base para iniciar el proyecto
