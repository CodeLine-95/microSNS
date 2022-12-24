/**
 * @Author: handsomejob
 * @WechatMp: 程序员小乔
 * @Version: 1.0.0
 * @IDE: GoLand
 * @Date: 2022/12/24 23:00
 */

package Crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

const IV = "mhytanglki872345"

// 使用aes进行对称加密
func AesEncrypt(ss string, key string) string {
	src := []byte(ss)
	//1. 创建并返回一个使用DES算法的cipher.Block接口
	block, _ := aes.NewCipher([]byte(key))
	//2. 对最后一个明文分组进行数据填充
	src = paddingBytes(src, block.BlockSize())
	//3. 创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	cbcDecrypter := cipher.NewCBCEncrypter(block, []byte(IV))
	//4. 加密连续的数据并返回
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)

	return hex.EncodeToString(dst)
}

// 使用aes进行解密
func AesDecrypt(ss string, key string) string {
	src, _ := hex.DecodeString(ss)
	//1. 创建并返回一个使用DES算法的cipher.Block接口
	block, _ := aes.NewCipher([]byte(key))
	//2. 创建一个密码分组为链接模式的，底层使用DES解密的BlockMode接口
	cbcDecrypter := cipher.NewCBCDecrypter(block, []byte(IV))
	//3. 数据块解密
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)
	//4. 去掉最后一组填充数据
	newBytes := unPaddingBytes(dst)
	return string(newBytes)
}

// 填充明文最后一个分组工具方法
// src - 原始数据
// blockSize - 每个分组的数据长度
func paddingBytes(src []byte, blockSize int) []byte {
	//1.求出最后一个分组要填充多个字节
	padding := blockSize - len(src)%blockSize
	//2.创建新的切片，切片的字节数为填充的字节数，并初始化化，每个字节的值为填充的字节数
	padBytes := bytes.Repeat([]byte{byte(padding)}, padding)
	//3.将创建出的新切片和原始数据进行连接
	newBytes := append(src, padBytes...)

	//4.返回新的字符串
	return newBytes
}

// 删除密文末尾分组填充的工具方法
func unPaddingBytes(src []byte) []byte {
	//1.求出要处理的切片的长度
	l := len(src)
	//2.取出最后一个字符，得到其整型值
	n := int(src[l-1])

	//3.将切片末尾的number个字节删除
	return src[:l-n]
}
