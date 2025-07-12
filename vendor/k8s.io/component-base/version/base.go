package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "31"
	gitVersion   = "v1.31.5-k3s1"
	gitCommit    = "e4005eb7bc4989bc63487e061af43698e771cef1"
	gitTreeState = "clean"
	buildDate = "2025-01-16T00:58:44Z"
)
