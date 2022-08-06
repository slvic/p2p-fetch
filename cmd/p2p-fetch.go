//package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/slvic/p2p-fetch/internal/app"
//	"os"
//	"os/signal"
//	"syscall"
//)
//
//func run(ctx context.Context) error {
//	newApp, err := app.Initialize(ctx)
//	if err != nil {
//		return err
//	}
//	err = newApp.Run(ctx)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func main() {
//	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGINT)
//	defer cancel()
//	if err := run(ctx); err != nil {
//		_, _ = fmt.Fprintf(os.Stderr, "app run: %s\n", err.Error())
//	}
//}

package main

func main() {

}
