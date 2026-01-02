<div style="text-align: center;">

# azhot

<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/3199/3199306.png" alt="Logo" width="128" height="128" />
</p>

[![Version Go](https://img.shields.io/badge/Go-%3E%3D1.18-blue)](https://golang.org/)
[![Licence](https://img.shields.io/github/license/maicarons/azhot)](LICENSE)
[![Statut de construction](https://img.shields.io/badge/build-passing-brightgreen)](https://golang.org/)
[![Rapport Go](https://goreportcard.com/badge/github.com/maicarons/azhot)](https://goreportcard.com/report/github.com/maicarons/azhot)

</div>

> Un service d'agrégation qui fournit des API de recherche populaire pour les principales plateformes

## 📖 Table des matières

- [Introduction du projet](#introduction-du-projet)
- [Fonctionnalités](#fonctionnalités)
- [Plateformes prises en charge](#plateformes-prises-en-charge)
- [Démarrage rapide](#démarrage-rapide)
- [Utilisation de l'API](#utilisation-de-lapi)
- [Serveur MCP](#serveur-mcp)
- [Développement et contribution](#développement-et-contribution)
- [Licence](#licence)
- [Retour d'information](#retour-dinformation)

## Introduction du projet

`azhot` est un service API qui agrège les données de recherche populaire des principales plateformes, fournissant une interface unifiée pour accéder aux contenus de recherche populaire de diverses plateformes. Le projet est développé en langage Go et construit sur le framework Fiber, prenant en charge la récupération en temps réel des données de classement des recherches populaires des principales plateformes.

## Fonctionnalités

- 🚀 Interface API unifiée pour récupérer les données de recherche populaire des principales plateformes
- ⚡ Haute performance, développé avec `Go`+`Fiber v2`, avec mécanisme de cache natif + contrôle d'accès
- 🔄 Mise à jour planifiée des données de recherche populaire vers la base de données [Prend en charge SQLite + MySQL + Extensible à d'autres bases de données]
- 📚 [Documentation API Swagger](https://github.com/maicarons/azhot/blob/main/docs/swagger.yaml)
- 🌐 Conception d'API RESTful
- 📦 Inclut un exemple de [Frontend](/frontend)
- 🔌 Prend en charge la diffusion de données en temps réel via WebSocket
- 🤖 **Nouveau** Prend en charge le protocole de contexte de modèle d'IA (MCP)

## Structure du projet
```
azhot/
├── all/                 # Code de toutes les fonctionnalités
├── app/                 # Code du programme principal
├── config/              # Lecture des fichiers de configuration
├── docs/                # Documentation de l'API Swagger
├── model/               # Modèles de base de données
├── mcp/                 # Serveur de protocole de contexte de modèle d'IA
├── router/              # Configuration du routage
├── service/             # Logique métier
├── websocket/           # Fonctionnalité WebSocket
├── frontend/            # Fichiers de modèle
├── .env                 # Variables d'environnement
├── Dockerfile           # Fichier de construction Docker
├── go.mod               # Définition du module Go
├── main.go              # Fichier du programme principal
└── README.md            # Documentation du projet
```

## Plateformes prises en charge

| Nom | Nom de la route | Disponibilité |
|:----:|:------:|:------:|
| 360doc | 360doc | ✅ |
| Recherche 360 | 360search | ✅ |
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
| Aujourd'hui dans l'histoire | historytoday | ✅ |
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

## Démarrage rapide

### Exigences d'environnement

- Go >= 1.18
- MySQL (Optionnel, pour le stockage des données)

### Étapes d'installation

1. Cloner le projet
```bash
git clone https://github.com/maicarons/azhot.git
cd azhot
```

2. Installer les dépendances
```bash
go mod tidy
```

3. Configurer l'environnement
```bash
# Copier le fichier de configuration
cp .env.example .env
# Éditer le fichier de configuration
vim .env
```

4. Générer la documentation de l'API
```bash
swag init
```

5. Exécuter le projet
```bash
# Exécuter en mode développement
make dev

# Ou construire puis exécuter
make run
```

### Exécution avec Docker

```bash
# Construire l'image
docker build -t azhot .

# Exécuter le conteneur
docker run -d -p 8080:8080 azhot
```

## Utilisation de l'API

### API HTTP

#### Obtenir la liste de toutes les plateformes

```
GET /list
```

Récupérer les informations de toutes les plateformes prises en charge.

#### Obtenir la recherche populaire pour une plateforme spécifique

```
GET /{plateforme}
```

Par exemple, pour obtenir la recherche populaire Zhihu :
```
GET /zhihu
```

### API WebSocket

Le projet prend en charge la diffusion de données en temps réel via WebSocket, fournissant la même structure de routage que l'API HTTP.

#### Point de terminaison WebSocket général

```
ws://localhost:8080/ws
```

Après la connexion, vous pouvez envoyer des messages pour vous abonner ou demander des données spécifiques à une plateforme.

#### Point de terminaison WebSocket spécifique à une plateforme

```
ws://localhost:8080/ws/{plateforme}
```

Par exemple, se connecter au WebSocket de recherche populaire Baidu :
```
ws://localhost:8080/ws/baidu
```

#### Format de message WebSocket

```json
{
  "type": "subscribe|request|ping",
  "source": "Nom de la plateforme, comme baidu, zhihu, etc.",
  "data": {}
}
```

- `subscribe`: S'abonner aux données en temps réel d'une plateforme spécifique
- `request`: Demander des données ponctuelles
- `ping`: Message de maintien de connexion

#### Liste des points de terminaison WebSocket

- Point de terminaison général : `ws://localhost:8080/ws`
- Baidu : `ws://localhost:8080/ws/{plateforme}`
- Agrégation de toutes les plateformes : `ws://localhost:8080/ws/all`
- Liste des plateformes : `ws://localhost:8080/ws/list`
- API de requête historique :
  - `ws://localhost:8080/ws/history/{source}` - Obtenir toutes les données historiques pour une plateforme spécifiée
  - `ws://localhost:8080/ws/history/{source}/{date}` - Obtenir toutes les données horaires pour une plateforme et une date spécifiées
  - `ws://localhost:8080/ws/history/{source}/{date}/{hour}` - Obtenir les données historiques pour une plateforme, une date et une heure spécifiées
- Et tous les autres points de terminaison WebSocket correspondant aux API HTTP

### Format de réponse de l'API

```json
{
  "code": 200,
  "icon": "https://static.zhihu.com/static/favicon.ico",
  "message": "zhihu",
  "obj": [
    {
      "index": 1,
      "title": "Souhaits de Nouvel An 2026",
      "url": "https://www.zhihu.com/search?q=Souhaits de Nouvel An 2026"
    },
    // ...
    {
      "index": 12,
      "title": "Les internautes du nord-est découvrent une souris 'Xiao Biga'",
      "url": "https://www.zhihu.com/search?q=Les internautes du nord-est découvrent une souris 'Xiao Biga'"
    }
  ]
}
```

## Serveur MCP

Le projet intègre désormais un serveur de protocole de contexte de modèle d'IA (MCP), permettant aux modèles d'IA et aux assistants intelligents d'accéder aux données de recherche populaire via un protocole standardisé.

### Fonctionnalités

- **Interface d'outils standardisée** : Fournit une liste d'outils MCP standardisée et une interface d'exécution
- **Accès aux données de recherche populaire** : Prend en charge la récupération des données de recherche populaire pour chaque plateforme via des outils
- **Requête de données historiques** : Prend en charge l'interrogation des données historiques de recherche populaire
- **Modes de déploiement multiples** : Prend en charge les modes de déploiement HTTP et STDIO

### Activation du serveur MCP

Configurez les options suivantes dans le fichier `.env` :

```env
MCP_STDIO_ENABLED=true      # Activer le serveur MCP STDIO
MCP_HTTP_ENABLED=true       # Activer le serveur MCP HTTP
MCP_PORT=8081               # Port du serveur MCP HTTP
```

### Liste des outils MCP

- `get_hot_search` : Obtenir les données de recherche populaire pour une plateforme spécifiée
- `get_all_hot_search` : Obtenir les données agrégées de recherche populaire pour toutes les plateformes
- `get_history_data` : Obtenir les données historiques de recherche populaire pour une plateforme spécifiée

### Points de terminaison MCP

- `/mcp/tools` - Obtenir la liste des outils disponibles
- `/mcp/tool/execute` - Exécuter l'outil spécifié
- `/mcp/prompts` - Obtenir la liste des invites disponibles
- `/mcp/ping` - Point de terminaison de vérification d'intégrité
- `/mcp/.well-known/mcp-info` - Métadonnées du serveur MCP

### Exemple d'utilisation

Appel de l'outil MCP via HTTP :
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

Pour plus de détails, veuillez vous référer à la [Documentation du serveur MCP](mcp/README.md).

## Développement et contribution

Nous accueillons toute forme de contribution ! Si vous souhaitez contribuer au projet, veuillez suivre ces étapes :

1. Forkez ce projet
2. Créez une branche de fonctionnalité (`git checkout -b feature/AmazingFeature`)
3. Validez vos modifications (`git commit -m 'Ajouter une fonctionnalité étonnante'`)
4. Poussez vers la branche (`git push origin feature/AmazingFeature`)
5. Créez une Pull Request

### Développement local

```bash
# Exécuter les tests
dev.sh # Utilisation d'Air comme outil de débogage avec rechargement à chaud
```

## Système de construction CMake

Le projet prend désormais en charge la construction avec CMake, prenant en charge les plateformes Windows et Linux.

### Commandes de construction

```bash
# Construire pour la plateforme actuelle
mkdir build && cd build
cmake ..
cmake --build . --target build

# Exécuter
cmake --build . --target run

# Exécuter en mode développement
cmake --build . --target dev

# Construction multiplateforme (plateformes prédéfinies)
cmake --build . --target build-platform-linux
cmake --build . --target build-platform-windows
cmake --build . --target build-platform-darwin
cmake --build . --target build-platform-linux-arm64
cmake --build . --target build-platform-windows-arm64

# Construction multiplateforme (en utilisant un script)
# Linux/macOS :
./build_platform.sh linux
./build_platform.sh windows
./build_platform.sh darwin

# Windows :
build_platform.bat linux
build_platform.bat windows
build_platform.bat darwin

# Empaqueter (créer des packages zip pour toutes les plateformes prises en charge)
cmake --build . --target package

# Nettoyer les artefacts de construction
cmake --build . --target azhot_clean

# Exécuter les tests
cmake --build . --target test

# Exécuter tous les tests
cmake --build . --target test-all

# Formater le code
cmake --build . --target fmt

# Gérer les dépendances
cmake --build . --target tidy

# Analyse statique
cmake --build . --target staticcheck

# Construire la version CI (sans générer la documentation swagger)
cmake --build . --target build-ci
```

## Licence

Ce projet est licencié sous la licence AGPL-3.0 - voir le fichier [LICENCE](LICENCE) pour plus de détails.

## Retour d'information

Si vous rencontrez des problèmes ou avez des suggestions pendant l'utilisation du projet, n'hésitez pas à soumettre un problème ou une Pull Request.

- 🐛 [Rapport de problème](https://github.com/maicarons/azhot/issues)
- ✨ [Demande de fonctionnalité](https://github.com/maicarons/azhot/issues)

---

> 🌟 Si ce projet vous est utile, veuillez nous donner une étoile ! Cela serait le plus grand soutien pour nous !