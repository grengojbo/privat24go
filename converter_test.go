package privat24go

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReadXlsFile(t *testing.T) {
	Convey("Загрузка файла с выписками", t, func() {
		cnt, rows, loads := LoadXlsFile("testdata/c2bstatements.xls", 1, true)
		row := rows[1]
		So(loads, ShouldBeNil)
		So(cnt, ShouldEqual, 8)
		So(row.NumTransaction, ShouldEqual, "1113499SB")

		r := rows[3]
		So(loads, ShouldBeNil)
		So(r.PaymentSrc, ShouldEqual, "row 3 НОМЕР IНДИВ КАРТ.ДОСТУПУ 3633872 МОБ ТЕЛ 0971234567 ВІРНІЙ _;;НОМЕР IНДИВ КАРТ.ДОСТУПУ3633872 МОБ ТЕЛ 0971234567;0971234567;")
		So(r.Payment, ShouldEqual, "ROW 3 НОМЕР IНДИВ КАРТ ДОСТУПУ 3633872 МОБ ТЕЛ 380971234567 ВІРНІЙНОМЕР IНДИВ КАРТ ДОСТУПУ3633872 МОБ ТЕЛ 380971234567 380971234567")

		Convey("Ошибка нет такого файла", func() {
			cnt, rows, loads = LoadXlsFile("testdata/noexists.xls", 1, true)
			So(loads, ShouldNotBeNil)
			So(cnt, ShouldEqual, 0)
		})
	})
}
