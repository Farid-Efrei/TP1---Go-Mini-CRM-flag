package main

import "fmt"

type Contact struct {
	ID int
	Nom string
	Email string
}
func main() {
	// contacts := make(map[int]Contact)
	// This is a placeholder for the main function.
// fmt.Println("Hello, World!")
printMenu()
}

func printMenu(){
	fmt.Println(" 🦋 === Menu Mini-CRM en CLI === 🦋")
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Lister Tous les contacts")
	fmt.Println("3. Supprimer un contact par ID")
	fmt.Println("4. Quitter le Mini-CRM")
}