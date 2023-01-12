# Protobuf 与 GPRC 使用
## 1. 编写 .proto 文件
命名要求: 
1. message 和 service, service 方法的名称要使用首字符大写的驼峰命名.
2. message 字段的名称使用蛇形命名.

service 中所有方法的参数和返回值都必须是一个 message 或者它的流.  
```golang
syntax = "proto3";

option go_package = "github.com/riicarus/grpclearn/rpc/route;route";

package route;

message Cpu {
    string brand = 1;
    string name = 2;
    string number_cores = 3;
    string number_threads = 4;
    double min_ghz = 5;
    double max_ghz = 6;
}

message CpuOfBrandRequest {
    string brand = 1;
}

message CpuOfNameRequest {
    string name = 1;
}

message CpuNumberResponse {
    int32 number = 1;
}

service CpuService {
    // unary
    rpc SearchCPU (CpuOfNameRequest) returns (Cpu);
    // server side stream
    rpc ListCPUOfOneBrand (CpuOfBrandRequest) returns (stream Cpu);
    // user side stream
    rpc CountNumber (stream CpuOfBrandRequest) returns (CpuNumberResponse);
    // bi-directional stream
    rpc ListCPUOfNames (stream CpuOfNameRequest) returns (stream Cpu);
}
```
我们在定义一个 .proto 文件时, 需要申明这个文件属于哪个包, 主要是为了规范整合以及避免重复, 这个概念在其他语言中也存在, 比如go中的 package.  
所以, 我们根据实际的分类情况, 给每1个proto文件都定义1个包, 一般这个包和 proto 所在的文件夹名称, 保持一致.  

`option go_package` 用于指定在生成 go 文件时, 对应的路径(`;` 前面部份)和 `package` 名称(`;` 后面部分)

