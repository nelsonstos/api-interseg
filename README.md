# Proyecto Interseguro Challenge - Matrices y Factorizacion QR

## Descripción

El proyecto **Interseguro Challenge** es una aplicación que incluye varios servicios basados en Docker, como `api-go-matrix` y `api-node-matrix`. El proyecto está diseñado para facilitar la integración y gestión de servicios utilizando Docker Compose, con soporte para bases de datos y configuraciones específicas.

## Estructura del Proyecto

El proyecto tiene la siguiente estructura:
```sh
interseg/
├── api-go-matrix/
│   ├── cmd/          # Comandos principales y configuración de la aplicación
│   ├── config/
│   ├── controllers/  # Controladores de las rutas de la API
│   ├── services/     # Lógica de negocio y servicios
│   ├── models/       # Definición de modelos y esquemas de datos
│   ├── repositories/ # Implementaciones concretas de las interfaces (bases de datos, APIs)
│   ├── .env          # Variables de entorno
│   ├── Dockerfile    # Dockerfile para contenerizar el servicio
│   └── ...
├── api-node-matrix/
│   ├── api/
│   │   ├── controllers/  # Controladores de las rutas de la API
│   │   ├── services/     # Lógica de negocio y servicios
│   │   ├── models/       # Definición de modelos y esquemas de datos
│   │   ├── repositories/ # Acceso a datos y operaciones CRUD
│   │   └── ...
│   ├── .env              # Variables de entorno
│   ├── Dockerfile        # Dockerfile para contenerizar el servicio
│   └── ...
├── docker-compose.yml    # Configuración de Docker Compose para orquestar los servicios
└── .env                  # Variables de entorno globales

```

## Requisitos

Para ejecutar el proyecto, asegúrate de tener las siguientes herramientas instaladas en tu sistema:

- **Go 1.22.5**: Necesario para construir y ejecutar el servicio `api-go-matrix`.
- **Node.js 20.15.1**: Requerido para el desarrollo y ejecución de `api-node-matrix`.
- **PostgreSQL**: Utilizado como base de datos en el proyecto. Asegúrate de tener acceso a una instancia de PostgreSQL.
- **Docker**: Utilizado para contenerizar los servicios y simplificar la gestión del entorno de desarrollo.
- **Docker Compose**: Herramienta para definir y ejecutar aplicaciones multi-contenedor. Utilizado para orquestar los servicios del proyecto.


## Configuración

Sigue estos pasos para configurar el proyecto en tu entorno local:

### 1. Clonar el Repositorio

Primero, clona el repositorio del proyecto en tu máquina local:

```sh
git clone  https://github.com/nelsonstos/interseg.git
cd interseg
```

### 2. Configuración de la variable de entorno
#### 2.1 Configuración global
- Copia el archivo .env de ejemplo al archivo .env real y ajusta las variables según sea necesario:
```sh
cp .env.example .env
```
- Edita el archivo .env con las configuraciones necesarias. Asegúrate de establecer correctamente las credenciales y configuraciones de la base de datos PostgreSQL, entre otras.
#### 2.2 Configuración Específica para Servicios
- api-go-matrix:
Copia el archivo .env de ejemplo dentro del directorio api-go-matrix y ajústalo:
```sh
cp api-go-matrix/.env.example api-go-matrix/.env
```
- api-node-matrix:
Copia el archivo .env de ejemplo dentro del directorio api-node-matrix y ajústalo:
```sh
cp api-node-matrix/.env.example api-node-matrix/.env
```
###  Construir las Imágenes de Docker
Construye las imágenes de Docker para los servicios del proyecto:
```sh
docker-compose up --build
```

## Uso
```sh
docker-compose up -d
```

## API Examples
### Matrix
Crea la matriz y facorización QR desde ´api-go-matrix´
- Request:
```sh
curl --location 'http://localhost:8080/api/v1/matrix' \
--header 'Content-Type: application/json' \
--data '{
    "matrix":[
        [1, 2, 6], 
        [3, 4, 8],
        [6, 7, 10]
        ]
}'
```
- Response:
```sh
{
    "status": "success",
    "message": "Factorization and register successfully!",
    "data": {
        "id": 37,
        "matrix_id": 37,
        "q_matrix": "[[-0.1474419561548972,-0.8610567718228911,-0.4866642633922886],[-0.44232586846469135,-0.38269189858795266,0.8111071056538122],[-0.8846517369293827,0.3348554112644583,-0.3244428422615245]]",
        "r_matrix": "[[-6.782329983125268,-8.25674954467424,-13.269776053940742],[0,-0.9088932591463842,-4.879321706996383],[0,0,0.32444284226151954]]",
        "created_at": "2024-07-30T13:58:48.255319184Z"
    }
}
```
### Statistics
Usa el servicio api-go-matrix para crear, factoriza la matriz y obtner las estadisticas.
- Request:
```sh
curl --location 'http://localhost:3000/api/v1/matrix/statistics' \
--header 'Content-Type: application/json' \
--data '{
    "matrix":[
        [1, 2, 6], 
        [3, 4, 8],
        [6, 7, 10]
        ]
}'
```
- Response:
```sh
{
    "status": "success",
    "message": "Statistic stored successfully!",
    "data": {
        "id": 37,
        "combinedStatistic": {
            "max": 0.8111071056538122,
            "min": -13.269776053940742,
            "avg": -2.008663362573159,
            "sum": -36.15594052631686
        },
        "isQMatrixDiagonal": false,
        "isRMatrixDiagonal": false,
        "factorization_id": 38,
        "created_at": "2024-07-30T14:08:43.443209425Z"
    }
}
```

## Contribución
If you wish to contribute to this project, follow these steps:
1. Fork the repository
2. Create a new branch: `git checkout -b feature-nueva`
3. Make your changes and commit them: `git commit -am 'Agrega una nueva característica'`
4. Push to the branch: `git push origin feature-nueva`
5. Submit a pull request
