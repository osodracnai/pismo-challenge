package version

// Default build-time variable.
// These values are overridden via ldflags
var (
	Version   = "unknown-version"
	GitCommit = "unknown-commit"
	BuildTime = "unknown-buildtime"
	OS        = "unknown-os"
	Arch      = "unknown-arch"
)
