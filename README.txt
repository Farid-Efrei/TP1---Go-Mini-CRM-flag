Mini-CRM CLI — TP1 GOlang (EFREI - Golang)

Description:
Ce petit projet est mon premier exercice en Go pour le cours de Golang à l'EFREI. Il s'agit d'un mini-CRM en ligne de commande permettant de gérer des contacts (ID, Nom, Email) en mémoire.

Fonctionnalités implémentées:
- Menu principal en boucle (interactive).
- Ajouter un contact (l'utilisateur saisit Nom + Email).
- Lister tous les contacts.
- Supprimer un contact par son ID.
- Mettre à jour un contact (par ID).
- Quitter l'application.
- Ajouter un contact directement via flags CLI :
  --ajouter --nom="Nom" --email="email@example.com"

Fichiers:
- main.go : implémentation complète du CLI et des opérations sur les contacts.

Concepts Go utilisés (à retenir):
- map[int]Contact pour stocker les contacts en mémoire.
- boucle for {} et switch pour le menu.
- bufio.NewReader(os.Stdin) + ReadString pour la saisie utilisateur.
- strconv.Atoi et strings.TrimSpace pour parser les entrées.
- package flag pour gérer les flags CLI.
- "comma ok idiom" (ex: v, ok := mymap[id]) pour vérifier l'existence d'une clé.

Limitations connues:
- Les contacts sont stockés uniquement en mémoire : toutes les données sont perdues à la fermeture.
- Les IDs sont générés de manière simple (len(contacts)+1) : après suppression, il peut y avoir des trous dans la numérotation.
- Pas de validation avancée d'email ni de persistance (JSON / fichier) pour le moment.

Comment exécuter (PowerShell, depuis le dossier TP1)
- Exécution interactive :
  go run main.go (ou go run .)

- Ajouter un contact via flags (sans entrer dans le menu) :
  go run main.go --ajouter --nom="Tanjiro" --email="tanjiro@kimetsu.jp"

- Compiler un exécutable :
  go build 

Objectifs pédagogiques:
- S'initier aux bases de Go : types, maps, entrées/sorties, parsing, flags et idiomes courants.
- Produire un outil CLI simple respectant les consignes du TP.

Améliorations possibles (si extension):
- Sauvegarde/chargement JSON pour persistance.
- Validation d'email et sanitization.
- Tri des contacts et pagination.
- Tests unitaires pour add/update/delete/list.

Auteur :
Fairytale-Dev (Farid-Efrei) 😸
Exercice réalisé pour le cours Golang (EFREI) — TP1.

🦋 Paix à tous 🦋