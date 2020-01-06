# lambda.provisioned.concurrency

To test differences in latency I used [ApacheBench tool](https://httpd.apache.org/docs/2.4/programs/ab.html)
It is installed on Mac. Example test commend

ab -n 10000 -c 100 https://some-url
Where n is a number of request and c param defines concurrent execution
