# 测试文档
## go benchmark 串行
goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	     456	   2256780 ns/op	   12719 B/op	     211 allocs/op
PASS
ok  	tests	3.262s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1328	   4087926 ns/op	   12764 B/op	     214 allocs/op
PASS
ok  	tests	5.537s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1321	    968097 ns/op	   12770 B/op	     214 allocs/op
PASS
ok  	tests	1.387s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1263	    941582 ns/op	   12767 B/op	     214 allocs/op
PASS
ok  	tests	1.300s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1324	   1050104 ns/op	   12763 B/op	     214 allocs/op
PASS
ok  	tests	1.497s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1296	    937481 ns/op	   12769 B/op	     214 allocs/op
PASS
ok  	tests	1.325s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1215	   1016223 ns/op	   12756 B/op	     214 allocs/op
PASS
ok  	tests	1.350s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1296	    957058 ns/op	   12764 B/op	     214 allocs/op
PASS
ok  	tests	1.350s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1203	    975120 ns/op	   12770 B/op	     214 allocs/op
PASS
ok  	tests	1.291s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1242	   1012930 ns/op	   12766 B/op	     214 allocs/op
PASS
ok  	tests	1.371s

## go benchmark 并发
goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	     568	   3224431 ns/op	   13573 B/op	     216 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	3.691s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    4312	   4805897 ns/op	   12920 B/op	     216 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	20.770s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    4722	    784564 ns/op	   13942 B/op	     220 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	3.749s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    5186	    803508 ns/op	   13751 B/op	     219 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	4.209s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    5264	    370425 ns/op	   13879 B/op	     220 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.992s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    4218	    276500 ns/op	   13928 B/op	     220 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.217s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    5137	    290197 ns/op	   13924 B/op	     220 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.535s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    4146	    286654 ns/op	   13794 B/op	     219 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.238s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    5662	    693730 ns/op	   13844 B/op	     220 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	3.969s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    5132	    294092 ns/op	   13872 B/op	     220 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.552s

## JMeter
### 10 thread
Label	# Samples	Average	Median	90% Line	95% Line	99% Line	Min	Max	Error %	Throughput	Received KB/sec	Sent KB/sec
HTTP Request	577238	2	1	3	4	23	0	124	0.00%	9628.1754	1673.65	3140.44
TOTAL	577238	2	1	3	4	23	0	124	0.00%	9628.1754	1673.65	3140.44

### 100 thread
Label	# Samples	Average	Median	90% Line	95% Line	99% Line	Min	Max	Error %	Throughput	Received KB/sec	Sent KB/sec
HTTP Request	287800	20	9	48	70	138	0	1540	0.00%	4779.06378	830.74	1558.8
TOTAL	287800	20	9	48	70	138	0	1540	0.00%	4779.06378	830.74	1558.8


