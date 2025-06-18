package metrics

import "scheduler/internal/common/health"

// Versioning constants
const (
	PackageHealthName       = "util/metrics"
	PackageHealthVersion    = "1.1.4"
	PackageHealthLastChange = "false duplicate reports removed"
)

func init() {
	health.SetPackageVersionInfo(PackageHealthName, PackageHealthVersion, PackageHealthLastChange)
}
