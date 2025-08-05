package main

import (
	rpcclient "abdulmajid/fileserver/cmd/client/rpc_client"
	"abdulmajid/fileserver/internal/types"
	"flag"
	"fmt"
	"log"
)

func main() {
	method := flag.String("method", "", "RPC method to call (e.g. CreateDir)")
	addr := flag.String("addr", "localhost:1234", "RPC server address")
	dirName := flag.String("name", "", "Directory name (for CreateDir)")
	flag.Parse()

	if *method == "" {
		log.Fatal("Missing required flag: -method")
	}

	switch *method {
	case "CreateDir":
		if *dirName == "" {
			log.Fatal("Missing -name flag for CreateDir")
		}
		args := rpcclient.NewDirRequest(*dirName)
		var reply types.GenericResponse

		err := rpcclient.Call(*method, *addr, args, &reply)
		if err != nil {
			log.Fatal("RPC call failed:", err)
		}

		if reply.Success {
			fmt.Println("✅ Success:", reply.Message)
		} else {
			fmt.Println("❌ Failed:", reply.Message)
		}

	default:
		log.Fatalf("Unsupported method: %s", *method)
	}
}
