// ysoftman
// 대상 문자의 인코딩, 해시값 생성기
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

func showUsage(file string) {
	fmt.Println("[usage] " + file + " string")
}
func main() {
	_, file := filepath.Split(os.Args[0])
	if len(os.Args) <= 1 {
		showUsage(file)
		return
	}
	fmt.Println("string =", os.Args[1])
	GetMD5(os.Args[1])
	GetSha1(os.Args[1])
	GetSha256(os.Args[1])
	GetEncURL(os.Args[1])
	GetDecURL(os.Args[1])
	GetEncBase32(os.Args[1])
	GetDecBase32(os.Args[1])
	GetEncBase64(os.Args[1])
	GetDecBase64(os.Args[1])
	GetEncHex(os.Args[1])
	GetDecHex(os.Args[1])
}

// GetMD5 : md5 생성하기
func GetMD5(str string) string {
	md5 := md5.New()
	data := []byte(str)
	// fmt.Printf("%s\n", data)

	// 해시값을 만들 데이터 설정
	md5.Write(data)
	// fmt.Printf("md5 sum = %x\n", md5.Sum(nil))
	// md5str := (string)(md5.Sum(nil))
	md5str := hex.EncodeToString(md5.Sum(nil))
	fmt.Printf("md5 = %s\n", md5str)

	return md5str
}

// GetSha1 : sha1 생성하기
func GetSha1(str string) string {
	// sha1 해시값 만들기
	sha1 := sha1.New()
	data := []byte(str)
	// fmt.Printf("%s\n", data)
	// 해시값을 만들 데이터 설정
	sha1.Write(data)
	// 해시값을 출력
	// 추가 data 를 넣어주면 기존 데이터에 sum 하는 방식으로 해시값을 리턴
	// data 1개인 경우 sum(nil) 사용
	sha1str := hex.EncodeToString(sha1.Sum(nil))
	fmt.Printf("sha1 = %s\n", sha1str)
	return sha1str
}

// GetSha256 : sha256 생성하기
func GetSha256(str string) string {
	// sha256 해시값 만들기
	sha256 := sha256.New()
	data := []byte(str)
	// fmt.Printf("%s\n", data)
	// 해시값을 만들 데이터 설정
	sha256.Write(data)
	// 해시값을 출력
	// 추가 data 를 넣어주면 기존 데이터에 sum 하는 방식으로 해시값을 리턴
	// data 1개인 경우 sum(nil) 사용
	sha256str := hex.EncodeToString(sha256.Sum(nil))
	fmt.Printf("sha256 = %s\n", sha256str)
	return sha256str
}

// GetEncURL : URL 인코딩
func GetEncURL(str string) string {
	encurl := url.QueryEscape(str)
	fmt.Printf("encode url = %s\n", encurl)
	return encurl
}

// GetDecURL : URL 디코딩
func GetDecURL(str string) string {
	decurl, err := url.QueryUnescape(str)
	if err != nil {
		fmt.Printf("decode url = %v\n", err.Error())
		return ""
	}
	fmt.Printf("decode url = %s\n", decurl)
	return decurl
}

// GetEncBase32 : base32 인코딩
func GetEncBase32(str string) string {
	data := []byte(str)
	encbase32 := base32.StdEncoding.EncodeToString(data)
	fmt.Printf("encode base32 = %s\n", encbase32)
	return encbase32
}

// GetDecBase32 : base32 디코딩
func GetDecBase32(str string) string {
	decbase32, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("decode base32 = %v\n", err.Error())
		return ""
	}
	fmt.Printf("decode base32 = %s\n", decbase32)
	return string(decbase32)
}

// GetEncBase64 :  base64 인코딩
func GetEncBase64(str string) string {
	data := []byte(str)
	encbase64 := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("encode base64 = %s\n", encbase64)
	return encbase64
}

// GetDecBase64 :  base64 디코딩
func GetDecBase64(str string) string {
	decbase64, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("decode base64 = %v\n", err.Error())
		return ""
	}
	fmt.Printf("decode base64 = %s\n", decbase64)
	return string(decbase64)
}

// GetEncHex : 스트링을 16진수로 인코딩
func GetEncHex(str string) string {
	data := []byte(str)
	hexdecimal := hex.EncodeToString(data)
	fmt.Printf("encode hexdecimal = %s\n", hexdecimal)
	return hexdecimal
}

// GetDecHex : 16진수를 스트링으로 디코딩
func GetDecHex(str string) string {
	data, err := hex.DecodeString(str)
	if err != nil {
		fmt.Printf("decode hexdecimal = %v\n", err.Error())
		return ""
	}
	decStr := string(data)
	fmt.Printf("decode hexdecimal = %s\n", decStr)
	return decStr
}
