#!/bin/bash

#Get all system logs
grpcurl -plaintext localhost:9089 aggregator.Aggregator.FetchNetworkLogs

#Get limited system logs
grpcurl -plaintext -d '{"limit":3}' localhost:9089 aggregator.Aggregator.FetchNetworkLogs

#Get total count of system logs
grpcurl -plaintext -d '{"count":true}' localhost:9089 aggregator.Aggregator.FetchNetworkLogs


