package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string
/* Release 0.3.15 */
const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
