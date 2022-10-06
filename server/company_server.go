package server

import (
	"companyRegistration-microservice/initializers"
	"companyRegistration-microservice/model"
	pb "companyRegistration-microservice/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	initializers.Migrate()
	initializers.Connect2DB()
}

type CompanyServer struct {
	pb.UnimplementedCompanyRegistrationServer
}

func (c *CompanyServer) RegisterCompany(ctx context.Context, req *pb.CompanyRegisterRequest) (*pb.RegisteredCompany, error) {
	registeredCompany := model.Company{
		Name:    req.GetName(),
		Address: req.GetAddress(),
		OIB:     req.GetOIB(),
	}

	tx := initializers.DB.Create(&registeredCompany)
	if tx.Error != nil {
		return nil, tx.Error
	}

	log.Printf("Company %v successfully registered", registeredCompany.Name)

	res := &pb.RegisteredCompany{
		ID:      registeredCompany.ID,
		Name:    registeredCompany.Name,
		Address: registeredCompany.Address,
		OIB:     registeredCompany.OIB,
	}

	return res, nil

}

func RunCompanyServer() {

	lis, err := net.Listen("tcp", "localhost:5050")
	if err != nil {
		log.Fatalf("could not start listening: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCompanyRegistrationServer(s, &CompanyServer{})

	log.Print("Server running on port:5050")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
