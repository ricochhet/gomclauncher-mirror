package launcher

import "time"

type Javaruntimesjson javaRuntimes

type javaRuntimes struct {
	Gamecore     JavaRuntimeTypes `json:"gamecore"`
	Linux        JavaRuntimeTypes `json:"linux"`
	LinuxI386    JavaRuntimeTypes `json:"linux-i386"`
	MacOs        JavaRuntimeTypes `json:"mac-os"`
	MacOsArm64   JavaRuntimeTypes `json:"mac-os-arm64"`
	WindowsArm64 JavaRuntimeTypes `json:"windows-arm64"`
	WindowsX64   JavaRuntimeTypes `json:"windows-x64"`
	WindowsX86   JavaRuntimeTypes `json:"windows-x86"`
}

type JavaRuntimeTypes struct {
	JavaRuntimeAlpha         []JavaRuntime `json:"java-runtime-alpha"`
	JavaRuntimeBeta          []JavaRuntime `json:"java-runtime-beta"`
	JavaRuntimeDelta         []JavaRuntime `json:"java-runtime-delta"`
	JavaRuntimeGamma         []JavaRuntime `json:"java-runtime-gamma"`
	JavaRuntimeGammaSnapshot []JavaRuntime `json:"java-runtime-gamma-snapshot"`
	JreLegacy                []JavaRuntime `json:"jre-legacy"`
	MinecraftJavaExe         []JavaRuntime `json:"minecraft-java-exe"`
}

type JavaRuntime struct {
	Availability javaDownloadAvailability `json:"availability"`
	Manifest     javaDownloadManifest     `json:"manifest"`
	Version      javaDownloadVersion      `json:"version"`
}

type javaDownloadAvailability struct {
	Group    int `json:"group"`
	Progress int `json:"progress"`
}

type javaDownloadVersion struct {
	Name     string    `json:"name"`
	Released time.Time `json:"released"`
}

type Javafilesjson struct {
	Files map[string]javaFiles `json:"files"`
}

type javaFiles struct {
	Downloads  javaDownloads `json:"downloads"`
	Executable bool          `json:"executable"`
	Type       string        `json:"type"`
}

type javaDownloads struct {
	Lzma javaDownloadManifest `json:"lzma"`
	Raw  javaDownloadManifest `json:"raw"`
}

type javaDownloadManifest struct {
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}
