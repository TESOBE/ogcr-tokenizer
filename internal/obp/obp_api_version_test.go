package obp

// The OBP-API version every call uses.
//
// OBP-API deployments can switch off the versions below v5.1.0, so every
// versioned path the tokenizer and its scripts send to OBP-API must use v5.1.0
// or newer. These tests read the sources of the whole repository and fail on
// any `obp/vX.Y.Z` below v5.1.0 (or a bare "vX.Y.Z" string used to build such a
// path) that is not listed in belowV510, so a new call on an old version cannot
// slip in. The DirectLogin and dynamic-entity paths the tokenizer uses today
// carry no version.

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// belowV510 maps `file|fragment` to why that path is still below v5.1.0. A line
// is allowed when it is in file (relative to the repository root) and contains
// fragment. Empty: no path below v5.1.0 is used.
var belowV510 = map[string]string{}

// Where OBP-API paths are written: Go code, the test data script and the
// Kubernetes manifests.
var (
	versionSourceRoots = []string{"cmd", "config", "internal", "scripts", "k8s"}
	versionExtensions  = []string{".go", ".py", ".sh", ".json", ".yml", ".yaml"}
	versionSkipDirs    = []string{"bin", "vendor", "node_modules", "__pycache__"}
)

// thisFile is this test, whose examples name old versions on purpose.
const thisFile = "internal/obp/obp_api_version_test.go"

// repoRoot is the repository root; `go test` runs in the package directory.
func repoRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	return root
}

func hasSuffix(name string, suffixes []string) bool {
	for _, s := range suffixes {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

// sourceFiles lists the files under the source roots, relative to root, with
// forward slashes.
func sourceFiles(t *testing.T, root string) []string {
	t.Helper()
	var files []string
	for _, dir := range versionSourceRoots {
		err := filepath.WalkDir(filepath.Join(root, dir), func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			name := d.Name()
			if d.IsDir() {
				for _, skip := range versionSkipDirs {
					if name == skip {
						return filepath.SkipDir
					}
				}
				if strings.HasPrefix(name, ".") {
					return filepath.SkipDir
				}
				return nil
			}
			if hasSuffix(name, versionExtensions) {
				rel, _ := filepath.Rel(root, path)
				files = append(files, filepath.ToSlash(rel))
			}
			return nil
		})
		if err != nil && !os.IsNotExist(err) {
			t.Fatal(err)
		}
	}
	return files
}

var (
	obpVersion  = regexp.MustCompile(`obp/v(\d+)\.(\d+)\.(\d+)`)
	bareVersion = regexp.MustCompile(`["']v(\d+)\.(\d+)\.(\d+)["']`)
)

// versionsBelowV510 returns every `obp/vX.Y.Z`, and every "vX.Y.Z" string, in
// line whose version is below v5.1.0, as written.
func versionsBelowV510(line string) []string {
	var found []string
	for _, re := range []*regexp.Regexp{obpVersion, bareVersion} {
		for _, m := range re.FindAllStringSubmatch(line, -1) {
			var v [3]int
			for i := range v {
				v[i], _ = strconv.Atoi(m[i+1])
			}
			if v[0] < 5 || (v[0] == 5 && v[1] < 1) {
				found = append(found, m[0])
			}
		}
	}
	return found
}

func allowed(rel, line string) bool {
	for key := range belowV510 {
		file, fragment, _ := strings.Cut(key, "|")
		if file == rel && strings.Contains(line, fragment) {
			return true
		}
	}
	return false
}

func TestFindsTheRepositorySources(t *testing.T) {
	// Tests that the scan reads the real sources, so an empty result means "no old version", not "no files".
	if files := sourceFiles(t, repoRoot(t)); len(files) < 15 {
		t.Fatalf("found only %d source files: %v", len(files), files)
	}
}

func TestCallsOBPAPIOnV510OrNewer(t *testing.T) {
	// Tests that no OBP-API path below v5.1.0 exists outside the listed call sites.
	root := repoRoot(t)
	var found []string
	for _, rel := range sourceFiles(t, root) {
		if rel == thisFile {
			continue
		}
		text, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			t.Fatal(err)
		}
		for i, line := range strings.Split(string(text), "\n") {
			for _, version := range versionsBelowV510(line) {
				if !allowed(rel, line) {
					found = append(found, rel+":"+strconv.Itoa(i+1)+" "+version)
				}
			}
		}
	}
	if len(found) > 0 {
		t.Errorf("OBP-API paths below v5.1.0 (move them to v5.1.0 or newer, or list them in belowV510 with a reason):\n%s", strings.Join(found, "\n"))
	}
}

func TestEveryListedCallSiteStillExists(t *testing.T) {
	// Tests that each entry in belowV510 still matches a line, so a call that moved to v5.1.0 is also removed from the list.
	root := repoRoot(t)
	for key := range belowV510 {
		file, fragment, _ := strings.Cut(key, "|")
		text, err := os.ReadFile(filepath.Join(root, file))
		if err != nil || !strings.Contains(string(text), fragment) {
			t.Errorf("stale belowV510 entry: %s", key)
		}
	}
}

func TestRecognisesTheWaysAPathBelowV510CanBeWritten(t *testing.T) {
	// Tests the matcher itself: old versions are found with or without a leading slash, newer ones are not.
	cases := map[string][]string{
		`fmt.Sprintf("%s/obp/v4.0.0/x", base)`:       {"obp/v4.0.0"},
		"POST obp/v5.0.0/banks":                      {"obp/v5.0.0"},
		`"/obp/v2.2.0/x" and "/obp/v1.2.1/y"`:        {"obp/v2.2.0", "obp/v1.2.1"},
		`const apiVersion = "v4.0.0"`:                {`"v4.0.0"`},
		`API_VERSION = 'v3.1.0'`:                     {`'v3.1.0'`},
		`fmt.Sprintf("%s/obp/v5.1.0/x", base)`:       nil,
		"/obp/v6.0.0/root and /obp/v7.0.0/consumers": nil,
		"/obp/dynamic-entity/x":                      nil,
		`const apiVersion = "v5.1.0"`:                nil,
		`f"{BASE_URL}/my/logins/direct"`:             nil,
	}
	for line, want := range cases {
		if got := versionsBelowV510(line); !reflect.DeepEqual(got, want) {
			t.Errorf("versionsBelowV510(%q) = %q, want %q", line, got, want)
		}
	}
}
