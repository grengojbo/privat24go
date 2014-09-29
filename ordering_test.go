package privat24go

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestOrderingFunction(t *testing.T) {
	Convey("Добавление описание", t, func() {
		o := Ordering{Amount: float64(49.50), PostingDate: "15.09.2014", TimePosting: "14:50:00"}
		o.AddDescription("#1001 ERR: ошибочка")
		o.Payment = "ROW 3 НОМЕР IНДИВ КАРТ ДОСТУПУ 3633872 МОБ ТЕЛ 380971234567 ВІРНІЙНОМЕР IНДИВ КАРТ ДОСТУПУ3633872 МОБ ТЕЛ 380971234567 380971234567"
		So(o.Payment, ShouldEqual, "ROW 3 НОМЕР IНДИВ КАРТ ДОСТУПУ 3633872 МОБ ТЕЛ 380971234567 ВІРНІЙНОМЕР IНДИВ КАРТ ДОСТУПУ3633872 МОБ ТЕЛ 380971234567 380971234567")
		So(o.Description, ShouldEqual, "#1001 ERR: ошибочка;")

		o.AddDescription("#1013 ERR: еще ошибочка")
		So(o.Description, ShouldEqual, "#1001 ERR: ошибочка; #1013 ERR: еще ошибочка;")
	})
}
