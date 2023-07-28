# 测试结果-优化后无连接池
## go benchmark 串行

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	     534	   2209578 ns/op	   12734 B/op	     212 allocs/op
PASS
ok  	tests	3.295s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1818	   3935195 ns/op	   12777 B/op	     215 allocs/op
PASS
ok  	tests	7.242s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1977	   1016544 ns/op	   12778 B/op	     215 allocs/op
PASS
ok  	tests	2.088s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1844	    666039 ns/op	   12775 B/op	     215 allocs/op
PASS
ok  	tests	1.314s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1936	    663659 ns/op	   12777 B/op	     215 allocs/op
PASS
ok  	tests	1.366s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    2056	    867575 ns/op	   12780 B/op	     215 allocs/op
PASS
ok  	tests	1.858s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1864	    615511 ns/op	   12773 B/op	     215 allocs/op
PASS
ok  	tests	1.232s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    2035	    627496 ns/op	   12778 B/op	     215 allocs/op
PASS
ok  	tests	1.353s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    2001	    641883 ns/op	   12777 B/op	     215 allocs/op
PASS
ok  	tests	1.362s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1756	    676288 ns/op	   12778 B/op	     215 allocs/op
PASS
ok  	tests	1.277s

## go benchmark 并发

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	     442	   3299661 ns/op	   14306 B/op	     218 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	2.638s


goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	   10000	   8516855 ns/op	   12923 B/op	     216 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	85.199s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	   10000	    132860 ns/op	   16646 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.358s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	   10000	    145383 ns/op	   16746 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.485s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9801	    146192 ns/op	   16699 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.465s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9624	    144175 ns/op	   16694 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.420s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9644	    143429 ns/op	   16683 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.418s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9513	    148575 ns/op	   16781 B/op	     232 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.445s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    8857	    146815 ns/op	   16607 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.336s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    8938	    146595 ns/op	   16692 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.346s



