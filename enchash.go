// ysoftman
// 대상 문자의 인코딩, 해시값 생성기
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	fmt.Println("~ enchash ~")
	if len(os.Args) <= 1 {
		fmt.Printf("ex)%s string\n", os.Args[0])
		return
	}
	log.Printf("target = %s", os.Args[1])
	GetMD5(os.Args[1])
	GetSha1(os.Args[1])
	GetURLEnc(os.Args[1])
	GetBase64(os.Args[1])
}

// GetMD5 : md5 생성하기
func GetMD5(str string) string {
	md5 := md5.New()
	data := []byte(str)
	// log.Printf("%s\n", data)

	// 해시값을 만들 데이터 설정
	md5.Write(data)
	// log.Printf("md5 sum = %x\n", md5.Sum(nil))
	// md5str := (string)(md5.Sum(nil))
	md5str := hex.EncodeToString(md5.Sum(nil))
	log.Printf("md5 = %s\n", md5str)

	return md5str
}

// GetSha1 : sha1 생성하기
func GetSha1(str string) string {
	// sha1 해시값 만들기
	sha1 := sha1.New()
	data := []byte(str)
	// log.Printf("%s\n", data)
	// 해시값을 만들 데이터 설정
	sha1.Write(data)
	// 해시값을 출력
	// 추가 data 를 넣어주면 기존 데이터에 sum 하는 방식으로 해시값을 리턴
	// data 1개인 경우 sum(nil) 사용
	sha1str := hex.EncodeToString(sha1.Sum(nil))
	log.Printf("sha1 = %s\n", sha1str)
	return sha1str
}

// GetURLEnc : URL 인코딩 생성하기
func GetURLEnc(str string) string {
	urlenc := url.QueryEscape(str)
	log.Printf("urlenc = %s\n", urlenc)
	return urlenc
}

// GetBase64 :  base64 생성하기
func GetBase64(str string) string {
	data := []byte(str)
	base64 := base64.StdEncoding.EncodeToString(data)
	log.Printf("base64 = %s\n", base64)
	return base64
}
