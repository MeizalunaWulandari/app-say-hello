package main

import (
	"fmt"
	"github.com/MeizalunaWulandari/go-say-hello/v2"
)

func main(){
	fmt.Println(go_say_hello.SayHello("Luna"))
}
/**
 * MENAMBAHKAN DEPENDENSI 
 * Untuk menambahkan depedensi pada project golang cukup dengan mejalankan go get namaModule
 * 
 * UPGRADE DEPENDENSI
 * Untuk upgrade dependensi cukup dengan mengubah nomor tag pada file go.mod
 * lalu jalankan perintah go get untuk mendowload update dependensinya
 * 
 * MAJOR UPGRADE
 * Pada saat ada major upgrade pada dependensi golang biasanya nama module juga berubah
 * oleh karena itu untuk melakukan major upgrade jalankan kembali perintah go get namaModule
*/