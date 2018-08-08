package mantis

import (
	"bytes"
	"io/ioutil"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"os/exec"
	"os"
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

//全角转换半角s
func DBCtoSBC(s string) string {
	res := ""
	for _, i := range s {
		insideCode := i
		if insideCode == 12288 {
			insideCode = 32
		} else {
			insideCode -= 65248
		}
		if insideCode < 32 || insideCode > 126 {
			res += string(i)
		} else {
			res += string(insideCode)
		}
	}
	return res
}

//convert GBK to UTF-8
func DecodeGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//convert BIG5 to UTF-8
func DecodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//convert UTF-8 to BIG5
func EncodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
// 简体中文转台湾繁体中文，需要安装opencc
func Trans2TW(s string) string {
	filepath := "/tmp/" + GetMD5Hash(s)
	writeToFile(filepath, s)
	cmd := exec.Command("opencc", "-c", "s2tw.json", "-i", filepath, "-o", filepath)
	cmd.Run()
	b, _ := ioutil.ReadFile(filepath)
	os.Remove(filepath)
	return string(b)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func writeToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()
	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return nil
	}
	return nil
}
