package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
    "testing"
)

// TODO Move duplicate code into its own setup function.
// TODO Re-compile the go script under test before all tests run.

func TestInvalidCharacterEncoding(t *testing.T) {
    envVars := getEnvVarsAsMap()

    scriptUnderTest := fmt.Sprintf("%s/bin/golang-weather-app",
        envVars["GOPATH"])

    invalidString := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    cmd := exec.Command(scriptUnderTest, "foo", invalidString)
    stdout, stderr := cmd.CombinedOutput()

    errMsg := getErrorTextFromOutput(stdout)

    expectedError := fmt.Sprintf("%s is not a valid utf8 string\n",
        invalidString)

    if string(errMsg) != expectedError {
        t.Error(fmt.Sprintf("Message '%s' doesn't match expected error\n",
            errMsg))
    }

    if stderr == nil {
        t.Error("Script should fail with exit code 1")
    }
}

func TestInvalidNumberOfArgs(t *testing.T) {
    envVars := getEnvVarsAsMap()

    scriptUnderTest := fmt.Sprintf("%s/bin/golang-weather-app",
        envVars["GOPATH"])

    cmd := exec.Command(scriptUnderTest, "foo")
    stdout, stderr := cmd.CombinedOutput()

    errMsg := getErrorTextFromOutput(stdout)

    if string(errMsg) != "Invalid number of positional arguments: 2 required\n" {
        t.Error(fmt.Sprintf("Message '%s' doesn't match expected error\n",
            errMsg))
    }

    if stderr == nil {
        t.Error("Script should fail with exit code 1")
    }
}

func TestInvalidCountryCode(t *testing.T) {
    envVars := getEnvVarsAsMap()

    scriptUnderTest := fmt.Sprintf("%s/bin/golang-weather-app",
        envVars["GOPATH"])

    cmd := exec.Command(scriptUnderTest, "foo", "bar")
    stdout, stderr := cmd.CombinedOutput()

    errMsg := getErrorTextFromOutput(stdout)

    if string(errMsg) != "Country code 'bar' is invalid\n" {
        t.Error(fmt.Sprintf("Message '%s' doesn't match expected error\n",
            errMsg))
    }

    if stderr == nil {
        t.Error("Script should fail with exit code 1")
    }
}


// The error message logged to stdout has the form:
// [yyyy/mm/dd] [hh:mm:ss] [message]
// We just want to check the message as the timestamp is variable.
func getErrorTextFromOutput(stdout []byte) (errMsg string) {
    errMsg = string(stdout)[20:]

    return errMsg
}

// Return the key/values listed in ENV as a map.
func getEnvVarsAsMap() (envVars map[string]string) {
    envVars = map[string]string{}

    for _, val := range os.Environ() {
        keyValSlice := strings.Split(val,"=")

        envVars[keyValSlice[0]] = keyValSlice[1]
    }

    return envVars
}
