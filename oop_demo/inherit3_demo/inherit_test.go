package inherit3_demo

import (
	"encoding/json"
	"reflect"
	"regexp"
)

type Info[T any] struct {
	T   T                      //固定字段，对于一般情况下会本身会读取到的字段写入这里
	Ext map[string]interface{} // 其它扩展数据，根据需要添加，key/value形式获取
}

func NewInfo[T any]() Info[T] { // 主要是为了初始化map
	res := Info[T]{}
	res.Ext = map[string]interface{}{}
	return res
}
func (s Info[T]) MarshalJSON() ([]byte, error) { // 将动态字段和静态字段组合放到一个结构体
	marshal, err := json.Marshal(s.T)
	if err != nil {
		return nil, err
	}
	if len(s.Ext) != 0 {
		extMarshal, err := json.Marshal(s.Ext)
		if err != nil {
			return nil, err
		}
		marshal = append(marshal[0:len(marshal)-1], []byte(`,"__ext__":{},`)...)
		marshal = append(marshal, extMarshal[1:]...)
	}
	return marshal, nil
}
func (s *Info[T]) UnmarshalJSON(data []byte) error { // 反序列化
	dataS := string(data)
	subInfo := regexp.MustCompile(`,"__ext__":{},`).Split(dataS, 2)
	if len(subInfo) == 1 { //未找到对象
		return json.Unmarshal(data, &s.T)
	} else {
		baseStr := subInfo[0] + "}"
		extStr := "{" + subInfo[1]
		err := json.Unmarshal([]byte(baseStr), &s.T)
		if err != nil {
			return err
		}
		ext := map[string]interface{}{}
		err = json.Unmarshal([]byte(extStr), &ext)
		if err != nil {
			return err
		}
		s.Ext = ext
		return nil
	}
}

func (s Info[T]) GetAttr(fieldName string) (v interface{}, find bool) { // 根据字符串获取值，这里甚至可以改造为忽略大小写
	refval := reflect.ValueOf(s)
	val := refval.FieldByName(fieldName)
	if val.IsValid() {
		return val.Interface(), true
	}
	dynamicStruct, ok := s.Ext[fieldName]
	return dynamicStruct, ok
}

func (s *Info[T]) SetAttr(fieldName string, value interface{}) bool { //根据字符串写入值，这里也可以忽略大小写
	refval := reflect.ValueOf(s).Elem()
	val := refval.FieldByName(fieldName)
	if val.IsValid() {
		val.Set(reflect.ValueOf(value))
		return true
	}
	s.Ext[fieldName] = value
	return false
}
