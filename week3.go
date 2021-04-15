package main

import (
	"fmt"
	"math/rand"
	"time"
	// "strings"
)

import JMT "github.com/mawir157/jmtcrypto"
import JMTR "github.com/mawir157/jmtcrypto/rand"

func Week3() {
	printDay(17)
	aes_key_17 := JMT.RandomBlock(8)
	aes_17 := JMT.MakeAES(aes_key_17)	
	
	var iv_17 [4]JMT.Word
	temp := JMT.RandomBlock(4)
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

	hack := CBCPadAttack(iv_17, aesCipherText_17, aes_17)
	fmt.Println(JMT.ParseToAscii(hack, true))

	secret, _ := JMT.CBCDecrypt(aes_17, iv_17, aesCipherText_17)
	fmt.Println(JMT.ParseToAscii(secret, true))

	printDay(18)
	CipherText_18 := JMT.ParseFromBase64(
		"L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==",
		false)
	key_18 := JMT.BytesToWords(JMT.ParseFromAscii("YELLOW SUBMARINE", false),
		                         false)
	aes_18 := JMT.MakeAES(key_18)
	nonce_18 := []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}

	DecodedText_18, _ := JMT.CTRDecrypt(aes_18, nonce_18, CipherText_18)

	fmt.Println(JMT.ParseToAscii(DecodedText_18, false))

	printDay(19)
	strings_19 := []string{
		"SSBoYXZlIG1ldCB0aGVtIGF0IGNsb3NlIG9mIGRheQ==",
		"Q29taW5nIHdpdGggdml2aWQgZmFjZXM=",
		"RnJvbSBjb3VudGVyIG9yIGRlc2sgYW1vbmcgZ3JleQ==",
		"RWlnaHRlZW50aC1jZW50dXJ5IGhvdXNlcy4=",
		"SSBoYXZlIHBhc3NlZCB3aXRoIGEgbm9kIG9mIHRoZSBoZWFk",
		"T3IgcG9saXRlIG1lYW5pbmdsZXNzIHdvcmRzLA==",
		"T3IgaGF2ZSBsaW5nZXJlZCBhd2hpbGUgYW5kIHNhaWQ=",
		"UG9saXRlIG1lYW5pbmdsZXNzIHdvcmRzLA==",
		"QW5kIHRob3VnaHQgYmVmb3JlIEkgaGFkIGRvbmU=",
		"T2YgYSBtb2NraW5nIHRhbGUgb3IgYSBnaWJl",
		"VG8gcGxlYXNlIGEgY29tcGFuaW9u",
		"QXJvdW5kIHRoZSBmaXJlIGF0IHRoZSBjbHViLA==",
		"QmVpbmcgY2VydGFpbiB0aGF0IHRoZXkgYW5kIEk=",
		"QnV0IGxpdmVkIHdoZXJlIG1vdGxleSBpcyB3b3JuOg==",
		"QWxsIGNoYW5nZWQsIGNoYW5nZWQgdXR0ZXJseTo=",
		"QSB0ZXJyaWJsZSBiZWF1dHkgaXMgYm9ybi4=",
		"VGhhdCB3b21hbidzIGRheXMgd2VyZSBzcGVudA==",
		"SW4gaWdub3JhbnQgZ29vZCB3aWxsLA==",
		"SGVyIG5pZ2h0cyBpbiBhcmd1bWVudA==",
		"VW50aWwgaGVyIHZvaWNlIGdyZXcgc2hyaWxsLg==",
		"V2hhdCB2b2ljZSBtb3JlIHN3ZWV0IHRoYW4gaGVycw==",
		"V2hlbiB5b3VuZyBhbmQgYmVhdXRpZnVsLA==",
		"U2hlIHJvZGUgdG8gaGFycmllcnM/",
		"VGhpcyBtYW4gaGFkIGtlcHQgYSBzY2hvb2w=",
		"QW5kIHJvZGUgb3VyIHdpbmdlZCBob3JzZS4=",
		"VGhpcyBvdGhlciBoaXMgaGVscGVyIGFuZCBmcmllbmQ=",
		"V2FzIGNvbWluZyBpbnRvIGhpcyBmb3JjZTs=",
		"SGUgbWlnaHQgaGF2ZSB3b24gZmFtZSBpbiB0aGUgZW5kLA==",
		"U28gc2Vuc2l0aXZlIGhpcyBuYXR1cmUgc2VlbWVkLA==",
		"U28gZGFyaW5nIGFuZCBzd2VldCBoaXMgdGhvdWdodC4=",
		"VGhpcyBvdGhlciBtYW4gSSBoYWQgZHJlYW1lZA==",
		"QSBkcnVua2VuLCB2YWluLWdsb3Jpb3VzIGxvdXQu",
		"SGUgaGFkIGRvbmUgbW9zdCBiaXR0ZXIgd3Jvbmc=",
		"VG8gc29tZSB3aG8gYXJlIG5lYXIgbXkgaGVhcnQs",
		"WWV0IEkgbnVtYmVyIGhpbSBpbiB0aGUgc29uZzs=",
		"SGUsIHRvbywgaGFzIHJlc2lnbmVkIGhpcyBwYXJ0",
		"SW4gdGhlIGNhc3VhbCBjb21lZHk7",
		"SGUsIHRvbywgaGFzIGJlZW4gY2hhbmdlZCBpbiBoaXMgdHVybiw=",
		"VHJhbnNmb3JtZWQgdXR0ZXJseTo=",
		"QSB0ZXJyaWJsZSBiZWF1dHkgaXMgYm9ybi4=",
	}

	cipherTexts_19 := [][]byte{}
	for _, s := range strings_19 {
		temp := JMT.ParseFromBase64(s, false)		
		cipherTexts_19 = append(cipherTexts_19, temp)                                
	}

	fixedNonceAttack(cipherTexts_19)

	printDay(20)
	input_20 ,_ := ReadStrFile("./inputs/20.txt")
	strings_20 := [][]byte{}
	for _, i_20 := range input_20 {
		temp := JMT.ParseFromBase64(i_20, false)		
		strings_20 = append(strings_20, temp) 
	}

	fixedNonceAttack(strings_20)

	printDay(21)

	printDay(22)
	seed, val := MT19937Pair(2)
	fmt.Printf("Find a seed that produces - %d\n", val)
	hack_22 := CrackMT19937(val)
	fmt.Printf("The seed is %d, our guess is %d\n", seed, hack_22)

	printDay(23)

	rng_23 := JMTR.Mersenne19937Init()
	rng_23.Seed(r1.Intn(2e8))

	cloneMersenne(rng_23)

	printDay(24)
	pcg := JMTR.PCGInit()
	pcg.Seed(5489)
	pcg2 := JMTR.PCGInit()
	pcg2.Seed(5489 + 1)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d|%d\n", pcg.Next(), pcg2.Next())
	}

}

