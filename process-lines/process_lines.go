package process_lines

import (
	"strings"
	"regexp"
)

var expression = regexp.MustCompile(`(?P<ipaddress>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) - - \[(?P<dateandtime>\d{2}\/[a-zA-Z]{3}\/\d{4}:\d{2}:\d{2}:\d{2} (\+|\-)\d{4})\] ((\"(GET|POST) )(?P<url>.+) (HTTP\/1\.1")) (?P<bytessent>\d+) (?P<bytesrec>\d+) (["](?P<refferer>(\-)|(.+))["]) (["](?P<useragent>.+)["])`)

//var names = expression.SubexpNames()

func Regex(src []string, indexs []int) [][]string{
	var res [][]string

	for _, line := range src {
		m := expression.FindStringSubmatch(line)

		wanted := make([]string, len(indexs))
		for i, index := range indexs {
			wanted[i] = m[index]
		}
		res = append(res, wanted)

	}
	return res
}

//Scan doesn't work because it stops at space, and we have many spaces
//67.248.219.84 - - [08/Dec/2017:13:05:16 +0200] "GET /a HTTP/1.1" 9670 1508 "-" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19"
//const apacheFormatPattern = "%s - - [%s] "%s %s %s" %d %d %.4f\n"
//const apacheFormatPattern = "%s - - [%s] \"%s %s %s\" %s %s %s"
//
//func Scanf(src []string, indexs []int) [][]string {
//	res := make([][]string, len(src))
//
//	for _, line := range src {
//		var a1, a2, a3, a4, a5,a6,a7,a8 string
//		fmt.Sscanf(line, apacheFormatPattern, &a1, &a2, &a3, &a4, &a5, &a6, &a7, &a8 )
//
//		s := [8]string {a1,a2,a3,a4,a5,a6,a7,a8}
//		wanted := make([]string, len(indexs))
//		for i, index := range indexs {
//			wanted[i] = s[index]
//		}
//		res = append(res, wanted)
//	}
//	return res
//}

// Splitn is not as powerful as regex, but a lot faster.
// probably in a production production multiple
// splitn and checkups can replace a regex and still
// being 15x faster.
func Splitn(src []string, separator string, indexs []int) [][]string{
	var res [][]string
	var bySpace []string

	for _, line := range src {
		if len(line) < 1 {
			continue
		}
		largestIndex := indexs[len(indexs)-1]+2
		bySpace = strings.SplitN(line, separator, largestIndex)
		wanted := make([]string, len(indexs))
		for i, index := range indexs {
			wanted[i] = bySpace[index]
		}
		res = append(res, wanted)
	}
	return res
}
