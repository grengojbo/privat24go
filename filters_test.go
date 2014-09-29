package privat24go

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFiltersFunction(t *testing.T) {
	Convey("Преобразуем номер телефона с National в International", t, func() {
		tel := "0971234567 НОМЕР IНДИВ 3633871 КАРТ НОМЕР IНДИВ КАРТ ДОСТУПУ 0003633872 MOБ ТЕЛ0971234567 380991234567"
		res := UpdatePhoneUa(tel)
		So(res, ShouldEqual, "380971234567 НОМЕР IНДИВ 3633871 КАРТ НОМЕР IНДИВ КАРТ ДОСТУПУ 0003633872 MOБ ТЕЛ380971234567 380991234567")
	})
}

func TestFilters(t *testing.T) {
	featureLiqPay := "LIQPAY 42733782DESCR Активацiя послуг Order"
	featureCardPhone := "НОМЕР IНДИВ КАРТ ДОСТУПУ 0003633872 MOБ ТЕЛ0971234567"
	Convey("Обработка Назначение платежа", t, func() {
		_, rows, _ := LoadXlsFile("testdata/c2bstatements.xls", 1, true)
		r := rows[3]
		So(r.Payment, ShouldEqual, "ROW 3 НОМЕР IНДИВ КАРТ ДОСТУПУ 3633872 МОБ ТЕЛ 380971234567 ВІРНІЙНОМЕР IНДИВ КАРТ ДОСТУПУ3633872 МОБ ТЕЛ 380971234567 380971234567")

		Convey("Проверям что платеж LiqPay", func() {
			res, ok := r.GetLiqpay()
			So(ok, ShouldBeFalse)

			r = rows[5]
			res, ok = r.GetLiqpay()
			So(ok, ShouldBeTrue)
			So(res, ShouldEqual, int(42733782))

			r.Payment = featureLiqPay
			res, ok = r.GetLiqpay()
			So(ok, ShouldBeTrue)
			So(res, ShouldEqual, int(42733782))
		})

		Convey("Поиск Теефона", func() {
			res, ok := r.GetPhone()
			So(ok, ShouldBeTrue)
			So(res, ShouldEqual, "380971234567")

			r := rows[6]
			res, ok = r.GetPhone()
			So(ok, ShouldBeTrue)
			So(res, ShouldEqual, "380981234567")

			r = rows[7]
			res, ok = r.GetPhone()
			So(ok, ShouldBeTrue)
			So(res, ShouldEqual, "380951234567")

			r = rows[2]
			_, ok = r.GetPhone()
			So(ok, ShouldBeFalse)
		})

		Convey("Заглушка Поиск номера счета", func() {
			_, ok := r.GetInvoice()
			So(ok, ShouldBeFalse)
		})

		Convey("Поиск номера индивидуальной карты доступа", func() {
			res, rev, ok := r.GetCard()
			So(ok, ShouldBeTrue)
			So(rev, ShouldEqual, int(100))
			So(res, ShouldEqual, "3633872")

			r.Payment = featureCardPhone
			res, rev, ok = r.GetCard()
			So(ok, ShouldBeTrue)
			So(res, ShouldEqual, "0003633872")

			r = rows[6]
			res, rev, ok = r.GetCard()
			So(ok, ShouldBeTrue)
			So(rev, ShouldEqual, int(100))
			So(res, ShouldEqual, "2477267")

			r = rows[7]
			res, rev, ok = r.GetCard()
			So(ok, ShouldBeTrue)
			So(res, ShouldEqual, "2802309")

			r = rows[4]
			res, rev, ok = r.GetCard()
			So(ok, ShouldBeTrue)
			So(rev, ShouldEqual, int(50))
			So(res, ShouldEqual, "0003601627")

			rf := rows[2]
			_, rif, okf := rf.GetCard()
			So(okf, ShouldBeFalse)
			So(rif, ShouldEqual, int(0))
		})

		Convey("Поиск услуги", func() {
			ok := r.FindCustom("СМС")
			So(ok, ShouldBeFalse)

			r := rows[1]
			ok = r.FindCustom("СМС")
			So(ok, ShouldBeTrue)

			r = rows[7]
			ok = r.FindCustom("SMS")
			So(ok, ShouldBeTrue)

			r = rows[0]
			ok = r.FindCustom("Вiдновлення картки")
			So(ok, ShouldBeTrue)
		})
	})
}
