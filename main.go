package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID int
	Nom string
	Email string
}
var contacts = make(map[int]Contact)


func main() {

	ajouter := flag.Bool("ajouter", false, "Ajouter un contact")
	nomFlag := flag.String("nom", "", "Nom du contact")
	emailFlag := flag.String("email", "", "Email du contact")
	flag.Parse()

	if *ajouter {
		if *nomFlag == "" || *emailFlag == "" {
			fmt.Println("Le nom et l'email sont requis pour ajouter un contact.")
			return
		}
		id := len(contacts) + 1
		contact := Contact{ID: id, Nom: *nomFlag, Email: *emailFlag}
		contacts[id] = contact
		fmt.Printf("Contact ajouté avec ID %d, Nom: %s, Email: %s\n", id, contact.Nom, contact.Email)

		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		printMenu()
		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Entrée invalide, veuillez entrer du choix correspondant.")

	}
		switch choice {
		case 1:
			ajouterContact(reader)
		case 2:
			listerContacts()
		case 3:
			supprimerContact(reader)
		case 4:
			modifierContact(reader)
		case 5:
			fmt.Println("A très bientôt ! 😸")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	// This is a placeholder for the main function.
// fmt.Println("Hello, World!")
}
	}

func printMenu(){
	fmt.Println(" 🦋 === Menu Mini-CRM en CLI === 🦋")
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Lister Tous les contacts")
	fmt.Println("3. Supprimer un contact par ID")
	fmt.Println("4. Modifier un contact par ID")
	fmt.Println("5. Quitter le Mini-CRM")
	fmt.Println("Choisissez une option (1-5): ")
}



func ajouterContact(reader *bufio.Reader){
	fmt.Print("Entrez le nom du contact: ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)

	fmt.Print("Entrez l'email du contact: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	if nom == "" || email == "" {
		fmt.Println("Le nom et l'email sont requis.")
		return
	}

	id := len(contacts) + 1
	contact := Contact{ID: id, Nom: nom, Email: email}
	contacts[id] = contact
	fmt.Printf("Contact ajouté avec ID %d\n", id)
} 


func listerContacts(){
	if len(contacts) == 0 {
		fmt.Println("Aucun contact disponible !!!")
		return
	}
	fmt.Println("Liste des contacts:")
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", contact.ID, contact.Nom, contact.Email)

	}
	fmt.Println("")

}

func supprimerContact(reader *bufio.Reader){
	fmt.Println("ID à supprimer : ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Entrée invalide, veuillez entrer un ID valide.")
		return
	}
	if _, ok := contacts[id]; !ok {
		fmt.Println("Aucun contact trouvé avec cet ID.")
		return
	}
	delete(contacts, id)
	fmt.Println("Contact supprimé.")
}

func modifierContact(reader *bufio.Reader){
	fmt.Println("ID à modifier : ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Entrée invalide, veuillez entrer un ID valide.")
		return
	}
	contact, ok := contacts[id]
	if !ok {
		fmt.Println("Aucun contact trouvé avec cet ID.")
		return
	}

	fmt.Printf("Nom actuel (%s), appuyez sur Entrée pour conserver: ", contact.Nom)
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	if nom != "" {
		contact.Nom = nom
	}

	fmt.Printf("Email actuel (%s), appuyez sur Entrée pour conserver: ", contact.Email)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		contact.Email = email
	}
	contacts[id] = contact
	fmt.Println("Contact modifié avec succès.")
}

