package flag

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

// EncryptProfile 将配置文件config.yaml加密后，将加密内容写入newConfig.yaml中
func EncryptProfile() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Errorf("read yamlConf file error: %v", err))
	}

	encryptedData, err := encrypt(data, []byte("MFwwDQYJKoZIhvcN"))
	if err != nil {
		panic(fmt.Errorf("encrypt yamlConf file error: %v", err))
	}

	err = ioutil.WriteFile("newConfig.yaml", encryptedData, 0644)
	if err != nil {
		panic(fmt.Errorf("write yamlConf file error: %v", err))
	}
}

// encrypt 使用AES加密算法对输入的数据进行加密，data为输入的数据，key是用于加密的密钥
func encrypt(data []byte, key []byte) ([]byte, error) {
	// 生成一个aes加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}
