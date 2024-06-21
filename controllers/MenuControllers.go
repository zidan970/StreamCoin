package controllers

import (
	"database/sql"
	"fmt"
	"os"
)

var status_login bool

func Menu(db *sql.DB) {

	fmt.Println("These are our menu:")
	fmt.Println("1. Add Account (Register)")
	fmt.Println("2. Login")
	fmt.Println("3. Read Account")
	fmt.Println("4. Update Account")
	fmt.Println("5. Delete Account")
	fmt.Println("6. Top-Up")
	fmt.Println("7. Transfer")
	fmt.Println("8. History Top-Up")
	fmt.Println("9. History Transfer")
	fmt.Println("10. Read User's Profile")
	fmt.Println("0. Quit")

	fmt.Println("\nChoose Your Menu:")

	var menu int
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		fmt.Println("Welcome to Add Account (Register)!")
		AddAccountController(db)
		//status_login = true
	case 2:
		if status_login == true {
			fmt.Println("Anda sudah login")
			fmt.Println()
			Menu(db)
		} else if status_login == false {
			fmt.Println("Welcome to Login!")
			LoginVerificationController(db)
		}
	case 3:
		if status_login == false {
			fmt.Println("Silakan login terlebih dahulu")
			fmt.Println()
			Menu(db)
		} else if status_login == true {
			fmt.Println("Welcome to Read Account!")
			ReadMyAccountController(db)
		}
	case 4:
		if status_login == false {
			fmt.Println("Silakan login terlebih dahulu")
			fmt.Println()
			Menu(db)
		} else if status_login == true {
			fmt.Println("Welcome to Update Account!")
			MenuProfil(db)
		}
	case 5:
		if status_login == false {
			fmt.Println("Silakan login terlebih dahulu")
			fmt.Println()
			Menu(db)
		} else if status_login == true {
			fmt.Println("Welcome to Delete Account!")
			DeleteMyAccountController(db)
		}
	case 6:
		if status_login == false {
			fmt.Println("Silakan login terlebih dahulu")
			fmt.Println()
			Menu(db)
		} else if status_login == true {
			fmt.Println("Welcome to Top-Up!")
			AddMyBalanceByInfController(db)
		}
	case 7:
		if status_login == false {
			fmt.Println("Silakan login terlebih dahulu")
			fmt.Println()
			Menu(db)
		} else if status_login == true {
			fmt.Println("Welcome to Transfer!")
			DoTransferController(db)
		}
	case 8:
		if status_login == false {
			fmt.Println("Silakan login terlebih dahulu")
			fmt.Println()
			Menu(db)
		} else if status_login == true {
			fmt.Println("Welcome to History Top-Up!")
			ReadHistoryTopUpController(db)
		}
	case 9:
		if status_login == false {
			fmt.Println("Silakan login terlebih dahulu")
			fmt.Println()
			Menu(db)
		} else if status_login == true {
			fmt.Println("Welcome to History Transfer!")
			ReadHistoryTransferController(db)
		}
	case 10:
		if status_login == false {
			fmt.Println("Silakan login terlebih dahulu")
			fmt.Println()
			Menu(db)
		} else if status_login == true {
			fmt.Println("Welcome to Read User's Profile")
			ReadOtherAccountsController(db)
		}
	case 0:
		fmt.Println("Terima kasih telah bertransaksi")
		os.Exit(0)
		status_login = false
	default:
		fmt.Println("Menu tidak tersedia")
	}
}

func MenuProfil(db *sql.DB) {
	fmt.Println("Silakan pilih data yang akan anda update:")
	fmt.Println("1. Username")
	fmt.Println("2. Display Name")
	fmt.Println("3. Phone")
	fmt.Println("4. Email")
	fmt.Println("5. Address")
	fmt.Println("6. Password")

	fmt.Println("\nChoose Your Menu:")

	var choose int
	fmt.Scanln(&choose)

	UpdateMyAccountController(db, choose)
}
