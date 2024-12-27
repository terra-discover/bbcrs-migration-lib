package lib

import (
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
)

func TestCustomFilters(t *testing.T) {
	replacer := strings.NewReplacer(`[`, ``, `]`, ``, `,`, ``, `"`, ``, ` `, ``, `1`, `?`, `20.5`, `?`, `value`, `?`, `null`, `NULL`, `true`, `?`, `false`, `?`)
	// filter single case
	qf := `["id","=",1]`
	queryFilters, whereFilters, _, _ := CustomFilters(qf, "", "")
	qf = replacer.Replace(qf)
	utils.AssertEqual(t, true, replacer.Replace(queryFilters) == qf && whereFilters != nil, "filter single case")

	// filter single case
	qf = `["id",1]`
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	utils.AssertEqual(t, true, queryFilters == "id = ?" && whereFilters != nil, "filter single case")

	// filter operator LIKE case
	qf = `["id","LIKE",1]`
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	qf = replacer.Replace(qf)
	utils.AssertEqual(t, true, replacer.Replace(queryFilters) == qf && whereFilters != nil, "filter operator LIKE case")

	// filter operator NOT LIKE case
	qf = `["name","NOT LIKE","value"]`
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	qf = replacer.Replace(qf)
	utils.AssertEqual(t, true, replacer.Replace(queryFilters) == qf && whereFilters != nil, "filter operator NOT LIKE case")

	// filter operator IS case
	qf = `["id","IS",null]`
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	qf = replacer.Replace(qf)
	utils.AssertEqual(t, true, replacer.Replace(queryFilters) == "("+qf+")" && whereFilters != nil, "filter operator IS case")

	// filter multiple case
	qf = `[["id","=",1],["AND"],["status",true],["OR"],["amount","=",20.5]]`
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	utils.AssertEqual(t, true, queryFilters == "id = ? AND status = ? OR amount = ?" && whereFilters != nil, "filter multiple case")

	// filter multiple OR case
	qf = `[["id","=",1],["status","=",true]]`
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	utils.AssertEqual(t, true, queryFilters == "id = ?  OR  status = ?" && whereFilters != nil, "filter multiple OR case")

	// filter operator IN case
	qf = `["payment_status","IN",["paid", "due", "partial"]]`
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	utils.AssertEqual(t, true, queryFilters == "payment_status IN ?" && whereFilters != nil, "filter operator IN case")

	// filter operator IN case
	qf = `["is_guest","IN",[true, false]]`
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	utils.AssertEqual(t, true, queryFilters == "is_guest IN ?" && whereFilters != nil, "filter operator IN case")

	// filter operator BETWEEN case
	qf = `["DATE(transaction_date)","BETWEEN",["date_start","date_end"]]` // date (YYYY-MM-DD hh:mm:ss)
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	utils.AssertEqual(t, true, queryFilters == "DATE(transaction_date) BETWEEN ? AND ?" && whereFilters != nil, "filter operator BETWEEN case")

	// filter nested object field case
	qf = `[["person.id","=","1"],["OR"],["person.status","=","value"]]`
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	qf = replacer.Replace(qf)
	utils.AssertEqual(t, true, replacer.Replace(queryFilters) == qf && whereFilters != nil, "filter nested object field case")

	// filter nested object ORM field case
	qf = `[["Person__id","=","1"],["OR"],["Person__status","=","value"]]`
	queryFilters, whereFilters, _, _ = CustomFilters(qf, "", "")
	qf = replacer.Replace(qf)
	utils.AssertEqual(t, true, replacer.Replace(queryFilters) == qf && whereFilters != nil, "filter nested object ORM field case")

	// search column like case
	q := `value`
	col := `["trx_id","id"]`
	_, _, querySearch, columnFilter := CustomFilters("", q, col)
	qs := replacer.Replace(querySearch)
	utils.AssertEqual(t, true, replacer.Replace(querySearch) == qs && columnFilter != nil, "search column like case")

	// filter negative case
	CreateFilter("unexpected json format")
}
