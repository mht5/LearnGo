package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func createQuery1(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q)
		n := t.Name()
		query := fmt.Sprintf("insert into %s (", n)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			if i == 0 {
				query = fmt.Sprintf("%s%s", query, t.Field(i).Name)
			} else {
				query = fmt.Sprintf("%s, %s", query, t.Field(i).Name)
			}
		}
		query = fmt.Sprintf("%s) values (", query)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery1(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery1(e)
	i := 90
	createQuery1(i)
}
