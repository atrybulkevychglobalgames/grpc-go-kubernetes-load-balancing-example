package main

import (
    "context"
    "fmt"
    "github.com/atrybulkevychglobalgames/grpc-go-kubernetes-load-balancing-example/app/responder"
    "google.golang.org/grpc"
    "google.golang.org/grpc/balancer/roundrobin"
    "log"
    "os"
    "strconv"
    "time"
)

func main() {
    var conn *grpc.ClientConn

    args := os.Args
    fmt.Println( "------------>>>>>",args[1])

    conn, err := grpc.Dial(args[1], grpc.WithInsecure(),
        grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
    )
    if err != nil {
        log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()

    service := responder.NewResponderServiceClient(conn)

    ctx := context.Background()
    var i int64 = 1;
    for {
        i ++
        time.Sleep(250 * time.Millisecond)
        if resp, err := service.GetName(ctx, &responder.GetNameRequest{
            Value: strconv.FormatInt(i, 10),
        }); err != nil {
            log.Fatalf("failed to get response: %s", err)
        }else{
            fmt.Println( "1",resp.Value, resp.Name)
        }
    }
}
