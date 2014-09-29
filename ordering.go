// Copyright 2014 Oleg Dolya. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package privat24go

import "strings"

// @Title Ordering
// @Description Выписка Privat24 Юр.лицо
type Ordering struct {
	ID             int64   `json:"id"`
	NumTransaction string  `json:"numTransaction"` // №
	PostingDate    string  `json:"postingDate"`    //   Дата проводки
	TimePosting    string  `json:"timePosting"`    //   Время проводки
	Amount         float64 `json:"amount"`         //   Сумма
	Currency       string  `json:"currency"`       //   Валюта
	PaymentSrc     string  `json:"paymentSrc"`     //   Назначение платежа необработаный
	Payment        string  `json:"payment"`        //   Назначение платежа
	EdrpouCompany  int     `json:"edrpouCompany"`  //   ЕГРПОУ контрагента
	NameCompany    string  `json:"nameCompany"`    //   Наименование контрагента
	AccountCompany int     `json:"accountCompany"` //   Счет контрагента
	MfoCompany     int     `json:"mfoCompany"`     //   МФО контрагента
	Reference      string  `json:"reference"`      //   Референс
	Invoice        string  `json:"invoice"`        //   Счет по которому платили(нет поля в выпеске)
	LiqPay         int     `json:"liqpay"`         //   LiqPay ID транзакции (нет поля в выпеске)
	Phone          string  `json:"phone"`          //   Номер телефона (нет поля в выпеске)
	Card           string  `json:"card"`           //   Номер карты (нет поля в выпеске)
	Description    string  `json:"description"`    //   пояснения при обработке (нет поля в выпеске)
}

func (o *Ordering) AddDescription(src string) {
	if len(o.Description) > 4 {
		o.Description = strings.Join([]string{o.Description, src + ";"}, " ")
	} else {
		o.Description = src + ";"
	}
}
