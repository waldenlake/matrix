package protoc

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const (
	getProtocGenGo = "go get -u github.com/golang/protobuf/protoc-gen-go"
	protocCommand  = "protoc -I%s -I%s --go_out=plugins=grpc,paths=source_relative:. --grpc-gateway_out=logtostderr=true,paths=source_relative:."
)

var ProtocCmd = &cobra.Command{
	Use:   "protoc",
	Short: "use protoc-gen-go generate protobuf code",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	protoFiles, err := getProtoFiles(args)
	if err != nil {
		log.Fatal(err)
	}

	if err = installProtocGenGoPlugin(); err != nil {
		log.Fatal(err)
	}

	if err = generate(protoFiles); err != nil {
		log.Fatal(err)
	}

	log.Printf("generate %s success.", strings.Join(protoFiles, " "))
}

func generate(protoFiles []string) error {
	pwd, _ := os.Getwd()
	extPath, err := getExtPath()
	if err != nil {
		return err
	}

	protocCommand := fmt.Sprintf(protocCommand, pwd, extPath)
	log.Println("protoc command:", protocCommand, strings.Join(protoFiles, " "))

	args := strings.Split(protocCommand, " ")
	args = append(args, protoFiles...)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = pwd
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getExtPath() (string, error) {
	gopath := strings.Split(os.Getenv("GOPATH"), string(filepath.ListSeparator))
	baseMod := path.Join(gopath[0], "pkg/mod/github.com/grpc-ecosystem")
	files, err := ioutil.ReadDir(baseMod)
	if err != nil {
		return "", err
	}
	var protoPath string
	for i := len(files) - 1; i >= 0; i-- {
		if strings.HasPrefix(files[i].Name(), "grpc-gateway@") {
			protoPath = path.Join(baseMod, files[i].Name(), "third_party/googleapis")
			break
		}
	}
	log.Println("ext path=", protoPath)
	return protoPath, nil
}

func installProtocGenGoPlugin() error {
	if path, err := exec.LookPath("protoc-gen-go"); err != nil {
		if err := installPlugin(); err != nil {
			return err
		}
		log.Println("protoc-gen-go path: ", path)
	}
	return nil
}

func installPlugin() error {
	args := strings.Split(getProtocGenGo, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(getProtocGenGo)
	return cmd.Run()
}

func getProtoFiles(args []string) ([]string, error) {
	if len(args) == 0 {
		protoFiles, err := filepath.Glob("*.proto")
		if err != nil {
			return nil, err
		}
		return protoFiles, nil
	}
	return args, nil
}
