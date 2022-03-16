package aggregator

import (
	"errors"
	"strconv"
	"strings"

	logger "github.com/accuknox/observability/src/logger"
	agg "github.com/accuknox/observability/src/proto/aggregator"
	"github.com/accuknox/observability/utils/constants"
	"github.com/accuknox/observability/utils/database"
	"github.com/accuknox/observability/utils/wrapper"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var log *zerolog.Logger = logger.GetInstance()

func GetSystemLogs(pbSystemRequest *agg.SystemLogsRequest) (agg.SystemLogsResponse, error) {
	var sysLogs []*agg.SystemLog
	var count int64
	var query string
	var filterQuery []string

	//Check Namespace Filter
	if len(pbSystemRequest.Namespace) != 0 {
		filterQuery = append(filterQuery, " namespace_name in ("+wrapper.ConvertFilterString(pbSystemRequest.Namespace)+")")
	}
	//Check Type filter
	if pbSystemRequest.Type != "" {
		filterQuery = append(filterQuery, " type = \""+pbSystemRequest.Type+"\"")
	}
	//Check Operation filter
	if len(pbSystemRequest.Operation) != 0 {
		filterQuery = append(filterQuery, " operation in ("+wrapper.ConvertFilterString(pbSystemRequest.Operation)+")")
	}
	//Check Pod filter
	if len(pbSystemRequest.Pod) != 0 {
		filterQuery = append(filterQuery, " pod_name in ("+wrapper.ConvertFilterString(pbSystemRequest.Pod)+")")
	}
	//Check Host filter
	if len(pbSystemRequest.Host) != 0 {
		filterQuery = append(filterQuery, " host_name in ("+wrapper.ConvertFilterString(pbSystemRequest.Host)+")")
	}
	//Check Source filter
	if pbSystemRequest.Source != "" {
		filterQuery = append(filterQuery, " source like \"%"+pbSystemRequest.Source+"%\"")
	}
	//Check Resource filter
	if pbSystemRequest.Resource != "" {
		filterQuery = append(filterQuery, " resource like \"%"+pbSystemRequest.Resource+"%\"")
	}
	//Check Container filter
	if len(pbSystemRequest.Container) != 0 {
		filterQuery = append(filterQuery, " container_name in ("+wrapper.ConvertFilterString(pbSystemRequest.Container)+")")
	}
	// TODO : Since Filter

	//Check Any filter exist
	if len(filterQuery) != 0 {
		query = " where" + strings.Join(filterQuery, " and")
	}
	//Check User want log or count of log
	if pbSystemRequest.Count {
		query = constants.SELECT_COUNT_KUBEARMOR + query
		//Fetch rows
		row := database.ConnectDB().QueryRow(query)
		if err := row.Scan(&count); err != nil {
			log.Error().Msg("Error in Connection in System Logs :" + err.Error())
			return agg.SystemLogsResponse{}, errors.New("error in Connecting system logs table")
		}
		return agg.SystemLogsResponse{Count: count}, nil
	} else {
		query = constants.SELECT_ALL_KUBEARMOR + query + constants.ORDER_BY_UPDATED_TIME
		//Check limit exist
		if pbSystemRequest.Limit != 0 {
			//query to fetch all logs with limit
			query = query + constants.LIMIT + strconv.FormatInt(pbSystemRequest.Limit, 10)
		}
		//Fetch rows
		rows, err := database.ConnectDB().Query(query)
		if err != nil {
			log.Error().Msg("Error in Connection in System Logs :" + err.Error())
			return agg.SystemLogsResponse{}, errors.New("error in Connecting system logs table")
		}
		defer rows.Close()
		for rows.Next() {
			var sysLog agg.SystemLog
			//Scan logs
			if err := rows.Scan(&sysLog.ClusterName, &sysLog.HostName,
				&sysLog.Namespace, &sysLog.PodName, &sysLog.ContainerID, &sysLog.ContainerName,
				&sysLog.Uid, &sysLog.Type, &sysLog.Source, &sysLog.Operation, &sysLog.Resource,
				&sysLog.Data, &sysLog.StartTime, &sysLog.UpdateTime, &sysLog.Result, &sysLog.Total); err != nil {
				log.Error().Msg("Error in Scan system Logs : " + err.Error())
				return agg.SystemLogsResponse{}, status.Errorf(codes.InvalidArgument, "Error in scanning system logs table")
			}
			//append record
			sysLogs = append(sysLogs, &sysLog)
		}
	}
	return agg.SystemLogsResponse{Logs: sysLogs, Count: count}, nil
}
