package main

import (
	"context"
	"fmt"

	"dagger.io/dagger"
)

func RunGoLint(ctx context.Context, c *dagger.Client) error {
	src := c.Host().Directory(".", dagger.HostDirectoryOpts{
		Exclude: []string{".git", "node_modules", "dist"},
	})

	// Use the same linter version as your current GHA (golangci-lint v2.0 maps to v1.58.0 in container)
	linter := c.Container().
		From("golangci/golangci-lint:v1.58.0").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"golangci-lint", "run", "--timeout=10m", "-v"})

	_, err := linter.Sync(ctx)
	if err != nil {
		return err
	}

	fmt.Println("golangci-lint passed via Dagger")
	return nil
}