package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

//将结构体序列化
func (m *Monster) Store(filePath string) error {
	//将m序列化
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("序列化错误 err=", err)
		return err
	}
	//输出序列化的结果
	fmt.Printf("序列化后的结果=%v\n", data)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()
	return nil
}

func (m *Monster) ReStore(filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
		return err
	}

	var monster Monster
	err = json.Unmarshal([]byte(content), &monster)
	if err != nil {
		fmt.Printf("unmarshal err =%v\n", err)
		return err
	}
	fmt.Printf("反序列化后 monster=%v\n monster.Name=%v\n", monster, monster.Name)
	return nil
}

func main() {
	m := Monster{
		"猪八戒",
		999,
		"睡觉",
	}
	filePath := ""
	fmt.Println("先保存到哪个文件：")
	fmt.Scanln(&filePath)
	m.Store(filePath)

	m.ReStore(filePath)
}
