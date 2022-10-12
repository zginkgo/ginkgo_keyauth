package utils

import (
	"testing"
)

func TestHash(t *testing.T) {
	v := Hash("abc")
	t.Log(v)
}

func TestPasswordHash(t *testing.T) {
	v := HashPassword("abc")
	t.Log(v)
	ok := CheckPasswordHash("abc", "$2a$14$0ZZjcMpa1SfPH55blNm1xOulFNaWtz67M5Tv8MeXvRA.8wqvRSNFC")
	t.Log(ok)
	ok = CheckPasswordHash("abc", "$2a$14$uK2YKOnmjLGV3lHpiXa6LONhvWd2ja9unZP0pEc7XmyeY2li5KwqC")
	t.Log(ok)
}
