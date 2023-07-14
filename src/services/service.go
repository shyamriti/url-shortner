package services

import "math/rand"

func ShortTheUrl(longUrl string) string {
	var runes = []rune("abcdefghijklmnopqrstuvwxyz!@#$%&*1234567890")
	str := make([]rune, 10)
	for i := 0; i < 10; i++ {
		str[i] = runes[rand.Intn(len(runes))]
	}
	return string(str)
}
