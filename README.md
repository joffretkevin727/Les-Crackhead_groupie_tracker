# 🦎 Lizard Crypto - Groupie Tracker

Une application web performante développée en **Go** / **HTML** / **CSS** / **JS** permettant de suivre les cours des cryptomonnaies en temps réel. Le projet exploite l'API CoinGecko pour fournir des données précises avec un système de gestion de favoris et de filtrage dynamique.

## 🚀 Fonctionnalités

* **Dashboard Live** : Visualisation en temps réel des prix, de la capitalisation boursière et des variations sur 24h.
* **Recherche Intelligente** : Système de recherche par nom ou symbole avec priorité aux correspondances par préfixe.
* **Filtres Avancés** : Tri dynamique par Market Cap (Supérieur/Inférieur à 1 Milliard $) et par performance (24h positif).
* **Session Guest** : Système de favoris persistant grâce à un stockage JSON local, sans besoin de création de compte traditionnel.
* **Fiches Détaillées** : Pages ressources complètes incluant l'offre en circulation (Supply), le volume d'échange et les descriptions.

## 📊 Technologies utilisées

* **Backend** : Go (Golang)
* **Frontend** : HTML5, CSS3 (Design modulaire), JavaScript (ES6+)
* **API externe** : [CoinGecko API](https://www.coingecko.com/en/api)
* **Persistance** : Fichiers JSON (Gestion des favoris et historique des sessions)

## 🛠️ Installation et Utilisation

1.  **Cloner le dépôt** :
    ```bash
    git clone [https://github.com/joffretkevin727/Les-Crackhead_groupie_tracker.git](https://github.com/joffretkevin727/Les-Crackhead_groupie_tracker.git)
    cd Les-Crackhead_groupie_tracker
    ```

2.  **Lancer l'application** :
    ```bash
    go run .
    ```

3.  **Accéder à l'interface** :
    Rendez-vous sur `http://localhost:8080` (ou le port défini dans votre configuration).

## 📁 Structure du Projet

```text
.
├── api/
│   └── api.go                                  # Gestion des appels à l'API CoinGecko
├── controller/
│   └── controller.go                           # Logique métier et handlers de requêtes
├── router/
│   └── router.go                               # Définition des routes du serveur HTTP
├── static/
│   ├── fonts/
│   │   └── unbounded-medium.ttf
│   ├── image/                                  # Logos et icônes (cœurs, recherche, etc.)
│   │   ├── avatar-defaut.png
│   │   ├── bitcoin.svg.png
│   │   ├── CG-Symbol.svg
│   │   ├── heart1.svg
│   │   ├── heartfull.svg
│   │   └── ...
│   ├── aboutus-style.css                       
│   ├── app.js                                  
│   ├── chart.js                                
│   ├── collection-style.css
│   ├── home-style.css
│   └── ressources-style.css                      
├── structure/                  
│   └── structure.go                            # Définition des types Go (Structs)
├── template/                                   # Templates HTML dynamiques
│   ├── aboutus.html                    
│   ├── collection.html                 
│   ├── home.html                   
│   ├── profil.html                 
│   ├── research.html                   
│   └── ressource.html                  
├── utils/                  
│   └── utils.go                                # Utilitaires (Recherche, Sync, Formats)
├── coins.json                                  # Cache de données locales
├── favorites.json                              # Sauvegarde des favoris
├── go.mod                                      # Gestion des modules Go
├── main.go                                     # Point d'entrée de l'application
├── README.md
└── userConnexion.json                          # Historique des sessions invité