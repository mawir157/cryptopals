package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

import JMT "github.com/mawir157/jmtcrypto"

func server12(bc JMT.BlockCipher, msg, secret string) []byte {
	plaintext := append(JMT.ParseFromAscii(msg, false),
										  JMT.ParseFromBase64(secret, false)...)

	cipherText := JMT.ECBEncrypt(bc, plaintext)

	return cipherText
}

func compareSlices(sl1, sl2 []byte) bool {
	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}

	return true
}

func PaddingAttack(bc JMT.BlockCipher, secretString string, blockSize int) {
	newWord := make([]byte, 0)
	for block := 0; block < 150; block += 16 {
		probeB := make([]byte, 0)
		for i := 0; i < blockSize - 1; i++ {
			probeB = append(probeB, byte('A'))
		}

		for i := 0; i < 16; i++ {	
			probe := JMT.ParseToAscii(probeB, false)
			target := server12(bc, probe, secretString)
			keyRune := byte(0)
			tempProbe := append(probeB, newWord...)

			for b:= byte(0); b < byte(255); b++ {
				probeW := JMT.ParseToAscii( append(tempProbe, b), false )
				probeOut := server12(bc, probeW, secretString)
				if (compareSlices(probeOut[block:(block + 16)],
				                  target[block:(block + 16)])) {
					keyRune = b
					break
				}
			}

			newWord = append(newWord, keyRune)
			if len(probeB) >= 1 {
				probeB = probeB[1:]
			} else {
				probeB = make([]byte, 0)
			}
		}
	}

	fmt.Println(JMT.ParseToAscii(newWord, false))
	return
}

// // foo=bar&baz=qux&zap=zazzle
func ParseCookie(s string) map[string]string {
	m := make(map[string]string)
	ps := strings.Split(s, "&")
	for _, p := range ps {
		qs := strings.Split(p, "=")
		m[qs[0]] = qs[1]
	}

	return m
}

// //email=foo@bar.com&uid=10&role=user
func profileFor(email string) string {
	email = strings.Replace(email, "=", "", -1)
	email = strings.Replace(email, "&", "", -1)

	return "email=" + string(email) + "&" +
	       "uid=" + "10" + "&" +
	       "role=" + "user"
}

func DeparseCookie(m map[string]string) string {
	s := ""
	for k, v := range m {
		s += (k + "=" + v + "&")
	}

	s = s[:len(s)-1] // drop final &

	return s
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
  b := make([]rune, n)
  for i := range b {
      b[i] = letters[rand.Intn(len(letters))]
  }
  return string(b)
}

func server14(bc JMT.BlockCipher, prefix, msg, secret string) []byte {
	// plaintext := append(JMT.ParseFromAscii(prefix, false),
	// 									  JMT.ParseFromAscii(msg, false)...)
	// plaintext = append(plaintext,
	// 								   JMT.ParseFromAscii(secret, false)...)

	plaintext := JMT.ParseFromAscii(prefix + msg + secret, false)

	cipherText := JMT.ECBEncrypt(bc, plaintext)

	return cipherText
}

func findPrefixLength(bc JMT.BlockCipher, prefix, target string) (int, int) {
	for i := 0; i < 1000; i++ {
		probe := ""
		for j := 0; j < i; j++ {
			probe += "A"
		}

		cipherText := server14(bc, prefix, probe, target)

  	if isECB2(cipherText) {
  		return i, (len(probe) - i)
  	}
	}

	return -1, -1
}

func byteSplit(s string) string {
	for i := 32; i < len(s); i += 33 {
		s = s[:i] + "-" + s[i:]
	}
	return s
}

func PaddingAttack2(bc JMT.BlockCipher, prefix, secretString string,
	                  probeSize int) {
	target_1 := server14(bc, prefix, "", secretString)
	target_2 := server14(bc, prefix, "A", secretString)

	firstDiff := 0
	for firstDiff = 0; firstDiff < len(target_1); firstDiff++ {
		if target_1[firstDiff] != target_2[firstDiff] {
			break
		}
	}
	
	k := firstDiff / 16

	newWord := make([]byte, 0)
	for block := (k+2)*16; block < (k+4)*16; block += 16 {
		probeB := make([]byte, 0)
		for i := 0; i < probeSize - 1; i++ {
			probeB = append(probeB, byte('A'))
		}

		for i := 0; i < 16; i++ {	
			probe := JMT.ParseToAscii(probeB, false)
			target := server14(bc, prefix, probe, secretString)


			keyRune := byte(0)
			tempProbe := append(probeB, newWord...)

			for b:= byte(0); b < byte(255); b++ {
				probeW := JMT.ParseToAscii( append(tempProbe, b), false )
				probeOut := server14(bc, prefix, probeW, secretString)

				if (compareSlices(probeOut[block:(block + 16)],
				                  target[block:(block + 16)])) {
					keyRune = b
					break
				}
			}

			newWord = append(newWord, keyRune)
			if len(probeB) >= 1 {
				probeB = probeB[1:]
			} else {
				probeB = make([]byte, 0)
			}
		}
	}

	fmt.Println(JMT.ParseToAscii(newWord, false))
	return
}

