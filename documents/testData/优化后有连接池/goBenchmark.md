# 测试结果-优化后有连接池

## go benchmark 串行

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	     394	   2715705 ns/op	   12710 B/op	     211 allocs/op
PASS
ok  	tests	2.295s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    2068	   4469513 ns/op	   12774 B/op	     215 allocs/op
PASS
ok  	tests	9.319s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    2035	    639748 ns/op	   12777 B/op	     215 allocs/op
PASS
ok  	tests	1.379s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1988	    651797 ns/op	   12776 B/op	     215 allocs/op
PASS
ok  	tests	1.373s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1800	    647580 ns/op	   12781 B/op	     215 allocs/op
PASS
ok  	tests	1.250s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1932	    654860 ns/op	   12775 B/op	     215 allocs/op
PASS
ok  	tests	1.345s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    2035	    646448 ns/op	   12781 B/op	     215 allocs/op
PASS
ok  	tests	1.391s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1874	    620651 ns/op	   12782 B/op	     215 allocs/op
PASS
ok  	tests	1.250s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    2031	    669946 ns/op	   12781 B/op	     215 allocs/op
PASS
ok  	tests	1.438s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentService-16    	    1860	    668561 ns/op	   12779 B/op	     215 allocs/op
PASS
ok  	tests	1.328s

## go benchmark 并发

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	     328	   3298967 ns/op	   14912 B/op	     219 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	2.312s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	   10000	   8977694 ns/op	   12927 B/op	     216 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	89.808s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9037	    136635 ns/op	   16618 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	2.257s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	   10000	    142540 ns/op	   16748 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.456s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9649	    130653 ns/op	   16505 B/op	     230 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.300s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9746	    135496 ns/op	   16756 B/op	     232 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.352s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	   10000	    135544 ns/op	   16781 B/op	     232 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.387s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    9936	    134104 ns/op	   16693 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.364s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    8864	    137327 ns/op	   16564 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.252s

goos: linux
goarch: amd64
pkg: tests
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkStudentServiceParallel-16    	    8230	    138234 ns/op	   16602 B/op	     231 allocs/op
testing: BenchmarkStudentServiceParallel-16 left GOMAXPROCS set to 8
PASS
ok  	tests	1.175s



