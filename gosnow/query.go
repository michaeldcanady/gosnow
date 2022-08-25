package gosnow

import (
	"strings"
)

//Filter is a item to base a filter on
func Filter(str string) string {
	return strings.ToLower(str)
}

//Item is deprecated, do not use
func Item(str ...string) (total string) {
	for _, num := range str {
		total += strings.ToLower(num) + ","
	}
	total = total[:len(total)-1]
	return
}

//AND is obvious
func AND() string {
	return "^"
}

//OR is obvious
func OR() string {
	return "^OR"
}

//IS is obvious
func IS(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return "=" + str
}

//ISNOT is obvious
func ISNOT(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return "!=" + str
}

//ISONEOF is obvious
func ISONEOF(str ...string) string {
	return "IN" + Item(str...)
}

//ISNOTONEOF is obvious
func ISNOTONEOF(str ...string) string {
	return "NOT%2520IN" + Item(str...)
}

//ISEMPTY is obvious
func ISEMPTY() string {
	return "ISEMPTY"
}

//ISNOTEMPTY is obvious
func ISNOTEMPTY() string {
	return "ISNOTEMPTY"
}

//ISLESSTHAN is obvious
func ISLESSTHAN(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return "<" + str
}

//ISLESSTHANOREQUALS is obvious
func ISLESSTHANOREQUALS(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return "<=" + str
}

//ISGREATERTHAN is obvious
func ISGREATERTHAN(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return ">" + str
}

//ISGREATERTHANOREQUALS is obvious
func ISGREATERTHANOREQUALS(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return ">=" + str
}

//ISBETWEEN is obvious
func ISBETWEEN(str1, str2 string) string {
	return "BETWEEN" + str1 + "@" + str2
}

//ISANYTHING is obvious
func ISANYTHING() string {
	return "ANYTHING"
}

//ISSAMEAS is obvious
func ISSAMEAS(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return "SAMEAS" + str
}

//ISDIFFERENTFROM is obvious
func ISDIFFERENTFROM(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return "NSAMEAS" + str
}

//ORDERBY is obvious
func ORDERBY(str string) string {
	return "^ORDERBY" + str
}

//ORDERBYDESC is obvious
func ORDERBYDESC(str string) string {
	return "^ORDERBYDESC" + str
}

//ISLIKE is obvious
func ISLIKE(str string) string {
	return "LIKE" + str
}

//ISNOTLIKE is obvious
func ISNOTLIKE(str string) string {
	return "NOT%20LIKE" + str
}

//ON is obvious
func ON(str string) string {
	return "ON" + str
}

//NOTON is obvious
func NOTON(str string) string {
	return "NOTON" + str
}

//BEFORE is obvious
func BEFORE(str string) string {
	return "<" + str
}

//ATORBEFORE is obvious
func ATORBEFORE(str string) string {
	return "<=" + str
}

//AFTER is obvious
func AFTER(str string) string {
	return ">" + str
}

//ATORAFTER is obvious
func ATORAFTER(str string) string {
	return ">=" + str
}

//BETWEEN is obvious
func BETWEEN(str string) string {
	return "BETWEEN" + str
}

//MORETHAN is for date-type fields only!
func MORETHAN(str string) string {
	return "MORETHAN" + str
}

//LESSTHAN is for date-type fields only!
func LESSTHAN(str string) string {
	return "LESSTHAN" + str
}
