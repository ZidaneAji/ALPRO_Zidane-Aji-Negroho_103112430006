package main

import "fmt"

func main() {
	var qirat, matauangfals, matauangdirham, matauangdinar int64

	fmt.Println("masukan uang dalam satuan qirat :")
	fmt.Scan(&qirat)

	matauangfals = qirat / 6
	matauangdirham = matauangfals / 10
	matauangdinar = matauangdirham / 10
	fmt.Println("konversi : ", matauangdinar, "dinar", matauangdirham,"dirham", matauangfals, "fals", qirat)
}