func server16(bc JMT.BlockCipher, iv [4]JMT.Word, msg string) []byte {
	// remove ; and = runes
	msg = strings.Replace(msg, "=", "", -1)
	msg = strings.Replace(msg, ";", "", -1)	  

	plaintext := JMT.ParseFromAscii("comment1=cooking%20MCs;userdata=" + 
		                              msg + 
		                              ";comment2=%20like%20a%20pound%20of%20bacon",
		                              false)

	cipherText := JMT.CBCEncrypt(bc, iv, plaintext)

	return cipherText
}

func decode16(bc JMT.BlockCipher, iv [4]JMT.Word, cipherText []byte) {
	plainText := JMT.ParseToAscii( JMT.CBCDecrypt(bc, iv, cipherText), true )

	ps := strings.Split(plainText, ";")
	fmt.Println("")
	for _, p := range ps {
		qs := strings.Split(p, "=")
		fmt.Printf("%s : %s\n", qs[0], qs[1])
	}

	return
}

func Week2() {
	// 2-10
 	key_10 := JMT.BytesToWords(JMT.ParseFromAscii("YELLOW SUBMARINE", false),
		                        false)
 	aes_10 := JMT.MakeAES(key_10)

	var iv_10 [4]JMT.Word

	cipherText_10, _ := ReadStrFile2("./inputs/10.txt")

	plaintext_10 := JMT.CBCDecrypt(aes_10, iv_10, cipherText_10)
	fmt.Println(JMT.ParseToAscii(plaintext_10, false))

	// 2-11
	for i := 0; i < 100; i++ {
		key_11 := JMT.RandomBlock(4)
		aes_11 := JMT.MakeAES(key_11)
		var iv_11 [4]JMT.Word
		copy(iv_11[:], JMT.RandomBlock(4))

		plaintext_Bytes := JMT.ParseFromAscii(bigTextMessage, false)
	  s1 := rand.NewSource(time.Now().UnixNano())
	  r1 := rand.New(s1)
	  extra := make([]byte, r1.Intn(6) + 5)

	  plaintext_Bytes = append(extra, plaintext_Bytes...)
	  plaintext_Bytes = append(plaintext_Bytes, extra...)
	  plaintext_11 := plaintext_Bytes //JMT.BytesToWords(plaintext_Bytes, true)
	  
		var cipherText_11 []byte
		flag := r1.Intn(2)
  	if flag == 0 {
  		cipherText_11 = JMT.ECBEncrypt(aes_11, plaintext_11)
  	} else {
  		cipherText_11 = JMT.CBCEncrypt(aes_11, iv_11, plaintext_11)
  	}

  	temp := JMT.ParseToAscii(cipherText_11, true)
  	if isECB(temp) != (flag == 0) {
  		fmt.Printf("ERROR!\n")
  	}
  }

  // cryptopals 2-12
	key_12 := JMT.RandomBlock(4)
	aes_12 := JMT.MakeAES(key_12)

  secretString := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
 	
  PaddingAttack(aes_12, secretString, 16)

  // cryptopals 13
  key_13 := JMT.RandomBlock(8)
	aes_13 := JMT.MakeAES(key_13)

	// email=Richard_Mayhew@cryptopals.com&uid=10&role=user
	// This email address was specifically chosen to fill 3 blocks of the cookie
	cookie_13 := profileFor("Richard_Mayhew@cryptopals.com")

	cipherText_13 := JMT.ECBEncrypt(aes_13,
		                              JMT.ParseFromAscii(cookie_13, false))

	plaintext_13_1 := JMT.ParseToAscii(JMT.ECBDecrypt(aes_13, cipherText_13),
	                                   true)

	// we want to swap out the last block with the encryption of
	// admin[0x11]x11 role=adminXXXXXX
	// role=01234567890 <- one block
	// role=01234567890||admin\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11
	secret := JMT.ParseFromAscii("0123456789admin", false)
	secret = append(secret, []byte{11,11,11,11,11,11,11,11,11,11,11}...)
	sString := JMT.ParseToAscii(secret, false)
	insert := profileFor(sString)

	dodgyCText := JMT.ECBEncrypt(aes_13, JMT.ParseFromAscii(insert, true))
	hackedCtext := append(cipherText_13[0:48], dodgyCText[16:32]...)
	hackedCookie := JMT.ECBDecrypt(aes_13, hackedCtext)

	fmt.Println(plaintext_13_1)
	fmt.Println(JMT.ParseToAscii(hackedCookie, true))


 	// cryptopal 14
  key_14 := JMT.RandomBlock(6)
	aes_14 := JMT.MakeAES(key_14)

 	// random string with between 32 and 128 characters
	s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)	
 	randomPrefix := randString(r1.Intn(96) + 32)

 	suffix := "Hello David, how are you?"

 	padSize, _ := findPrefixLength(aes_14, randomPrefix, suffix)

 	PaddingAttack2(aes_14, randomPrefix, suffix, padSize)


 	// cryptopal 16
  key_16 := JMT.RandomBlock(6)
	aes_16 := JMT.MakeAES(key_16)

	var iv_16 [4]JMT.Word
	temp := JMT.RandomBlock(4)
	copy(iv_16[:], temp) 

	encypted_16 := server16(aes_16, iv_16, "_WORD_role_admin")
	decode16(aes_16, iv_16, encypted_16)

	encypted_16[16+0] = byte('=') ^ byte('_') ^ encypted_16[16+0]
	encypted_16[16+5] = byte(';') ^ byte('_') ^ encypted_16[16+5]
	encypted_16[16+10] = byte('=') ^ byte('_') ^ encypted_16[16+10]

	decode16(aes_16, iv_16, encypted_16)
}
