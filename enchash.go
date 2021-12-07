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

	"github.com/fatih/color"
)

var currentColorNum int = 0

func GetNextColorString(i int, str string) string {
	n := i % 6
	switch n {
	case 0:
		yellow := color.New(color.FgYellow).SprintFunc()
		return yellow(str)
	case 1:
		green := color.New(color.FgGreen).SprintFunc()
		return green(str)
	case 2:
		red := color.New(color.FgRed).SprintFunc()
		return red(str)
	case 3:
		blue := color.New(color.FgBlue).SprintFunc()
		return blue(str)
	case 4:
		magenta := color.New(color.FgMagenta).SprintFunc()
		return magenta(str)
	case 5:
		cyan := color.New(color.FgCyan).SprintFunc()
		return cyan(str)
	default:
		white := color.New(color.FgWhite).SprintFunc()
		return white(str)
	}
}

func showUsage(file string) {
	fmt.Println("[usage] " + file + " string")
}
func main() {
	_, file := filepath.Split(os.Args[0])
	if len(os.Args) <= 1 {
		showUsage(file)
		return
	}
	fmt.Printf("%30v = %v\n", "string", os.Args[1])
	GetMD5(os.Args[1])
	GetSha1(os.Args[1])
	GetSha256(os.Args[1])
	GetEncURL(os.Args[1])
	GetDecURL(os.Args[1])
	GetEncBase32(os.Args[1])
	GetDecBase32(os.Args[1])
	GetEncBase64Std(os.Args[1])
	GetDecBase64Std(os.Args[1])
	GetEncBase64URL(os.Args[1])
	GetDecBase64URL(os.Args[1])
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
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "md5", md5str)))
	return md5str
}

// GetSha1 : sha1 생성하기
func GetSha1(str string) string {
	data := []byte(str)

	// sha1 해시값 만들기
	// 방법1
	sha1new := sha1.New()
	// fmt.Printf("%s\n", data)
	// 해시값을 만들 데이터 설정
	sha1new.Write(data)
	// 해시값을 출력
	// 추가 data 를 넣어주면 기존 데이터에 sum 하는 방식으로 해시값을 리턴
	// data 1개인 경우 sum(nil) 사용
	sha1str := hex.EncodeToString(sha1new.Sum(nil))
	currentColorNum++
	fmt.Printf("%30v = %s\n", "sha1", sha1str)
	sha1str = base64.URLEncoding.EncodeToString(sha1new.Sum(nil))
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "sha1(base64.URLEncoding)", sha1str)))

	// 방법2
	// new 생성없이 바로 사용하기
	// EncodeToString대신 %x 로 출력해서 스트링으로 저장해도됨
	// sha1str2 := fmt.Sprintf("%x", sha1.Sum(data))
	// currentColorNum++
	// fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "sha1", sha1str2)))

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
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "sha256", sha256str)))
	sha256str = base64.URLEncoding.EncodeToString(sha256.Sum(nil))
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "sha256(base64.URLEncoding)", sha256str)))

	return sha256str
}

// GetEncURL : URL 인코딩
func GetEncURL(str string) (string, string) {
	encurl1 := url.QueryEscape(str)
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "encode url(+)", encurl1)))

	t := &url.URL{Path: str}
	encurl2 := t.String()
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "encode url(%20)", encurl2)))
	return encurl1, encurl2
}

// GetDecURL : URL 디코딩
func GetDecURL(str string) string {
	decurl, err := url.QueryUnescape(str)
	if err != nil {
		fmt.Printf("%30v = %v\n", "decode url", err.Error())
		return ""
	}
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "decode url", decurl)))
	return decurl
}

// GetEncBase32 : base32 인코딩
func GetEncBase32(str string) string {
	data := []byte(str)
	encbase32 := base32.StdEncoding.EncodeToString(data)
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "encode base32", encbase32)))
	return encbase32
}

// GetDecBase32 : base32 디코딩
func GetDecBase32(str string) string {
	decbase32, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("%30v = %v\n", "decode base32", err.Error())
		return ""
	}
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "decode base32", decbase32)))
	return string(decbase32)
}

// GetEncBase64Std : base64 인코딩
func GetEncBase64Std(str string) string {
	data := []byte(str)
	encbase64std := base64.StdEncoding.EncodeToString(data)
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "encode base64(std)", encbase64std)))
	return encbase64std
}

// GetDecBase64Std : base64 디코딩
func GetDecBase64Std(str string) string {
	// 패딩이 없는 경우 4의 배수까지 = 로 패딩
	remainder := len(str) % 4
	for i := 0; i < remainder; i++ {
		str += "="
	}
	decbase64std, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("%30v = %v\n", "decode base64(std)", err.Error())
		return ""
	}
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "decode base64(std)", decbase64std)))
	return string(decbase64std)
}

// GetEncBase64URL : base64 URL 인코딩
// base64 인코딩은 RFC 4648 에 정의 되어 있다.
// StdEncoding 기본 base64 인코딩 방법이고
// URLEncoding 기본 base64 의 + 과 / 문자가 URL과file name에서 문제가 되어
// +, / 를 각각 -,_ 로 변경해 URL과file name에 사용할 수 있다.
func GetEncBase64URL(str string) string {
	data := []byte(str)
	encbase64url := base64.URLEncoding.EncodeToString(data)
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "encode base64(url)", encbase64url)))
	return encbase64url
}

// GetDecBase64URL : base64 URL 디코딩
func GetDecBase64URL(str string) string {
	// 패딩이 없는 경우 4의 배수까지 = 로 패딩
	remainder := len(str) % 4
	for i := 0; i < remainder; i++ {
		str += "="
	}
	decbase64url, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("%30v = %v\n", "decode base64(url)", err.Error())
		return ""
	}
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "decode base64(url)", decbase64url)))
	return string(decbase64url)
}

// GetEncHex : 스트링을 16진수로 인코딩
func GetEncHex(str string) string {
	data := []byte(str)
	hexdecimal := hex.EncodeToString(data)
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "encode hexdecimal", hexdecimal)))
	return hexdecimal
}

// GetDecHex : 16진수를 스트링으로 디코딩
func GetDecHex(str string) string {
	data, err := hex.DecodeString(str)
	if err != nil {
		fmt.Printf("%30v = %v\n", "decode hexdecimal", err.Error())
		return ""
	}
	decStr := string(data)
	currentColorNum++
	fmt.Print(GetNextColorString(currentColorNum, fmt.Sprintf("%30v = %s\n", "decode hexdecimal", decStr)))
	return decStr
}
