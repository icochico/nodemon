# nodemon
An aggregation server for statistics and sensor data
```
                   _
 _ __    ___    __| |  ___  _ __ ___    ___   _ __
| '_ \  / _ \  / _` | / _ \| '_ ` _ \  / _ \ | '_ \
| | | || (_) || (_| ||  __/| | | | | || (_) || | | |
|_| |_| \___/  \__,_| \___||_| |_| |_| \___/ |_| |_|
```

<b>NodeMon</b> is a server application that collects, shapes and aggregates data from local sensors. NodeMon uses NATS (http://nats.io/) to distribute to potential clients along the network.

<b>Dependencies</b>

Go (>= 1.8) https://golang.org/dl/<br/>
GNU make (suggested) https://www.gnu.org/software/make/<br/>

<b>Build</b>

With GNU make:

```make```

Without make:

```go get ihmc.us/nodemon``` <br/>
```go install ihmc.us/nodemon``` <br/>

<b>Run</b>

```Usage:
  anglova [command]
```
