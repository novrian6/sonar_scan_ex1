package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

func main() {
	// Contoh bug: Potensi pembagian dengan nol tanpa pengecekan
	divider := 0
	result := 100 / divider // Akan memicu temuan karena pembagian dengan nol
	fmt.Println("Result:", result)

	// Contoh code smell: Hard-coded credentials
	username := "admin"
	password := "12345" // Ini akan dianggap sebagai hard-coded password

	fmt.Println("Credentials:", username, password)

	// Contoh vulnerability: Menggunakan algoritma hashing yang tidak aman (MD5)
	data := []byte("sensitive data")
	hash := md5.Sum(data)
	fmt.Printf("MD5 Hash: %x\n", hash)

	// Contoh duplicative code (code smell)
	duplicateCode()

	// Fungsi yang tidak digunakan
	unusedFunction()
}

func duplicateCode() {
	data, err := ioutil.ReadFile("file.txt") // Contoh penggunaan API deprecated
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("File content:", string(data))

	// Baris kode ini diulang
	data, err = ioutil.ReadFile("file.txt") // Contoh penggunaan API deprecated
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("File content:", string(data))
}

func unusedFunction() {
	fmt.Println("This function is unused!")
}
