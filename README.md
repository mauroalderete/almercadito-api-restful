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
