// +build ignore

package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	version string = "v1"
	// deb & rpm does not support semver so have to handle their version a little differently
	linuxPackageVersion   string = "v1"
	linuxPackageIteration string = ""
	race                  bool

	workingDir string
	binaries   []string = []string{"alertmanager"}
)

const minGoVersion = 1.7

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	ensureGoPath()
	findVersion()

	log.Printf("Version: %s, Linux Version: %s, Package Iteration: %s\n", version, linuxPackageVersion, linuxPackageIteration)
	flag.BoolVar(&race, "race", race, "Use race detector")
	flag.Parse()

	if flag.NArg() == 0 {
		log.Println("Usage: go run build.go build")
		return
	}

	workingDir, _ = os.Getwd()

	for _, cmd := range flag.Args() {
		switch cmd {

		case "package":
			createLinuxPackages()
			sha1FilesInDist()

		case "pkg-rpm":
			createRpmPackages()
			sha1FilesInDist()
		case "latest":
			makeLatestDistCopies()
			sha1FilesInDist()

		case "sha1-dist":
			sha1FilesInDist()
		case "clean":
			clean()

		default:
			log.Fatalf("Unknown command %q", cmd)
		}
	}
}

func makeLatestDistCopies() {
	rpmIteration := "-1"
	if linuxPackageIteration != "" {
		rpmIteration = linuxPackageIteration
	}

	runError("cp", fmt.Sprintf("dist/alertmanager_%v-%v_amd64.deb", linuxPackageVersion, linuxPackageIteration), "dist/alertmanager_latest_amd64.deb")
	runError("cp", fmt.Sprintf("dist/alertmanager-%v-%v.x86_64.rpm", linuxPackageVersion, rpmIteration), "dist/alertmanager-latest-1.x86_64.rpm")
	runError("cp", fmt.Sprintf("dist/alertmanager-%v-%v.linux-x64.tar.gz", linuxPackageVersion, linuxPackageIteration), "dist/alertmanager-latest.linux-x64.tar.gz")
}

func findVersion() {
	var files = []string{"VERSION"}
	for _, file := range files {
		if fileExists(file) {
			version = readFile(file)
		}
	}
	linuxPackageVersion = version
	linuxPackageIteration = ""

	// handle pre version stuff (deb / rpm does not support semver)
	parts := strings.Split(version, "-")

	if len(parts) > 1 {
		linuxPackageVersion = parts[0]
		linuxPackageIteration = parts[1]
	}

	// add timestamp to iteration
	linuxPackageIteration = fmt.Sprintf("%d%s", time.Now().Unix(), linuxPackageIteration)
}

type linuxPackageOptions struct {
	packageType            string
	homeDir                string
	binPath                string
	serverBinPath          string
	cliBinPath             string
	configDir              string
	configFilePath         string
	etcDefaultPath         string
	etcDefaultFilePath     string
	initdScriptFilePath    string
	systemdServiceFilePath string
	logrotatePath          string

	postinstSrc      string
	initdScriptSrc   string
	defaultFileSrc   string
	systemdFileSrc   string
	logrotateFileSrc string
	configDirSrc     string
	depends          []string
}

func createRpmPackages() {
	createPackage(linuxPackageOptions{
		packageType:            "rpm",
		homeDir:                "/usr/share/alertmanager",
		binPath:                "/usr/sbin",
		configDir:              "/etc/alertmanager",
		etcDefaultPath:         "/etc/sysconfig",
		etcDefaultFilePath:     "/etc/sysconfig/alertmanager",
		initdScriptFilePath:    "/etc/init.d/alertmanager",
		systemdServiceFilePath: "/usr/lib/systemd/system/alertmanager.service",
		logrotatePath:          "/etc/logrotate.d",

		postinstSrc:      "packaging/rpm/control/postinst",
		initdScriptSrc:   "packaging/rpm/init.d/alertmanager",
		defaultFileSrc:   "packaging/rpm/sysconfig/alertmanager",
		systemdFileSrc:   "packaging/rpm/systemd/alertmanager.service",
		logrotateFileSrc: "packaging/rpm/log/alertmanager",
		configDirSrc:     "packaging/rpm/conf",

		depends: []string{"initscripts"},
	})
}

func createLinuxPackages() {
	createRpmPackages()
}

