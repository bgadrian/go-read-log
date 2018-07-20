package process_lines

import "testing"

var lines = []string{
	"67.248.219.84 - - [08/Dec/2017:13:05:16 +0200] \"GET /a HTTP/1.1\" 9670 1508 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
"133.141.46.13 - - [08/Dec/2017:13:05:16 +0200] \"GET /a/c/d HTTP/1.1\" 738 9493 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
"97.215.90.134 - - [08/Dec/2017:13:05:16 +0200] \"GET /a HTTP/1.1\" 2329 7628 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
"38.183.7.237 - - [08/Dec/2017:13:05:16 +0200] \"GET /a/c/d HTTP/1.1\" 4963 1715 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
"145.165.238.133 - - [08/Dec/2017:13:05:16 +0200] \"GET /a/c/d HTTP/1.1\" 2950 2224 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
"26.122.48.212 - - [08/Dec/2017:13:05:16 +0200] \"GET /a/c/d HTTP/1.1\" 9343 9250 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
"87.71.152.3 - - [08/Dec/2017:13:05:16 +0200] \"GET /a HTTP/1.1\" 3338 5299 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
"54.246.114.76 - - [08/Dec/2017:13:05:16 +0200] \"GET /a/c HTTP/1.1\" 9701 534 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
"14.207.111.110 - - [08/Dec/2017:13:05:16 +0200] \"GET /a/c/d HTTP/1.1\" 1585 1127 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
"10.152.106.40 - - [08/Dec/2017:13:05:16 +0200] \"GET / HTTP/1.1\" 4672 9792 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
}

func TestRegex(t *testing.T) {
	res :=	Regex([]string{	"67.248.219.84 - - [08/Dec/2017:13:05:16 +0200] \"GET /a HTTP/1.1\" 9670 1508 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
	}, []int{1, 7, 9})

	if len(res) != 1 {
		t.Errorf("failed, found these results: '%v'", res)
	}

	should := []string{
		"67.248.219.84",
		"/a",
		"9670",
	}

	for i, v := range should{
		got := res[0][i]
		if got == v {
			continue
		}

		t.Errorf("exp '%s', got '%v'",
			v, got)
	}
}

func TestSplitn(t *testing.T) {
	res :=	Splitn([]string{	"67.248.219.84 - - [08/Dec/2017:13:05:16 +0200] \"GET /a HTTP/1.1\" 9670 1508 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
	}, " ", []int{0, 6, 8})

	if len(res) != 1 {
		t.Errorf("failed, found these results: '%v'", res)
	}

	should := []string{
		"67.248.219.84",
		"/a",
		"9670",
	}

	for i, v := range should{
		got := res[0][i]
		if got == v {
			continue
		}

		t.Errorf("exp '%s', got '%v'",
			v, got)
	}
}


func BenchmarkRegex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Regex([]string{	"67.248.219.84 - - [08/Dec/2017:13:05:16 +0200] \"GET /a HTTP/1.1\" 9670 1508 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
		}, []int{1, 7, 9})
	}
}
func BenchmarkSplitn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Splitn([]string{	"67.248.219.84 - - [08/Dec/2017:13:05:16 +0200] \"GET /a HTTP/1.1\" 9670 1508 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.5 Safari/535.19\"",
		}, " ", []int{0, 6, 8})
	}
}