/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/jonboulle/clockwork"
	"gopkg.in/h2non/gock.v1"
)

// TESTAPICLIENT is the api client to use for tests
var TESTAPICLIENT *datadog.APIClient

// TESTAUTH is the authentication context to use with each test API call
var TESTAUTH context.Context

// TESTCLOCK is the time module to use in tests
var TESTCLOCK clockwork.FakeClock

func setClock(t *testing.T) {
	os.MkdirAll("cassettes", 0755)
	f, err := os.Create(fmt.Sprintf("cassettes/%s.freeze", t.Name()))
	if err != nil {
		t.Fatalf("Could not set clock: %v", err)
	}
	defer f.Close()
	now := clockwork.NewRealClock().Now()
	f.WriteString(now.Format(time.RFC3339Nano))
	TESTCLOCK = clockwork.NewFakeClockAt(now)
}

func restoreClock(t *testing.T) {
	data, err := ioutil.ReadFile(fmt.Sprintf("cassettes/%s.freeze", t.Name()))
	if err != nil {
		t.Fatalf("Could not load clock: %v", err)
	}
	now, err := time.Parse(time.RFC3339Nano, string(data))
	if err != nil {
		t.Fatalf("Could not parse clock date: %v", err)
	}
	TESTCLOCK = clockwork.NewFakeClockAt(now)
}

func removeURLSecrets(u *url.URL) *url.URL {
	query := u.Query()
	query.Del("api_key")
	query.Del("application_key")
	u.RawQuery = query.Encode()
	return u
}

func setupTest(t *testing.T) func(t *testing.T) {
	// SETUP testing
	TESTAUTH = context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			},
		},
	)
	config := datadog.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"

	// Configure recorder
	var mode recorder.Mode
	if os.Getenv("RECORD") == "true" {
		mode = recorder.ModeRecording
	} else {
		mode = recorder.ModeReplaying
	}
	r, err := recorder.NewAsMode(fmt.Sprintf("cassettes/%s", t.Name()), mode, nil)
	if err != nil {
		log.Fatal(err)
	}

	r.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		return r.Method == i.Method && removeURLSecrets(r.URL).String() == i.URL
	})

	r.AddFilter(func(i *cassette.Interaction) error {
		u, err := url.Parse(i.URL)
		if err != nil {
			return err
		}
		i.URL = removeURLSecrets(u).String()
		i.Request.Headers.Del("Dd-Api-Key")
		i.Request.Headers.Del("Dd-Application-Key")
		return nil
	})

	if os.Getenv("RECORD") == "true" {
		setClock(t)
	} else {
		restoreClock(t)
	}

	config.HTTPClient = &http.Client{
		Transport: r, // Inject as transport!
	}

	TESTAPICLIENT = datadog.NewAPIClient(config)

	return func(t *testing.T) {
		r.Stop()
	}
}

func setupUnitTest(t *testing.T) func(t *testing.T) {
	// SETUP testing
	TESTAUTH = context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: "FAKE_KEY",
			},
			"appKeyAuth": {
				Key: "FAKE_KEY",
			},
		},
	)
	config := datadog.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"
	TESTAPICLIENT = datadog.NewAPIClient(config)
	return func(t *testing.T) {
		// TEARDOWN testing
	}
}

func setupGock(t *testing.T, fixtureFile string, method string, uriPath string) []byte {
	fixturePath, _ := filepath.Abs(fmt.Sprintf("fixtures/%s", fixtureFile))
	dat, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}
	x := gock.New("https://api.datadoghq.com/api/v1")
	switch strings.ToLower(method) {
	case "get":
		x.Get(uriPath)
	case "post":
		x.Post(uriPath)
	case "put":
		x.Put(uriPath)
	case "delete":
		x.Put(uriPath)
	}

	x.Reply(200).JSON(dat)
	return dat
}