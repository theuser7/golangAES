package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io/ioutil"
	"io"
	"fmt"
)


func Encrypter() {
	key := []byte("abdefghjklmnopars123456789123456") //key bytes 
	block,err := aes.NewCipher(key) //new key
	if err != nil {
		panic(err.Error()) 
	}

	content,err := ioutil.ReadFile("plaintext.txt") //read small file
	buffer := make([]byte,aes.BlockSize+len(content))
	iv := buffer[:aes.BlockSize]
	if _,err := io.ReadFull(rand.Reader, iv); err != nil { //new padding
		panic(err.Error())
	}

	stream := cipher.NewCTR(block,iv) //counter blocks
	stream.XORKeyStream(buffer[aes.BlockSize:],content) //xorkeystream
	err = ioutil.WriteFile("ciphertext.bin",buffer,0644) //write new file encrypted
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Encrypted!")

}



func Decrypter() {
	key := []byte("abdefghjklmnopars123456789123456")
	block,err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	content,err := ioutil.ReadFile("ciphertext.bin")
	buffer := make([]byte,len(content[aes.BlockSize:]))
	stream := cipher.NewCTR(block,content[:aes.BlockSize])
	stream.XORKeyStream(buffer,content[aes.BlockSize:])
	err = ioutil.WriteFile("decrypted.txt",buffer,0644)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Decrypted!")


}


func main() {
	Encrypter()
	Decrypter()
}
