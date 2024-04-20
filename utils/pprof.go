package utils

import (
	"fmt"
	"itachi/cairo/config"
	"net/http"
	_ "net/http/pprof"
)

func StartUpPprof(cfg *config.Config) {
	if cfg.EnablePprof {
		go func() {
			fmt.Println(http.ListenAndServe(cfg.PprofAddr, nil))
		}()
	}
}
