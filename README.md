
# Prueba técnica Enerbit

API rest para administrar medidores de energía.


## Variables de entorno

Para correr este proyecto, necesita añadir las siguientes variables de entorno a su archivo .env

`HOST`

`PORT`

`USER`

`PASS`

`DB`


## API

#### Swagger doc

```http
  GET /docs/index.html
```

#### Trae todos los registros

```http
  GET /api/v1/listemeter
```

#### Trae todos los registros deshabilidados

```http
  GET /api/v1/listdisableemeter
```

#### Trae el ultimo medidor instalado

```http
  GET /api/v1/lastemeter/${brand}/${serial}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `brand`      | `string` | **Required**. Marca del medidor |
| `serial`      | `string` | **Required**. Serial del medidor |

#### Crea un nuevo medidor

```http
  POST /api/v1/newemeter
```

```json
{
  "address": "string",
  "brand": "string",
  "installationDate": "string",
  "lines": 0,
  "serial": "string"
}
```

#### Elimina un medidor

```http
  POST /api/v1/deleteemeter
```

```json
{
  "id": "string"
}
```

#### Actualiza estado de un medidor

```http
  POST /api/v1/updateemeterstatus
```

```json
{
  "id": "string",
  "status": true
}
```
## Anexo

No se pudo implementar base de datos Redis.

