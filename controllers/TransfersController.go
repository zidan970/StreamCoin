package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"zidan/AccountServiceAppProject/entities"
)

func DoTransferController(db *sql.DB) {
	var akun1 entities.Accounts
	var akun2 entities.Accounts
	//var newTransfer entities.Transfers
	var phone_target string
	var my_phone string
	var amount uint
	var money string

	fmt.Println("Masukkan nomor anda:")
	fmt.Scanln(&my_phone)
	fmt.Println("Masukkan nomor tujuan transfer:")
	fmt.Scanln(&phone_target)
	fmt.Println("Masukkan jumlah transfer:")
	fmt.Scanln(&amount)

	money = strconv.Itoa(int(amount))

	res := FormatCheckController(db, money)

	if res == true {
		//cek jumlah amount pada balance. Apa cukup untuk transfer?
		row := db.QueryRow("select user_id, user_phone, user_balance from accounts where user_phone = ?", my_phone)
		if err := row.Scan(&akun1.User_id, &akun1.User_phone, &akun1.User_balance); err != nil {
			fmt.Println("cannot read balance", err)
		}

		if amount <= akun1.User_balance {
			row := db.QueryRow("select user_id, user_phone, user_balance from accounts where user_phone = ?", phone_target)
			if err := row.Scan(&akun2.User_id, &akun2.User_phone, &akun2.User_balance); err != nil {
				fmt.Println("cannot read balance", err)
			}

			balance2 := akun2.User_balance + amount
			fmt.Println(balance2)
			balance1 := akun1.User_balance - amount
			fmt.Println(balance1)

			result, errExec := db.Exec("UPDATE accounts SET user_balance = ? WHERE user_balance = ? and user_phone = ?", balance2, akun2.User_balance, akun2.User_phone)
			if errExec != nil {
				fmt.Println("saldo target tidak bisa update", errExec)
			}

			resultRow, _ := result.RowsAffected()
			fmt.Println("resultRow:", resultRow)

			result2, errExec2 := db.Exec("UPDATE accounts SET user_balance = ? WHERE user_balance = ? and user_phone = ?", balance1, akun1.User_balance, akun1.User_phone)
			if errExec2 != nil {
				fmt.Println("saldo anda tidak bisa diupdate", errExec2)
			}

			resultRow2, _ := result2.RowsAffected()
			fmt.Println("resultRow2:", resultRow2)

			AddHistoryTransferController(db, akun1.User_id, akun2.User_id, amount)

			fmt.Println()
			fmt.Println("Want to do another transaction?")
			Menu(db)
		} else if amount > akun1.User_balance {
			fmt.Println("Your balance is not enough. Please top up.")
			fmt.Println()
			fmt.Println("Want to do another transaction?")
			Menu(db)
		}
	} else if res == false {
		fmt.Println("Format angka tidak valid")
		fmt.Println()
		fmt.Println("Want to do another transaction?")
		Menu(db)
	}
}

func AddHistoryTransferController(db *sql.DB, id_sdr uint, id_rcv uint, amount uint) {
	var newTransfer entities.Transfers

	statement, errPrepare := db.Prepare("insert into transfers (transfer_id, user_id_sdr, user_id_rcv, transfer_amount) VALUES (?, ?, ?, ?)")
	if errPrepare != nil {
		log.Fatal("err prepare: ", errPrepare)
	}

	result, errExec := statement.Exec(&newTransfer.Transfer_id, id_sdr, id_rcv, amount)
	if errExec != nil {
		log.Fatal("insert history transfer is failed: ", errExec)
	}

	_, errID := result.LastInsertId()
	_, errRows := result.RowsAffected()
	if errID != nil || errRows != nil {
		log.Println("errID:", errID)
		log.Println("errRows:", errRows)
	}
}

func ReadHistoryTransferController(db *sql.DB) {
	var allHistory []entities.TransfersJoin
	var rowTransfer entities.TransfersJoin
	var username string
	var id_sdr uint

	fmt.Println("Masukkan username anda:")
	fmt.Scanln(&username)

	rowID := db.QueryRow("select user_id from accounts where username = ?", username)
	if err := rowID.Scan(&id_sdr); err != nil {
		if err == sql.ErrNoRows {
			//return member, fmt.Errorf("member dengan ID %d tidak ditemukan", id)
		}
		//return member, fmt.Errorf("gagal membaca data member: %v", err)
	}

	rowFull, errQuery := db.Query("select transfers.transfer_id, accounts.username, accounts.user_phone, transfers.transfer_amount, transfers.transfer_time from transfers inner join accounts on accounts.user_id = transfers.user_id_rcv where user_id_sdr = ?", id_sdr)
	if errQuery != nil {
		log.Fatal("cannot do select: ", errQuery)
	}

	for rowFull.Next() {
		errScan := rowFull.Scan(&rowTransfer.Transfer_id, &rowTransfer.Username_rcv, &rowTransfer.Phone_rcv, &rowTransfer.Transfer_amount, &rowTransfer.Transfer_time)
		if errScan != nil {
			log.Fatal("cannot do scan from row: ", errScan.Error())
		}
		allHistory = append(allHistory, rowTransfer)
	}

	fmt.Println("This is your Top-up History")
	for _, value := range allHistory {
		fmt.Printf("\nTransfer ID: %v \nUsername penerima: %v \nNo. Telp Penerima: %v \nTransfer Amount: %v \nTransfer Time: %v\n", value.Transfer_id, value.Username_rcv, value.Phone_rcv, value.Transfer_amount, value.Transfer_time)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Want to do another transaction?")
	Menu(db)
}
