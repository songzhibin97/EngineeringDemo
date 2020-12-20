/******
** @创建时间 : 2020/12/17 12:52
** @作者 : SongZhiBin
******/
package main

import (
	"Songzhibin/EngineeringDemo/internal/pkg"
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	app, clear, err := pkg.NewApp()
	if err != nil {
		zap.L().Info(err.Error())
		return
	}
	defer clear()
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		zap.L().Info("grpc Start")
		go func() {
			<-ctx.Done()
			pkg.Service.Stop()
		}()
		return pkg.Service.Serve(app.Listener)

	})
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
		sig := make(chan os.Signal, 1)
		// 注册信号
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				zap.L().Info("signal ctx done")
				return nil
			case <-sig:
				zap.L().Info("signal signal done")
				return errors.New("signal signal done")
			}
		}
	})
	zap.L().Info(g.Wait().Error())
}
