// Package health contains common information about service
package health

import (
	"encoding/json"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
	// L "gitlab.gametechlabs.net/ppbet/backend-services/logger"
)

// VersionInfo contains version info
var VersionInfo = make(map[string]string)

// PkgVersionInfo folds package version and short info about last change
type PkgVersionInfo struct {
	Version    string `json:"version"`
	LastChange string `json:"lastChange,omitempty"`
}

// PackageInfo contains version info of packages used in the service
var PackageInfo = make(map[string]PkgVersionInfo)

// SetVersionInfo sets version info
func SetVersionInfo(values map[string]string) {
	getValue := func(v string) string {
		if v == "" || (v[0:2] == "${" && v[len(v)-1:] == "}") {
			return "n/a"
		}

		return v
	}
	for k, v := range values {
		values[k] = getValue(v)
	}

	VersionInfo = values

	if VersionInfo["serviceName"] == "" {
		VersionInfo["serviceName"] = strings.Trim(path.Base(os.Args[0]), "./")
	}
	if VersionInfo["startedAt"] == "" {
		VersionInfo["startedAt"] = time.Now().UTC().String()
	}
}

// SetPackageVersionInfo is used to register a package into version data
func SetPackageVersionInfo(PackageName, PackageVersion, LastChange string) {
	PackageInfo[PackageName] = PkgVersionInfo{Version: PackageVersion, LastChange: LastChange}
}

// Handler is a http handler
func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// L.L.Debug("health.Handler request", L.String("URI", r.RequestURI), L.String("Method", r.Method), L.String("Host", r.Host), L.String("ClientIP", httputil.GetIP(r)), L.String(L.XRequestIDKey, L.ReqIDFromContext(r.Context())))

		// avoid caching of a http response
		w.Header().Add("Expires", time.Now().UTC().Format(http.TimeFormat))
		w.Header().Add("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		w.Header().Add("Cache-Control", "max-age=0, no-cache, must-revalidate, proxy-revalidate")

		b, err := json.Marshal(VersionInfo)
		if err != nil {
			// L.L.Error("health.Handler error marshalling version info", L.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(PackageInfo) > 0 {
			p, err := json.Marshal(PackageInfo)
			if err != nil {
				// L.L.Error("health.Handler error marshalling package info", L.Error(err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			b = append(b[:len(b)-1], `,"packages":`...)
			b = append(b, p...)
			b = append(b, '}')
		}

		w.Header().Add("Content-type", "application/json")
		w.Write(b)
	}
}
