package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/riicarus/grpclearn/rpc/route"
)

type cpuServiceServer struct {
	// db simulation
	Cpus []*pb.Cpu

	pb.UnimplementedCpuServiceServer
}

// unary
func (s *cpuServiceServer) SearchCPU(ctx context.Context, nameRequest *pb.CpuOfNameRequest) (*pb.Cpu, error) {
	for _, cpu := range s.Cpus {
		if cpu.Name == nameRequest.Name {
			return cpu, nil
		}
	}
	return nil, nil
}
// server side stream
func (s *cpuServiceServer) ListCPUOfOneBrand(brandRequest *pb.CpuOfBrandRequest, stream pb.CpuService_ListCPUOfOneBrandServer) error {
	for _, cpu := range s.Cpus {
		if cpu.Brand == brandRequest.Brand {
			if err := stream.Send(cpu); err != nil {
				log.Fatalln(err)
			}
		}
	}

	return nil
}
// user side stream
func (s *cpuServiceServer) CountNumber(stream pb.CpuService_CountNumberServer) error {
	count := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		count ++
	}
	stream.SendAndClose(&pb.CpuNumberResponse{
		Number: int32(count),
	})
	return nil
}
// bi-directional stream
func (s *cpuServiceServer) ListCPUOfNames(bistream pb.CpuService_ListCPUOfNamesServer) error {
	for {
		nameRequest, err := bistream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(nameRequest)

		for _, cpu := range s.Cpus {
			if cpu.Name == nameRequest.Name {
				if err := bistream.Send(cpu); err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
}

func NewServer() *cpuServiceServer {
	return &cpuServiceServer{
		Cpus: []*pb.Cpu{
			{
				Brand: "INTEL",
				Name: "INTEL-I7-10875F",
				NumberCores: "8",
				NumberThreads: "16",
				MinGhz: 2.3,
				MaxGhz: 4.7,
			},
			{
				Brand: "INTEL",
				Name: "INTEL-I7-10800H",
				NumberCores: "8",
				NumberThreads: "16",
				MinGhz: 2.3,
				MaxGhz: 4.5,
			},
			{
				Brand: "AMD",
				Name: "AMD-R5000",
				NumberCores: "8",
				NumberThreads: "16",
				MinGhz: 2.2,
				MaxGhz: 4.4,
			},
		},
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCpuServiceServer(grpcServer, NewServer())
	log.Fatalln(grpcServer.Serve(lis))
}