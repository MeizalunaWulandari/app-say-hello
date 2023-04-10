package main

import (
	"fmt"
	"github.com/MeizalunaWulandari/go-say-hello"
)

func main(){
	fmt.Println(go_say_hello.SayHello())
}
/**
 * MENAMBAHKAN DEPENDENSI 
 * Untuk menambahkan depedensi pada project golang cukup dengan mejalankan go get namaModule
 * 
 * UPGRADE DEPENDENSI
 * Untuk upgrade dependensi cukup dengan mengubah nomor tag pada file go.mod
 * lalu jalankan perintah go get untuk mendowload update dependensinya
*/