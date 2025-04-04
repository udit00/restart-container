package models

type DockerHubWebhook struct {
	CallbackURL string `json:"callback_url"`
	PushData    struct {
		Tag      string `json:"tag"`
		PushedAt int64  `json:"pushed_at"`
		Images   []struct {
			Architecture string `json:"architecture"`
			Features     string `json:"features"`
			Variant      string `json:"variant"`
			Digest       string `json:"digest"`
		} `json:"images"`
		Pusher string `json:"pusher"`
	} `json:"push_data"`
	Repository struct {
		RepoName     string `json:"repo_name"`
		Name         string `json:"name"`
		Namespace    string `json:"namespace"`
		DateCreated  int64  `json:"date_created"`
		DateModified int64  `json:"date_modified"`
		Status       string `json:"status"`
		Description  string `json:"description"`
		IsPrivate    bool   `json:"is_private"`
		IsAutomated  bool   `json:"is_automated"`
		CanEdit      bool   `json:"can_edit"`
		StarCount    int    `json:"star_count"`
		PullCount    int    `json:"pull_count"`
		LastUpdated  string `json:"last_updated"`
	} `json:"repository"`
	// Add other fields as needed based on your webhook configuration
}
