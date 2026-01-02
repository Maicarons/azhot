<div style="text-align: center;">

# azhot

<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/3199/3199306.png" alt="Logo" width="128" height="128" />
</p>

<p align="center">
  <img src="banner.jpg" alt="Banner" style="max-width:100%;height:auto;" />
</p>

[![Versão Go](https://img.shields.io/badge/Go-%3E%3D1.18-blue)](https://golang.org/)
[![Licença](https://img.shields.io/github/license/maicarons/azhot)](LICENSE)
[![Status da compilação](https://img.shields.io/badge/build-passing-brightgreen)](https://golang.org/)
[![Relatório Go](https://goreportcard.com/badge/github.com/maicarons/azhot)](https://goreportcard.com/report/github.com/maicarons/azhot)

</div>

> Um serviço de agregação que fornece APIs de pesquisas populares para as principais plataformas

## 📖 Índice

- [Introdução do projeto](#introdução-do-projeto)
- [Recursos](#recursos)
- [Plataformas suportadas](#plataformas-suportadas)
- [Início rápido](#início-rápido)
- [Uso da API](#uso-da-api)
- [Servidor MCP](#servidor-mcp)
- [Desenvolvimento e contribuição](#desenvolvimento-e-contribuição)
- [Licença](#licença)
- [Feedback de problemas](#feedback-de-problemas)

## Introdução do projeto

`azhot` é um serviço API que agrega dados de pesquisas populares das principais plataformas, fornecendo uma interface unificada para acessar conteúdo de pesquisas populares de várias plataformas. O projeto é desenvolvido na linguagem Go e construído com base no framework Fiber, suportando a obtenção em tempo real de dados de classificação de pesquisas populares das principais plataformas.

## Recursos

- 🚀 Interface API unificada para obter dados de pesquisas populares das principais plataformas
- ⚡ Alto desempenho, desenvolvido com `Go`+`Fiber v2`, com mecanismo de cache nativo + controle de acesso
- 🔄 Atualização agendada de dados de pesquisas populares para o banco de dados [Suporta SQLite + MySQL + Extensível para outros bancos de dados]
- 📚 [Documentação API Swagger](https://github.com/maicarons/azhot/blob/main/docs/swagger.yaml)
- 🌐 Design de API RESTful
- 📦 Inclui exemplo de [Frontend](/frontend)
- 🔌 Suporta envio de dados em tempo real via WebSocket
- 🤖 **Novo** Suporta Protocolo de Contexto de Modelo de IA (MCP)

## Estrutura do projeto
```
azhot/
├── all/                 # Código de todas as funcionalidades
├── app/                 # Código do programa principal
├── config/              # Leitura de arquivos de configuração
├── docs/                # Documentação API Swagger
├── model/               # Modelos de banco de dados
├── mcp/                 # Servidor de Protocolo de Contexto de Modelo de IA
├── router/              # Configuração de roteamento
├── service/             # Lógica de negócios
├── websocket/           # Funcionalidade WebSocket
├── frontend/            # Arquivos de modelo
├── .env                 # Variáveis de ambiente
├── Dockerfile           # Arquivo de construção Docker
├── go.mod               # Definição do módulo Go
├── main.go              # Arquivo do programa principal
└── README.md            # Documentação do projeto
```

## Plataformas suportadas

| Nome | Nome da rota | Disponibilidade |
|:----:|:------:|:------:|
| 360doc | 360doc | ✅ |
| Busca 360 | 360search | ✅ |
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
| Hoje na História | historytoday | ✅ |
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

## Início rápido

### Requisitos de ambiente

- Go >= 1.18
- MySQL (Opcional, para armazenamento de dados)

### Passos de instalação

1. Clonar o projeto
```bash
git clone https://github.com/maicarons/azhot.git
cd azhot
```

2. Instalar dependências
```bash
go mod tidy
```

3. Configurar ambiente
```bash
# Copiar arquivo de configuração
cp .env.example .env
# Editar arquivo de configuração
vim .env
```

4. Gerar documentação da API
```bash
swag init
```

5. Executar o projeto
```bash
# Executar em modo de desenvolvimento
make dev

# Ou compilar e executar
make run
```

### Executando com Docker

```bash
# Compilar imagem
docker build -t azhot .

# Executar contêiner
docker run -d -p 8080:8080 azhot
```

## Uso da API

### API HTTP

#### Obter lista de todas as plataformas

```
GET /list
```

Recuperar informações de todas as plataformas suportadas.

#### Obter pesquisas populares para uma plataforma específica

```
GET /{platform}
```

Por exemplo, para obter pesquisas populares do Zhihu:
```
GET /zhihu
```

### API WebSocket

O projeto suporta envio de dados em tempo real via WebSocket, fornecendo a mesma estrutura de roteamento da API HTTP.

#### Ponto de extremidade WebSocket geral

```
ws://localhost:8080/ws
```

Após a conexão, você pode enviar mensagens para se inscrever ou solicitar dados específicos de plataforma.

#### Ponto de extremidade WebSocket específico de plataforma

```
ws://localhost:8080/ws/{platform}
```

Por exemplo, conectando ao WebSocket de pesquisas populares do Baidu:
```
ws://localhost:8080/ws/baidu
```

#### Formato de mensagem WebSocket

```json
{
  "type": "subscribe|request|ping",
  "source": "Nome da plataforma, como baidu, zhihu, etc.",
  "data": {}
}
```

- `subscribe`: Inscrever-se nos dados em tempo real de uma plataforma específica
- `request`: Solicitar dados únicos
- `ping`: Mensagem de heartbeat

#### Lista de pontos de extremidade WebSocket

- Ponto de extremidade geral: `ws://localhost:8080/ws`
- Baidu: `ws://localhost:8080/ws/{platform}`
- Agregação de todas as plataformas: `ws://localhost:8080/ws/all`
- Lista de plataformas: `ws://localhost:8080/ws/list`
- API de consulta histórica:
  - `ws://localhost:8080/ws/history/{source}` - Obter todos os dados históricos para uma plataforma especificada
  - `ws://localhost:8080/ws/history/{source}/{date}` - Obter todos os dados horários para uma plataforma e data especificadas
  - `ws://localhost:8080/ws/history/{source}/{date}/{hour}` - Obter dados históricos para uma plataforma, data e hora especificadas
- E todos os outros pontos de extremidade WebSocket correspondentes às APIs HTTP

### Formato de resposta da API

```json
{
  "code": 200,
  "icon": "https://static.zhihu.com/static/favicon.ico",
  "message": "zhihu",
  "obj": [
    {
      "index": 1,
      "title": "Saudações de Ano Novo 2026",
      "url": "https://www.zhihu.com/search?q=Saudações de Ano Novo 2026"
    },
    // ...
    {
      "index": 12,
      "title": "Usuários do nordeste descobrem o rato 'Xiao Biga'",
      "url": "https://www.zhihu.com/search?q=Usuários do nordeste descobrem o rato 'Xiao Biga'"
    }
  ]
}
```

## Servidor MCP

O projeto agora integra um servidor de Protocolo de Contexto de Modelo de IA (MCP), permitindo que modelos de IA e assistentes inteligentes acessem dados de pesquisas populares através de um protocolo padronizado.

### Recursos

- **Interface de ferramentas padronizada**: Fornece lista de ferramentas MCP padronizada e interface de execução
- **Acesso a dados de pesquisas populares**: Suporta a obtenção de dados de pesquisas populares para cada plataforma através de ferramentas
- **Consulta de dados históricos**: Suporta consulta de dados históricos de pesquisas populares
- **Múltiplos modos de implantação**: Suporta modos de implantação HTTP e STDIO

### Habilitar servidor MCP

Configure as seguintes opções no arquivo `.env`:

```env
MCP_STDIO_ENABLED=true      # Habilitar servidor MCP STDIO
MCP_HTTP_ENABLED=true       # Habilitar servidor MCP HTTP
MCP_PORT=8081               # Porta do servidor MCP HTTP
```

### Lista de ferramentas MCP

- `get_hot_search`: Obter dados de pesquisas populares para uma plataforma especificada
- `get_all_hot_search`: Obter dados agregados de pesquisas populares para todas as plataformas
- `get_history_data`: Obter dados históricos de pesquisas populares para uma plataforma especificada

### Pontos de extremidade MCP

- `/mcp/tools` - Obter lista de ferramentas disponíveis
- `/mcp/tool/execute` - Executar ferramenta especificada
- `/mcp/prompts` - Obter lista de prompts disponíveis
- `/mcp/ping` - Ponto de extremidade de verificação de integridade
- `/mcp/.well-known/mcp-info` - Metadados do servidor MCP

### Exemplo de uso

Chamando a ferramenta MCP via HTTP:
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

Para mais detalhes, por favor consulte a [Documentação do Servidor MCP](mcp/README.md).

## Desenvolvimento e contribuição

Nós damos boas-vindas a qualquer forma de contribuição! Se você quiser contribuir para o projeto, por favor siga estes passos:

1. Faça um fork deste projeto
2. Crie um branch de funcionalidade (`git checkout -b feature/AmazingFeature`)
3. Faça commit das alterações (`git commit -m 'Add some AmazingFeature'`)
4. Faça push para o branch (`git push origin feature/AmazingFeature`)
5. Crie um Pull Request

### Desenvolvimento local

```bash
# Executar testes
dev.sh # Usando Air como ferramenta de depuração com recarga automática
```

## Sistema de compilação CMake

O projeto agora suporta compilação com CMake, suportando as plataformas Windows e Linux.

### Comandos de compilação

```bash
# Compilar para plataforma atual
mkdir build && cd build
cmake ..
cmake --build . --target build

# Executar
cmake --build . --target run

# Executar em modo de desenvolvimento
cmake --build . --target dev

# Compilação multiplataforma (plataformas predefinidas)
cmake --build . --target build-platform-linux
cmake --build . --target build-platform-windows
cmake --build . --target build-platform-darwin
cmake --build . --target build-platform-linux-arm64
cmake --build . --target build-platform-windows-arm64

# Compilação multiplataforma (usando script)
# Linux/macOS:
./build_platform.sh linux
./build_platform.sh windows
./build_platform.sh darwin

# Windows:
build_platform.bat linux
build_platform.bat windows
build_platform.bat darwin

# Empacotar (criar pacotes zip para todas as plataformas suportadas)
cmake --build . --target package

# Limpar artefatos de compilação
cmake --build . --target azhot_clean

# Executar testes
cmake --build . --target test

# Executar todos os testes
cmake --build . --target test-all

# Formatar código
cmake --build . --target fmt

# Organizar dependências
cmake --build . --target tidy

# Análise estática
cmake --build . --target staticcheck

# Compilar versão CI (sem gerar documentação swagger)
cmake --build . --target build-ci
```

## Licença

Este projeto está licenciado sob a licença AGPL-3.0 - veja o arquivo [LICENÇA](LICENSE) para detalhes.

## Feedback de problemas

Se você encontrar problemas ou tiver sugestões durante o uso do projeto, sinta-se à vontade para submeter uma Issue ou Pull Request.

- 🐛 [Relatório de problemas](https://github.com/maicarons/azhot/issues)
- ✨ [Solicitação de funcionalidades](https://github.com/maicarons/azhot/issues)

---

> 🌟 Se este projeto foi útil para você, por favor nos dê uma estrela! Isso seria o maior apoio para nós!