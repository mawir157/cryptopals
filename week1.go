package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

import JMT "github.com/mawir157/jmtcrypto"

func printDay(day int) {
	fmt.Printf("\n----------- PART %d ----------\n", day)
}

func ReadStrFile(fname string) (strs []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil { return nil, err }

	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 { continue }
		strs = append(strs, l)
	}

	return strs, nil
}

func ReadStrFile2(fname string) (strs []byte, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil { return nil, err }
	strs = JMT.ParseFromBase64(strings.Replace(string(b), "\n", "", -1), false)
	return strs, nil
}

// func ReadStrFile3(fname string) (strs []byte, err error) {
// 	b, err := ioutil.ReadFile(fname)
// 	if err != nil { return nil, err }
// 	strs = JMT.ParseFromBase64(strings.Replace(string(b), "\n", "", -1), false)
// 	return strs, nil
// }

func isEnglish(test []byte, spaces bool) bool {
	count := 0
	seenSpace := !spaces
	for _, p := range test {
		if ((0x41 <= p) && (p <= 0x5a)) ||
		   ((0x61 <= p) && (p <= 0x7a)) ||
		   (p == 0x20) {
			count++
		}
		seenSpace = seenSpace || (p == 0x20)
	}

	return (10*count > 8*len(test)) && seenSpace
}
// ETAOIN SHRDLU
func englishCount(test []byte) int {
	count := 0
	for _, p := range test {
		if ((0x41 <= p) && (p <= 0x5a)) ||
		   ((0x61 <= p) && (p <= 0x7a)) ||
		   (p == 0x20) {
			count++
		} else {
			count--
		}
	}
	return count
}

func bitsSetCount(b byte) (count int) {
	count = 0
	for i := 0; i < 8; i++ {
		count += int(b & byte(1))
		b >>= 1
	}

	return
}

func hammingDistance(bs1, bs2 []byte) (hd int) {
	hd = 0
	for i, _ := range bs1 {
		temp := bs1[i] ^ bs2[i]
		hd += bitsSetCount(temp)
	}

	return
}

func solveVigenere(cTxt []byte) (key, message string) {
	bestKey := -1
	bestHD := 10000.0
	for keySize := 2; keySize < 100; keySize++ {
		hamDist := 0
		for k := 0; k < 10; k++ {
			hamDist += hammingDistance(cTxt[(k*keySize):((k+1)*keySize)],
			                           cTxt[((k+1)*keySize):((k+2)*keySize)])
		}
		normHD := float64(hamDist) / (float64(10) * float64(keySize))

		if normHD < bestHD {
			bestHD = normHD
			bestKey = keySize
		}
	}

	key_bytes := make([]byte, bestKey)
	for i := 0; i < bestKey; i++ {
		subWord := make([]byte, 0)
		for j := i; j < len(cTxt); j += bestKey {
			subWord = append(subWord, cTxt[j])
		}
		bestChar := byte(0)
		bestScore := -1000
		for j := 0; j < 256; j++ {
			bytes := make([]byte, 0)
			for _, char := range subWord {
				b := byte(char) ^ byte(j)
				bytes = append(bytes, b)
			}

			if englishCount(bytes) > bestScore {
				bestScore = englishCount(bytes)
				bestChar = byte(j)
			}
		}		
		key_bytes[i] = bestChar
	}

	j := 0
	for i := 0; i < len(cTxt); i++ {
		cTxt[i] = cTxt[i] ^ key_bytes[j]

		j++
		if j == len(key_bytes) {
			j = 0
		}
	}


	key = JMT.ParseToAscii(key_bytes, false)
	message = JMT.ParseToAscii(cTxt, false)

	return 
}

func isECB(cipherText string) bool {
	table := make(map[string]int)

	for i := 0; i < len(cipherText) - 16; i += 16 {
		temp := cipherText[i:(i+16)]
		table[temp] += 1
	}

	for _,v := range table {
		if v > 1 {
			return true
		}
	}

	return false
}

