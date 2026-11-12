package obp

// The contract between this repository and OBP-API.
//
// obp-api-contract/contract.json lists every OBP-API call the tokenizer makes,
// with the response fields the code reads. obp-api-contract/snapshot.json is what an OBP-API documents for those
// calls, written by obp-api-contract/snapshot.py from its resource docs. These
// tests need no server: they check that the list matches the code, and that
// every call reaches an endpoint OBP-API documents, with every field the code
// reads. Retake the snapshot from a new OBP-API build to learn whether that
// build breaks this repository.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
)

type contractCall struct {
	Verb                 string   `json:"verb"`
	Path                 string   `json:"path"`
	Source               string   `json:"source"`
	Needles              []string `json:"needles"`
	Reads                []string `json:"reads"`
	UndocumentedFields   []string `json:"undocumented_fields"`
	UndocumentedEndpoint string   `json:"undocumented_endpoint"`
}

type snapshotEndpoint struct {
	Verb        string   `json:"verb"`
	Path        string   `json:"path"`
	Missing     bool     `json:"missing"`
	OperationID string   `json:"operation_id"`
	Ambiguous   []string `json:"ambiguous"`
	Fields      []string `json:"fields"`
}

func (c contractCall) key() string     { return c.Verb + " " + c.Path }
func (e snapshotEndpoint) key() string { return e.Verb + " " + e.Path }

func loadContract(t *testing.T, name string, out any) {
	t.Helper()
	text, err := os.ReadFile(filepath.Join(repoRoot(t), "obp-api-contract", name))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(text, out); err != nil {
		t.Fatalf("%s: %v", name, err)
	}
}

func calls(t *testing.T) []contractCall {
	var c struct {
		Calls []contractCall `json:"calls"`
	}
	loadContract(t, "contract.json", &c)
	return c.Calls
}

func snapshot(t *testing.T) map[string]snapshotEndpoint {
	var s struct {
		Endpoints []snapshotEndpoint `json:"endpoints"`
	}
	loadContract(t, "snapshot.json", &s)
	byKey := make(map[string]snapshotEndpoint, len(s.Endpoints))
	for _, e := range s.Endpoints {
		byKey[e.key()] = e
	}
	return byKey
}

// undocumentedReads returns the reads that fields does not document. A read
// written `a|b` is satisfied by either field.
func undocumentedReads(reads []string, fields map[string]bool) []string {
	var out []string
	for _, read := range reads {
		ok := false
		for _, f := range strings.Split(read, "|") {
			ok = ok || fields[f]
		}
		if !ok {
			out = append(out, read)
		}
	}
	return out
}

// obpPathLine is a line of Go source that builds an OBP-API path.
type obpPathLine struct {
	file string
	n    int
	line string
}

// obpPathLines returns the lines of non-test Go source that build an OBP-API
// path (an /obp/ path or DirectLogin), skipping comments.
func obpPathLines(t *testing.T) []obpPathLine {
	t.Helper()
	root := repoRoot(t)
	var out []obpPathLine
	for _, rel := range sourceFiles(t, root) {
		if !strings.HasSuffix(rel, ".go") || strings.HasSuffix(rel, "_test.go") {
			continue
		}
		text, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			t.Fatal(err)
		}
		for i, line := range strings.Split(string(text), "\n") {
			code := strings.TrimSpace(line)
			if strings.HasPrefix(code, "//") {
				continue
			}
			if strings.Contains(code, `"`) && (strings.Contains(code, "/obp/") || strings.Contains(code, "/my/logins/")) {
				out = append(out, obpPathLine{rel, i + 1, line})
			}
		}
	}
	return out
}

func TestEveryCallInTheContractIsInItsSourceFile(t *testing.T) {
	// Tests that each listed call still exists: every needle of an entry occurs in the entry's source file.
	root := repoRoot(t)
	for _, c := range calls(t) {
		text, err := os.ReadFile(filepath.Join(root, c.Source))
		if err != nil {
			t.Errorf("%s: %v", c.key(), err)
			continue
		}
		if len(c.Needles) == 0 {
			t.Errorf("%s: no needles", c.key())
		}
		for _, needle := range c.Needles {
			if !strings.Contains(string(text), needle) {
				t.Errorf("%s: %s no longer contains %s", c.key(), c.Source, needle)
			}
		}
	}
}

