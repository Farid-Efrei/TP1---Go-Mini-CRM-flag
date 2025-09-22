package main

import "fmt"

type Contact struct {
	ID int
	Nom string
	Email string
}
contacts := make(map[int]Contact)
func main() {
	// This is a placeholder for the main function.
// fmt.Println("Hello, World!")
printMenu()
}

func printMenu(){
	fmt.Println(" 🦋 === Menu Mini-CRM en CLI === 🦋")
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Lister Tous les contacts")
	fmt.Println("3. Supprimer un contact par ID")
	fmt.Println("4. Modifier un contact par ID")
	fmt.Println("5. Quitter le Mini-CRM")
}

func ajouterContact(reader *bufio.Reader){
	fmt.Print("Entrez le nom du contact: ")
	nom := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)

	fmt.Print("Entrez l'email du contact: ")
	email := reader.ReadString('\n')
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