# Sujet

Ce sujet en deux parties doit se réaliser en binome. Chaque partie aura une note indépendante. L'idée est de mettre en place une API (Application Programming Interface) et la CLI (Command Line Interface) qui correspond. Ces deux types de programmes sont régulièrement créés par les équipes DevOps pour mettre à dispo des informations ou pour en récupérer auprès de différents fournisseurs de services.

## Partie 1

Vous êtes chargé(e) de développer une API REST en Golang pour la gestion d'utilisateurs, de rôles, de groupes d'utilisateurs, d'un système d'authentification basé sur JWT et d'un système de licence. C'est le socle de base de toute application SaaS.(Software As A Service)

### Spécifications de l'API

L'API doit implémenter les endpoints suivants :

* `/users`: permet de gérer les utilisateurs (GET, POST, PUT, DELETE).
* `/roles`: permet de gérer les rôles pour les utilisateurs (GET, POST, PUT, DELETE).
* `/groups`: permet de gérer les groupes d'utilisateurs (GET, POST, PUT, DELETE).
* `/auth`: permet de gérer l'authentification des utilisateurs en utilisant JWT (POST).
* `/licenses`: permet de gérer les licences (GET, POST, PUT, DELETE).

Chaque endpoint doit être protégé par l'authentification JWT, sauf le endpoint `/auth`.

### Contraintes techniques

* Vous devez utiliser le langage Golang pour implémenter l'API.
* Vous devez utiliser une base de données PostgreSQL pour stocker les données.
* Vous devez utiliser la bibliothèque Gin pour implémenter l'API.
* Vous devez utiliser la bibliothèque Gorm pour gérer les objets contenus dans la base de données PostgreSQL.
* Vous devez utiliser la bibliothèque jwt-go pour implémenter l'authentification JWT.
* L'application doit être déployable à l'aide de Docker et Docker Compose.

### Livrables

Vous devez fournir les livrables suivants :

* Le code source de l'API en Golang.
* Un script SQL pour créer les tables dans la base de données.
* Un fichier README expliquant comment construire et exécuter l'API, ainsi que les endpoints disponibles.
* Un fichier docker-compose.yml pour déployer l'application avec Docker Compose.
* Un fichier .env contenant les variables d'environnement nécessaires pour l'exécution de l'application.
* Ajouter une CI est un plus

Vous devrez également fournir une documentation détaillée de l'API, y compris les modèles de données et les schémas de réponse, ainsi qu'une description des erreurs qui peuvent être renvoyées. La documentation doit être au format **OpenAPI** (Swagger).

Vous devrez soumettre vos livrables sous la forme d'un dépôt GitLab auquel vous m'ajoutez (@tsaquet) en `mainteneur`. Le dépôt doit être bien structuré et facile à naviguer. Vous devrez également fournir un rapport détaillé décrivant les choix que vous avez faits lors de la conception et de la mise en œuvre de l'application, ainsi que les difficultés que vous avez rencontrées et comment vous les avez résolues.

### Attributs des objets

* User: 
  * id
  * name
  * email
  * password
  * roles
  * groups
  * created\_at
  * updated\_at
  * deleted\_at
  * auth\_tokens
	
* AuthToken:
  * id
  * token
  * expires\_at
	
* RefreshToken:
  * token
  * expires\_at
	
* Role:
  * id
  * name
  * description
  * permissions
  * created\_at
  * updated\_at
  * deleted\_at
	
* Group:
  * id
  * name
  * parent\_group\_id
  * child\_group\_ids
  * user\_ids
  * created\_at
  * updated\_at
  * deleted\_at
	
* License:
  * id
  * name
  * max\_users
  * valid\_until
  * created\_at
  * updated\_at
  * deleted\_at

## Partie 2

En plus de l'API, vous devez également développer une CLI pour communiquer avec l'API en question.

### Spécifications de la CLI

La CLI doit permettre d'interagir avec l'API et doit supporter les commandes suivantes :

* `login`: connecter un utilisateur et recevoir un jeton d'authentification JWT et un jeton de rafraîchissement (à stocker chiffré localement) (refresh token)
  * `email`: l'adresse email de l'utilisateur
  * `password`: le mot de passe de l'utilisateur
* `refresh`: renouveler un jeton d'authentification JWT à l'aide d'un jeton de rafraîchissement
  * `refresh_token`: le jeton de rafraîchissement
* `logout`: supprimer un jeton d'authentification JWT et de rafraîchissement
  * `access_token`: le jeton d'authentification à supprimer
  * `refresh_token`: le jeton de rafraîchissement à supprimer
* `users`: gérer les utilisateurs
  * `list`: lister tous les utilisateurs
  * `get`: récupérer un utilisateur spécifique
    * `user_id`: l'ID de l'utilisateur à récupérer
  * `create`: créer un nouvel utilisateur
    * `name`: le nom de l'utilisateur
    * `email`: l'adresse email de l'utilisateur
    * `password`: le mot de passe de l'utilisateur
  * `update`: mettre à jour un utilisateur existant
    * `user_id`: l'ID de l'utilisateur à mettre à jour
    * `name`: le nouveau nom de l'utilisateur
    * `email`: la nouvelle adresse email de l'utilisateur
    * `password`: le nouveau mot de passe de l'utilisateur
  * `delete`: supprimer un utilisateur existant
    * `user_id`: l'ID de l'utilisateur à supprimer
* `roles`: gérer les rôles
  * `list`: lister tous les rôles
  * `get`: récupérer un rôle spécifique
    * `role_id`: l'ID du rôle à récupérer
  * `create`: créer un nouveau rôle
    * `name`: le nom du rôle
    * `description`: la description du rôle
    * `permissions`: les permissions accordées au rôle
  * `update`: mettre à jour un rôle existant
    * `role_id`: l'ID du rôle à mettre à jour
    * `name`: le nouveau nom du rôle
    * `description`: la nouvelle description du rôle
    * `permissions`: les nouvelles permissions accordées au rôle
  * `delete`: supprimer un rôle existant
    * `role_id`: l'ID du rôle à supprimer
* `groups`: gérer les groupes d'utilisateurs
  * `list`: lister tous les groupes
  * `get`: récupérer un groupe spécifique
    * `group_id`: l'ID du groupe à récupérer
  * `create`: créer un nouveau groupe
    * `name`: le nom du groupe
    * `parent_group_id`: l'ID du groupe parent
    * `user_ids`: les ID des utilisateurs faisant partie du groupe
  * `update`: mettre à jour un groupe existant
    * `group_id`: l'ID du groupe à mettre à jour
    * `name`: le nouveau nom du groupe

### Contraintes techniques

* Vous devez utiliser le langage Golang pour implémenter la CLI.
* Vous devez utiliser la bibliothèque Cobra pour implémenter la CLI.
* L'application doit être déployable à l'aide de Docker et Docker Compose.

### Livrables


* Le code source de la CLI en Golang.
* Un fichier README expliquant comment construire et exécuter la CLI, ainsi que les commandes disponibles.
* Un fichier docker-compose.yml pour déployer l'application avec Docker Compose.
* Un fichier .env contenant les variables d'environnement nécessaires pour l'exécution de l'application.
* Ajouter une CI est un plus

Vous devrez soumettre vos livrables sous la forme d'un dépôt GitLab auquel vous m'ajoutez (@tsaquet) en `mainteneur`. Le dépôt doit être bien structuré et facile à naviguer. Vous devrez également fournir un rapport détaillé décrivant les choix que vous avez faits lors de la conception et de la mise en œuvre de l'application, ainsi que les difficultés que vous avez rencontrées et comment vous les avez résolues.
