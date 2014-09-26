// Copyright 2014 Oleg Dolya. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package privat24go

import (
	"strings"
	"time"
)

// @Title Ordering
// @Description Выписка Privat24 Юр.лицо
type Ordering struct {
	ID             int64     `orm:"column(id);auto;pk" json:"id"`
	NumTransaction string    `orm:"size(255);index;null" json:"numTransaction"`         // №
	PostingDate    string    `orm:"size(10);index;null" json:"postingDate"`             //   Дата проводки
	TimePosting    string    `orm:"size(8);null" json:"timePosting"`                    //   Время проводки
	Amount         float64   `orm:"null;digits(12);decimals(2)" json:"amount"`          //   Сумма
	Currency       string    `orm:"size(3);default(UAH)" json:"currency"`               //   Валюта
	PaymentSrc     string    `orm:"size(255);null" json:"paymentSrc"`                   //   Назначение платежа необработаный
	Payment        string    `orm:"size(255);null" json:"payment"`                      //   Назначение платежа
	EdrpouCompany  int       `orm:"null" json:"edrpouCompany"`                          //   ЕГРПОУ контрагента
	NameCompany    string    `orm:"size(255);null" json:"nameCompany"`                  //   Наименование контрагента
	AccountCompany int       `orm:"null" json:"accountCompany"`                         //   Счет контрагента
	MfoCompany     int       `orm:"null" json:"mfoCompany"`                             //   МФО контрагента
	Reference      string    `orm:"size(255);null" json:"reference"`                    //   Референс
	Phone          string    `orm:"size(12);null" json:"phone"`                         //   Номер телефона (нет поля в выпеске)
	Description    string    `orm:"size(255);null" json:"description"`                  //   пояснения при обработке (нет поля в выпеске)
	Created        time.Time `orm:"auto_now_add;type(datetime);null"  json:"created"`   // Дата добавления (нет поля в выпеске)
	Updated        time.Time `orm:"auto_now;type(datetime);index;null"  json:"updated"` // Дата обновления (нет поля в выпеске)
}

// PostingDate    time.Time `orm:"type(datetime)" json:"postingDate"`         //   Дата проводки
// TimePosting    time.Time `orm:"type(datetime)" json:"timePosting"`         //   Время проводки
// regexp.Compile(`[0-9]+`)

// type Phones struct {
// 	Phone      string    `orm:"-" json:"phone"`            //   Номер телефона

// }

// Добавляем описание
func (o *Ordering) AddDescription(src string) {
	if len(o.Description) > 4 {
		o.Description = strings.Join([]string{o.Description, src}, "; ")
	} else {
		o.Description = src
	}
}
