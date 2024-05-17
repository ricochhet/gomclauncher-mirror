package launcher

type Quiltjsonv3 quiltv3Version

type quiltv3Loader struct {
	Separator string `json:"separator"`
	Build     int    `json:"build"`
	Maven     string `json:"maven"`
	Version   string `json:"version"`
}

type quiltv3Hashed struct {
	Maven   string `json:"maven"`
	Version string `json:"version"`
}

type quiltv3Intermediary struct {
	Maven   string `json:"maven"`
	Version string `json:"version"`
}

type quiltv3Library struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type quiltv3LauncherMeta struct {
	Version        int `json:"version"`
	MinJavaVersion int `json:"min_java_version"`
	Libraries      struct {
		Client      []quiltv3Library `json:"client"`
		Common      []quiltv3Library `json:"common"`
		Server      []quiltv3Library `json:"server"`
		Development []quiltv3Library `json:"development"`
	} `json:"libraries"`
	MainClass interface{} `json:"mainClass"`
}

type quiltv3Version struct {
	Loader       quiltv3Loader       `json:"loader"`
	Hashed       quiltv3Hashed       `json:"hashed"`
	Intermediary quiltv3Intermediary `json:"intermediary"`
	LauncherMeta quiltv3LauncherMeta `json:"launcherMeta"`
}
