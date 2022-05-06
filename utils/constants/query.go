package constants

//Query constant
const (
	CREATE_CILIUM_TABLE = `CREATE TABLE if not exists cilium_logs (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"verdict" TEXT,
		"ip_source" TEXT,
		"ip_destination" TEXT,
		"ip_version" TEXT,
		"ip_encrypted" BOOLEAN,
		"l4_tcp_source_port" INTEGER,
		"l4_tcp_destination_port" INTEGER,
		"l4_udp_source_port" INTEGER,
		"l4_udp_destination_port" INTEGER,
		"l4_icmpv4_type" INTEGER,
		"l4_icmpv4_code" INTEGER,
		"l4_icmpv6_type" INTEGER,
		"l4_icmpv6_code" INTEGER,
		"source_namespace" TEXT,
		"source_labels" TEXT,
		"source_pod_name" TEXT,
		"destination_namespace" TEXT,
		"destination_labels" TEXT,
		"destination_pod_name" TEXT,
		"type" TEXT,
		"node_name" TEXT,

		"l7_type" TEXT,
		
		"l7_dns_cnames" TEXT,
		"l7_dns_observation_source" TEXT,

		"l7_http_code" INTEGER,
		"l7_http_method" TEXT,
		"l7_http_url" TEXT,
		"l7_http_protocol" TEXT,
		"l7_http_headers" TEXT,

		"event_type_type" INTEGER,
		"event_type_sub_type" INTEGER,

		"source_service_name" TEXT,
		"source_service_namespace" TEXT,
		"destination_service_name" TEXT,
		"destination_service_namespace" TEXT,
		"traffic_direction" TEXT,

		"trace_observation_point" TEXT,

		"drop_reason_desc" INTEGER,
		
		"is_reply" BOOLEAN,

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
		"labels" TEXT,
		"data" TEXT,
		"start_time" INTEGER,
		"updated_time" INTEGER,
		"result" TEXT,
		"total" INTEGER		
				
  	)`
	INSERT_KUBEARMOR = `INSERT INTO kubearmor_logs (cluster_name,host_name,namespace_name,
		pod_name,container_id,container_name,uid,type,source,operation,
		resource,labels,data,start_time,updated_time,result,total) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1)`
	INSERT_CILIUM = `INSERT INTO cilium_logs (verdict,
		ip_source,ip_destination,ip_version,ip_encrypted,
		l4_tcp_source_port,l4_tcp_destination_port,
		l4_udp_source_port,l4_udp_destination_port,
		l4_icmpv4_type,l4_icmpv4_code,
		l4_icmpv6_type,l4_icmpv6_code,
		source_namespace,source_labels,source_pod_name,
		destination_namespace,destination_labels,destination_pod_name,
		type,node_name,
		l7_type,l7_dns_cnames,l7_dns_observation_source,
		l7_http_code,l7_http_method,l7_http_url,l7_http_protocol,l7_http_headers,
		event_type_type,event_type_sub_type,
		source_service_name,source_service_namespace,
		destination_service_name,destination_service_namespace,
		traffic_direction,
		trace_observation_point,
		drop_reason_desc,
		is_reply,
		start_time,updated_time,total) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1)`

	SELECT_KUBEARMOR = `SELECT ID FROM kubearmor_logs WHERE cluster_name = ? and host_name = ? and namespace_name = ? and 
	pod_name = ? and container_id = ? and container_name = ? and uid = ? and type = ? and source = ? and operation = ? and 
	resource = ? and label = ? and data  = ? and result = ?`

	SELECT_CILIUM = `SELECT ID FROM cilium_logs WHERE verdict = ? and 
		ip_source = ? and ip_destination = ? and ip_version = ? and ip_encrypted = ? and 
		l4_tcp_source_port = ? and l4_tcp_destination_port = ? and 
		l4_udp_source_port = ? and l4_udp_destination_port = ? and 
		l4_icmpv4_type = ? and l4_icmpv4_code = ? and 
		l4_icmpv6_type = ? and l4_icmpv6_code = ? and 
		source_namespace = ? and source_labels = ? and source_pod_name = ? and 
		destination_namespace = ? and destination_labels = ? and destination_pod_name = ? and 
		type = ? and node_name = ? and 
		l7_type = ? and 
		l7_dns_cnames = ? and l7_dns_observation_source = ? and 
		l7_http_code = ? and l7_http_method = ? and l7_http_url = ? and l7_http_protocol = ? and l7_http_headers = ? and
		event_type_type = ? and event_type_sub_type = ? and 
		source_service_name = ? and source_service_namespace = ? and 
		destination_service_name = ? and destination_service_namespace = ? and 
		traffic_direction = ? and
		trace_observation_point = ? and drop_reason_desc = ? and is_reply = ?`

	UPDATE_CILIUM = `UPDATE cilium_logs SET total = total+1,updated_time = ?  where id = ?`

	UPDATE_KUBEARMOR = `UPDATE kubearmor_logs SET total = total+1,updated_time = ?  where id = ?`

	SELECT_ALL_KUBEARMOR = `SELECT cluster_name,host_name,namespace_name,pod_name,
	container_id,container_name,uid,type,source,operation,resource,data,start_time,updated_time,result,total FROM kubearmor_logs`

	SELECT_COUNT_KUBEARMOR = `SELECT count(id) from kubearmor_logs`

	SELECT_COUNT_CILIUM = `SELECT count(id) from cilium_logs`

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

	SELECT_ALL_CILIUM = `SELECT verdict,ip_source,ip_destination,
	ip_version,ip_encrypted,l4_tcp_source_port,l4_tcp_destination_port,l4_udp_source_port,l4_udp_destination_port,
	l4_icmpv4_type,l4_icmpv4_code,l4_icmpv6_type,l4_icmpv6_code,source_namespace,
	source_labels,source_pod_name,destination_namespace,destination_labels,
	destination_pod_name,type,node_name,l7_type,
	l7_dns_cnames,l7_dns_observation_source,
	l7_http_code,l7_http_method,l7_http_url,l7_http_protocol,l7_http_headers,
	event_type_type,event_type_sub_type,source_service_name,
	source_service_namespace,destination_service_name,destination_service_namespace,traffic_direction,
	trace_observation_point,drop_reason_desc,is_reply,
	start_time,updated_time,total from cilium_logs`
)
