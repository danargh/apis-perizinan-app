package internal

import (
	"fmt"
	"net/http"
)

// goroutine terdapat sinkronasi agar aplikasi menunggu tugas latar belakang selesai
func (app *Application) backgroundTask(r *http.Request, fn func() error) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		// menangani panik yang mungkin terjadi di goroutine
		defer func() {
			err := recover()
			if err != nil {
				app.reportServerError(r, fmt.Errorf("%s", err))
			}
		}()

		// eksekusi fungsi yang dimasukan sebagai parameter
		err := fn()
		if err != nil {
			app.reportServerError(r, err)
		}
	}()
}
