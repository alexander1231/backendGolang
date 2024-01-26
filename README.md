# Getting Started

## Paso 1: Instalar la dependependencia de golang

```bash
go mod download
```

## Paso 2: Instalar Docker y instalar imagen de Postgres

```bash
docker run --name some-postgres -e POSTGRES_USER=fazt -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

## Paso 3: Creatar la base de datos llamada gorm

```bash
# ingresamos a la imagen de postgres en modo interantivo
docker exec -it some-postgres bash

# luego ingresamos a la base de datos y agregamos el passowrd previamente seteado
psql -U fazt --password

# creamos la base de datos
CREATE DATABASE gorm;
```

