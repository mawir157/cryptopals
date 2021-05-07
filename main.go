package main

import (
	"fmt"
	// "math/big"
	"math/rand"
	"time"
)

import JMT "github.com/mawir157/jmtcrypto"
import JMTR "github.com/mawir157/jmtcrypto/rand"
// import DH "github.com/mawir157/jmtcrypto/dh"

func blockTitle(s string) {
	fmt.Printf("###############################################################################\n")
	fmt.Printf("##\n")
	fmt.Printf("## %s\n", s)
	fmt.Printf("##\n")

	return
}

func main() {
	// testString := "CRIwqt4+szDbqkNY+I0qbDe3LQz0wiw0SuxBQtAM5TDdMbjCMD/venUDW9BLPEXODbk6a48oMbAY6DDZsuLbc0uR9cp9hQ0QQGATyyCESq2NSsvhx5zKlLtzdsnfK5ED5srKjK7Fz4Q38/ttd+stL/9WnDzlJvAo7WBsjI5YJc2gmAYayNfmCW2lhZE/ZLG0CBD2aPw0W417QYb4cAIOW92jYRiJ4PTsBBHDe8o4JwqaUac6rqdi833kbyAOV/Y2RMbN0oDb9Rq8uRHvbrqQJaJieaswEtMkgUt3P5Ttgeh7J+hE6TR0uHot8WzHyAKNbUWHoi/5zcRCUipvVOYLoBZXlNu4qnwoCZRSBgvCwTdz3Cbsp/P2wXB8tiz6l9rL2bLhBt13Qxyhhu0H0+JKj6soSeX5ZD1Rpilp9ncR1tHW8+uurQKyXN4xKeGjaKLOejr2xDIw+aWF7GszU4qJhXBnXTIUUNUfRlwEpS6FZcsMzemQF30ezSJHfpW7DVHzwiLyeiTJRKoVUwo43PXupnJXDmUysCa2nQz/iEwyor6kPekLv1csm1Pa2LZmbA9Ujzz8zb/gFXtQqBAN4zA8/wt0VfoOsEZwcsaLOWUPtF/Ry3VhlKwXE7gGH/bbShAIKQqMqqUkEucZ3HPHAVp7ZCn3Ox6+c5QJ3Uv8V7L7SprofPFN6F+kfDM4zAc59do5twgDoClCbxxG0L19TBGHiYP3CygeY1HLMrX6KqypJfFJW5O9wNIF0qfOC2lWFgwayOwq41xdFSCW0/EBSc7cJw3N06WThrW5LimAOt5L9c7Ik4YIxu0K9JZwAxfcU4ShYu6euYmWLP98+qvRnIrXkePugS9TSOJOHzKUoOcb1/KYd9NZFHEcp58Df6rXFiz9DSq80rR5Kfs+M+Vuq5Z6zY98/SP0A6URIr9NFu+Cs9/gf+q4TRwsOzRMjMQzJL8f7TXPEHH2+qEcpDKz/5pE0cvrgHr63XKu4XbzLCOBz0DoFAw3vkuxGwJq4Cpxkt+eCtxSKUzNtXMn/mbPqPl4NZNJ8yzMqTFSODS4bYTBaN/uQYcOAF3NBYFd5x9TzIAoW6ai13a8h/s9i5FlVRJDe2cetQhArrIVBquF0L0mUXMWNPFKkaQEBsxpMCYh7pp7YlyCNode12k5jY1/lc8jQLQJ+EJHdCdM5t3emRzkPgND4a7ONhoIkUUS2R1oEV1toDj9iDzGVFwOvWyt4GzA9XdxT333JU/n8m+N6hs23MBcZ086kp9rJGVxZ5f80jRz3ZcjU6zWjR9ucRyjbsuVn1t4EJEm6A7KaHm13m0vwN/O4KYTiiY3aO3siayjNrrNBpn1OeLv9UUneLSCdxcUqjRvOrdA5NYv25Hb4wkFCIhC/Y2ze/kNyis6FrXtStcjKC1w9Kg8O25VXB1Fmpu+4nzpbNdJ9LXahF7wjOPXN6dixVKpzwTYjEFDSMaMhaTOTCaqJig97624wv79URbCgsyzwaC7YXRtbTstbFuEFBee3uW7B3xXw72mymM2BS2uPQ5NIwmacbhta8aCRQEGqIZ078YrrOlZIjar3lbTCo5o6nbbDq9bvilirWG/SgWINuc3pWl5CscRcgQQNp7oLBgrSkQkv9AjZYcvisnr89TxjoxBO0Y93jgp4T14LnVwWQVx3l3d6S1wlscidVeaM24E/JtS8k9XAvgSoKCjyiqsawBMzScXCIRCk6nqX8ZaJU3rZ0LeOMTUw6MC4dC+aY9SrCvNQub19mBdtJUwOBOqGdfd5IoqQkaL6DfOkmpnsCs5PuLbGZBVhah5L87IY7r6TB1V7KboXH8PZIYc1zlemMZGU0o7+etxZWHgpdeX6JbJIs3ilAzYqw/Hz65no7eUxcDg1aOaxemuPqnYRGhW6PvjZbwAtfQPlofhB0jTHt5bRlzF17rn9q/6wzlc1ssp2xmeFzXoxffpELABV6+yj3gfQ/bxIB9NWjdZK08RX9rjm9CcBlRQeTZrD67SYQWqRpT5t7zcVDnx1s7ZffLBWm/vXLfPzMaQYEJ4EfoduSutjshXvR+VQRPs2TWcF7OsaE4csedKUGFuo9DYfFIHFDNg+1PyrlWJ0J/X0PduAuCZ+uQSsM/ex/vfXp6Z39ngq4exUXoPtAIqafrDMd8SuAtyEZhyY9V9Lp2qNQDbl6JI39bDz+6pDmjJ2jlnpMCezRK89cG11IqiUWvIPxHjoiT1guH1uk4sQ2Pc1J4zjJNsZgoJDcPBbfss4kAqUJvQyFbzWshhtVeAv3dmgwUENIhNK/erjpgw2BIRayzYw001jAIF5c7rYg38o6x3YdAtU3d3QpuwG5xDfODxzfL3yEKQr48C/KqxI87uGwyg6H5gc2AcLU9JYt5QoDFoC7PFxcE3RVqc7/Um9Js9X9UyriEjftWt86/tEyG7F9tWGxGNEZo3MOydwX/7jtwoxQE5ybFjWndqLp8DV3naLQsh/Fz8JnTYHvOR72vuiw/x5D5PFuXV0aSVvmw5Wnb09q/BowS14WzoHH6ekaWbh78xlypn/L/M+nIIEX1Ol3TaVOqIxvXZ2sjm86xRz0EdoHFfupSekdBULCqptxpFpBshZFvauUH8Ez7wA7wjL65GVlZ0f74U7MJVu9SwsZdgsLmnsQvr5n2ojNNBEv+qKG2wpUYTmWRaRc5EClUNfhzh8iDdHIsl6edOewORRrNiBay1NCzlfz1cj6VlYYQUM9bDEyqrwO400XQNpoFOxo4fxUdd+AHmCBhHbyCR81/C6LQTG2JQBvjykG4pmoqnYPxDyeiCEG+JFHmP1IL+jggdjWhLWQatslrWxuESEl3PEsrAkMF7gt0dBLgnWsc1cmzntG1rlXVi/Hs2TAU3RxEmMSWDFubSivLWSqZj/XfGWwVpP6fsnsfxpY3d3h/fTxDu7U8GddaFRQhJ+0ZOdx6nRJUW3u6xnhH3mYVRk88EMtpEpKrSIWfXphgDUPZ0f4agRzehkn9vtzCmNjFnQb0/shnqTh4Mo/8oommbsBTUKPYS7/1oQCi12QABjJDt+LyUan+4iwvCi0k0IUIHvk21381vC0ixYDZxzY64+xx/RNID+iplgzq9PDZgjc8L7jMg+2+mrxPS56e71m5E2zufZ4d+nFjIg+dHD/ShNPzVpXizRVUERztLuak8Asah3/yvwOrH1mKEMMGC1/6qfvZUgFLJH5V0Ep0n2K/Fbs0VljENIN8cjkCKdG8aBnefEhITdV7CVjXcivQ6efkbOQCfkfcwWpaBFC8tD/zebXFE+JshW16D4EWXMnSm/9HcGwHvtlAj04rwrZ5tRvAgf1IR83kqqiTvqfENcj7ddCFwtNZrQK7EJhgB5Tr1tBFcb9InPRtS3KYteYHl3HWR9t8E2YGE8IGrS1sQibxaK/C0kKbqIrKpnpwtoOLsZPNbPw6K2jpko9NeZAx7PYFmamR4D50KtzgELQcaEsi5aCztMg7fp1mK6ijyMKIRKwNKIYHagRRVLNgQLg/WTKzGVbWwq6kQaQyArwQCUXo4uRtyzGMaKbTG4dns1OFB1g7NCiPb6s1lv0/lHFAF6HwoYV/FPSL/pirxyDSBb/FRRA3PIfmvGfMUGFVWlyS7+O73l5oIJHxuaJrR4EenzAu4Avpa5d+VuiYbM10aLaVegVPvFn4pCP4U/Nbbw4OTCFX2HKmWEiVBB0O3J9xwXWpxN1Vr5CDi75FqNhxYCjgSJzWOUD34Y1dAfcj57VINmQVEWyc8Tch8vg9MnHGCOfOjRqp0VGyAS15AVD2QS1V6fhRimJSVyT6QuGb8tKRsl2N+a2Xze36vgMhw7XK7zh//jC2H"

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
	// blockTitle("McEliese")
	// public, private := JMT.GenerateKeyPair(2, 11)

	// public.Write("mce.pub")
	// private.Write("mce.pri")

	// public2 := JMT.ReadPublic("mce.pub")
	// private2 := JMT.ReadPrivate("mce.pri")

	// cipherText := public2.Encrypt(textMessage)
	// // PrintHex(cipherText, true)
	
	// plaintext := private2.Decrypt(cipherText)
	// JMT.PrintAscii(plaintext, true)
	// fmt.Println("")


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