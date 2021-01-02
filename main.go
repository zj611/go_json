package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type ServerInfo struct {
	//标签有string，ServerName的值会进行二次JSON编码 “”
	ServerName  string `json:"serverName"`
	//ServerName2 string `json:"server_name_2,string"`
	//如果ServerIP为空，则不输出到JSON中
	ServerIP    string `json:"serverIP,omitempty"`
}

type Server struct {
	//ID不会导出到JSON
	ID int `json:"-"`
	ServerInfos []ServerInfo `json:"server_infos"`
	Description string `json:"description,string"`
	//tag中带有“,string”选项，在输出到JSON的时候会把该字段对应的值转换成JSON字符串
	Score  float32 `json:"score,string"`
	Age    int `json:"age"`
}

func main(){
	s01 := ServerInfo{
		ServerName: "s1",
		ServerIP: "10.0.12.1",
	}
	s02 := ServerInfo{
		ServerName: "s2",
		ServerIP: "",
	}

	s := Server{
		ID:          3,
		Description: `描述信息`,
		Score: 		87.2,
		Age: 25,
	}

	s.ServerInfos = append(s.ServerInfos, s01,s02)
	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
	fmt.Println()

	//-------------------
	data := `{"server_infos":[{"serverName":"s1","serverIP":"10.0.12.1"},{"serverName":"s2"}],"description":"\"描述信息\"","score":"87.2","age":25}`
	var tmp map[string]interface{}
	json.Unmarshal([]byte(data), &tmp)

	description := GetItemString(tmp,"description")
	fmt.Println(description)
	fmt.Println(reflect.TypeOf(description))

	score := GetItemString(tmp,"score")
	fmt.Println(score)
	fmt.Println(reflect.TypeOf(score))

	age := GetItemFloat64(tmp,"age")
	fmt.Println(age)
	fmt.Println(reflect.TypeOf(age))

	//繁琐解析法
	//fmt.Println(len(tmp["server_infos"].([]interface {})))
	//mapTmp1 := (tmp["server_infos"].([]interface {}))[0].
	//	(map[string]interface {})
	//fmt.Println(mapTmp1["serverName"])
	//fmt.Println(mapTmp1["serverIP"])

	//函数封装解析
	server_infos := GetItemArray(tmp,"server_infos")
	fmt.Println(server_infos)
	fmt.Println(reflect.TypeOf(server_infos[0]))

	serverName := GetItemString(server_infos[0].(map[string]interface{}),"serverName")
	fmt.Println(serverName)
	serverIP := GetItemString(server_infos[0].(map[string]interface{}),"serverIP")
	fmt.Println(serverIP)


	data1 := `{"data":{"t1":"sdd","t2":"sdsd","t3":{"y1":"sd","y2":"12"}}}`
	json.Unmarshal([]byte(data1), &tmp)
	test := GetItemMap(tmp,"data")
	fmt.Println(test)
	//fmt.Println(reflect.TypeOf(test))
	t3 := GetItemMap(test,"t3")
	fmt.Println(GetItemString(t3,"y2"))
	//fmt.Println(reflect.TypeOf(t3))

}

func GetItemArray(parent map[string]interface{}, key string) []interface{} {
	val, ok := parent[key].([]interface{})
	if ok {
		return val
	} else {
		return nil
	}
}
func GetItemMap(parent map[string]interface{}, key string) map[string]interface{} {
	val, ok := parent[key].(map[string]interface{})
	if ok {
		return val
	} else {
		return nil
	}
}
func GetItemString(parent map[string]interface{}, key string) string {
	val, ok := parent[key].(string)
	if ok {
		return val
	}
	return ""
}
func GetItemFloat64(parent map[string]interface{}, key string) float64 {
	val, ok := parent[key].(float64)
	if ok {
		return val
	}
	return 0
}