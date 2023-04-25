package v1

type ArtifactManifest struct{}

type Artifact struct {
	MediaType    string
	Subject      *Descriptor
	Blobs        []Descriptor
	ArtifactType string
	Annotations  map[string]string
}

const MediaTypeArtifactManifest = ""
