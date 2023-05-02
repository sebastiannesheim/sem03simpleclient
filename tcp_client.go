package main

import (
	"log"
	"net"
	"os"

	"github.com/sebastiannesheim/is105sem03/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.2:8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("os.Args[1] = ", os.Args[1])

	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4) //Dette kryptere meldingen

	_, err = conn.Write([]byte(string(kryptertMelding))) //Dette sender den krypterte meldingen
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	response := string(buf[:n])                                                                                //Dette gjør om fra byte til string
	log.Printf("reply from proxy: %s", response)                                                               //Printer response
	dekryptertMelding := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)+4) //Meldingen som komme tilbake blir dekrypt her
	log.Println("Dekrypter melding: ", string(dekryptertMelding))                                              //Printer den dekrypte meldingen

}
