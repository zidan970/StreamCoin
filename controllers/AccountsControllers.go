package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"zidan/AccountServiceAppProject/entities"
)

func AddAccountController(db *sql.DB) {
	var newAccount entities.Accounts
	newAccount.User_balance = 0

	fmt.Println("Masukkan Username:")
	fmt.Scanln(&newAccount.Username)
	fmt.Println("Masukkan Nama:")
	fmt.Scanln(&newAccount.User_nama)
	fmt.Println("Masukkan No. Telpon:")
	fmt.Scanln(&newAccount.User_phone)
	fmt.Println("Masukkan Email:")
	fmt.Scanln(&newAccount.User_email)
	fmt.Println("Masukkan Address:")
	fmt.Scanln(&newAccount.User_address)
	fmt.Println("Atur Password:")
	fmt.Scanln(&newAccount.User_pswd)

	fmt.Println()
	statement, errPrepare := db.Prepare("insert into accounts (username, user_nama, user_phone, user_email, user_address, user_balance, user_pswd) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if errPrepare != nil {
		log.Fatal("err prepare: ", errPrepare)
	}

	result, errExec := statement.Exec(&newAccount.Username, &newAccount.User_nama, &newAccount.User_phone, &newAccount.User_email, &newAccount.User_address, &newAccount.User_balance, &newAccount.User_pswd)
	if errExec != nil {
		log.Fatal("insert data is failed: ", errExec)
	}

	//fmt.Println(result)

	hasilID, errID := result.LastInsertId()
	hasilRows, errRows := result.RowsAffected()
	if errID != nil || errRows != nil {
		log.Println(errID)
		log.Println(errRows)
	} else {
		fmt.Println("Succesfully insert data. Last inserted ID:", hasilID)
		fmt.Println("Succesfully insert data. Row affected ID:", hasilRows)
	}
	fmt.Println()
	fmt.Println("Want to do another transaction?")
	Menu(db)
}

func ReadAllAccountsController(db *sql.DB) []entities.Accounts {
	// *** read / select data all_accounts *** //
	var All_accounts []entities.Accounts

	rows, errSelect := db.Query("select * from accounts")
	if errSelect != nil {
		log.Fatal("cannot run select query: ", errSelect)
	}

	for rows.Next() {
		var Row_accounts entities.Accounts
		errScan := rows.Scan(&Row_accounts.User_id, &Row_accounts.Username, &Row_accounts.User_nama, &Row_accounts.User_phone, &Row_accounts.User_email, &Row_accounts.User_address, &Row_accounts.User_balance, &Row_accounts.User_pswd)
		if errScan != nil {
			log.Fatal("cannot run scan query: ", errScan.Error())
		}
		All_accounts = append(All_accounts, Row_accounts)
	}

	return All_accounts
}

func ReadMyAccountController(db *sql.DB) {
	// *** read / select data my account *** //
	var my_account entities.Accounts

	rowID := db.QueryRow("select accounts.user_id from accounts inner join login on accounts.user_id = login.user_id where accounts.user_id = (select user_id from login where login_id = (select max(login_id) from login)) limit 1;")

	if err := rowID.Scan(&my_account.User_id); err != nil {
		log.Fatal("cannot read user id. Please register: ", err)
	}

	fmt.Println(my_account.User_id)

	// // query select
	rowFull := db.QueryRow("select * from accounts where user_id = ?", my_account.User_id)

	if err := rowFull.Scan(&my_account.User_id, &my_account.Username, &my_account.User_nama, &my_account.User_phone, &my_account.User_email, &my_account.User_address, &my_account.User_balance, &my_account.User_pswd); err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}

	fmt.Println("This is your profile")
	fmt.Printf("\nUser ID: %v \nUsername: %v \nDisplay Name: %v \nPhone: %v \nEmail: %v \nAddress: %v \nBalance: %v \nPassword: %v\n", my_account.User_id, my_account.Username, my_account.User_nama, my_account.User_phone, my_account.User_email, my_account.User_address, my_account.User_balance, my_account.User_pswd)
	fmt.Println()
	fmt.Println("Want to do another transaction?")
	Menu(db)
}

