package bus

import (
	"advisor/advisor/event"
	monitor2 "advisor/advisor/event/monitor"
	"advisor/internal/redact"

	"github.com/wagoodman/go-partybus"
	"github.com/wagoodman/go-progress"

	"github.com/anchore/clio"
)

func Exit() {
	Publish(clio.ExitEvent(false))
}

func ExitWithInterrupt() {
	Publish(clio.ExitEvent(true))
}

func Report(report string) {
	if len(report) == 0 {
		return
	}
	report = redact.Apply(report)
	Publish(partybus.Event{
		Type:  event.CLIReport,
		Value: report,
	})
}

func Notify(message string) {
	Publish(partybus.Event{
		Type:  event.CLINotification,
		Value: message,
	})
}

func StartCatalogerTask(info monitor2.GenericTask, size int64, initialStage string) *monitor2.CatalogerTaskProgress {
	t := &monitor2.CatalogerTaskProgress{
		AtomicStage: progress.NewAtomicStage(initialStage),
		Manual:      progress.NewManual(size),
	}

	Publish(partybus.Event{
		Type:   event.CatalogerTaskStarted,
		Source: info,
		Value:  progress.StagedProgressable(t),
	})

	return t
}
