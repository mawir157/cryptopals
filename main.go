package main

import (
	"fmt"
	"math/rand"
	"time"
)

import JMT "github.com/mawir157/jmtcrypto"
import JMTR "github.com/mawir157/jmtcrypto/rand"
import MCe "github.com/mawir157/jmtcrypto/mce"

func blockTitle(s string) {
	fmt.Printf("###############################################################################\n")
	fmt.Printf("##\n")
	fmt.Printf("## %s\n", s)
	fmt.Printf("##\n")

	return
}

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
	blockTitle("McEliese")
	public, private := MCe.GenerateKeyPair(2, 11)

	public.Write("mce.pub")
	private.Write("mce.pri")

	public2 := MCe.ReadPublic("mce.pub")
	private2 := MCe.ReadPrivate("mce.pri")

	cipherText := public2.Encrypt(textMessage)
	
	plaintext := private2.Decrypt(cipherText)
	MCe.PrintAscii(plaintext, true)
	fmt.Println("")


	//////////////////////////////////////////////////////////////////////////////
	blockTitle("ECB w/ AES")
	aes_key := make([]byte, 16)
	rand.Read(aes_key)
	aes := JMT.MakeAES(aes_key)

	aesPlainText, _   := JMT.ParseFromAscii(textMessage, true)
	aesCipherText     := JMT.ECBEncrypt(aes, aesPlainText)
	aesDecodedText, _ := JMT.ECBDecrypt(aes, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("CBC w/ AES")
	rand.Read(aes_key)
	aes = JMT.MakeAES(aes_key)
	iv := make([]byte, 16)
	rand.Read(iv)

	aesCipherText     = JMT.CBCEncrypt(aes, iv, aesPlainText)
	aesDecodedText, _ = JMT.CBCDecrypt(aes, iv, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("PCB w/ AES")
	rand.Read(aes_key)
	aes = JMT.MakeAES(aes_key)
	rand.Read(iv)

	aesCipherText     = JMT.PCBCEncrypt(aes, iv, aesPlainText)
	aesDecodedText, _ = JMT.PCBCDecrypt(aes, iv, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("OFB w/ AES")
	rand.Read(aes_key)
	aes = JMT.MakeAES(aes_key)
	rand.Read(iv)

	aesCipherText     = JMT.OFBEncrypt(aes, iv, aesPlainText)
	aesDecodedText, _ = JMT.OFBDecrypt(aes, iv, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("CTR w/ AES")
	rand.Read(aes_key)
	aes = JMT.MakeAES(aes_key)
	nonce := make([]byte, 8)
	rand.Read(nonce)

	aesPlainTextNoPad, _   := JMT.ParseFromAscii(textMessage, false)
	aesCipherText     = JMT.CTREncrypt(aes, nonce, aesPlainTextNoPad)
	aesDecodedText, _ = JMT.CTRDecrypt(aes, nonce, aesCipherText)
	
	// we don't need to remove padding for a stream cipher
	fmt.Println(JMT.ParseToAscii(aesDecodedText, false))
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("CFB w/ AES")
	rand.Read(aes_key)
	aes = JMT.MakeAES(aes_key)
	rand.Read(iv)

	aesCipherText     = JMT.CFBEncrypt(aes, iv, aesPlainText)
	aesDecodedText, _ = JMT.CFBDecrypt(aes, iv, aesCipherText)

	fmt.Println(JMT.ParseToAscii(aesDecodedText, true))
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("PRNG Stream Encyption w/ Mersenne Twister")
	mt := JMTR.Mersenne19937Init()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	seed := r1.Intn(2000000000)
	_, streamCipherText := JMT.PRNGStreamEncode(seed, mt, aesPlainTextNoPad)
	streamDecodedText := JMT.PRNGStreamDecode(seed, mt, streamCipherText)

	// we don't need to remove padding for a stream cipher
	fmt.Println(JMT.ParseToAscii(streamDecodedText, false))
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("PRNG Stream Encyption w/ Permuted congruential generator")
	pcg := JMTR.PCGInit()
	seed, streamCipherText = JMT.PRNGStreamEncode(0, pcg, aesPlainTextNoPad)
	streamDecodedText = JMT.PRNGStreamDecode(seed, pcg, streamCipherText)

	// we don't need to remove padding for a stream cipher
	fmt.Println(JMT.ParseToAscii(streamDecodedText, false))
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("Encrypt then MAC w/ AES+ECB")
	rand.Read(aes_key)
	aes = JMT.MakeAES(aes_key)
	hash := JMT.MakeSHA256()
	key2 := make([]byte, 32)
	rand.Read(key2)
	extras := make(map[string]([]byte))
	aeCipher := JMT.EtMEncrypt(aesPlainText, aes, hash, JMT.ECB, key2, extras)
	aeDecoded, err := JMT.EtMDecrypt(aeCipher, aes, hash, JMT.ECB, key2, extras)

	if err != nil { 
		fmt.Println(err)
	} else {
		fmt.Println(JMT.ParseToAscii(aeDecoded, true))
	}
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("Encrypt-and-MAC w/ AES+ECB")
	rand.Read(aes_key)
	aes = JMT.MakeAES(aes_key)
	aeCipher = JMT.EaMEncrypt(aesPlainText, aes, hash, JMT.ECB, extras)
	aeDecoded, err = JMT.EaMDecrypt(aeCipher, aes, hash, JMT.ECB, extras)

	if err != nil { 
		fmt.Println(err)
	} else {
		fmt.Println(JMT.ParseToAscii(aeDecoded, true))
	}
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("MAC-then-Encrypt w/ AES+ECB")
	rand.Read(aes_key)
	aes = JMT.MakeAES(aes_key)
	aeCipher = JMT.MtEEncrypt(aesPlainText, aes, hash, JMT.ECB, extras)
	aeDecoded, err = JMT.MtEDecrypt(aeCipher, aes, hash, JMT.ECB, extras)

	if err != nil { 
		fmt.Println(err)
	} else {
		fmt.Println(JMT.ParseToAscii(aeDecoded, true))
	}
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("Encrypt then MAC w/ Camellia+CBC")
	rand.Read(aes_key)
	rand.Read(iv)
	rand.Read(key2)
	cam := JMT.MakeCamellia(aes_key)
	extras["iv"] = iv
	aeCipher = JMT.EtMEncrypt(aesPlainText, cam, hash, JMT.CBC, key2, extras)
	aeDecoded, err = JMT.EtMDecrypt(aeCipher, cam, hash, JMT.CBC, key2, extras)

	if err != nil { 
		fmt.Println(err)
	} else {
		fmt.Println(JMT.ParseToAscii(aeDecoded, true))
	}
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("Encrypt-and-MAC w/ Camellia+CBC")
	rand.Read(aes_key)
	cam = JMT.MakeCamellia(aes_key)
	rand.Read(extras["iv"])
	aeCipher = JMT.EaMEncrypt(aesPlainText, cam, hash, JMT.CBC, extras)
	aeDecoded, err = JMT.EaMDecrypt(aeCipher, cam, hash, JMT.CBC, extras)

	if err != nil { 
		fmt.Println(err)
	} else {
		fmt.Println(JMT.ParseToAscii(aeDecoded, true))
	}
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////
	blockTitle("MAC-then-Encrypt w/ Camellia+CBC")
	rand.Read(aes_key)
	cam = JMT.MakeCamellia(aes_key)
	rand.Read(extras["iv"])
	aeCipher = JMT.MtEEncrypt(aesPlainText, cam, hash, JMT.CBC, extras)
	aeDecoded, err = JMT.MtEDecrypt(aeCipher, cam, hash, JMT.CBC, extras)

	if err != nil { 
		fmt.Println(err)
	} else {
		fmt.Println(JMT.ParseToAscii(aeDecoded, true))
	}
	fmt.Println("")

	// Week1()
	// Week2()
	// Week3()
/*
	cipher_25_a, _ := ReadStrFile2("./inputs/25.txt")
	key_ecb_25 := JMT.BytesToWords(JMT.ParseFromAscii("YELLOW SUBMARINE", false),
		                         false)
	aes_ecb_25 := JMT.MakeAES(key_ecb_25)
	plainText_25, _ := JMT.ECBDecrypt(aes_ecb_25, cipher_25_a)

	aes_ctr_key := JMT.RandomBlock(8)
	aes_ctr_25 := JMT.MakeAES(aes_ctr_key)
	// this implentation of ctr assumes the nonce is zero
	nonce_25 := make([]byte, 8)

	ctrCipherText_25 := JMT.CTREncrypt(aes_ctr_25, nonce_25, plainText_25)
	test_25 := JMT.CTREncrypt(aes_ctr_25, nonce_25, ctrCipherText_25)
	fmt.Println(JMT.ParseToAscii(test_25, false))
	secret := JMT.ParseFromAscii("---]]]Hack the Planet[[[---", false)
	HackedCipherText_25 := ctrEdit(ctrCipherText_25, aes_ctr_key, 250, secret)
	fmt.Println(ctrCipherText_25[:80])
	qqq, _ := JMT.CTRDecrypt(aes_ctr_25, nonce_25, HackedCipherText_25)
	fmt.Println(JMT.ParseToAscii(qqq, false))
*/

	// dh := DH.DiffHell(*big.NewInt(77377), *big.NewInt(5))

/*
	dh := DH.DiffHell(DH.ParseToBigIntHex("ffffffffffffffffc90fdaa22168c234c4c6628b80dc1cd129024e088a67cc74020bbea63b139b22514a08798e3404ddef9519b3cd3a431b302b0a6df25f14374fe1356d6d51c245e485b576625e7ec6f44c42e9a637ed6b0bff5cb6f406b7edee386bfb5a899fa5ae9f24117c4b1fe649286651ece45b3dc2007cb8a163bf0598da48361c55d39a69163fa8fd24cf5f83655d23dca3ad961c62f356208552bb9ed529077096966d670c354e4abc9804f1746c08ca237327ffffffffffffffff"),
										DH.ParseToBigIntHex("2"))

	a, b := big.NewInt(int64(r1.Intn(37))), big.NewInt(int64(r1.Intn(37)))

	A := dh.ToPublic(a)
	B := dh.ToPublic(b)

	fmt.Printf("Public keys A:%s, B:%s\n", A.String(), B.String())

	S_a := dh.ToShared(&B, a)
	S_b := dh.ToShared(&A, b)

	diff := big.NewInt(0)
	diff.Sub(&S_a, &S_b)

	fmt.Printf("Shared keys A:%s, B:%s\n", S_a.String(), S_b.String())
	fmt.Printf("Diff = %s\n", diff.String())
*/
 	return
}

func ctrEdit(ciphertext []byte, key []byte, offset int, newtext []byte) []byte {
	bc := JMT.MakeAES(key)
	// each block is 16 bytes, so we need to increment the counter 'clicks' time
	clicks := offset / 16
	step := offset % 16

	out := make([]byte, len(ciphertext))
	copy(out, ciphertext)

	// this implentation of ctr assumes the nonce is zero
	nonce := make([]byte, 8)

	counter := []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}

	// increment the counter 'offset' times
	for i := 0; i < clicks; i++ {
		counter = dupIncrementCTR(counter)
	}

	bytePad := JMT.ECBEncrypt(bc, append(nonce, counter...))
	for i := 0; i < len(newtext); i++ {
		by := ciphertext[offset + i] // grab the offset+ith byte
		by ^= bytePad[step] // xor it with the step byte of the CTRPad
		// at this point by should be the plain text value
		mask := by ^ newtext[i]
		out[offset + i] ^= mask

		// go to next byte
		step++
		if step == 16 { // we have gone over the end of the pad, so make a new one
			counter = dupIncrementCTR(counter)
			bytePad = JMT.ECBEncrypt(bc, append(nonce, counter...))

			step = 0
		}
	}

	return out
}

func dupIncrementCTR(counter []byte) []byte {
	pos := 0

	counter[pos] += 1
	for counter[pos] == 0 {
		pos++
		counter[pos] += 1

		if pos > len(counter) {
			pos = len(counter) - 1
		}
	}

	return counter
}