func isECB2(cipherText []byte) bool {
	table := make(map[[16]byte]int)
	temp := [16]byte{}
	for i := 0; i < len(cipherText) - 16; i += 16 {
		// temp := cipherText[i:(i+16)]
		copy(temp[:], cipherText[i:(i+16)])
		table[temp] += 1
	}

	for _,v := range table {
		if v > 1 {
			return true
		}
	}

	return false
}

func Week1() {
	// cryptopals 1-1
	printDay(1)
	str_1   := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	str_1_2 := JMT.ParseToBase64(JMT.ParseFromHex(str_1, false))

	fmt.Println(str_1_2)
	fmt.Println("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")
	fmt.Println("")

	// cryptopals 1-2
	printDay(2)
	str_2_1 := "1c0111001f010100061a024b53535009181c"
	str_2_2 := "686974207468652062756c6c277320657965"
	str_2_3 := JMT.ParseToHex(
		JMT.ByteStreamXOR( JMT.ParseFromHex(str_2_1, false),
		                   JMT.ParseFromHex(str_2_2, false) ) )

	fmt.Println(str_2_3)
	fmt.Println("746865206b696420646f6e277420706c6179")
	fmt.Println("")

	// cryptopals 1-3
	printDay(3)
	str_3 := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bytes_3 := JMT.ParseFromHex(str_3, false)
	for i:= 0; i < 256; i++ {
		bytes := make([]byte, 0)
		for _, char := range bytes_3 {
			b := byte(char) ^ byte(i)
			bytes = append(bytes, b)
		}
		// if (ok && seenSpace) {
		if isEnglish(bytes, true) {
			fmt.Printf("%02X: %s\n", i, JMT.ParseToAscii(bytes, false))
		}
	}
	fmt.Println("")

	// cryptopals 1-4
	printDay(4)
	cipherTexts_4,_ := ReadStrFile("./inputs/4.txt")
	for j, ct := range cipherTexts_4 {
		bytes_4 := JMT.ParseFromHex(ct, false)
		for i := 0; i < 256; i++ {
			bytes := make([]byte, 0)
			for _, char := range bytes_4 {
				b := byte(char) ^ byte(i)
				bytes = append(bytes, b)
			}
			if isEnglish(bytes, true) {
				fmt.Printf("%03d||%02X: %s\n", j, i, JMT.ParseToAscii(bytes, false))
			}
		}		
	}
	fmt.Println("")

	// cryptopals 1-5
	printDay(5)
	plaintext_5 := JMT.ParseFromAscii(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`, false)
	key_5 := JMT.ParseFromAscii("ICE", false)
	cipherText_5 := make([]byte, len(plaintext_5))
	j := 0
	for i := 0; i < len(plaintext_5); i++ {
		cipherText_5[i] = plaintext_5[i] ^ key_5[j]

		j++
		if j == len(key_5) {
			j = 0
		}
	}
	fmt.Println(JMT.ParseToHex(cipherText_5))
	fmt.Println("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	fmt.Println("")

	// cryptopals 1-6
	printDay(6)
	cipherText_6, _ := ReadStrFile2("./inputs/6.txt")
	key_6, plaintext_6 := solveVigenere(cipherText_6)
	fmt.Println(key_6)
	fmt.Println(plaintext_6)
	fmt.Println("")

	// cryptopals 1-7
	printDay(7)
	cipherText_7,_ := ReadStrFile2("./inputs/7.txt")
	key_7 := JMT.BytesToWords(JMT.ParseFromAscii("YELLOW SUBMARINE", false),
		                        false)
	aes_7 := JMT.MakeAES(key_7)
	aesDecodedText_7, _ := JMT.ECBDecrypt(aes_7, cipherText_7)
	fmt.Println(JMT.ParseToAscii(aesDecodedText_7, false))
	fmt.Println("")

	// cryptopals 1-8
	printDay(8)
	input_8,_ := ReadStrFile("./inputs/8.txt")
	for _, cTxt := range input_8 {
		if isECB(cTxt) {
			fmt.Println(cTxt)
		}
	}
}