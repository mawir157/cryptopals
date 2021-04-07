package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

import JMT "github.com/mawir157/jmtcrypto"


func main() {
	textMessage :=
`It was the best of times, it was the worst of times, it was the age of wisdom,
it was the age of foolishness, it was the epoch of belief, it was the epoch of
incredulity, it was the season of Light, it was the season of Darkness, it was
the spring of hope, it was the winter of despair, we had everything before us,
we had nothing before us, we were all going direct to Heaven, we were all going
direct the other way – in short, the period was so far like the present period,
that some of its noisiest authorities insisted on its being received, for good
or for evil, in the superlative degree of comparison only.`

  //////////////////////////////////////////////////////////////////////////////
	//
	// McEliese
	//
	// public, private := generateKeyPair(2, 11)

	// public.Write("mce.pub")
	// private.Write("mce.pri")

	// public2 := ReadPublic("mce.pub")
	// private2 := ReadPrivate("mce.pri")

	// cipherText := public2.Encrypt(textMessage)
	// // PrintHex(cipherText, true)
	
	// plaintext := private2.Decrypt(cipherText)
	// PrintAscii(plaintext, true)
	// fmt.Println("")


	//////////////////////////////////////////////////////////////////////////////
	//
	// ECB w/ AES
	//
	aes_key := JMT.RandomBlock(8)
	aes := JMT.MakeAES(aes_key)

	aesPlainText   := JMT.ParseFromAscii(textMessage, false)
	aesCipherText  := JMT.ECBEncrypt(aes, aesPlainText)
	aesDecodedText := JMT.ECBDecrypt(aes, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")
	//////////////////////////////////////////////////////////////////////////////
	//
	// CBC w/ AES
	//
	aes_key = JMT.RandomBlock(8)
	aes = JMT.MakeAES(aes_key)

	var iv [4]JMT.Word
	temp := JMT.RandomBlock(4)
	copy(iv[:], temp)

	aesCipherText  = JMT.CBCEncrypt(aes, iv, aesPlainText)
	aesDecodedText = JMT.CBCDecrypt(aes, iv, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")
	//////////////////////////////////////////////////////////////////////////////
	//
	// PCB w/ AES
	//
	aes_key = JMT.RandomBlock(8)
	aes = JMT.MakeAES(aes_key)
	temp = JMT.RandomBlock(4)
	copy(iv[:], temp)

	aesCipherText  = JMT.PCBCEncrypt(aes, iv, aesPlainText)
	aesDecodedText = JMT.PCBCDecrypt(aes, iv, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")
	//////////////////////////////////////////////////////////////////////////////
	//
	// OFB w/ AES
	//
	aes_key = JMT.RandomBlock(8)
	aes = JMT.MakeAES(aes_key)
	temp = JMT.RandomBlock(4)
	copy(iv[:], temp)

	aesCipherText  = JMT.OFBEncrypt(aes, iv, aesPlainText)
	aesDecodedText = JMT.OFBDecrypt(aes, iv, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")
	//////////////////////////////////////////////////////////////////////////////
	//
	// CFB w/ AES
	//
	aes_key = JMT.RandomBlock(8)
	aes = JMT.MakeAES(aes_key)
	temp = JMT.RandomBlock(4)
	copy(iv[:], temp)

	aesCipherText  = JMT.CFBEncrypt(aes, iv, aesPlainText)
	aesDecodedText = JMT.CFBDecrypt(aes, iv, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")

	Week1()
	Week2()

 	return
}
