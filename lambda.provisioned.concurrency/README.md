# lambda.provisioned.concurrency

To test differences in latency I used [ApacheBench tool](https://httpd.apache.org/docs/2.4/programs/ab.html)
It is installed on Mac. Example test commend

ab -n 1000 -c 10 https://some-url
Where n is a number of request and c param defines concurrent execution

## Example results

### Normal lambda

Connection Times (ms)
min mean[+/-sd] median max
Connect: 154 227 158.7 187 1344
Processing: 70 108 53.8 95 973
Waiting: 70 104 49.9 93 972
Total: 229 335 173.8 290 1819

Percentage of the requests served within a certain time (ms)
50% 290
66% 312
75% 326
80% 340
90% 437
95% 555
98% 1154
99% 1378
100% 1819 (longest request)

### Provisioned concurrency lambda

Connection Times (ms)
min mean[+/-sd] median max
Connect: 155 193 33.1 185 574
Processing: 71 106 25.2 102 279
Waiting: 70 105 24.0 100 279
Total: 231 299 42.4 292 648

Percentage of the requests served within a certain time (ms)
50% 292
66% 309
75% 318
80% 325
90% 351
95% 372
98% 399
99% 433
100% 648 (longest request)
