package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/riicarus/grpclearn/rpc/route"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewCpuServiceClient(conn)

	go runFirst(client)
	go runSecond(client)
	go runThird(client)

	runForth(client)
}

func runFirst(client pb.CpuServiceClient) {
	cpu, err := client.SearchCPU(context.Background(), &pb.CpuOfNameRequest{
		Name: "INTEL-I7-10875F",
	})

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(cpu)
	}
}

func runSecond(client pb.CpuServiceClient) {
	serverStream, err := client.ListCPUOfOneBrand(context.Background(), &pb.CpuOfBrandRequest{
		Brand: "INTEL",
	})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		cpu, err := serverStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(cpu)
	}
}

func runThird(client pb.CpuServiceClient) {
	clientStream, err := client.CountNumber(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	clientStream.Send(&pb.CpuOfBrandRequest{
		Brand: "INTEL",
	})
	clientStream.Send(&pb.CpuOfBrandRequest{
		Brand: "AMD",
	})
	clientStream.Send(&pb.CpuOfBrandRequest{
		Brand: "AMD",
	})

	numberResponse, err2 := clientStream.CloseAndRecv()
	if err2 != nil {
		log.Fatalln(err2)
	}

	fmt.Println(numberResponse)
}

func runForth(client pb.CpuServiceClient) {
	stream, err := client.ListCPUOfNames(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// this routine listen to the server stream
	go func() {
		for {
			cpu, err2 := stream.Recv()
			if err2 != nil {
				log.Fatalln(err2)
			}
	
			fmt.Println(cpu)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		request := pb.CpuOfNameRequest{}
		var name string
		fmt.Print("input name of cpu: ")

		readStringFromCommandLine(reader, &name)
		request.Name = name

		if err2 := stream.Send(&request); err2 != nil {
			log.Fatalln(err2)
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func readStringFromCommandLine(reader *bufio.Reader, target *string) {
	_, err := fmt.Fscanf(reader, "%s\n", target)
	if err != nil {
		log.Fatalln(err)
	}
}