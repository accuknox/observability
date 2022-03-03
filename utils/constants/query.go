package constants

//Query constant
const (
	CREATE_CILIUM_TABLE = `CREATE TABLE if not exists cilium_logs (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"verdict" INTEGER,
		"ethernet_source" TEXT,
		"ethernet_destination" TEXT,
		"ip_source" TEXT,
		"ip_destination" TEXT,
		"ip_version" INTEGER,
		"ip_encrypted" BOOLEAN,
		"l4_tcp_source_port" INTEGER,
		"l4_tcp_destination_port" INTEGER,
		"l4_udp_source_port" INTEGER,
		"l4_udp_destination_port" INTEGER,
		"l4_icmpv4_type" INTEGER,
		"l4_icmpv4_code" INTEGER,
		"l4_icmpv6_type" INTEGER,
		"l4_icmpv6_code" INTEGER,
		"source_id" INTEGER,
		"source_identity" INTEGER,
		"source_namespace" TEXT,
		"source_labels" TEXT,
		"source_pod_name" TEXT,
		"destination_id" INTEGER,
		"destination_identity" INTEGER,
		"destination_namespace" TEXT,
		"destination_labels" TEXT,
		"destination_pod_name" TEXT,
		"type" INTEGER,
		"node_name" TEXT,
		"source_names" TEXT,
		"destination_names" TEXT,

		"l7_type" INTEGER,
		"l7_latency_ns" INTEGER,
		
		"l7_dns_query" TEXT,
		"l7_dns_ips" TEXT,
		"l7_dns_ttl" INTEGER,
		"l7_dns_cnames" TEXT,
		"l7_dns_observation_source" TEXT,
		"l7_dns_rcode" INTEGER,
		"l7_dns_qtypes" TEXT,
		"l7_dns_rrtypes" TEXT,

		"l7_http_code" INTEGER,
		"l7_http_method" TEXT,
		"l7_http_url" TEXT,
		"l7_http_protocol" TEXT,
		"l7_http_headers" TEXT,

		"l7_kafka_error_code" INTEGER,
		"l7_kafka_api_version" INTEGER,
		"l7_kafka_api_key" TEXT,
		"l7_kafka_correlation_id" INTEGER,
		"l7_kafka_topic" TEXT,

		"event_type_type" INTEGER,
		"event_type_sub_type" INTEGER,

		"source_service_name" TEXT,
		"source_service_namespace" TEXT,
		"destination_service_name" TEXT,
		"destination_service_namespace" TEXT,
		"traffic_direction" INTEGER,

		"policy_match_type" INTEGER,

		"trace_observation_point" INTEGER,

		"drop_reason_desc" INTEGER,
		
		"is_reply" BOOLEAN,
		"debug_capture_point" INTEGER,
		"interface_index" INTEGER,
		"interface_name" TEXT,
		"proxy_port" INTEGER,

		"start_time" INTEGER,
		"updated_time" INTEGER,
		"total" INTEGER		
  	)`
	CREATE_KUBEARMOR_TABLE = `CREATE TABLE if not exists kubearmor_logs (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"cluster_name" TEXT,
		"host_name" TEXT,
		"namespace_name" TEXT,
		"pod_name" TEXT,
		"container_id" TEXT,
		"container_name" TEXT,
		"uid" INTEGER,
		"type" TEXT,
		"source" TEXT,
		"operation" TEXT,
		"resource" TEXT,
		"data" TEXT,
		"start_time" INTEGER,
		"updated_time" INTEGER,
		"result" TEXT,
		"total" INTEGER		
				
  	)`
	INSERT_KUBEARMOR = `INSERT INTO kubearmor_logs (cluster_name,host_name,namespace_name,
		pod_name,container_id,container_name,uid,type,source,operation,
		resource,data,start_time,updated_time,result,total) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1)`
	INSERT_CILIUM = `INSERT INTO cilium_logs (verdict,ethernet_source,ethernet_destination,
		ip_source,ip_destination,ip_version,ip_encrypted,
		l4_tcp_source_port,l4_tcp_destination_port,
		l4_udp_source_port,l4_udp_destination_port,
		l4_icmpv4_type,l4_icmpv4_code,
		l4_icmpv6_type,l4_icmpv6_code,
		source_id,source_identity,source_namespace,source_labels,source_pod_name,
		destination_id,destination_identity,destination_namespace,destination_labels,destination_pod_name,
		type,node_name,source_names,destination_names,
		l7_type,l7_latency_ns,
		l7_dns_query,l7_dns_ips,l7_dns_ttl,l7_dns_cnames,l7_dns_observation_source,l7_dns_rcode,l7_dns_qtypes,l7_dns_rrtypes,
		l7_http_code,l7_http_method,l7_http_url,l7_http_protocol,l7_http_headers,
		l7_kafka_error_code,l7_kafka_api_version,l7_kafka_api_key,l7_kafka_correlation_id,l7_kafka_topic,
		event_type_type,event_type_sub_type,
		source_service_name,source_service_namespace,
		destination_service_name,destination_service_namespace,
		traffic_direction,
		policy_match_type,
		trace_observation_point,
		drop_reason_desc,
		is_reply,
		debug_capture_point,interface_index,interface_name,proxy_port,

		start_time,updated_time,total) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1)`

	SELECT_KUBEARMOR = `SELECT ID FROM kubearmor_logs WHERE cluster_name = ? and host_name = ? and namespace_name = ? and 
	pod_name = ? and container_id = ? and container_name = ? and uid = ? and type = ? and source = ? and operation = ? and 
	resource = ? and data  = ? and result = ?`

	SELECT_CILIUM = `SELECT ID FROM cilium_logs WHERE verdict = ? and ethernet_source = ? and ethernet_destination = ? and 
		ip_source = ? and ip_destination = ? and ip_version = ? and ip_encrypted = ? and 
		l4_tcp_source_port = ? and l4_tcp_destination_port = ? and 
		l4_udp_source_port = ? and l4_udp_destination_port = ? and 
		l4_icmpv4_type = ? and l4_icmpv4_code = ? and 
		l4_icmpv6_type = ? and l4_icmpv6_code = ? and 
		source_id = ? and source_identity = ? and source_namespace = ? and source_labels = ? and source_pod_name = ? and 
		destination_id = ? and destination_identity = ? and destination_namespace = ? and destination_labels = ? and destination_pod_name = ? and 
		type = ? and node_name = ? and source_names = ? and destination_names = ? and 
		l7_type = ? and l7_latency_ns = ? and 
		l7_dns_query = ? and l7_dns_ips = ? and l7_dns_ttl = ? and l7_dns_cnames = ? and l7_dns_observation_source = ? and l7_dns_rcode = ? and l7_dns_qtypes = ? and l7_dns_rrtypes = ? and 
		l7_http_code = ? and l7_http_method = ? and l7_http_url = ? and l7_http_protocol = ? and l7_http_headers = ? and
		l7_kafka_error_code = ? and l7_kafka_api_version = ? and l7_kafka_api_key = ? and l7_kafka_correlation_id = ? and l7_kafka_topic = ? and 
		event_type_type = ? and event_type_sub_type = ? and 
		source_service_name = ? and source_service_namespace = ? and 
		destination_service_name = ? and destination_service_namespace = ? and 
		traffic_direction = ? and policy_match_type = ? and
		trace_observation_point = ? and drop_reason_desc = ? and is_reply = ? and 
		debug_capture_point = ? and interface_index = ? and interface_name = ? and proxy_port = ?`

	UPDATE_CILIUM = `UPDATE cilium_logs SET total = total+1,updated_time = ?  where id = ?`

	UPDATE_KUBEARMOR = `UPDATE kubearmor_logs SET total = total+1,updated_time = ?  where id = ?`

	SELECT_ALL_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,uid,type,source,operation,resource,data,start_time,updated_time,result,total FROM kubearmor_logs`

	SELECT_HostName_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE host_name in (?)`

	SELECT_Namespace_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE namespace_name in (?)`

	SELECT_Pod_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE pod_name in (?)`

	SELECT_Container_ID_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE container_id in (?)`

	SELECT_Container_Name_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE container_name in (?)`

	SELECT_UID_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE uid in (?)`

	SELECT_type_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE type in (?)`

	SELECT_Source_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE source in (?)`

	SELECT_Operation_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE operation in (?)`

	SELECT_Resource_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE resource in (?)`

	SELECT_Data_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,total FROM kubearmor_logs WHERE data  in (?)`

	SELECT_ALL_CILIUM = `SELECT verdict,ethernet_source,ethernet_destination,ip_source,ip_destination,
	ip_version,ip_encrypted,l4_tcp_source_port,l4_tcp_destination_port,l4_udp_source_port,l4_udp_destination_port,
	l4_icmpv4_type,l4_icmpv4_code,l4_icmpv6_type,l4_icmpv6_code,source_id,source_identity,source_namespace,
	source_labels,source_pod_name,destination_id,destination_identity,destination_namespace,destination_labels,
	destination_pod_name,type,node_name,source_names,destination_names,l7_type,l7_latency_ns,l7_dns_query,
	l7_dns_ips,l7_dns_ttl,l7_dns_cnames,l7_dns_observation_source,l7_dns_rcode,l7_dns_qtypes,l7_dns_rrtypes,
	l7_http_code,l7_http_method,l7_http_url,l7_http_protocol,l7_http_headers,l7_kafka_error_code,l7_kafka_api_version,
	l7_kafka_api_key,l7_kafka_correlation_id,l7_kafka_topic,event_type_type,event_type_sub_type,source_service_name,
	source_service_namespace,destination_service_name,destination_service_namespace,traffic_direction,policy_match_type,
	trace_observation_point,drop_reason_desc,is_reply,debug_capture_point,interface_index,interface_name,proxy_port,
	start_time,updated_time,total from cilium_logs`
)
