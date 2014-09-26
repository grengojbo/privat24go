// Copyright 2014 Oleg Dolya. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package privat24go

import (
	"regexp"
	"strconv"
	"strings"
)

// type Pattern string

// func (p Pattern) Compile() (*regexp.Regexp, error) {
// 	return regexp.Compile(string(p))
// }

// Преобразуем номер телефона с National в International
func UpdatePhoneUa(national string) (international string) {
	re := regexp.MustCompile(`(039|050|063|066|067|068|091|092|093|094|095|096|097|098|099([0-9]{7}){10})`)
	r := regexp.MustCompile(`(38380)`)
	international = re.ReplaceAllString(national, "38$0")
	return r.ReplaceAllString(international, "380")
}

// Очищаем от табуляций, лишних пробелов
func CleanSpace(src string) (res string) {
	res = strings.Replace(src, "\n", " ", -1)
	res = strings.Replace(res, "\t", " ", -1)
	res = strings.TrimSpace(res)

	r, _ := regexp.Compile(`\s(\s)+`)
	res = r.ReplaceAllString(res, "")
	return res
}

// Убираем все лишнее и переводим все в верхний регистр
func Clean(src string) (res string) {
	res = strings.Replace(src, "\n", " ", -1)
	res = strings.Replace(res, "\t", " ", -1)
	res = strings.Replace(res, "+", " ", -1)
	res = strings.Replace(res, "_", " ", -1)
	res = strings.Replace(res, ";", " ", -1)
	res = strings.Replace(res, ",", " ", -1)
	res = strings.Replace(res, ".", " ", -1)
	res = strings.Replace(res, "?", "і", -1)
	res = strings.TrimSpace(res)
	res = strings.ToUpper(res)

	r, _ := regexp.Compile(`\s(\s)+`)
	res = r.ReplaceAllString(res, "")
	return res
}

// Поиск номера индивидуальной карты доступа
// Return: card - номер карты
// 				 rev - Ревелентность результата если 100 это точно карта доступа
// 							 50 определил на основании длины числа от 6 до 10 чисел
func (t *Ordering) GetCard() (card string, rev int, ok bool) {
	rev = 100
	r, _ := regexp.Compile(`ДОСТУПУ (\d+)`)
	re := regexp.MustCompile(`([0-9]{6,10})`)
	// r := regexp.MustCompilePOSIX(`[ДОСТУПУ |ДОСТУПУ]([0-9]+)`)
	f := r.FindStringSubmatch(t.Payment)
	if len(f) > 0 {
		// fmt.Println(">>>>>>>>>>", strings.TrimSpace(f[1]))
		return strings.TrimSpace(f[1]), rev, true
		// } else {
		// r := regexp.MustCompilePOSIX(`(00[0-9]{10}|)`)
	} else {
		res := re.FindStringSubmatch(t.Payment)
		if len(res) > 0 {
			return strings.TrimSpace(res[1]), int(50), true
		}
	}
	return "", int(0), false
}

// Поиск номера телефона
// TODO: сейчас возвращает первый найденый, добавить все найденые телефоны без дубликатов
func (t *Ordering) GetPhone() (phone string, ok bool) {
	re := regexp.MustCompile(`(38[0-9]{10})`)
	res := re.FindStringSubmatch(t.Payment)
	if len(res) > 0 {
		return strings.TrimSpace(res[1]), true
	}
	return "", false
}

func (t *Ordering) GetLiqpay() (res int, ok bool) {
	res = 0
	ok = false
	r, _ := regexp.Compile(`LIQPAY (\d+)`)
	f := r.FindStringSubmatch(t.Payment)
	if len(f) > 0 {
		v, err := strconv.ParseInt(f[1], 10, 32)
		if err != nil {
			return res, ok
		}
		return int(v), true
	}
	// res = f[1]
	return res, ok
}

// произвольный поиск например название услуги
func (t *Ordering) FindCustom(m string) (ok bool) {
	m = strings.TrimSpace(strings.ToUpper(m))
	r, _ := regexp.Compile(m)
	return r.MatchString(t.Payment)
}

// TODO: Поиск номера счета
func (t *Ordering) GetInvoice() (res string, ok bool) {
	return "", false
}
