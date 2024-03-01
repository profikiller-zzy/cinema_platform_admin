package core

import (
	"AfterEnd/config"
	"AfterEnd/global"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const ConfigFile = "newConfig.yaml"

func InitConfig() *config.Config {
	// 使用ioutil导入配置文件，使用yaml.Unmarshal将配置文件反序列化读取到结构体中
	config := &config.Config{}
	key := []byte("MFwwDQYJKoZIhvcN")

	// 读取配置文件
	encryptedData, err := readConfig(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("read yamlConf file error: %v", err))
	}

	// 解密配置文件
	decryptedData, err := decrypt(encryptedData, key)
	if err != nil {
		panic(fmt.Errorf("error decrypting config file: %v", err))
	}

	err = yaml.Unmarshal(decryptedData, config)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err) // log.Fatalf()用于记录一条严重的错误消息，并且终止程序运行
	}
	//fmt.Println("config yamlFile Init success.") // 输出语句，仅作调试使用
	return config
}

func SetYaml() error {

	// 将结构体编码为yaml格式
	siteInfoYaml, err := yaml.Marshal(&global.Config)
	if err != nil {
		return err
	}

	// 将yaml内容写入文件
	err = ioutil.WriteFile(ConfigFile, siteInfoYaml, os.FileMode(0644))
	return err
}

// decrypt 使用AES加密算法对输入的数据进行解密，data为输入的数据，key是密钥
func decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)

	return data, nil
}

func readConfig(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(fmt.Errorf("get yamlConf file error: %v", err))
	}
	return data, nil
}

func writeConfig(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
