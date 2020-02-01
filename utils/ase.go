package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"strconv"
)
var Key = []byte("12345678abcdefgh")

//填充字符串（末尾）
func PaddingText1(str []byte, blockSize int) []byte {
	//需要填充的数据长度
	paddingCount := blockSize - len(str)%blockSize
	//填充数据为：paddingCount ,填充的值为：paddingCount
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)
	//fmt.Println(newPaddingStr)
	return newPaddingStr
}

//去掉字符（末尾）
func UnPaddingText1(str []byte) []byte {
	n := len(str)
	count := int(str[n-1])
	newPaddingText := str[:n-count]
	return newPaddingText
}
//---------------DES加密  解密--------------------
func EncyptogAES(src, key []byte) []byte {
	block,err:=aes.NewCipher(key)
	if err!= nil{
		fmt.Println(nil)
		return nil
	}
	src=PaddingText1(src,block.BlockSize())
	blockMode:=cipher.NewCBCEncrypter(block,key)
	blockMode.CryptBlocks(src,src)
	return src

}
func DecrptogAES(src,key[]byte) []byte {
	block,err:=aes.NewCipher(key)
	if err!= nil{
		fmt.Println(nil)
		return nil
	}
	blockMode:=cipher.NewCBCDecrypter(block,key)
	blockMode.CryptBlocks(src,src)
	src=UnPaddingText1(src)
	return  src
}

func demo() {
	id := 3
	str := strconv.Itoa(id)
	fmt.Println("编码的数据为：", str)
	src := EncyptogAES([]byte(str), Key)
	DecrptogAES(src, Key)
	fmt.Println("解码之后的数据为：", string(src))
}

