package monitor

import (
	"github.com/wagoodman/go-progress"
)

const (
	TopLevelCatalogingTaskID = "cataloging"
	CICatalogingTaskID       = "CI-cataloging"
)

type CatalogerTaskProgress struct {
	*progress.AtomicStage
	*progress.Manual
}
