package inbox

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Listening(ctx context.Context, path string) {
	// to consume messages
	topic := "company"
	partition := 0

	conn, err := kafka.DialLeader(ctx, "tcp", path, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	go func() {
		defer func() {
			if err := conn.Close(); err != nil {
				log.Print("failed to close connection:", err)
			}
			fmt.Println("end..")
		}()
		for {
			fmt.Println("start..")
			conn.SetReadDeadline(time.Now().Add(10 * time.Second))
			batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

			b := make([]byte, 10e3) // 10KB max per message
			for {
				select {
				case <-ctx.Done():
					return
				case <-time.After(1 * time.Second):
				default:
					n, err := batch.Read(b)
					if err == nil {
						fmt.Println("read:", err)
					} else {
						log.Print(string(b[:n]))
					}
				}
			}

		}
	}()
}
