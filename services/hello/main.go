package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/nikunjy/go/protos/hello"
	"github.com/nikunjy/go/services/hello/server"

	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc"
)

type Options struct {
	ServerPort int `long:"proxy-port" descriptiont:"port which proxy listens on" default:"24689"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	srv := &server.Server{}
	server := grpc.NewServer()

	pb.RegisterHelloServer(server, srv)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.ServerPort))
	if err != nil {
		log.Fatal(err)
	}
	// Start Server
	go func() {
		log.Println("Setting proxy server port", opts.ServerPort)
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(server)
}

func waitForShutdown(server *grpc.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	server.GracefulStop()

	log.Println("Shutting down")
	os.Exit(0)
}