func TestEveryOBPAPICallInTheSourceIsInTheContract(t *testing.T) {
	// Tests coverage: a line that builds an OBP-API path must contain a needle of a call listed for that file,
	// so a new call cannot reach OBP-API without its fields being recorded.
	all := calls(t)
	lines := obpPathLines(t)
	if len(lines) == 0 {
		t.Fatal("found no line that builds an OBP-API path; the scan is broken")
	}
	for _, l := range lines {
		listed := false
		for _, c := range all {
			if c.Source != l.file {
				continue
			}
			for _, n := range c.Needles {
				listed = listed || strings.Contains(l.line, strings.TrimSpace(n))
			}
		}
		if !listed {
			t.Errorf("not in obp-api-contract/contract.json: %s:%d %s", l.file, l.n, strings.TrimSpace(l.line))
		}
	}
}

func TestTheSnapshotWasTakenForThisContract(t *testing.T) {
	// Tests that snapshot.json holds exactly the calls of contract.json; otherwise rerun snapshot.py.
	var listed, snapped []string
	for _, c := range calls(t) {
		listed = append(listed, c.key())
	}
	for k := range snapshot(t) {
		snapped = append(snapped, k)
	}
	sort.Strings(listed)
	sort.Strings(snapped)
	if !reflect.DeepEqual(listed, snapped) {
		t.Errorf("contract.json and snapshot.json differ; rerun obp-api-contract/snapshot.py\ncontract: %q\nsnapshot: %q", listed, snapped)
	}
}

func TestEveryCallReachesAnEndpointOBPAPIDocuments(t *testing.T) {
	// Tests that each call's version and path name a documented endpoint, unless the entry records why not
	// (undocumented_endpoint). A recorded gap that OBP-API has since documented fails too, so the note goes.
	snap := snapshot(t)
	for _, c := range calls(t) {
		e, ok := snap[c.key()]
		missing := !ok || e.Missing
		switch {
		case missing && c.UndocumentedEndpoint == "":
			t.Errorf("%s: no documented endpoint", c.key())
		case !missing && c.UndocumentedEndpoint != "":
			t.Errorf("%s: now documented as %s; drop undocumented_endpoint", c.key(), e.OperationID)
		case missing:
			t.Logf("known gap: %s: %s", c.key(), c.UndocumentedEndpoint)
		}
		if len(e.Ambiguous) > 0 {
			t.Errorf("%s: matches several endpoints %q", c.key(), e.Ambiguous)
		}
	}
}

func TestEveryFieldTheCodeReadsIsDocumented(t *testing.T) {
	// Tests that each field the code reads is in the documented response of its endpoint, unless the entry
	// lists it in undocumented_fields; a listed field that OBP-API now documents fails, so the list shrinks.
	snap := snapshot(t)
	for _, c := range calls(t) {
		e, ok := snap[c.key()]
		if !ok || e.Missing {
			continue
		}
		fields := map[string]bool{}
		for _, f := range e.Fields {
			fields[f] = true
		}
		known := map[string]bool{}
		for _, f := range c.UndocumentedFields {
			known[f] = true
			if fields[f] {
				t.Errorf("%s: %s is now documented; drop it from undocumented_fields", c.key(), f)
			}
		}
		for _, read := range undocumentedReads(c.Reads, fields) {
			if known[read] {
				t.Logf("known gap: %s: %s is not documented", c.key(), read)
			} else {
				t.Errorf("%s (%s): reads %s, not documented", c.key(), e.OperationID, read)
			}
		}
	}
}

func TestARenamedOrRemovedFieldIsReported(t *testing.T) {
	// Tests the check itself: with token renamed in the documented fields, the DirectLogin read is
	// reported, and a read with an alternative (`a|b`) passes on either field.
	renamed := map[string]bool{"access_token": true}
	if got := undocumentedReads([]string{"token"}, renamed); !reflect.DeepEqual(got, []string{"token"}) {
		t.Errorf("undocumentedReads = %q, want [token]", got)
	}
	read := []string{"ENTITY.ENTITY_id|ENTITY_id"}
	if got := undocumentedReads(read, map[string]bool{"ENTITY_id": true}); len(got) != 0 {
		t.Errorf("undocumentedReads = %q, want none", got)
	}
	if got := undocumentedReads(read, map[string]bool{}); !reflect.DeepEqual(got, read) {
		t.Errorf("undocumentedReads = %q, want %q", got, read)
	}
}
