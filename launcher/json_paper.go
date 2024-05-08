package launcher

type Paperjsonv2 paperv2Project

type paperv2Change struct {
	Commit  string `json:"commit"`
	Summary string `json:"summary"`
	Message string `json:"message"`
}

type paperv2Download struct {
	Name   string `json:"name"`
	SHA256 string `json:"sha256"`
}

type paperv2Build struct {
	Build     int             `json:"build"`
	Time      string          `json:"time"`
	Channel   string          `json:"channel"`
	Promoted  bool            `json:"promoted"`
	Changes   []paperv2Change `json:"changes"`
	Downloads struct {
		Application    paperv2Download `json:"application"`
		MojangMappings paperv2Download `json:"mojang-mappings"`
	} `json:"downloads"`
}

type paperv2Project struct {
	ProjectID   string         `json:"project_id"`
	ProjectName string         `json:"project_name"`
	Version     string         `json:"version"`
	Builds      []paperv2Build `json:"builds"`
}
