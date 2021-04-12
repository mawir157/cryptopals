package main

import (
	"fmt"
	"math/rand"
	"time"
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

	aesPlainText      := JMT.ParseFromAscii(textMessage, false)
	aesCipherText     := JMT.ECBEncrypt(aes, aesPlainText)
	aesDecodedText, _ := JMT.ECBDecrypt(aes, aesCipherText)

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

	aesCipherText     = JMT.CBCEncrypt(aes, iv, aesPlainText)
	aesDecodedText, _ = JMT.CBCDecrypt(aes, iv, aesCipherText)

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

	aesCipherText     = JMT.PCBCEncrypt(aes, iv, aesPlainText)
	aesDecodedText, _ = JMT.PCBCDecrypt(aes, iv, aesCipherText)

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

	aesCipherText     = JMT.OFBEncrypt(aes, iv, aesPlainText)
	aesDecodedText, _ = JMT.OFBDecrypt(aes, iv, aesCipherText)

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

	aesCipherText     = JMT.CFBEncrypt(aes, iv, aesPlainText)
	aesDecodedText, _ = JMT.CFBDecrypt(aes, iv, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")

	// Week1()
	// Week2()
	printDay(17)
	aes_key_17 := JMT.RandomBlock(8)
	aes_17 := JMT.MakeAES(aes_key_17)	
	
	var iv_17 [4]JMT.Word
	copy(iv_17[:], temp)
	strings_17 := []string{
		"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
		"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
		"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
		"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
		"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
		"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
		"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
		"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
		"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
		"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
	}

	s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  secret_string_17 := JMT.ParseFromBase64(strings_17[r1.Intn(len(strings_17))],
                                          false)
	aesCipherText_17  := JMT.CBCEncrypt(aes_17, iv_17, secret_string_17)

	k := full_attack(iv_17, aesCipherText_17, aes_17)
	fmt.Println(JMT.ParseToAscii(k, true))

	secret, _ := JMT.CBCDecrypt(aes_17, iv_17, aesCipherText_17)
	fmt.Println(JMT.ParseToAscii(secret, true))

 	return
}

func singleBlockAttack(block []byte, bc JMT.BlockCipher) []byte {
	zeroing_iv := make([]byte, 16)

	for pad_val := 1; pad_val <= 16; pad_val++ {
		padding_iv := make([]byte, 16)
		for i := 0; i < len(padding_iv); i++ {
			padding_iv[i] = zeroing_iv[i] ^ byte(pad_val)
		}

		for candidate := byte(0x00); candidate < byte(0xFF); candidate++ {
			padding_iv[16 - pad_val] = candidate

			var temp_iv [4]JMT.Word
			copy(temp_iv[:], JMT.BytesToWords(padding_iv, false))
			_, err := JMT.CBCDecrypt(bc, temp_iv, block)
			if (err == nil) {
				// fmt.Println(strange)
				// fmt.Printf("* - %d - %d\n", candidate, pad_val)
				zeroing_iv[16 - pad_val] = candidate ^ byte(pad_val)
				break
			}
		}
	}

	return zeroing_iv
}

func full_attack(iv [4]JMT.Word, ct []byte, bc JMT.BlockCipher) []byte {
	noBlock := len(ct) / 16

	iv_bytes := []byte{}
	for _, w := range iv {
		for _, b := range w {
			iv_bytes = append(iv_bytes, b)
		}
	}

	ct = append(iv_bytes, ct...)
	result := []byte{}
	for block := 1; block <= noBlock; block++ {
		ct_block := ct[block*16:(block + 1)*16]
		dec := singleBlockAttack(ct_block, bc)

		pt := make([]byte, 16)
		for i := 0; i < 16; i++ {
			pt[i] = ct[(block - 1)*16 + i] ^ dec[i]
		}

		result = append(result, pt...)
	}

	return result
}