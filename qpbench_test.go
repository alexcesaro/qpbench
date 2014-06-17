package qpbench

import (
	"fmt"
	"github.com/alexcesaro/qp"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func ExampleQpDecode() {
	var err error
	inputs := []string{
		"=0 ",
		"=0\r",
		"=0 \n",
	}
	for _, in := range inputs {
		fmt.Printf("Decoding %q\n", in)
		_, err = ioutil.ReadAll(NewQuotedPrintableReader(strings.NewReader(in)))
		fmt.Println("currently:", err)
		_, err = ioutil.ReadAll(qp.NewDecoder(strings.NewReader(in)))
		fmt.Println("  mime/qp:", err)
		fmt.Print("\n")
	}

	// Output:
	// Decoding "=0 "
	// currently: unexpected EOF
	//   mime/qp: qp: invalid quoted-printable hex byte 0x20
	//
	// Decoding "=0\r"
	// currently: unexpected EOF
	//   mime/qp: qp: invalid quoted-printable hex byte 0x0d
	//
	// Decoding "=0 \n"
	// currently: multipart: invalid quoted-printable hex byte 0x0a
	//   mime/qp: qp: invalid quoted-printable hex byte 0x20
}

func BenchmarkOldQpDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		io.Copy(ioutil.Discard, NewQuotedPrintableReader(strings.NewReader(mail)))
	}
}

func BenchmarkNewQpDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		io.Copy(ioutil.Discard, qp.NewDecoder(strings.NewReader(mail)))
	}
}

func BenchmarkOldQEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeRFC2047Word(text)
	}
}

func BenchmarkNewQEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qp.StdWordEncoder.EncodeHeader(text)
	}
}

func BenchmarkOldQDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecodeRFC2047Word(encodedWord)
	}
}

func BenchmarkNewQDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qp.DecodeWord(encodedWord)
	}
}

var text = "¡Hola, señor!"

var encodedWord = "=?UTF-8?Q?=C2=A1Hola,_se=C3=B1or!?="

var mail = `Cher ami,

Je suis toute =C3=A9mue de vous dire que j'ai
bien compris l'autre jour que vous aviez
toujours une envie folle de me faire
danser. Je garde le souvenir de votre
baiser et je voudrais bien que ce soit
une preuve que je puisse =C3=AAtre aim=C3=A9e
par vous. Je suis pr=C3=AAte =C3=A0 montrer mon
affection toute d=C3=A9sint=C3=A9ress=C3=A9e et sans cal-
cul, et si vous voulez me voir ainsi
vous d=C3=A9voiler, sans artifice, mon =C3=A2me
toute nue, daignez me faire visite,
nous causerons et en amis franchement
je vous prouverai que je suis la femme
sinc=C3=A8re, capable de vous offrir l'affection
la plus profonde, comme la plus =C3=A9troite
amiti=C3=A9, en un mot : la meilleure =C3=A9pouse
dont vous puissiez r=C3=AAver. Puisque votre
=C3=A2me est libre, pensez que l'abandon ou je
vis est bien long, bien dur et souvent bien
insupportable. Mon chagrin est trop
gros. Accourrez bien vite et venez me le
faire oublier. =C3=80 vous je veux me sou-
mettre enti=C3=A8rement.

Votre poup=C3=A9e`
