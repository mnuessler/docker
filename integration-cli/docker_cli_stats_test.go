package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestStatsNoArgs(t *testing.T) {
	runCmd := exec.Command(dockerBinary, "run", "-i", "-d", "busybox", "sleep", "20")
	out, _, err := runCommandWithOutput(runCmd)
	if err != nil {
		t.Fatalf("failed to start the container: %s, %v", out, err)
	}

	cleanedContainerID := stripTrailingCharacters(out)
	defer deleteContainer(cleanedContainerID)

	statsCmd := exec.Command(dockerBinary, "stats")
	out, _, err = runCommandWithOutput(statsCmd)
	if err != nil {
		t.Fatalf("failed to run stats: %s, %v", out, err)
	}

	if !strings.Contains(out, "FOO") {
		t.Fatalf("did not see FOO after top -o pid: %s", out)
	}

	logDone("stats - no arguments")
}
