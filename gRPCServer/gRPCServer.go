package gRPCserver

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	config "github.com/AnoopKV/GoExercise23/configs"
	"github.com/AnoopKV/GoExercise23/data"
	"github.com/AnoopKV/GoExercise23/entities"
	pb "github.com/AnoopKV/GoExercise23/gRPCClient/proto/output/proto"
	"github.com/AnoopKV/GoExercise23/interfaces"
	"github.com/AnoopKV/GoExercise23/services"
	"github.com/AnoopKV/GoExercise23/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
	userCollection     *mongo.Collection
	userService        interfaces.IUser
	productCollection  *mongo.Collection
	productService     interfaces.IProduct
	mongoClient        *mongo.Client
	genericKey         string
)

type ecServer struct {
	pb.UnimplementedProductServiceServer
	pb.UnimplementedUserServiceServer
}

func Start(serverPort string, key string) {
	genericKey = key
	var err error
	if mongoClient, err = config.Connect2DB(utils.GetEnvVal("MONGO_CONNECTION_STRING")); err != nil {
		log.Panic(err.Error())
	}
	initializeUser()
	initializeProduct()

	_port, err := utils.ParseInt(serverPort)
	if err != nil {
		log.Panicln("Port not found!")
	}
	var port = flag.Int("port", _port, "the port to serve on")
	flag.Parse()
	fmt.Printf("server starting on port %d...\n", *port)

	cert, err := tls.LoadX509KeyPair(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}
	opts := []grpc.ServerOption{
		// The following grpc.ServerOption adds an interceptor for all unary RPCs. To configure an interceptor for streaming RPCs, see:
		// https://godoc.org/google.golang.org/grpc#StreamInterceptor
		grpc.UnaryInterceptor(ensureValidToken),
		// Enable TLS for all incoming connections.
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(s, &ecServer{})
	pb.RegisterProductServiceServer(s, &ecServer{})
	lis, _err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if _err != nil {
		log.Fatalf("failed to listen: %v", _err)
	}
	if __err := s.Serve(lis); __err != nil {
		log.Fatalf("failed to serve: %v", __err)
	}
}

func initializeUser() {
	userCollection = config.GetCollection(mongoClient, utils.GetEnvVal("USER_COLLECTION_NAME"), utils.GetEnvVal("DB_NAME"))
	userService = services.InitUserService(userCollection)
}

func initializeProduct() {
	productCollection = config.GetCollection(mongoClient, utils.GetEnvVal("PRODUCT_COLLECTION_NAME"), utils.GetEnvVal("DB_NAME"))
	productService = services.InitProductService(productCollection)
}

func (s *ecServer) Register(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {
	res, err := userService.Register(&entities.User{FirstName: req.FirstName, LastName: req.LastName, Age: int(req.Age), Email: req.Email, Password: req.Password, ConfirmPassword: req.ConfirmPassword})
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: res.Id.Hex(), FirstName: res.FirstName, LastName: res.LastName, Age: int32(res.Age), Email: res.Email, User_Type: res.User_Type}, nil
}

func (s *ecServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res, err := userService.Login(&entities.Login{Email: req.EmailId, Password: req.Password})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: res.TokenId}, nil
}

func (s *ecServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	err := userService.Logout(req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.LogoutResponse{Message: "Logged Out Successfully"}, nil
}

func (s *ecServer) AddProduct(ctx context.Context, req *pb.Product) (*pb.ProductCreateResponse, error) {
	res, err := productService.AddProduct(&entities.Product{Name: req.Name, Category: req.Category, Quantity: int(req.Quantity)})
	if err != nil {
		return nil, err
	}

	log.Println("Product ID: ", res)
	return &pb.ProductCreateResponse{Id: res.Hex()}, nil
}

func (s *ecServer) GetProductById(ctx context.Context, req *pb.ProductValue) (*pb.Product, error) {
	log.Println("Reached")
	if primitiveId, err := utils.SetId(req.Val); err != nil {
		log.Println("Exception in main.go, GetProductById", err.Error())
		return nil, err
	} else {
		if product, error := productService.GetProductById(*primitiveId); error == nil {
			return &pb.Product{Id: product.Id.Hex(), Name: product.Name, Category: product.Category, Quantity: int32(product.Quantity)}, nil
		} else {
			return nil, error
		}
	}
}

func (s *ecServer) SearchProduct(ctx context.Context, req *pb.ProductValue) (*pb.ProductsResponse, error) {
	if products, err := productService.SearchProducts(req.Val); err == nil {
		if products != nil && len(*products) > 0 {
			response := pb.ProductsResponse{}
			var datas []*pb.Product
			for _, val := range *products {
				datas = append(datas, &pb.Product{Id: val.Id.Hex(), Name: val.Name, Category: val.Category, Quantity: int32(val.Quantity)})
			}
			response.Product = datas
			return &response, nil
		}
		return nil, fmt.Errorf("no record found")
	}
	return nil, fmt.Errorf("no record found")
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == genericKey //"some-secret-token"
}

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}
