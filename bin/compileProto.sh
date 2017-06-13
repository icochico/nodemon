#!/usr/bin/env bash

# sources
GO_DIR=../src/ihmc.us/nodemon
# .proto files
PROTO_DIR=../include/ihmc.us/nodemon

# measure
MEASURE=measure
# sensors data
NETSENSOR=sensors/netsensor
# subjects
SUB_TRAFFIC=subjects/traffic
SUB_HOST=subjects/host
SUB_NETWORK=subjects/network
SUB_MOCKETS=subjects/mockets
SUB_CPU=subjects/cpu
SUB_MEMORY=subjects/memory
SUB_GEOLOCATION=subjects/geolocation
SUB_BATTERY=subjects/battery

# measure
protoc --go_out=$GO_DIR/$MEASURE -I=$PROTO_DIR/$MEASURE $PROTO_DIR/$MEASURE/*.proto;
# sensors: netsensor
protoc --go_out=$GO_DIR/$NETSENSOR -I=$PROTO_DIR/$NETSENSOR $PROTO_DIR/$NETSENSOR/*.proto;
# subjects: traffic
protoc --go_out=$GO_DIR/$SUB_TRAFFIC -I=$PROTO_DIR/$SUB_TRAFFIC $PROTO_DIR/$SUB_TRAFFIC/*.proto;
# subjects: host
protoc --go_out=$GO_DIR/$SUB_HOST -I=$PROTO_DIR/$SUB_HOST $PROTO_DIR/$SUB_HOST/*.proto;
# subjects: network
protoc --go_out=$GO_DIR/$SUB_NETWORK -I=$PROTO_DIR/$SUB_NETWORK $PROTO_DIR/$SUB_NETWORK/*.proto;
# subjects: mockets
protoc --go_out=$GO_DIR/$SUB_MOCKETS -I=$PROTO_DIR/$SUB_MOCKETS $PROTO_DIR/$SUB_MOCKETS/*.proto;
# subjects: memory
protoc --go_out=$GO_DIR/$SUB_MEMORY -I=$PROTO_DIR/$SUB_MEMORY $PROTO_DIR/$SUB_MEMORY/*.proto;
# subjects: cpu
protoc --go_out=$GO_DIR/$SUB_CPU -I=$PROTO_DIR/$SUB_CPU $PROTO_DIR/$SUB_CPU/*.proto;
# subjects: geolocation
protoc --go_out=$GO_DIR/$SUB_GEOLOCATION -I=$PROTO_DIR/$SUB_GEOLOCATION $PROTO_DIR/$SUB_GEOLOCATION/*.proto;
# subjects: battery
protoc --go_out=$GO_DIR/$SUB_BATTERY -I=$PROTO_DIR/$SUB_BATTERY $PROTO_DIR/$SUB_BATTERY/*.proto;