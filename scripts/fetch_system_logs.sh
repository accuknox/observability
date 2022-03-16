#!/bin/bash

#Get all system logs
grpcurl -plaintext localhost:9089 aggregator.Aggregator.FetchSystemLogs

#Get limited system logs
grpcurl -plaintext -d '{"limit":3}' localhost:9089 aggregator.Aggregator.FetchSystemLogs

#Get total count of system logs
grpcurl -plaintext -d '{"count":true}' localhost:9089 aggregator.Aggregator.FetchSystemLogs

#Get Namespace specific system logs 
grpcurl -plaintext -d '{"namespace":[<namespace_name>]}' localhost:9089 aggregator.Aggregator.FetchSystemLogs

#Get type(ContainerLog/HostLog) specific system logs
grpcurl -plaintext -d '{"type":"ContainerLog"}' localhost:9089 aggregator.Aggregator.FetchSystemLogs

#Get Operation specific system logs 
grpcurl -plaintext -d '{"operation":[<operation_name>]}' localhost:9089 aggregator.Aggregator.FetchSystemLogs

#Get Pod Name specific system logs 
grpcurl -plaintext -d '{"pod":[<pod_name>]}' localhost:9089 aggregator.Aggregator.FetchSystemLogs

#Get Host Name specific system logs 
grpcurl -plaintext -d '{"host":[<host_name>]}' localhost:9089 aggregator.Aggregator.FetchSystemLogs

#Get source based system logs
grpcurl -plaintext -d '{"source":"/bin/"}' localhost:9089 aggregator.Aggregator.FetchSystemLogs

#Get resource based system logs
grpcurl -plaintext -d '{"resource":"/bin/"}' localhost:9089 aggregator.Aggregator.FetchSystemLogs

#Get Container specific system logs 
grpcurl -plaintext -d '{"container":[<container_name>]}' localhost:9089 aggregator.Aggregator.FetchSystemLogs
