package privat24go

import (
	"bufio"
	"strconv"
	"strings"
	// "time"

	"os"

	// github.com/opesun/goquery
	// github.com/opesun/goquery/exp/html
	// "github.com/kr/pretty"
	"github.com/puerkitobio/goquery"
	// "code.google.com/p/go.net/html"
)

// Load Privat24 ordering file -> c2bstatements.xls
// f - 0: вседанные, 1: только поступления, 2: только выплаты
// c - true очищать и преобразовывать номера телефонов
func LoadXlsFile(name string, f int, cl bool) (int, []*Ordering, error) {
	cnt := 0
	objects := make([]*Ordering, 0)
	v0, _ := strconv.ParseFloat("0.00", 64)
	xlsFile, err := os.Open(name)
	if err != nil {
		return cnt, objects, err
	}
	defer xlsFile.Close()
	r := bufio.NewReader(xlsFile)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return cnt, objects, err
	}

	doc.Find("tr[class=xl24]").Each(func(_ int, s *goquery.Selection) {
		// fmt.Printf("%#v\n", s)
		td := false
		row := new(Ordering)
		s.Find("td").Each(func(n int, c *goquery.Selection) {
			switch n {
			case 0:
				row.NumTransaction = strings.TrimSpace(c.Text())
				// fmt.Printf("%s|", c.Text())
				if len(row.NumTransaction) > 0 {
					td = true
				}
			case 1:
				// dateLayout := "30.12.2014"
				// d, _ := time.Parse(dateLayout, c.Text())
				// fmt.Println(d)
				row.PostingDate = strings.TrimSpace(c.Text())
				// fmt.Printf("%s|", row.PostingDate)
			case 2:
				// timeLayout := "23:50:59"
				// t, _ := time.Parse(timeLayout, strings.TrimSpace(c.Text()))
				// fmt.Printf("%s |", c.Text())
				row.TimePosting = strings.TrimSpace(c.Text())
				// fmt.Printf("%s|", row.TimePosting)
				// fmt.Printf("(%s) %d:%d:%d", c.Text(), row.TimePosting.Hour(), row.TimePosting.Minute(), row.TimePosting.Second())
			case 3:
				v, _ := strconv.ParseFloat(strings.TrimSpace(c.Text()), 64)
				row.Amount = float64(v)
				// fmt.Printf("%.2f|", row.Amount)
			case 4:
				row.Currency = strings.TrimSpace(c.Text())
				// fmt.Printf("(%d) %s |", n, strings.TrimSpace(c.Text()))
			case 5:
				row.PaymentSrc = CleanSpace(c.Text())
				// fmt.Printf("(%d) %s |", n, strings.TrimSpace(c.Text()))
			case 6:
				v, _ := strconv.ParseInt(strings.TrimSpace(c.Text()), 10, 32)
				row.EdrpouCompany = int(v)
				// fmt.Printf("%d|", row.EdrpouCompany)
			case 7:
				row.NameCompany = strings.TrimSpace(c.Text())
				// fmt.Printf("(%d) %s |", n, strings.TrimSpace(c.Text()))
			case 8:
				v, _ := strconv.ParseInt(strings.TrimSpace(c.Text()), 10, 32)
				row.AccountCompany = int(v)
				// fmt.Printf("%d|", row.AccountCompany)
			case 9:
				v, _ := strconv.ParseInt(strings.TrimSpace(c.Text()), 10, 32)
				row.MfoCompany = int(v)
				// fmt.Printf("%d|", row.MfoCompany)
			case 10:
				row.Reference = strings.TrimSpace(c.Text())
				// fmt.Printf("%s\n", strings.TrimSpace(c.Text()))
			}
		})
		if td {
			if cl {
				row.Payment = UpdatePhoneUa(row.PaymentSrc)
				row.Payment = Clean(row.Payment)
			}
			switch {
			case f == 0:
				objects = append(objects, row)
				cnt++
			case f == 1 && row.Amount > float64(v0):
				// fmt.Println("....................................................................................")
				objects = append(objects, row)
				cnt++
			case f == 2 && row.Amount < float64(v0):
				objects = append(objects, row)
				cnt++
			}
		}
		// fmt.Println(objects)
		// fmt.Println("")
	})
	return cnt, objects, nil
}
