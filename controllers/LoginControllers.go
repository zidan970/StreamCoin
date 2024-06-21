package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"zidan/AccountServiceAppProject/entities"
)

func AddLoginController(db *sql.DB, user_id uint) {
	var newLogin entities.Login

	statement, errPrepare := db.Prepare("insert into login (login_id, user_id) VALUES (?, ?)")
	if errPrepare != nil {
		log.Fatal("err prepare: ", errPrepare)
	}

	result, errExec := statement.Exec(&newLogin.Login_id, user_id)
	if errExec != nil {
		log.Fatal("insert data is failed: ", errExec)
	}

	_, errID := result.LastInsertId()
	_, errRows := result.RowsAffected()
	if errID != nil || errRows != nil {
		log.Println("errID:", errID)
		log.Println("errRows:", errRows)
	}
	// else {
	// 	fmt.Println("Succesfully insert data. Last inserted ID:", hasilID)
	// 	fmt.Println("Succesfully insert data. Row affected ID:", hasilRows)
	// }
}

func LoginVerificationController(db *sql.DB) {
	var phone_number string
	var password string
	var id uint

	fmt.Println("Masukkan No. Telpon:")
	fmt.Scanln(&phone_number)
	fmt.Println("Masukkan Password:")
	fmt.Scanln(&password)

	row := db.QueryRow("SELECT user_id FROM accounts WHERE user_phone = ?", phone_number)
	if err := row.Scan(&id); err != nil {
		fmt.Println("Login gagal. Anda belum terdaftar")
		fmt.Println()
		Menu(db)
	}

	fmt.Printf("Login berhasil.\n")
	fmt.Println("** Silakan Menikmati Layanan Kami **")
	status_login = true
	AddLoginController(db, id)
	fmt.Println()
	Menu(db)
}
