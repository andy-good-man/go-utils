package gostructs

import (
	"fmt"
	"reflect"
)

// ToMgoBson 将{结构体对象}转换为{字典}键值对，然后通过 insert/update 操作入库至：mongodb
func ToMgoBson(obj interface{}, filterFields []string) map[string]interface{} {
	var err error
	data := make(map[string]interface{}, 0)

	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()

	for i := 0; i < t.NumField(); i++ {
		bsonVal, ok := t.Field(i).Tag.Lookup("bson")
		if ok && bsonVal != "-" {
			if bsonVal == "_id" {
				if !v.Field(i).IsZero() {
					data["_id"] = v.Field(i).Interface()
				} else {
					err = fmt.Errorf(fmt.Sprintf("ToMgomap.将对象转化mgo字典.主键id值不能为空"))
					panic(err)
				}
			} else {
				data[bsonVal] = v.Field(i).Interface()
			}
		}
	}

	if filterFields == nil || len(filterFields) == 0 {
		return data
	}
	filterData := make(map[string]interface{}, 0)

	for _, field := range filterFields {
		if val, ok := data[field]; ok {
			filterData[field] = val
		}
	}

	return filterData
}