func ReadOtherAccountsController(db *sql.DB) {
	// *** read / select data my account *** //
	var my_account entities.Accounts
	var phone string
	fmt.Println("Masukkan No. Telp yang ingin Anda lihat Profilnya:")
	fmt.Scanln(&phone)

	row := db.QueryRow("select user_nama, user_address, user_balance from accounts where user_phone = ?", phone)
	if err := row.Scan(&my_account.User_nama, &my_account.User_address, &my_account.User_balance); err != nil {
		fmt.Println("no user's profile found: ", err)
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	}

	fmt.Println("\nIni adalah profil yang Anda cari:")
	fmt.Printf("Display Name: %v \nAlamat: %v \nSaldo: %v\n", my_account.User_nama, my_account.User_address, my_account.User_balance)
	fmt.Println()
	fmt.Println("Want to do another transaction?")
	Menu(db)
}

func UpdateMyAccountController(db *sql.DB, choose int) {
	// *** coba update *** //
	var profil entities.Accounts
	var old_username string
	var old_name string
	var old_phone string
	var old_email string
	var old_address string
	var old_pswd string

	switch choose {
	case 1:
		fmt.Println("Masukkan Username Baru Anda:")
		fmt.Scanln(&profil.Username)
		fmt.Println("Masukkan Username Lama Anda:")
		fmt.Scanln(&old_username)

		result1, err := db.Exec("UPDATE accounts SET username = ? WHERE username = ?", profil.Username, old_username)
		if err != nil {
			log.Fatal("cannot update data: ", err)
		}

		hasil1ID, err := result1.LastInsertId()
		hasil1Row, err := result1.RowsAffected()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("berhasil update. Last inserted ID:", hasil1ID)
			fmt.Println("berhasil update. Row affected ID:", hasil1Row)
		}
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	case 2:
		fmt.Println("Masukkan Username Anda:")
		fmt.Scanln(&profil.Username)
		fmt.Println("Masukkan Nama Baru Anda:")
		fmt.Scanln(&profil.User_nama)
		fmt.Println("Masukkan Nama Lama Anda:")
		fmt.Scanln(&old_name)

		result1, err := db.Exec("UPDATE accounts SET user_nama = ? WHERE user_nama = ? and username = ?", profil.User_nama, old_name, profil.Username)
		if err != nil {
			log.Fatal("cannot update data: ", err)
		}

		hasil1ID, err := result1.LastInsertId()
		hasil1Row, err := result1.RowsAffected()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("berhasil update. Last inserted ID:", hasil1ID)
			fmt.Println("berhasil update. Row affected ID:", hasil1Row)
		}
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	case 3:
		fmt.Println("Masukkan No Telepon Baru Anda:")
		fmt.Scanln(&profil.User_phone)
		fmt.Println("Masukkan No Telepon Lama Anda:")
		fmt.Scanln(&old_phone)

		result1, err := db.Exec("UPDATE accounts SET user_phone = ? WHERE user_phone = ?", profil.User_phone, old_phone)
		if err != nil {
			log.Fatal("cannot update data: ", err)
		}

		hasil1ID, err := result1.LastInsertId()
		hasil1Row, err := result1.RowsAffected()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("berhasil update. Last inserted ID:", hasil1ID)
			fmt.Println("berhasil update. Row affected ID:", hasil1Row)
		}
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	case 4:
		fmt.Println("Masukkan Username Anda:")
		fmt.Scanln(&profil.Username)
		fmt.Println("Masukkan Email Baru Anda:")
		fmt.Scanln(&profil.User_email)
		fmt.Println("Masukkan Email Lama Anda:")
		fmt.Scanln(&old_email)

		result1, err := db.Exec("UPDATE accounts SET user_email = ? WHERE user_email = ? and username = ?", profil.User_email, old_email, profil.Username)
		if err != nil {
			log.Fatal("cannot update data: ", err)
		}

		hasil1ID, err := result1.LastInsertId()
		hasil1Row, err := result1.RowsAffected()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("berhasil update. Last inserted ID:", hasil1ID)
			fmt.Println("berhasil update. Row affected ID:", hasil1Row)
		}
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	case 5:
		fmt.Println("Masukkan Username Anda:")
		fmt.Scanln(&profil.Username)
		fmt.Println("Masukkan Alamat Baru Anda:")
		fmt.Scanln(&profil.User_address)
		fmt.Println("Masukkan Alamat Lama Anda:")
		fmt.Scanln(&old_address)

		result1, err := db.Exec("UPDATE accounts SET user_address = ? WHERE user_address = ? and username = ?", profil.User_address, old_address, profil.Username)
		if err != nil {
			log.Fatal("cannot update data: ", err)
		}

		hasil1ID, err := result1.LastInsertId()
		hasil1Row, err := result1.RowsAffected()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("berhasil update. Last inserted ID:", hasil1ID)
			fmt.Println("berhasil update. Row affected ID:", hasil1Row)
		}
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	case 6:
		fmt.Println("Masukkan Username Anda:")
		fmt.Scanln(&profil.Username)
		fmt.Println("Masukkan Password Baru Anda:")
		fmt.Scanln(&profil.User_pswd)
		fmt.Println("Masukkan Password Lama Anda:")
		fmt.Scanln(&old_pswd)

		result1, err := db.Exec("UPDATE accounts SET user_pswd = ? WHERE user_pswd = ? and username = ?", profil.User_pswd, old_pswd, profil.Username)
		if err != nil {
			log.Fatal("cannot update data: ", err)
		}

		hasil1ID, err := result1.LastInsertId()
		hasil1Row, err := result1.RowsAffected()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("berhasil update. Last inserted ID:", hasil1ID)
			fmt.Println("berhasil update. Row affected ID:", hasil1Row)
		}
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	}
}

