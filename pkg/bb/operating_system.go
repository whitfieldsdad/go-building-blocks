package bb

import (
	"runtime"

	"github.com/elastic/go-sysinfo"
)

type OperatingSystem struct {
	Type     string `json:"type"`               // OS Type (one of linux, macos, unix, windows).
	Arch     string `json:"arch"`               // OS architecture (e.g. x86_64).
	Family   string `json:"family"`             // OS Family (e.g. redhat, debian, freebsd, windows).
	Platform string `json:"platform"`           // OS platform (e.g. centos, ubuntu, windows).
	Name     string `json:"name"`               // OS Name (e.g. Mac OS X, CentOS).
	Version  string `json:"version"`            // OS version (e.g. 10.12.6).
	Major    *int   `json:"major,omitempty"`    // Major release version.
	Minor    *int   `json:"minor,omitempty"`    // Minor release version.
	Patch    *int   `json:"patch,omitempty"`    // Patch release version.
	Build    string `json:"build,omitempty"`    // Build (e.g. 16G1114).
	Codename string `json:"codename,omitempty"` // OS codename (e.g. jessie).
}

func GetOperatingSystem() *OperatingSystem {
	o := &OperatingSystem{
		Type: runtime.GOOS,
		Arch: runtime.GOARCH,
	}
	host, err := sysinfo.Host()
	if err == nil {
		info := host.Info()

		o.Family = info.OS.Family
		o.Platform = info.OS.Platform
		o.Name = info.OS.Name
		o.Version = info.OS.Version
		o.Build = info.OS.Build
		o.Codename = info.OS.Codename

		major := int(info.OS.Major)
		minor := int(info.OS.Minor)
		patch := int(info.OS.Patch)
		o.Major = &major
		o.Minor = &minor
		o.Patch = &patch
	}
	return o
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}
