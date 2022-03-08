package assets

import "embed"

var (
	//go:embed *
	AssetsFileSystem embed.FS
)
