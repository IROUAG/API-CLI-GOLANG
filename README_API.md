# Projet Golang: Ilies & Sylvain

Ce dépôt contient une API Golang avec une base de données PostgreSQL, organisée comme suit :

- `app/` : Contient le code de l'application Go et son Dockerfile.
- `postgres-setup/` : Contient le script SQL pour la configuration de la base de données PostgreSQL et son Dockerfile.
- `docker-compose.yml` : Fichier de configuration Docker Compose pour construire et exécuter l'ensemble de l'application.
- `.env` : Fichier des variables d'environnement pour stocker les paramètres de connexion à la base de données.

## Prérequis

- Docker
- Docker Compose

## Comment construire et exécuter l'API

1. Clonez le dépôt :

```bash
git clone https://gitlab.com/votre-username/golang_project_ilies_sylvain.git
cd golang_project_ilies_sylvain
```

2. Construisez et exécutez l'application en utilisant Docker Compose :

```bash
docker-compose up -d --build
```

Cette commande construira les conteneurs de l'application Go et de PostgreSQL et les exécutera ensemble. Votre application sera accessible à l'adresse http://localhost:8080.

Pour arrêter les conteneurs, appuyez sur Ctrl+C ou exécutez :

```bash
docker-compose down
```
## EndPoint de l'API

* `/users`: Gérer les utilisateurs (GET, POST, PUT, DELETE).
* `/roles`: Gérer les rôles pour les utilisateurs (GET, POST, PUT, DELETE).
* `/groups`: Gérer les groupes d'utilisateurs (GET, POST, PUT, DELETE).
* `/auth`: Gérer l'authentification des utilisateurs en utilisant JWT (POST).

### /users

* `GET /users`: Récupérer la liste des utilisateurs.
* `POST /users`: Créer un nouvel utilisateur.
* `PUT /users/:id`: Mettre à jour un utilisateur existant avec l'ID spécifié.
* `DELETE /users/:id`: Supprimer un utilisateur avec l'ID spécifié.

### /roles

* `GET /roles`: Récupérer la liste des rôles.
* `POST /roles`: Créer un nouveau rôle.
* `PUT /roles/:id`: Mettre à jour un rôle existant avec l'ID spécifié.
* `DELETE /roles/:id`: Supprimer un rôle avec l'ID spécifié.

### /roles

* `GET /groups`: Récupérer la liste des groupes.
* `POST /groups`: Créer un nouveau groupes.
* `PUT /groups/:id`: Mettre à jour un groupe existant avec l'ID spécifié.
* `DELETE /groups/:id`: Supprimer un groupe avec l'ID spécifié.

### /auth

* `POST /auth`: Authentifier un utilisateur et retourner un jeton JWT.
* `POST /signup`: création d'un utilisateur dans la DB avec email + password
* `POST /login`: authentification d'un utilisateur + retour du token JWT
* `POST /validate`: récupération du token JWT pour analyse et sécurisation des routes d'accès