## 2. 生成 go 代码
使用如下命令, 可以将其写入一个 shell 或者 bat 文件中, 便于使用.
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./rpc/route/*.proto
```

## 3. 使用 grpc
grpc 支持四种调用模式:  
1. unary: client 发送一条消息, server 返回一条消息.
2. server stream: client 发送一条消息, server 使用流进行返回. 
3. client stream: client 使用流发送消息, client 发送完后, server 返回一条消息. 
4. bi-direction stream: client 和 server 都使用流进行消息传递.

### 3.1 Server 基础实现
在 grpc 中, protoc-gen-go-grpc 为我们生成了我们在 .proto 中定义的 service 的接口, 分别包括 server 端和 client 端. client 端的方法调用已经为我们实现好了(即调用 grpc 提供的远程调用去调用 server 端的方法), 而 server 端的方法都是空实现, 所以需要我们自己来对 service 进行实现. 如: 
```golang
func (UnimplementedCpuServiceServer) SearchCPU(context.Context, *CpuOfNameRequest) (*Cpu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCPU not implemented")
}
```

要实现 service 的 server 端接口, 我们需要先定义一个实现了 server 方法的 struct:
```golang
// protoc-gen-go-grpc 生成的接口
type CpuServiceServer interface {
	// unary
	SearchCPU(context.Context, *CpuOfNameRequest) (*Cpu, error)
	// server side stream
	ListCPUOfOneBrand(*CpuOfBrandRequest, CpuService_ListCPUOfOneBrandServer) error
	// user side stream
	CountNumber(CpuService_CountNumberServer) error
	// bi-directional stream
	ListCPUOfNames(CpuService_ListCPUOfNamesServer) error

    // 必须继承这个类
	mustEmbedUnimplementedCpuServiceServer()
}

// UnimplementedCpuServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCpuServiceServer struct {
}

type cpuServiceServer struct {
    // 模拟数据库
	Cpus []*pb.Cpu

    // 必须继承这个类, 保证如果没有接口方法的实现, 可以调用父类的方法实现(即上面的空实现)
	pb.UnimplementedCpuServiceServer
}
```

struct 定义好之后, 需要实现接口的方法, 如:  
```golang
// unary
func (s *cpuServiceServer) SearchCPU(ctx context.Context, nameRequest *pb.CpuOfNameRequest) (*pb.Cpu, error) {
	for _, cpu := range s.Cpus {
		if cpu.Name == nameRequest.Name {
			return cpu, nil
		}
	}
	return nil, nil
}
```

到这里我们的 struct 的基础实现就已经做好了, 下面需要进行方法具体逻辑的编写.  

### 3.2 Server 启动
```golang
func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}

    // 创建一个新的 gprc 服务器
	grpcServer := grpc.NewServer()

    // 向服务器中注册自定义的服务
	pb.RegisterCpuServiceServer(grpcServer, NewServer())

    // 启动 grpc 服务
	log.Fatalln(grpcServer.Serve(lis))
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
```

### 3.2 Client 连接
```golang
conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
if err != nil {
	log.Fatalln(err)
}
defer conn.Close()

// 获取到 client, 后续的方法调用都是用 client 来调用.
client := pb.NewCpuServiceClient(conn)
```

### 3.3 Unary 模式编写
Unary 模式是最简单的模式, 无论是 server 还是 client 都只需要简单的进行单一参数的收发.  
```golang
// server unary
func (s *cpuServiceServer) SearchCPU(ctx context.Context, nameRequest *pb.CpuOfNameRequest) (*pb.Cpu, error) {
	for _, cpu := range s.Cpus {
		if cpu.Name == nameRequest.Name {
			return cpu, nil
		}
	}
	return nil, nil
}

// client unary
func runFirst(client pb.CpuServiceClient) {
    // 调用 client 的方法
	cpu, err := client.SearchCPU(context.Background(), &pb.CpuOfNameRequest{
		Name: "INTEL-I7-10875F",
	})

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(cpu)
	}
}
```

### 3.4 Server Stream 模式编写
在 Server Stream 模式中, Server 会使用 stream 的 `Send()` 来进行连续的数据发送, Client 调用方法获取到 stream , 使用 `Recv` 来接收 Server 发送的数据.  
> 需要了解 Server Stream 模式下, grpc 判断 stream 发送完毕的逻辑.

```golang
// Server: server side stream
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

// Client
func runSecond(client pb.CpuServiceClient) {
    // 向 Server 发送一个数据并获取 Server 返回数据的流
	serverStream, err := client.ListCPUOfOneBrand(context.Background(), &pb.CpuOfBrandRequest{
		Brand: "INTEL",
	})
	if err != nil {
		log.Fatalln(err)
	}

    // 循环从 stream 中读取数据
	for {
		cpu, err := serverStream.Recv()
        // 当读取到 io.EOF 时, 表示 Server Stream 已经发送完毕了
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(cpu)
	}
}
```

### 3.5 Client Stream 模式编写
在 Client Stream 模式中, Client 以流的形式连续发送数据, 当 Client 数据发送完毕, Server 再返回一个数据.  

```golang
// Client
func runThird(client pb.CpuServiceClient) {
    // Client 调用函数直接获取一个 ClientStream
	clientStream, err := client.CountNumber(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

    // 可以通过 ClientStream 多次发送数据
	clientStream.Send(&pb.CpuOfBrandRequest{
		Brand: "INTEL",
	})
	clientStream.Send(&pb.CpuOfBrandRequest{
		Brand: "AMD",
	})
	clientStream.Send(&pb.CpuOfBrandRequest{
		Brand: "AMD",
	})

    // 调用 clientStream.CloseAndRecv() 来表示 Client 的数据传输已经结束, 并且阻塞等待 Server 的返回数据.
    // 在此之后 Server 端读取会读取到一个 io.EOF err 表示当前流的数据输入已经结束.
	numberResponse, err2 := clientStream.CloseAndRecv()
	if err2 != nil {
		log.Fatalln(err2)
	}

	fmt.Println(numberResponse)
}

// Server: client side stream
func (s *cpuServiceServer) CountNumber(stream pb.CpuService_CountNumberServer) error {
	count := 0

    // 可以循环使用 stream.Recv() 函数来阻塞获取 Client 的数据输入, 直到读取到 io.EOF 的 err 来表示 Client 数据传输结束.
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

    // 在 Client 数据传输结束后, 调用 stream.SendAndClose() 来发送一次消息并且关闭 Server 端的输出流.
	stream.SendAndClose(&pb.CpuNumberResponse{
		Number: int32(count),
	})
	return nil
}
```

### 3.6 Bi-direction Stream 模式编写
在 Bi-direction Stream 模式中, Server 和 Client 系统都使用 stream 来相互发送消息.  
在同一个进程中, stream.Recv() 和 stream.Send() 方法可以同时进行调用. 但是不要在不同的进程中对 stream.Recv() 或 stream.Send() 方法进行调用.  

```golang
// Servcer: bi-directional stream
func (s *cpuServiceServer) ListCPUOfNames(bistream pb.CpuService_ListCPUOfNamesServer) error {
    // 在循环中调用 stream.Recv() 来多次获取 Client 传入的数据.
    // 当 读到 io.EOF 时, 表示 Client 的输入已经结束.
	for {
		nameRequest, err := bistream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(nameRequest)

        // 调用 stream.Send() 来向 Client 传递数据
		for _, cpu := range s.Cpus {
			if cpu.Name == nameRequest.Name {
				if err := bistream.Send(cpu); err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
}

// Client
func runForth(client pb.CpuServiceClient) {
    // Client 调用方法来获取 stream
	stream, err := client.ListCPUOfNames(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// this routine listen to the server stream
    // 可能 接收和发送会在不同的 routine 中, 进行异步的处理.
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
```