package src

import "math/rand"

func RandomString() string {
	var str = []string{"Hello", "Kamusta", "Xin chào", "Bonjour", "你好", "こんにちは", "안녕하십니까", "Hallo", "Merhaba", "Ciao", "Hola", "Привет", "Olá", "नमस्ते"}
	return str[rand.Intn(13)]

}
