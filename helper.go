package tripay

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)


func structToQuery(d interface{}) (query string, err error) {
	dType := reflect.TypeOf(d)
	if dType.Elem().Kind() != reflect.Struct {
		return "", errors.New("Parameter must be an Struct")
	}
	dValue := reflect.ValueOf(d)
	u := url.Values{}
	for  i:=0;i<dType.Elem().NumField();i++{
		field := dType.Elem().Field(i)
		key := field.Tag.Get("query")
		kind := field.Type.Kind()
		v :=dValue.Elem().Field(i)
		switch kind{
		case reflect.Int:
			k := strings.Split(key, ",")
			for j:=0;j<len(k);j++{
				if k[j] == "omitempty"|| k[j] == "omitzero"{
					if v.IsZero(){
						goto next
					}
				}
			}
			u.Add(k[0],strconv.Itoa(int(v.Int())))
		case reflect.String:
			k := strings.Split(key, ",")
			for j:=0;j<len(k);j++{
				if k[j] == "omitempty"{
					if v.String() == ""{
						goto next
					}
				}
			}
			u.Add(k[0], v.String())
		}
		next:
	}
	return u.Encode(), nil


}