package v2

import (
	"fmt"
	"github.com/goharbor/harbor/src/pkg/distribution"
	"io"
)

func (a *adapter) ManifestExist(repository, reference string) (bool, *distribution.Descriptor, error) {
	if a.Client.RepoPrefix != "" {
		repository = fmt.Sprintf("%s/%s", a.Client.RepoPrefix, repository)
	}

	return a.Adapter.ManifestExist(repository, reference)
}

func (a *adapter) PullManifest(repository, reference string, accepttedMediaTypes ...string) (manifest distribution.Manifest, digest string, err error) {
	if a.Client.RepoPrefix != "" {
		repository = fmt.Sprintf("%s/%s", a.Client.RepoPrefix, repository)
	}

	return a.Adapter.PullManifest(repository, reference, accepttedMediaTypes...)
}

func (a *adapter) BlobExist(repository, digest string) (exist bool, err error) {
	if a.Client.RepoPrefix != "" {
		repository = fmt.Sprintf("%s/%s", a.Client.RepoPrefix, repository)
	}

	return a.Adapter.BlobExist(repository, digest)
}

func (a *adapter) PullBlob(repository, digest string) (size int64, blob io.ReadCloser, err error) {
	if a.Client.RepoPrefix != "" {
		repository = fmt.Sprintf("%s/%s", a.Client.RepoPrefix, repository)
	}

	return a.Adapter.PullBlob(repository, digest)
}

func (a *adapter) PullBlobChunk(repository, digest string, blobSize, start, end int64) (size int64, blob io.ReadCloser, err error) {
	if a.Client.RepoPrefix != "" {
		repository = fmt.Sprintf("%s/%s", a.Client.RepoPrefix, repository)
	}

	return a.Adapter.PullBlobChunk(repository, digest, blobSize, start, end)
}
