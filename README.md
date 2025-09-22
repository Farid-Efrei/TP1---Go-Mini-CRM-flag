# Mini-CRM CLI — TP1 (Golang, EFREI)

**Auteur :** Fairytale-Dev (Farid-Efrei)

---

## Description

Petit projet d'initiation à Go réalisé pour le cours de M2 de Golang à l'EFREI. Ce mini-CRM en ligne de commande permet de gérer des contacts simples (ID, Nom, Email) en mémoire.

## Fonctionnalités implémentées

- Menu principal en boucle (mode interactif).
- Ajouter un contact (l'utilisateur saisit Nom + Email).
- Lister tous les contacts.
- Supprimer un contact par son ID.
- Mettre à jour un contact (par ID).
- Quitter l'application.
- Ajouter un contact directement via flags CLI : `--ajouter --nom="Nom" --email="email@example.com"`

## Fichiers

- `main.go` : implémentation complète du CLI et des opérations CRUD sur les contacts.
- `README.md` : documentation formatée (ce fichier).

## Concepts Go utilisés

- `map[int]Contact` pour stocker les contacts en mémoire.
- Boucle `for {}` et `switch` pour le menu principal.
- `bufio.NewReader(os.Stdin)` + `ReadString` pour la saisie utilisateur.
- `strconv.Atoi` et `strings.TrimSpace` pour parser les entrées.
- Package `flag` pour gérer les options CLI.
- "Comma ok idiom" (ex : `v, ok := mymap[id]`) pour vérifier l'existence d'une clé dans une map.

## Limitations connues

- Les contacts sont stockés uniquement en mémoire : toutes les données sont perdues à la fermeture de l'application.
- Les IDs sont générés de manière simple (`len(contacts)+1`) : après suppression, il peut y avoir des trous dans la numérotation.
- Pas de validation avancée d'email ni de persistance (JSON / fichier) pour le moment.

## Comment exécuter

Ouvrir PowerShell (ou autre terminal) et se placer dans le dossier du projet :

- Exécution interactive :

```powershell
go run .
# ou
go run main.go
```

- Ajouter un contact directement via flags (sans entrer dans le menu) :

```powershell
go run . --ajouter --nom="Tanjiro" --email="tanjiro@kimetsu.jp"
```

- Compiler un exécutable :

```powershell
go build -o mini-crm.exe .
.\mini-crm.exe
```

## Objectifs pédagogiques

- S'initier aux bases de Go : types, maps, I/O, parsing, flags et idiomes courants.
- Produire un outil CLI simple respectant les consignes du TP.

## Améliorations possibles

- Sauvegarde/chargement JSON pour persistance des contacts.
- Validation d'email et sanitization des entrées.
- Tri des contacts et pagination pour l'affichage.
- Tests unitaires pour les opérations CRUD.

## Notes personnelles

Exercice réalisé pour le cours Golang (EFREI) — TP1.

🦋 Paix à tous 🦋
