package main

type typeName struct {
	//...
}

var varName typeName         //①
varName := new(typeName) //②
varName := typeName{[初始化值]}  //③
varName := &typeName{[初始化值]} //④

// ①③返回 typeName 类型变量；
//②④返回 *typeName 类型变量；
//③④[]可省略；若无初始化值，则默认为零值
