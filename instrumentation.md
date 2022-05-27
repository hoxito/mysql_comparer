1 install libraries via:

go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp

2 import libraries in main.go and add this 2 lines:

        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2000", nil)

3 Add custom metrics :

a- add a counter

create a promCounter in controller: 

var DiffCounter = prometheus.NewCounter(
   prometheus.CounterOpts{
       Name: "ping_request_count",
       Help: "No of request handled by Diff handler",
   },
)

add the incrementer in controller GetDiff():

   metrics.DiffCounter.Inc()


Register the counter to the default register 

   prometheus.MustRegister(metrics.DiffCounter)
