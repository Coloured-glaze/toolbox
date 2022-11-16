package check

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"log"
	"os"

	s "github.com/Coloured-glaze/toolbox/str"
)

func MD5(str string) string {
	m := md5.Sum(s.StrtoByte(str))
	return hex.EncodeToString(m[:])
}

func MD5File(fileName string) string {
	f, e := os.Open(fileName)
	if e != nil {
		log.Fatal(e)
	}
	h := md5.New()
	_, e = io.Copy(h, f)
	if e != nil {
		log.Fatal(e)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func SHA1(str string) string {
	r := sha1.Sum(s.StrtoByte(str))
	return hex.EncodeToString(r[:])
}

func SHA1File(fileName string) string {
	f, e := os.Open(fileName)
	if e != nil {
		log.Fatal(e)
	}
	h := sha1.New()
	_, e = io.Copy(h, f)
	if e != nil {
		log.Fatal(e)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func SHA256(str string) string {
	r := sha256.New().Sum(s.StrtoByte(str))
	return hex.EncodeToString(r[:])
}

func SHA256File(fileName string) string {
	N, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	sha := sha256.New()
	if _, err := io.Copy(sha, N); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(sha.Sum(nil))
}

func SHA512(str string) string {
	r := sha512.New().Sum(s.StrtoByte(str))
	return hex.EncodeToString(r[:])
}

func SHA512File(fileName string) string {
	N, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	sha := sha512.New()
	if _, err := io.Copy(sha, N); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(sha.Sum(nil))
}
