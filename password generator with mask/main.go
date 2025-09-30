package main

import (
	"fmt"
	"math/rand"
	"time"
)

// сборник всех допустимых символов
var (
	nums       = "0123456789"
	lowbukva   = "abcdefghijklmnopqrstuvwxyz"
	upperbukva = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specsymb   = "!@#$%&*?-_"
	all        = nums + lowbukva + upperbukva + specsymb //суммируем все символы
)

// создание своей пароля
// генератор символов по маске
func rand_bukvi(maskabukv rune) rune {
	switch maskabukv {
	case '#':
		return rune(nums[rand.Intn(len(nums))])
	case 'a':
		return rune(lowbukva[rand.Intn(len(lowbukva))])
	case 'A':
		return rune(upperbukva[rand.Intn(len(upperbukva))])
	case '@':
		return rune(specsymb[rand.Intn(len(specsymb))])
	case '*':
		return rune(all[rand.Intn(len(all))])
	default:
		return maskabukv //из маски другие символы копируются и переводятся в результ
	}
}

// жеский генератор пароля
func ExGenPasswd(maska string) string {
	var passwd []rune              //создание среза для символов
	for _, repeat := range maska { //итерация символов
		passwd = append(passwd, rand_bukvi(repeat)) //динамичный срез через append
	}
	return string(passwd) //после возвращает то что собрал
}
func main() {
	rand.Seed(time.Now().UnixMilli()) //установка времени перебора и создания пароля в миллисекундах

	var maska string //данные о маске
	var paroli int   //данные о количестве паролей

	//запросы пошли
	fmt.Println("сначало введи маску для генерации пароля(например: Aa##@*): ")
	fmt.Scanln(&maska)
	fmt.Println("сколько паролей надо сгенерировать?(до 64):")
	fmt.Scanln(&paroli)

	fmt.Println("\nвот ваши сгенеренные пароли:")

	if paroli <= 64 { //накладывание ограничения на количество паролей
		for i := 0; i < paroli; i++ { //цикл чтобы выводил сразу несколько паролей
			time.Sleep(time.Second) //пауза каждую секунду
			passwd := ExGenPasswd(maska)
			fmt.Printf("%d) %s\n", i+1, passwd)
		}
	} else {
		fmt.Println("много хочешь!! >:(")
	}
	fmt.Println("\n\nпрограмма завершена!")

	comment()
}
func comment() {
	/*
		Написать программу,
		которая генерирует пароли по заданной маске

		генератор паролей:
		постоянный(var) набор символов
		|
		функция генерирует случайный символ по маске
		|
		функция- генерирует пароль на основе тела маски
		|
		введите маску = текст
		сколько паролей сгенерировать
		|
		мейн
		объявлять главные переменные (тело маски, количество паролей)
		|
		если число <= 10 то выводить цикл с сгенерированными паролями
		если больше, то nuh uh...
	*/
}
