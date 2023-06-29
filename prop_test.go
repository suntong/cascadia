package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	dirTest = "test/"
	cmdTest = "../cascadia"
	extRef  = ".ref" // extension for reference file
	extGot  = ".got" // extension for generated file
)

// testIt runs @cmdTest with @argv and compares the generated
// output for @name with the corresponding @extRef
func testIt(t *testing.T, name string, argv ...string) {
	var (
		diffOut         bytes.Buffer
		generatedOutput = name + extGot
		cmd             = exec.Command(cmdTest, argv...)
	)

	t.Logf("Testing %s:\n\t%s %s\n", name, cmdTest, strings.Join(argv, " "))

	// open the out file for writing
	outfile, err := os.Create(generatedOutput)
	if err != nil {
		t.Errorf("write error [%s: %s] %s.", name, argv, err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		t.Errorf("start error [%s: %s] %s.", name, argv, err)
	}
	err = cmd.Wait()
	if err != nil {
		t.Errorf("exit error [%s: %s] %s.", name, argv, err)
	}

	cmd = exec.Command("diff", "-U1", name+extRef, generatedOutput)
	cmd.Stdout = &diffOut

	err = cmd.Start()
	if err != nil {
		t.Errorf("start error %s [%s: %s]", err, name, argv)
	}
	err = cmd.Wait()
	if err != nil {
		t.Errorf("cmp error %s [%s: %s]\n%s", err, name, argv, diffOut.String())
	}
	//os.Remove(generatedOutput)
}

type testCase struct {
	name string
	args []string
}

func testCases(t *testing.T, name string, testData []testCase) {
	t.Logf("\n\n== Testing %s\n\n", name)
	os.Chdir(dirTest)

	for _, tc := range testData {
		testIt(t, tc.name, tc.args...)
	}

}
