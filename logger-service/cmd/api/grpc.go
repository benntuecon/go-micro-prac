package main

import (
	"context"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	logEntry := data.LogEntry{
		UserId: input.UserId,
		Action: input.Action,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		log.Panic("error inserting log entry ", err)
		res := &logs.LogResponse{Result: "failed logging"}
		return res, err
	}

	res := &logs.LogResponse{Result: "success logged"}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", gRPCPort)
	if err != nil {
		log.Panic("failed to listen on GRPC port", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

	log.Print("gRPC server started on port ", gRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Panic("failed to serve gRPC", err)
	}

}
