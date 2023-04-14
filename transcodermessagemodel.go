package messagemodels

import "encoding/json"

type TranscoderMessageModel struct {
	JobID               string
	RenditionID         string
	InputMediaFileName  string
	OutputMediaFileName string
	TranscodeOption     string
}

// Unmarshal will attempt to load a json.RawMessage into this object.
func (tmm *TranscoderMessageModel) Unmarshal(raw json.RawMessage) error {
	if err := json.Unmarshal(raw, tmm); err != nil {
		return err
	}
	return nil
}
