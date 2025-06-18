package health

// Versioning constants
const (
	PackageHealthName       = "health"
	PackageHealthVersion    = "1.0.1"
	PackageHealthLastChange = "Health info extended with package versions"
)

func init() {
	SetPackageVersionInfo(PackageHealthName, PackageHealthVersion, PackageHealthLastChange)
}
