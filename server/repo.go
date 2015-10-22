package main

import "fmt"

var currentId int

var accounts Accounts

func RepoFindAccount(id int) Account {
	for _, ac := range accounts {
		if ac.Id == id {
			return ac
		}
	}

	return Account{}
}

func RepoCreateAccount(ac Account) Account {
	currentId += 1
	ac.Id = currentId

	accounts = append(accounts, ac)

	return ac
}

func RepoDestroyAccount(id int) error {
	for i, ac := range accounts {
		if ac.Id == id {
			accounts = append(accounts[:i], accounts[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
