/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// The client demonstrates how to supply an OAuth2 token for every RPC.
package main

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

func Register(client pb.UserServiceClient, user *pb.User) (*pb.UserResponse, error) { //
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Register(ctx, user)
	if err != nil {
		log.Fatalf("client.Register(_) = _, %v: ", err)
	}
	fmt.Printf("Register: %#v", resp)
	return nil, nil
}

func Login(client pb.UserServiceClient, req *pb.LoginRequest) (*pb.LoginResponse, error) { //
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Login(ctx, req)
	if err != nil {
		log.Fatalf("client.Register(_) = _, %v: ", err)
	}
	fmt.Printf("Register: %#v", resp)
	return nil, nil
}

func Logout(client pb.UserServiceClient, req *pb.LogoutRequest) (*pb.LoginResponse, error) { //
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Logout(ctx, req)
	if err != nil {
		log.Fatalf("client.Register(_) = _, %v: ", err)
	}
	fmt.Printf("Register: %#v", resp)
	return nil, nil
}

// Product related
func AddProduct(client pb.ProductServiceClient, req *pb.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.AddProduct(ctx, req)
	if err != nil {
		log.Fatalf("client.GetProductById(_) = _, %v: ", err)
	}
	fmt.Printf("GetProductById: %#v\n", resp)
	return nil
}

func GetProductById(client pb.ProductServiceClient, req *pb.ProductValue) (*pb.Product, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.GetProductById(ctx, req)
	if err != nil {
		log.Fatalf("client.GetProductById(_) = _, %v: ", err)
	}
	fmt.Printf("GetProductById: %#v\n", resp)
	return nil, nil
}

func SearchProduct(client pb.ProductServiceClient, req *pb.ProductValue) (*[]pb.ProductsResponse, error) {
	fmt.Printf("SearchProduct")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SearchProduct(ctx, req)
	if err != nil {
		log.Fatalf("client.SearchProduct(_) = _, %v: ", err)
	}
	fmt.Printf("SearchProduct: %#v\n", resp)
	return nil, nil
}

func main() {
	serverPort := "50051"
	key := "S3cr3tVA1U3"
	globalKey = key

	log.Println("Port::", serverPort)
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
	defer conn.Close()
	//userRPC := pb.NewUserServiceClient(conn)
	productRPC := pb.NewProductServiceClient(conn)
	//Register(userRPC, &pb.User{FirstName: "Anoop", LastName: "Varambally", Age: 30, Email: "anoop@noreply.com", Password: "Password", ConfirmPassword: "Password"})
	//Login(userRPC, &pb.LoginRequest{EmailId: "anoop.kv@noreply.com", Password: "654789"})
	//GetProductById(productRPC, &pb.ProductValue{Val: "65142632c827b0e57141f246"})
	SearchProduct(productRPC, &pb.ProductValue{Val: "SamsungS9"})
}

// fetchToken simulates a token lookup and omits the details of proper token
// acquisition. For examples of how to acquire an OAuth2 token, see:
// https://godoc.org/golang.org/x/oauth2
func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: globalKey,
	}
}
