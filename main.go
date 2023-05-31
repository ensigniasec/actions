package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	dsselib "github.com/secure-systems-lab/go-securesystemslib/dsse"
	gha "github.com/sethvargo/go-githubactions"
)

var Version = "dev"
var Commit = "unknown"
var CommitDate = "unknown"
var TreeState = "unknown"

func main() {
	ctx := context.Background()
	if err := realMain(ctx); err != nil {
		gha.Fatalf("error: %s", err)
	}
}

func EnvelopeFromBytes(payload []byte) (env *dsselib.Envelope, err error) {
	env = &dsselib.Envelope{}
	err = json.Unmarshal(payload, env)
	return
}

func realMain(ctx context.Context) error {
	gha.Infof("Ensignia Action Version: %s", Version)

	apiKey := gha.GetInput("api-key")
	if apiKey == "" {
		gha.Fatalf("api-key input param is required")
	}

	provenanceName := gha.GetInput("provenance-name")
	if provenanceName == "" {
		gha.Fatalf("provenance-name input param is required (e.g. 'needs.build.outputs.provenance-name')")
	}

	prov, err := os.ReadFile(provenanceName)
	if err != nil {
		return err
	}

	env, err := EnvelopeFromBytes(prov)
	if err != nil {
		return err
	}

	if env.PayloadType != "application/vnd.in-toto+json" {
		return fmt.Errorf("invalid payload type: %s", env.PayloadType)
	}

	gha.Infof("Payload", env.Payload)

	return nil
}

func setOutput(key, value string) {
	fmt.Printf("%s=%s >> $GITHUB_OUTPUT\n", key, value)
}
