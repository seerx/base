package base

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

// DecodeJSONFile 解析 JOSN 文件
func DecodeJSONFile(file string, v interface{}) error {
	jsonFile, err := os.Open(file)
	if err != nil {
		return err
	}
	return json.NewDecoder(jsonFile).Decode(v)
	// if err != nil {
	// 	return err
	// }
	// return nil
}

//EncodeJSONFile 写入 json 文件
func EncodeJSONFile(file string, v interface{}) error {
	infoData, err := json.Marshal(v)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	err = json.Indent(&out, infoData, "", "\t")
	if err != nil {
		return err
	}

	// out.WriteTo(file)
	return ioutil.WriteFile(file, out.Bytes(), 0666)
}
