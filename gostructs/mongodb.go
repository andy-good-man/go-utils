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

// ToMgoBsonSweet ToMgoBsonSweet
func ToMgoBsonSweet(obj interface{}, checks, dels []string) map[string]interface{} {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		bsonVal, ok := t.Field(i).Tag.Lookup("bson")
		if ok && bsonVal != "-" {
			if bsonVal == "_id" {
				if !v.Field(i).IsZero() {
					data["_id"] = v.Field(i).Interface()
				} else {
					errmsg := fmt.Errorf("在对象转换键值对时，发现_id为空，请核查")
					panic(errmsg)
				}
			} else {
				data[bsonVal] = v.Field(i).Interface()
			}
		}
	}

	checkData := CheckMap(data, checks)
	delData := DeleteMap(checkData, dels)

	return delData
}

// CheckMap CheckMap
func CheckMap(intput map[string]interface{}, checkeys []string) map[string]interface{} {
	if checkeys == nil {
		return intput
	}
	output := map[string]interface{}{}
	for _, key := range checkeys {
		if val, ok := intput[key]; ok {
			output[key] = val
		}
	}
	return output
}

// DeleteMap DeleteMap
func DeleteMap(intput map[string]interface{}, dels []string) map[string]interface{} {
	if dels == nil {
		return intput
	}
	for _, key := range dels {
		delete(intput, key)
	}
	return intput
}
