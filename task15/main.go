package main

//К каким негативным последствиям может привести данный фрагмент кода, и как
//это исправить? Приведите корректный пример реализации.
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//func main() {
//	someFunc()
//}

var justString string

func main() {

	// при выполнении кода может возникнуть ошибка, если в строке нет 100 символов - выход за пределы массива
	someFunc()
	// для решения проблемы нужно добавить проверку на количество символов
}

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) > 100 {
		justString = v[:100]
	} else {
		justString = v
	}
}
