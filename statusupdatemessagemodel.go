package messagemodels

type JobStatusMessage struct {
	Listener     string
	JobID        string
	RenditionID  string
	InputKey     string
	InputBucket  string
	OutputKey    string
	OutputBucket string
	Status       string
}
