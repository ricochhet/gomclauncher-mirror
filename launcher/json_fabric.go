package launcher

type Fabricjsonv2 fabricv2Version

type fabricv2Loader struct {
	Separator string `json:"separator"`
	Build     int    `json:"build"`
	Maven     string `json:"maven"`
	Version   string `json:"version"`
	Stable    bool   `json:"stable"`
}

type fabricv2Intermediary struct {
	Maven   string `json:"maven"`
	Version string `json:"version"`
	Stable  bool   `json:"stable"`
}

type fabricv2Library struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	MD5    string `json:"md5"`
	SHA1   string `json:"sha1"`
	SHA256 string `json:"sha256"`
	SHA512 string `json:"sha512"`
	Size   int    `json:"size"`
}

type fabricv2LauncherMeta struct {
	Version        int `json:"version"`
	MinJavaVersion int `json:"min_java_version"`
	Libraries      struct {
		Client      []fabricv2Library `json:"client"`
		Common      []fabricv2Library `json:"common"`
		Server      []fabricv2Library `json:"server"`
		Development []fabricv2Library `json:"development"`
	} `json:"libraries"`
	MainClass interface{} `json:"mainClass"`
}

type fabricv2Version struct {
	Loader       fabricv2Loader       `json:"loader"`
	Intermediary fabricv2Intermediary `json:"intermediary"`
	LauncherMeta fabricv2LauncherMeta `json:"launcherMeta"`
}
