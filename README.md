# NodeMon
A sensor fusion service for statistics and sensor data
```
                   _
 _ __    ___    __| |  ___  _ __ ___    ___   _ __
| '_ \  / _ \  / _` | / _ \| '_ ` _ \  / _ \ | '_ \
| | | || (_) || (_| ||  __/| | | | | || (_) || | | |
|_| |_| \___/  \__,_| \___||_| |_| |_| \___/ |_| |_|
```

<b>NodeMon</b> is a sensor fusion service that collects, shapes and distributes data from local sensors.<br/>
NodeMon uses NATS (http://nats.io/) to distribute the aggregated data to potential clients along the network.

<b>Dependencies</b>

Go (>= 1.8) https://golang.org/dl/<br/>
gNATSd (>= 0.9.6) http://nats.io/download/nats-io/gnatsd/<br />
GNU make (suggested) https://www.gnu.org/software/make/<br/>

<b>Build</b>

With GNU make:

```
cd nodemon
make
```

Without make:

```
cd nodemon
go get ihmc.us/nodemon
go install ihmc.us/nodemon
```

<b>Run</b>

Run gNATSd (NATS Server):

Unzip in a directory and launch:
```
  gnatsd (starts the NATS server)
```

Run NodeMon:
```
  nodemon (starts the NodeMon service)
```


