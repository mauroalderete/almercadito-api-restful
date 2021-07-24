# alMercadito API RESTful

API RESTful del sistema alMercadito desarrollado con la plataforma AppSheet y Googlesheets

## Objetivo

Desarrollar una API RESTful en golang que se conecte a la api v4 de googlesheets y permita consultar la hoja de datos.
La información leida se brindara en por verbos webapi permitiendo ampliar las posibilidades de explotación.

### Detalles

Se utilizara la libreria para la authenticación del servicio de google con oauth2, https://gitlab.com/rayquen-google/golang/auth

La api tendra dos modos de ejecución. Uno para generar las credenciales y otro para escuchar las solicitudes.

El formato de salida será en JSON.

El modelos de datos se correspondera con el planteado en las planillas para los maestros y como agregados para las entidades compuestas.

Se mantendra en cache una copia de los datos desde la última actualización.

La actualización se realizará de forma periodica.

Se dispodra de una solicitud para forzar la actualización de la cache.

El entorno hará uso de variables de entorno, y será diseñado para funcionar con docker.

# Implementación con Docker

Se adjunta un ```Dockerfile``` que compila y genera una imagen con el servicio ejecutable.
Se configura para que el servicio busque las credenciales dentro de la carpeta ```credential``` en el home del usuario ```appuser``` ```/app```

Se adjunta un archivo ```docker-compose.yml``` con un detalle de las caracteristicas necesarias para desplegar el servicio.

Se fijan algunas varaibles de entorno para indicar los id de las spreadsheets, las paginas y los rangos que cada modulo requiere para parsear la informacion.

El servidor debe tener:

- git
- Docker

> Se debe tener un par de credencial-token listos para correr el servicio

El procedimiento es:

ir a un directorio ddesde donde se ejecutara el servicio, por ejemplo:

```bash
cd /home/user/Programas
```

clonar el proyecto desde [Gitlab](https://gitlab.com/vyra/almercadito/almercadito-api-restful)

```bash
git clone https://gitlab.com/vyra/almercadito/almercadito-api-restful
```

Ingresar

```bash
cd almercadito-api-restful
```

editar el archivo ```docker-compose.yml``` segun sus necesidades

```bash
nano docker-compose.yml
```

> Recuerde ajustar las varaibles de entorno y el puerto a utilizar. Internamente el servicio interactuara con el puerto 8080, eso queda siempre igual.

Guardar los cambios. Antes de ejecutar asegurece de que tenga las credenciales dentro de la carpeta credential. Una forma de hacerlo es creando los archivos por defecto y copiando el contenido con los valores necesarios

```bash
cd credential
touch credential.json
touch token.json
echo "........" > credential.json
echo "........" > token.json
cd ..
```

Por ultimo solo queda ejecutar el servicio con docker compose

```bash
docker-compose -d up
```

## Credenciales y acceso

Antes de iniciar el servicio, se debe contar con las credenciales y token validos. Para esto se brinda un binario ´´´login´´´ que permite generar el token de acceso necesario.

Para usarlo es necesario contar con un archivo de credencial de cliente otorgado por google. Se puede obtener gestionando las credenciales con Google Console.

Para generar el token se debe ejecutar por consola login, por ejemplo

```bash
./login -credential credential.json -token token.json -workdir ./
```

Donde

- *credential.json* es el nombre del archivo de la credencial de cliente otorgada por google
- *token* es el nombre del archivo con el que se guardara el token solicitado
- *workdir* es la carpeta donde se encuentra el archivo de credencial y donde se guardara el archivo con el token

Se mostrará en la salida del promtp un enlace que debera abrir en un navegador para otorgar los permisos necesarios con una cuenta de Google. Al final del proceso se le entregara un hash que debera copiar y pegar en el promtp. Tras presionar enter, se realizaran las comprobaciones necesarias y se generara un archivo token con el nombre dado en el argumento y dentro de la carpeta de trabajo.

Para mas información del uso de login ejecute

```bash
./login -help
```
