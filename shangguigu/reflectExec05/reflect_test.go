package main

import (
	"reflect"
	"testing"
)

type Student struct {
	Name string
	Age  int
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *Student
		tye   reflect.Type
		val   reflect.Value
	)

	tye = reflect.TypeOf(model)                           //获取类型*Student
	t.Log("reflect.TypeOf", tye.Kind().String())          //Kind是ptr
	tye = tye.Elem()                                      //tye指向的类型
	t.Log("reflect.TypeOf.Elem", tye.Kind().String())     //指针地址指向的Kind是struct
	val = reflect.New(tye)                                //New返回一个Value类型值，该值持有一个指向类型为tye的新申请的零值的指针
	t.Log("reflect.New", val.Kind().String())             //ptr
	t.Log("reflect.New.Elem", val.Elem().Kind().String()) //struct

	//model就是创建的Student结构体变量(实例)
	model = val.Interface().(*Student)      //通过类型断言，将model转成*Student，它的指向和elem是一样的
	val = val.Elem()                        //取得val指向的值
	val.FieldByName("Name").SetString("猪猪") //赋值
	val.FieldByName("Age").SetInt(2)
	t.Log("model model.Name", model, model.Name)

}
