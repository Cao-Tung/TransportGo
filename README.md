# TransportGo

## Test perform by go-wrk

### 1 worker

Running 5s test @ http://127.0.0.1:8080/v1/message/
  10 goroutine(s) running concurrently
43379 requests in 4.956300126s, 6.29MB read
Requests/sec:		8752.29
Transfer/sec:		1.27MB
Avg Req Time:		1.142557ms
Fastest Request:	106.031µs
Slowest Request:	32.267381ms
Number of Errors:	0

### 2 worker 

Running 5s test @ http://127.0.0.1:8080/v1/message/
  10 goroutine(s) running concurrently
43773 requests in 4.956018466s, 6.35MB read
Requests/sec:		8832.29
Transfer/sec:		1.28MB
Avg Req Time:		1.132209ms
Fastest Request:	99.165µs
Slowest Request:	19.171356ms
Number of Errors:	0

### 3 worker 

Running 5s test @ http://127.0.0.1:8080/v1/message/
  10 goroutine(s) running concurrently
44857 requests in 4.958756811s, 6.50MB read
Requests/sec:		9046.02
Transfer/sec:		1.31MB
Avg Req Time:		1.105458ms
Fastest Request:	107.22µs
Slowest Request:	18.671736ms
Number of Errors:	0