func DeleteMyAccountController(db *sql.DB) {
	var my_account entities.Accounts
	var pilih int

	fmt.Println("Anda yakin?")
	fmt.Println("Pilih 1 jika yakin")
	fmt.Println("Pilih 2 jika ingin membatalkan")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		rowID := db.QueryRow("select accounts.user_id from accounts inner join login on accounts.user_id = login.user_id where accounts.user_id = (select user_id from login where login_id = (select max(login_id) from login)) limit 1;")

		if err := rowID.Scan(&my_account.User_id); err != nil {
			//
		}

		fmt.Println(my_account.User_id)

		// // query select
		result, errExec := db.Exec("delete from accounts where user_id = ?", my_account.User_id)
		if errExec != nil {
			log.Fatal("cannot delete data: ", errExec)
		}

		hasilID, errID := result.LastInsertId()
		hasilRow, errRow := result.RowsAffected()
		if errID != nil || errRow != nil {
			log.Fatal("errID: ", errID)
			log.Fatal("errRow: ", errRow)
		} else {
			fmt.Println("berhasil delete. Last inserted ID:", hasilID)
			fmt.Println("berhasil delete. Row affected ID:", hasilRow)
		}

		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	case 2:
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	}
}

func AddMyBalanceByInfController(db *sql.DB) {
	var my_account entities.Accounts
	var amount string
	fmt.Println("Masukkan Jumlah Top-Up:")
	fmt.Scanln(&amount)

	status := FormatCheckController(db, amount)

	if status == true {
		real_amount, _ := strconv.Atoi(amount)

		rowID := db.QueryRow("select accounts.user_id from accounts inner join login on accounts.user_id = login.user_id where accounts.user_id = (select user_id from login where login_id = (select max(login_id) from login)) limit 1;")

		if err := rowID.Scan(&my_account.User_id); err != nil {
			log.Fatal("cannot read user id. Please register: ", err)
		}

		fmt.Println(my_account.User_id)

		// // query select
		rowFull := db.QueryRow("select user_balance from accounts where user_id = ?", my_account.User_id)

		if err := rowFull.Scan(&my_account.User_balance); err != nil {
			if err == sql.ErrNoRows {
				log.Fatal(err)
			}
		}

		total_balance := my_account.User_balance + uint(real_amount)

		result, errExec := db.Exec("update accounts set user_balance = ? where user_balance = ? and user_id = ?", total_balance, my_account.User_balance, my_account.User_id)
		if errExec != nil {
			log.Fatal("cannot update balance: ", errExec)
		}

		AddHistoryTopUpController(db, my_account.User_id, uint(real_amount))

		hasilID, errID := result.LastInsertId()
		hasilRow, errRow := result.RowsAffected()
		if errID != nil || errRow != nil {
			log.Fatal("errID: ", errID)
			log.Fatal("errRow: ", errRow)
		} else {
			fmt.Println("berhasil top up. Last inserted ID:", hasilID)
			fmt.Println("berhasil top up. Row affected ID:", hasilRow)
		}
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	} else if status == false {
		fmt.Println("Format angka tidak valid")
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	}
}

// func AddMyBalanceByPhoneController(db *sql.DB) {

// }
