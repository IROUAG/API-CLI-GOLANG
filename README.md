# Projet Golang: Ilies & Sylvain

Ce dépôt contient une API Golang avec une base de données PostgreSQL et une CLI pour interagir avec l'API, organisée comme suit :

- `app/` : Contient le code de l'application Go, son Dockerfile et le fichier des variables d'environnement.
- `postgres-setup/` : Contient le script SQL pour la configuration de la base de données PostgreSQL et son Dockerfile.
- `cli/` : Contient le code source de la CLI, son Dockerfile et les fichiers nécessaires à son fonctionnement.
- `docker-compose.yml` : Fichier de configuration Docker Compose pour construire et exécuter l'ensemble de l'application.
- `.env` : Fichier des variables d'environnement pour stocker les paramètres de connexion à la base de données.

## Prérequis

- Docker
- Docker Compose

## Comment construire et exécuter l'API et la CLI

1. Clonez le dépôt :

```bash
git clone https://gitlab.com/votre-username/golang_project_ilies_sylvain.git
cd golang_project_ilies_sylvain
```

2. Construisez et exécutez l'application en utilisant Docker Compose :

```bash
docker-compose up -d --build
```

Cette commande construira les conteneurs de l'application Go, de PostgreSQL et de la CLI, puis les exécutera ensemble. Votre application sera accessible à l'adresse http://localhost:8080.

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

## Utilisation de la CLI

Pour utiliser la CLI, il est fortement recommander de créer un alias:

```bash
alias cli="docker exec -it cli ./cli"
```

Pour pouvoir exécuter les commandes à l'intérieur du conteneur:

```bash
cli your-command [args]
```

Sinon vous pouvez exécuter les commandes à l'intérieur du conteneur en utilisant docker exec :

```bash
docker exec -it cli ./cli your-command [args]
```
Remplacez your-command et [args] par la commande appropriée et les arguments de votre application CLI.

### Commandes disponibles

* `login`: Se connecter en tant qu'utilisateur et récupérer un JWT d'authentification et un jeton d'actualisation.
    * Flags:
        * `--email`: Adresse e-mail de l'utilisateur.
        * `---password`: Mot de passe de l'utilisateur.
* `refresh`: Actualiser un JWT d'authentification à l'aide d'un jeton d'actualisation.
    * Flags:
        * `--refresh_token`: Le jeton d'actualisation.
* `logout`: Se déconnecter et supprimer un JWT d'authentification et un jeton d'actualisation.
    * Flags:
        * `--access_token`: Le jeton d'authentification.
* `users list`: Lister tous les utilisateurs.
* `users get [user_id]`: Récupérer un utilisateur spécifique.
* `users create`: Créer un nouvel utilisateur.
    * Flags:
        * `--email`: Adresse e-mail de l'utilisateur.
        * `--password`: Mot de passe de l'utilisateur.
        * `--name`: Nom complet de l'utilisateur.
* `users update [user_id]`: Mettre à jour un utilisateur existant.
    * Flags:
        * `--email`: Nouvelle adresse e-mail de l'utilisateur.
        * `--password`: Nouveau mot de passe de l'utilisateur.
        * `--name`: Nouveau nom complet de l'utilisateur.