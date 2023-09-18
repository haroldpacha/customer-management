# API Rest de Gestión de Clientes

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Linkedin](https://img.shields.io/badge/linkedin-Harold%20Pacha-blue)](https://discord.gg/TFEZ7j4MZd)

Prueba Técnica - CRUD de Clientes en GO

## Cómo utilizar
Esta guía describe los pasos a seguir para poder ejecutar el proyecto en un entorno local.

Renombre el archivo `.env.example` por `.env` con la siguiente información: 

```bash
DB_USER=root
DB_PASS=test
DB_HOST=db
DB_NAME=golang_api

JWT_SECRET=jwtsecret123

REDIS_HOST=cache
REDIS_PORT=6379
REDIS_PASS=""
```
Luego ejecute el siguiente comando para levantar el proyecto en local y le mostrará los end-points creados.

```bash
docker-compose up -d --build
```

> Use el archivo `docs/Rest-Api-Customer-Management.postman_collection.json` probar con postman los servicios.

> Por defecto se crea la migracion de departamentos

## Informes

Este proyecto ha sido ejecutado con las siguientes versiones:

| Programa  | Version |
| ------------- |:-------------:|
| GO      | v1.21.1-alpine     |
| mariadb      | latest     |
| redis      | alpine     |

## Autor

[Harold Pacha](mailto:haroldpacha.rm@gmail.com)

## Licencia

Este proyecto está bajo la licencia MIT.