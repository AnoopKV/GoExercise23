package grpcclient

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/AnoopKV/GoExercise23/data"
	pb "github.com/AnoopKV/GoExercise23/gRPCClient/proto/output/proto"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

var (
	globalKey string
)

type GRPCCLientService struct {
	UserGRPCClient    *pb.UserServiceClient
	ProductGRPCClient *pb.ProductServiceClient
}

func InitGRPCService(serverPort string, key string) *GRPCCLientService {
	globalKey = key

	log.Println("InitGRPCService Port::", serverPort)
	addr := flag.String("addr", "localhost:"+serverPort, "the address to connect to")
	flag.Parse()

	// Set up the credentials for the connection.
	perRPC := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(fetchToken())}
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	userRPC := pb.NewUserServiceClient(conn)
	productRPC := pb.NewProductServiceClient(conn)
	return &GRPCCLientService{UserGRPCClient: &userRPC, ProductGRPCClient: &productRPC}
}

func (g *GRPCCLientService) Register(user *pb.User) (*pb.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := (*(g.UserGRPCClient)).Register(ctx, user)
	if err != nil {
		log.Fatalf("client.Register(_) = _, %v: ", err)
	}
	fmt.Printf("Register: %#v\n", resp)
	return nil, nil
}

func (g *GRPCCLientService) Login(req *pb.LoginRequest) (*pb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := (*(g.UserGRPCClient)).Login(ctx, req)
	if err != nil {
		log.Fatalf("client.Login(_) = _, %v: ", err.Error())
	}
	fmt.Printf("Register: %#v", resp)
	return nil, nil
}

func (g *GRPCCLientService) Logout(client pb.UserServiceClient, req *pb.LogoutRequest) (*pb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := (*(g.UserGRPCClient)).Logout(ctx, req)
	if err != nil {
		log.Fatalf("client.Logout(_) = _, %v: ", err)
	}
	fmt.Printf("Register: %#v", resp)
	return nil, nil
}

// Product related
func (g *GRPCCLientService) AddProduct(req *pb.Product) (*pb.ProductCreateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := (*(g.ProductGRPCClient)).AddProduct(ctx, req)
	if err != nil {
		log.Fatalf("client.AddProduct(_) = _, %v: ", err)
		return nil, err
	}
	fmt.Printf("GetProductById: %#v\n", resp)
	return resp, nil
}

func (g *GRPCCLientService) GetProductById(req *pb.ProductValue) (*pb.Product, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := (*(g.ProductGRPCClient)).GetProductById(ctx, req)
	if err != nil {
		log.Fatalf("client.GetProductById(_) = _, %v: ", err)
	}
	fmt.Printf("GetProductById: %#v\n", resp)
	return nil, nil
}

func (g *GRPCCLientService) SearchProduct(req *pb.ProductValue) (*[]pb.ProductsResponse, error) {
	fmt.Printf("SearchProduct")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := (*(g.ProductGRPCClient)).SearchProduct(ctx, req)
	if err != nil {
		log.Fatalf("client.SearchProduct(_) = _, %v: ", err)
	}
	fmt.Printf("SearchProduct: %#v\n", resp)
	return nil, nil
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: globalKey,
	}
}
