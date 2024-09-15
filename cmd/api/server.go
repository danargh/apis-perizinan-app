package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultIdleTimeout    = time.Minute      // batas waktu server tetapi aktif tanpa koneksi 1 menit
	defaultReadTimeout    = 5 * time.Second  // batas waktu server membaca permintaan client adlah 5 detik
	defaultWriteTimeout   = 10 * time.Second // batas waktu server menulis respons untuk client adalah 10 detik
	defaultShutdownPeriod = 30 * time.Second // waktu maksimum server untuk shutdown adalah 30 detik apabila lebih maka force!
)

func (app *application) serveHTTP() error {
	// membuat instance dari sruct menggunakan pointer
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.httpPort),
		Handler:      app.routes(), // di app terdapat method routes() file routes.go
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelWarn),
		IdleTimeout:  defaultIdleTimeout,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	// Membuat channel untuk mengirim error yang mungkin terjadi selama proses shutdown
	shutdownErrorChan := make(chan error)

	go func() {
		// membuat channel khusus untuk menerima sinyal dari sistem operasi bertipe os.Signal. Hanya menampung 1 sinyal saja
		quitChan := make(chan os.Signal, 1)

		// Ketika pengguna menekan Ctrl+C (interrupt signal) atau sistem mengirim sinyal penghentian ke aplikasi, sinyal tersebut akan dikirim ke quitChan
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan

		// Membuat context baru dengan batas waktu (timeout) menggunakan context.WithTimeout. Context ini akan mengatur waktu maksimal untuk proses shutdown yang tertib.
		ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownPeriod)

		// proses shutdown berhasil atau tidak yang pasti context with timeout dibersihkan
		defer cancel()

		// menjalankan proses gracefull shutdown pada server
		shutdownErrorChan <- srv.Shutdown(ctx)
	}()

	app.logger.Info("starting server", slog.Group("server", "addr", srv.Addr))

	// memblokir eksekusi program sampai server dihentikan. Apabila error bukan serverClose maka return err
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// menunggu error dari channel
	err = <-shutdownErrorChan
	if err != nil {
		return err
	}

	app.logger.Info("stopped server", slog.Group("server", "addr", srv.Addr))

	// menunggu seluruh goroutine selesai
	app.wg.Wait()
	return nil
}
