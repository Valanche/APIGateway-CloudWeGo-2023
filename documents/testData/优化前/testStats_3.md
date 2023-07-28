# 测试结果-优化前
## go benchmark 串行
goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	     628	   2981209 ns/op	   12750 B/op	     214 allocs/op
PASS
ok  	tests	3.888s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1902	   4871456 ns/op	   12796 B/op	     216 allocs/op
PASS
ok  	tests	9.351s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1938	    961537 ns/op	   12793 B/op	     216 allocs/op
PASS
ok  	tests	1.942s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1646	    672467 ns/op	   12785 B/op	     216 allocs/op
PASS
ok  	tests	1.205s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1927	    724708 ns/op	   12791 B/op	     216 allocs/op
PASS
ok  	tests	1.477s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1960	    911382 ns/op	   12793 B/op	     216 allocs/op
PASS
ok  	tests	1.866s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1944	    704206 ns/op	   12792 B/op	     216 allocs/op
PASS
ok  	tests	1.449s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1928	    717533 ns/op	   12792 B/op	     216 allocs/op
PASS
ok  	tests	1.463s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1923	    711209 ns/op	   12795 B/op	     216 allocs/op
PASS
ok  	tests	1.448s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1910	    693159 ns/op	   12796 B/op	     216 allocs/op
PASS
ok  	tests	1.406s

## go benchmark 并发
goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	     142	  12503369 ns/op	   14432 B/op	     216 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	3.183s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9570	   7439096 ns/op	   12818 B/op	     216 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	71.224s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9223	    146248 ns/op	   16549 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.383s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9196	    139633 ns/op	   16438 B/op	     230 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.317s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    7658	    150962 ns/op	   16599 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.193s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9490	    154406 ns/op	   16514 B/op	     230 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.497s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    8011	    151808 ns/op	   16617 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.256s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    8716	    144403 ns/op	   16497 B/op	     230 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.293s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    8180	    151715 ns/op	   16521 B/op	     230 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.278s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    8526	    154533 ns/op	   16466 B/op	     230 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.358s