func singleBlockAttack(block []byte, bc JMT.BlockCipher) []byte {
	zeroedIV := make([]byte, 16)

	for padByte := 1; padByte <= 16; padByte++ {
		paddingIV := make([]byte, 16)
		for i := 0; i < len(paddingIV); i++ {
			paddingIV[i] = zeroedIV[i] ^ byte(padByte)
		}

		for b := byte(0x00); b < byte(0xFF); b++ {
			paddingIV[16 - padByte] = b

			var temp_iv [4]JMT.Word
			copy(temp_iv[:], JMT.BytesToWords(paddingIV, false))
			_, err := JMT.CBCDecrypt(bc, temp_iv, block)
			if (err == nil) {
				zeroedIV[16 - padByte] = b ^ byte(padByte)
				break
			}
		}
	}

	return zeroedIV
}

func CBCPadAttack(iv [4]JMT.Word, ct []byte, bc JMT.BlockCipher) []byte {
	noBlock := len(ct) / 16

	ivBytes := []byte{}
	for _, w := range iv {
		for _, b := range w {
			ivBytes = append(ivBytes, b)
		}
	}

	ct = append(ivBytes, ct...)
	result := []byte{}
	for block := 1; block <= noBlock; block++ {
		ctBlock := ct[block*16:(block + 1)*16]
		dec := singleBlockAttack(ctBlock, bc)

		pt := make([]byte, 16)
		for i := 0; i < 16; i++ {
			pt[i] = ct[(block - 1)*16 + i] ^ dec[i]
		}

		result = append(result, pt...)
	}

	return result
}

