<div style="text-align: center;">

# azhot

<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/3199/3199306.png" alt="Logo" width="128" height="128" />
</p>

<p align="center">
  <img src="banner.jpg" alt="Banner" style="max-width:100%;height:auto;" />
</p>

[![Versión de Go](https://img.shields.io/badge/Go-%3E%3D1.18-blue)](https://golang.org/)
[![Licencia](https://img.shields.io/github/license/maicarons/azhot)](LICENSE)
[![Estado de compilación](https://img.shields.io/badge/build-passing-brightgreen)](https://golang.org/)
[![Informe Go](https://goreportcard.com/badge/github.com/maicarons/azhot)](https://goreportcard.com/report/github.com/maicarons/azhot)

</div>

> Un servicio de agregación que proporciona APIs de búsqueda popular para las principales plataformas

## 📖 Tabla de Contenidos

- [Introducción del proyecto](#introducción-del-proyecto)
- [Características](#características)
- [Plataformas soportadas](#plataformas-soportadas)
- [Inicio rápido](#inicio-rápido)
- [Uso de la API](#uso-de-la-api)
- [Servidor MCP](#servidor-mcp)
- [Desarrollo y contribución](#desarrollo-y-contribución)
- [Licencia](#licencia)
- [Comentarios sobre problemas](#comentarios-sobre-problemas)

## Introducción del proyecto

`azhot` es un servicio API que agrega datos de búsqueda popular de las principales plataformas, proporcionando una interfaz unificada para acceder al contenido de búsqueda popular de varias plataformas. El proyecto está desarrollado en lenguaje Go y construido sobre el framework Fiber, soportando la obtención en tiempo real de datos de clasificación de búsquedas populares de las principales plataformas.

## Características

- 🚀 Interfaz API unificada para obtener datos de búsqueda popular de las principales plataformas
- ⚡ Alta performance, desarrollado con `Go`+`Fiber v2`, con mecanismo de caché nativo + control de acceso
- 🔄 Actualización programada de datos de búsqueda popular a la base de datos [Soporta SQLite + MySQL + Extensible a otras bases de datos]
- 📚 [Documentación API Swagger](https://github.com/maicarons/azhot/blob/main/docs/swagger.yaml)
- 🌐 Diseño de API RESTful
- 📦 Incluye ejemplo de [Frontend](/frontend)
- 🔌 Soporta envío de datos en tiempo real mediante WebSocket
- 🤖 **Nuevo** Soporta el protocolo de contexto de modelo de IA (MCP)

## Estructura del proyecto
```
azhot/
├── all/                 # Código de todas las funcionalidades
├── app/                 # Código del programa principal
├── config/              # Lectura de archivos de configuración
├── docs/                # Documentación API Swagger
├── model/               # Modelos de base de datos
├── mcp/                 # Servidor de protocolo de contexto de modelo de IA
├── router/              # Configuración de enrutamiento
├── service/             # Lógica de negocio
├── websocket/           # Funcionalidad WebSocket
├── frontend/            # Archivos de plantilla
├── .env                 # Variables de entorno
├── Dockerfile           # Archivo de construcción Docker
├── go.mod               # Definición del módulo Go
├── main.go              # Archivo del programa principal
└── README.md            # Documentación del proyecto
```

## Plataformas soportadas

| Nombre | Nombre de ruta | Disponibilidad |
|:----:|:------:|:------:|
| 360doc | 360doc | ✅ |
| Búsqueda 360 | 360search | ✅ |
| AcFun | acfun | ✅ |
| Baidu | baidu | ✅ |
| Bilibili | bilibili | ✅ |
| CCTV | cctv | ✅ |
| CSDN | csdn | ✅ |
| Dongqiudi | dongqiudi | ✅ |
| Douban | douban | ✅ |
| Douyin | douyin | ✅ |
| GitHub | github | ✅ |
| National Geographic | guojiadili | ✅ |
| Hoy en la historia | historytoday | ✅ |
| Hupu | hupu | ✅ |
| IT Home | ithome | ✅ |
| Pear Video | lishipin | ✅ |
| Southern Weekly | nanfang | ✅ |
| Pengpai News | pengpai | ✅ |
| Tencent News | qqnews | ✅ |
| Quark | quark | ✅ |
| People's Daily Online | renmin | ✅ |
| Sogou | sougou | ✅ |
| Sohu | souhu | ✅ |
| Toutiao | toutiao | ✅ |
| V2EX | v2ex | ✅ |
| NetEase News | wangyinews | ✅ |
| Weibo | weibo | ✅ |
| Xinjing Daily | xinjingbao | ✅ |
| Zhihu | zhihu | ✅ |

## Inicio rápido

### Requisitos del entorno

- Go >= 1.18
- MySQL (Opcional, para almacenamiento de datos)

### Pasos de instalación

1. Clonar el proyecto
```bash
git clone https://github.com/maicarons/azhot.git
cd azhot
```

2. Instalar dependencias
```bash
go mod tidy
```

3. Configurar entorno
```bash
# Copiar archivo de configuración
cp .env.example .env
# Editar archivo de configuración
vim .env
```

4. Generar documentación de la API
```bash
swag init
```

5. Ejecutar el proyecto
```bash
# Ejecutar en modo desarrollo
make dev

# O construir y ejecutar
make run
```

### Ejecución con Docker

```bash
# Construir imagen
docker build -t azhot .

# Ejecutar contenedor
docker run -d -p 8080:8080 azhot
```

### Configuración de variables de entorno

El proyecto utiliza el archivo `.env` para la configuración. Aquí están las variables de entorno disponibles:

#### Configuración del servidor
- `SERVER_HOST`: Dirección del host del servidor, por defecto es `localhost`
- `SERVER_PORT`: Puerto del servidor, por defecto es `8080`
- `TLS_ENABLED`: Si habilitar TLS/HTTPS, por defecto es `false`
- `TLS_CERT_FILE`: Ruta del archivo de certificado TLS, requerido cuando `TLS_ENABLED` es `true`
- `TLS_KEY_FILE`: Ruta del archivo de clave privada TLS, requerido cuando `TLS_ENABLED` es `true`

#### Configuración de la base de datos
- `DB_TYPE`: Tipo de base de datos, soporta `sqlite` y `mysql`, por defecto es `sqlite`
- `MYSQL_DSN`: Cadena de conexión a la base de datos MySQL, efectiva cuando `DB_TYPE` es `mysql`

#### Configuración MCP
- `MCP_STDIO_ENABLED`: Si habilitar el servidor MCP STDIO, por defecto es `false`
- `MCP_HTTP_ENABLED`: Si habilitar el servidor MCP HTTP, por defecto es `false`
- `MCP_PORT`: Puerto del servidor MCP HTTP, por defecto es `8081`

#### Configuración de depuración
- `DEBUG`: Si habilitar el modo depuración, por defecto es `false`

#### Configuración CORS
- `CORS_ALLOW_ORIGINS`: Orígenes permitidos para solicitudes cross-origin, múltiples orígenes separados por comas, por defecto está vacío para permitir todos los orígenes (recomendado configurar orígenes específicos solo en entorno de producción)

## Uso de la API

### API HTTP

#### Obtener lista de todas las plataformas

```
GET /list
```

Recuperar información de todas las plataformas soportadas.

#### Obtener búsqueda popular para una plataforma específica

```
GET /{platform}
```

Por ejemplo, para obtener la búsqueda popular de Zhihu:
```
GET /zhihu
```

### API WebSocket

El proyecto soporta envío de datos en tiempo real mediante WebSocket, proporcionando la misma estructura de enrutamiento que la API HTTP.

#### Punto final WebSocket general

```
ws://localhost:8080/ws
```

Después de la conexión, puedes enviar mensajes para suscribirte o solicitar datos específicos de una plataforma.

#### Punto final WebSocket específico de plataforma

```
ws://localhost:8080/ws/{platform}
```

Por ejemplo, conectarse al WebSocket de búsqueda popular de Baidu:
```
ws://localhost:8080/ws/baidu
```

#### Formato de mensaje WebSocket

```json
{
  "type": "subscribe|request|ping",
  "source": "Nombre de la plataforma, como baidu, zhihu, etc.",
  "data": {}
}
```

- `subscribe`: Suscribirse a los datos en tiempo real de una plataforma específica
- `request`: Solicitar datos puntuales
- `ping`: Mensaje de latido de corazón

#### Lista de puntos finales WebSocket

- Punto final general: `ws://localhost:8080/ws`
- Baidu: `ws://localhost:8080/ws/{platform}`
- Agregación de todas las plataformas: `ws://localhost:8080/ws/all`
- Lista de plataformas: `ws://localhost:8080/ws/list`
- API de consulta histórica:
  - `ws://localhost:8080/ws/history/{source}` - Obtener todos los datos históricos para una plataforma especificada
  - `ws://localhost:8080/ws/history/{source}/{date}` - Obtener todos los datos horarios para una plataforma y fecha especificadas
  - `ws://localhost:8080/ws/history/{source}/{date}/{hour}` - Obtener datos históricos para una plataforma, fecha y hora especificadas
- Y todos los demás puntos finales WebSocket correspondientes a las APIs HTTP

### Formato de respuesta de la API

```json
{
  "code": 200,
  "icon": "https://static.zhihu.com/static/favicon.ico",
  "message": "zhihu",
  "obj": [
    {
      "index": 1,
      "title": "Saludos de Año Nuevo 2026",
      "url": "https://www.zhihu.com/search?q=Saludos de Año Nuevo 2026"
    },
    // ...
    {
      "index": 12,
      "title": "Usuarios del noreste descubren un ratón 'Xiao Biga'",
      "url": "https://www.zhihu.com/search?q=Usuarios del noreste descubren un ratón 'Xiao Biga'"
    }
  ]
}
```

## Servidor MCP

El proyecto ahora integra un servidor de protocolo de contexto de modelo de IA (MCP), permitiendo que los modelos de IA y asistentes inteligentes accedan a datos de búsqueda popular a través de un protocolo estandarizado.

### Características

- **Interfaz de herramientas estandarizada**: Proporciona lista de herramientas MCP estandarizada e interfaz de ejecución
- **Acceso a datos de búsqueda popular**: Soporta la obtención de datos de búsqueda popular para cada plataforma mediante herramientas
- **Consulta de datos históricos**: Soporta la consulta de datos históricos de búsqueda popular
- **Múltiples modos de despliegue**: Soporta modos de despliegue HTTP y STDIO

### Habilitar servidor MCP

Configure las siguientes opciones en el archivo `.env`:

```env
MCP_STDIO_ENABLED=true      # Habilitar servidor MCP STDIO
MCP_HTTP_ENABLED=true       # Habilitar servidor MCP HTTP
MCP_PORT=8081               # Puerto del servidor MCP HTTP
```

### Lista de herramientas MCP

- `get_hot_search`: Obtener datos de búsqueda popular para una plataforma especificada
- `get_all_hot_search`: Obtener datos agregados de búsqueda popular para todas las plataformas
- `get_history_data`: Obtener datos históricos de búsqueda popular para una plataforma especificada

### Puntos finales MCP

- `/mcp/tools` - Obtener lista de herramientas disponibles
- `/mcp/tool/execute` - Ejecutar herramienta especificada
- `/mcp/prompts` - Obtener lista de indicaciones disponibles
- `/mcp/ping` - Punto final de verificación de salud
- `/mcp/.well-known/mcp-info` - Metadatos del servidor MCP

### Ejemplo de uso

Llamada a la herramienta MCP mediante HTTP:
```bash
curl -X POST http://localhost:8080/mcp/tool/execute \
  -H "Content-Type: application/json" \
  -d '{
    "method": "tool/execute",
    "params": {
      "name": "get_hot_search",
      "arguments": {
        "platform": "zhihu"
      }
    },
    "id": "req-1",
    "jsonrpc": "2.0"
  }'
```

Para más detalles, por favor consulte la [Documentación del servidor MCP](mcp/README.md).

## Desarrollo y contribución

¡Damos la bienvenida a cualquier forma de contribución! Si desea contribuir al proyecto, por favor siga estos pasos:

1. Haga fork de este proyecto
2. Cree una rama de funcionalidad (`git checkout -b feature/AmazingFeature`)
3. Confirme los cambios (`git commit -m 'Add some AmazingFeature'`)
4. Haga push a la rama (`git push origin feature/AmazingFeature`)
5. Cree una Pull Request

### Desarrollo local

```bash
# Ejecutar pruebas
dev.sh # Usando Air como herramienta de depuración con recarga en caliente
```

## Sistema de construcción CMake

El proyecto ahora soporta la construcción con CMake, soportando las plataformas Windows y Linux.

### Comandos de construcción

```bash
# Construir para la plataforma actual
mkdir build && cd build
cmake ..
cmake --build . --target build

# Ejecutar
cmake --build . --target run

# Ejecutar en modo desarrollo
cmake --build . --target dev

# Construcción multiplataforma (plataformas predefinidas)
cmake --build . --target build-platform-linux
cmake --build . --target build-platform-windows
cmake --build . --target build-platform-darwin
cmake --build . --target build-platform-linux-arm64
cmake --build . --target build-platform-windows-arm64

# Construcción multiplataforma (usando script)
# Linux/macOS:
./build_platform.sh linux
./build_platform.sh windows
./build_platform.sh darwin

# Windows:
build_platform.bat linux
build_platform.bat windows
build_platform.bat darwin

# Empaquetar (crear paquetes zip para todas las plataformas soportadas)
cmake --build . --target package

# Limpiar artefactos de construcción
cmake --build . --target azhot_clean

# Ejecutar pruebas
cmake --build . --target test

# Ejecutar todas las pruebas
cmake --build . --target test-all

# Formatear código
cmake --build . --target fmt

# Organizar dependencias
cmake --build . --target tidy

# Análisis estático
cmake --build . --target staticcheck

# Construir versión CI (sin generar documentación swagger)
cmake --build . --target build-ci
```

## Licencia

Este proyecto está licenciado bajo la licencia AGPL-3.0 - consulte el archivo [LICENCIA](LICENSE) para más detalles.

## Comentarios sobre problemas

Si encuentra problemas o tiene sugerencias mientras usa el proyecto, no dude en enviar un Issue o Pull Request.

- 🐛 [Reporte de problemas](https://github.com/maicarons/azhot/issues)
- ✨ [Solicitud de funcionalidades](https://github.com/maicarons/azhot/issues)

---

> 🌟 ¡Si este proyecto le ha sido útil, por favor danos una estrella! ¡Esto sería el mayor apoyo para nosotros!