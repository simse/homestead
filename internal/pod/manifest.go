package pod

// Manifest represents a Docker image manifest.
type Manifest struct {
	Ref        string `json:"Ref"`
	Descriptor struct {
		MediaType string `json:"mediaType"`
		Size      int64  `json:"size"`
		Digest    string `json:"digest"`
		Platform  struct {
			Architecture string `json:"architecture"`
			OS           string `json:"os"`
		} `json:"platform"`
	}
	SchemaV2Manifest struct {
		Layers []struct {
			Digest string `json:"digest"`
			Size   int    `json:"size"`
		} `json:"layers"`
	} `json:"SchemaV2Manifest"`
}

// Size returns the size of an image
func (m *Manifest) Size() int {
	totalSize := 0
	for _, layer := range m.SchemaV2Manifest.Layers {
		totalSize += layer.Size
	}
	return totalSize
}

// Architechture returns the architecture of an image
func (m *Manifest) Architecture() string {
	return m.Descriptor.Platform.Architecture
}
