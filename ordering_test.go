package privat24go

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestOrderingFunction(t *testing.T) {
	Convey("Добавление описание", t, func() {
		o := Ordering{Amount: float64(49.50), PostingDate: "15.09.2014"}
		o.Payment = "ROW 3 НОМЕР IНДИВ КАРТ ДОСТУПУ 3633872 МОБ ТЕЛ 380971234567 ВІРНІЙНОМЕР IНДИВ КАРТ ДОСТУПУ3633872 МОБ ТЕЛ 380971234567 380971234567"
		So(o.Payment, ShouldEqual, "ROW 3 НОМЕР IНДИВ КАРТ ДОСТУПУ 3633872 МОБ ТЕЛ 380971234567 ВІРНІЙНОМЕР IНДИВ КАРТ ДОСТУПУ3633872 МОБ ТЕЛ 380971234567 380971234567")
	})
}