func isIn(b byte, bs []byte) bool {
	for _, v := range bs {
		if b == v {
			return true
		}
	}
	return false
}

func isSystem(b byte) bool {
	return (b < 0x20)
}

// ETAOIN SHRDLU
func isCommonLetter(b byte) bool {
	return  (b == 0x20) ||                // [SPACE]
	        (b == 0x45) || (b == 0x65) || // E
	        (b == 0x54) || (b == 0x74) || // T
	        (b == 0x41) || (b == 0x61) || // A
	        (b == 0x4F) || (b == 0x6F) || // O
	        (b == 0x49) || (b == 0x69) || // I
	        (b == 0x4E) || (b == 0x6E) || // N
	        (b == 0x53) || (b == 0x73) || // S
	        (b == 0x48) || (b == 0x68) || // H
	        (b == 0x52) || (b == 0x72) || // R
	        (b == 0x44) || (b == 0x64) || // D
	        (b == 0x4C) || (b == 0x6C) || // L
	        (b == 0x55) || (b == 0x75)    // U
}

func mostCommon(m map[byte]int) byte {
	bestK := byte(0)
  bestV := 0
	for k, v := range m {
		if v > bestV {
			bestV = v
			bestK = k
		}
	}

	return bestK
}

func fixedNonceAttack(cTexts [][]byte) {
	bigN := 0
	for _, bs := range cTexts {
		if len(bs) > bigN {
			bigN = len(bs)
		}
	}

	block := make([]byte, bigN)
	for i := 0; i < bigN; i++ {
		freqChar := make(map[byte]int)

		for b := 0; b < 256; b++ {
			for _, s := range cTexts {
				if i >= len(s) {
					continue
				}

				if isCommonLetter(byte(b) ^ s[i]) {
					freqChar[byte(b)]++
				} else if isSystem(byte(b) ^ s[i]) {
					freqChar[byte(b)] -= 40
				}
			}
		}
		block[i] = mostCommon(freqChar)
	}

	for j, s := range cTexts {
		newBytes := make([]byte, bigN)
		for i := 0; i < bigN; i++ {
			if i >= len(s) {
				continue
			}
			newBytes[i] = s[i] ^ block[i]
		}
		fmt.Printf("%d NEW: %s\n", j, JMT.ParseToAscii(newBytes, false))
	}
	return	
}

func MT19937Pair(delay int) (int, int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// wait for some time
	fmt.Println("Waiting...")
	time.Sleep(time.Duration(5 + r1.Intn(delay)) * time.Second)
	fmt.Println("Done.")

	t := time.Now()
	seed := int(t.Unix())
	rng := JMTR.Mersenne19937Init()
	rng.Seed(seed)

	// wait for some time
	fmt.Println("Waiting...")
	time.Sleep(time.Duration(5 + r1.Intn(delay)) * time.Second)
	fmt.Println("Done.")

	val := rng.Next()
	return seed, val
}

func CrackMT19937(target int) int {
	// we know the rng is seeded with a time stamp
	// so we start at now and run back through time until we get a hit
	t := time.Now()
	seedGuess := int(t.Unix())

	for true {
		rng := JMTR.Mersenne19937Init()
		rng.Seed(seedGuess)
		val := rng.Next()

		if val == target {
			break
		}
		seedGuess-- // try one second earlier
	}
	return seedGuess
}

// we don't know anything about the seed used in rng
func cloneMersenne(rng *JMTR.Mersenne19937) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// run the rng for some number of ticks
	for i := 0; i < r1.Intn(1000); i++ {
		rng.Next()
	}

	arr := []int{}
	for i := 0; i < 624; i++ {
		nxt := rng.Next()
		arr = append(arr, JMTR.UnTwist(nxt))
	}

	// make a new rng and splice the untempered arr into it
	clonedRNG := JMTR.Mersenne19937Init()
	clonedRNG.Seed(0)
	clonedRNG.Splice(arr)

	// tick the cloned rng forward 624 steps
	for i := 0; i < 624; i++ {
		clonedRNG.Next()
	}

	// compare the clone to the original
	for i := 0; i < 25; i++ {
		lhs := rng.Next()
		rhs := clonedRNG.Next()
		fmt.Printf("%d ", lhs ^ rhs)
	}
	fmt.Println()
}