func createPackage(options linuxPackageOptions) {
	packageRoot, _ := ioutil.TempDir("", "alertmanager-linux-pack")
	fmt.Println(packageRoot)
	mkdir("dist")
	// create directories
	runPrint("mkdir", "-p", filepath.Join(packageRoot, options.homeDir))
	runPrint("mkdir", "-p", filepath.Join(packageRoot, options.configDir))
	runPrint("mkdir", "-p", filepath.Join(packageRoot, "/etc/init.d"))
	runPrint("mkdir", "-p", filepath.Join(packageRoot, options.etcDefaultPath))
	runPrint("mkdir", "-p", filepath.Join(packageRoot, "/usr/lib/systemd/system"))
	runPrint("mkdir", "-p", filepath.Join(packageRoot, options.logrotatePath))
	runPrint("mkdir", "-p", filepath.Join(packageRoot, "/usr/sbin"))

	// copy binary
	for _, binary := range binaries {
		runPrint("cp", "-p", filepath.Join(workingDir, ".build/linux-amd64/"+binary), filepath.Join(packageRoot, "/usr/sbin/"+binary))
	}

	// copy conf files
	runPrint("cp", "-r", options.configDirSrc, filepath.Join(packageRoot, options.homeDir))

	// copy init.d script
	runPrint("cp", "-p", options.initdScriptSrc, filepath.Join(packageRoot, options.initdScriptFilePath))
	// copy environment var file
	runPrint("cp", "-p", options.defaultFileSrc, filepath.Join(packageRoot, options.etcDefaultFilePath))
	// copy systemd file
	runPrint("cp", "-p", options.systemdFileSrc, filepath.Join(packageRoot, options.systemdServiceFilePath))
	// copy logrotate file
	runPrint("cp", "-p", options.logrotateFileSrc, filepath.Join(packageRoot, options.logrotatePath))
	// remove bin path
	runPrint("rm", "-rf", filepath.Join(packageRoot, options.homeDir, "bin"))
	println("Rinku ")
	args := []string{
		"-s", "dir",
		"--description", "alertmanager",
		"-C", packageRoot,
		"--vendor", "AlertManager",
		"--url", "https://github.ibm.com/cds-delivery/alertmanager.git",
		"--license", "IBM",
		"--maintainer", "rsamal@us.ibm.com",
		"--config-files", options.initdScriptFilePath,
		"--config-files", options.etcDefaultFilePath,
		"--config-files", options.systemdServiceFilePath,
		"--config-files", options.logrotatePath,
		"--after-install", options.postinstSrc,
		"--name", "alertmanager",
		"--version", linuxPackageVersion,
		"--rpm-os", "linux",
		"-p", "./dist",
	}

	// add dependenciesj
	for _, dep := range options.depends {
		args = append(args, "--depends", dep)
	}

	args = append(args, ".")

	fmt.Println("Creating package: ", options.packageType)
	runPrint("fpm", append([]string{"-t", options.packageType}, args...)...)
}

func ensureGoPath() {
	if os.Getenv("GOPATH") == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		gopath := filepath.Clean(filepath.Join(cwd, "../../../../"))
		log.Println("GOPATH is", gopath)
		os.Setenv("GOPATH", gopath)
	}
}

func rmr(paths ...string) {
	for _, path := range paths {
		log.Println("rm -r", path)
		os.RemoveAll(path)
	}
}

func mkdir(paths ...string) {
	for _, path := range paths {
		log.Println("mkdir -p", path)
		os.Mkdir(path, 0777)
	}
}

func clean() {
	rmr("dist")
	rmr("tmp")
}

func buildStamp() int64 {
	bs, err := runError("git", "show", "-s", "--format=%ct")
	if err != nil {
		return time.Now().Unix()
	}
	s, _ := strconv.ParseInt(string(bs), 10, 64)
	return s
}

func run(cmd string, args ...string) []byte {
	bs, err := runError(cmd, args...)
	if err != nil {
		log.Println(cmd, strings.Join(args, " "))
		log.Println(string(bs))
		log.Fatal(err)
	}
	return bytes.TrimSpace(bs)
}

func runError(cmd string, args ...string) ([]byte, error) {
	ecmd := exec.Command(cmd, args...)
	bs, err := ecmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return bytes.TrimSpace(bs), nil
}

func runPrint(cmd string, args ...string) {
	log.Println(cmd, strings.Join(args, " "))
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	err := ecmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func md5File(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	h := md5.New()
	_, err = io.Copy(h, fd)
	if err != nil {
		return err
	}

	out, err := os.Create(file + ".md5")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "%x\n", h.Sum(nil))
	if err != nil {
		return err
	}

	return out.Close()
}

func sha1FilesInDist() {
	filepath.Walk("./dist", func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, ".sha1") == false {
			sha1File(path)
		}
		return nil
	})
}

func sha1File(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	h := sha1.New()
	_, err = io.Copy(h, fd)
	if err != nil {
		return err
	}

	out, err := os.Create(file + ".sha1")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "%x\n", h.Sum(nil))
	if err != nil {
		return err
	}

	return out.Close()
}

// fileExists checks if a file exists
func fileExists(path ...string) bool {
	finfo, err := os.Stat(filepath.Join(path...))
	if err == nil && !finfo.IsDir() {
		return true
	}
	if os.IsNotExist(err) || finfo.IsDir() {
		return false
	}
	if err != nil {
		fatal(err)
	}
	return true
}

// readFile reads a file and return the trimmed output
func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.Trim(string(data), "\n\r ")
}

// fatal prints a error and exit
func fatal(err error) {
	fmt.Fprintln(os.Stderr, "!!", err)
	os.Exit(1)
}

// fatalMsg prints a fatal message alongside the error and exit
func fatalMsg(msg string, err error) {
	fmt.Fprintf(os.Stderr, "!! %s: %s\n", msg, err)
	os.Exit(1)
}
