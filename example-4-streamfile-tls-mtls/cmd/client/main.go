package main

import (
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/elisardofelix/grpc-examples/example-4-streamfile-tls-mtls/proto"
)

func main() {

	certPool := x509.NewCertPool()
	cert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatalf("failed to read server certificate: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatal("failed to append server certificate")
	}

	// Create TLS credentials
	creds := credentials.NewClientTLSFromCert(certPool, "")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := proto.NewFileUploadServiceClient(conn)

	http.HandleFunc("/", downloadHandler(client))

	log.Printf("starting http server on address: %s", ":8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func downloadHandler(client proto.FileUploadServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// make request to gRPC server and initialise server stream
		stream, err := client.DownloadFile(ctx, &proto.DownloadFileRequest{Name: "gopher.png"})
		if err != nil {
			// check status code returned from server
			st := status.Convert(err)
			switch st.Code() {
			case codes.NotFound:
				http.Error(w, "File not found.", 404)
				return
			case codes.InvalidArgument:
				http.Error(w, "Bad request.", 400)
				return
			}

			http.Error(w, err.Error(), 500)
			return
		}

		log.Println("server stream started")

		// create slice of file contents
		var fileContents []byte

		for {
			// receive file chunk
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break // stream done
				}

				http.Error(w, err.Error(), 500)
				return
			}

			log.Println("received file chunk")

			// append file chunk to slice
			fileContents = append(fileContents, res.Content...)
		}

		log.Println("server stream done")

		// return file contents to user
		if _, err := w.Write(fileContents); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
}
