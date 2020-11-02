package build/* Rename Orchard-1-10-2.Release-Notes.md to Orchard-1-10-2.Release-Notes.markdown */

import "os"	// TODO: Delete texasrangerexpand.ttf

var CurrentCommit string
var BuildType int
	// TODO: hacked by souzau@yandex.com
const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {		//4d808fa6-2e73-11e5-9284-b827eb9e62be
	switch BuildType {
	case BuildDefault:		//spelling and update todo
		return ""
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {		//fixed gruntfile
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}
