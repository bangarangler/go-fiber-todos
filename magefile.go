// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Migrate() {
	fmt.Print("running migrations")
	sh.Run("migrate", "-source", "file://postgres/migrations", "-database", "postgres://jonp:jonp@localhost:5432/go_fiber_todos_db?sslmode=disable", "up")
}

func Rollback() {
	fmt.Print("rolling back migrations")
	sh.Run("migrate", "-source", "file://postgres/migrations", "-database", "postgres://jonp:jonp@localhost:5432/go_fiber_todos_db?sslmode=disable", "down")
}

func Drop() {
	fmt.Print("rolling back migrations")
	sh.Run("migrate", "-source", "file://postgres/migrations", "-database", "postgres://jonp:jonp@localhost:5432/go_fiber_todos_db?sslmode=disable", "drop")
}

func Migration() {
	var name string
	fmt.Print("Enter Migation name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Creating migration files...")
	sh.Run("migrate", "create", "-ext", "sql", "-dir", "postgres/migrations", name)
}

func SQLCGen() error {
	fmt.Println("sqlc generating queries...")
	system := runtime.GOOS
	switch system {
	case "windows":
		println("No Thank You, Switch to Linux ; )")
	case "darwin":
		println("Running on mac")
		return sh.Run("docker", "run", "--rm", "-v", "/Users/jonathanpalacio/Desktop/go-fiber-todos:/src", "-w", "/src", "kjconroy/sqlc", "generate")
	case "linux":
		println("Linux ; )")
		return sh.Run("docker", "run", "--rm", "-v", "/home/jonathan/Desktop/go-gqlgen-sqlc-example:/src", "-w", "/src", "kjconroy/sqlc", "generate")
	}
	return nil
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	return cmd.Run()
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./MyApp", "/usr/bin/MyApp")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("MyApp")
}
