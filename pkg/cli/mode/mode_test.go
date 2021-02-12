package mode

import (
	"testing"

	"src.elv.sh/pkg/tt"
	"src.elv.sh/pkg/ui"
)

func TestModeLine(t *testing.T) {
	testModeLine(t, tt.Fn("Line", ModeLine))
}

func TestModePrompt(t *testing.T) {
	testModeLine(t, tt.Fn("Prompt",
		func(s string, b bool) ui.Text { return ModePrompt(s, b)() }))
}

func testModeLine(t *testing.T, fn *tt.FnToTest) {
	tt.Test(t, fn, tt.Table{
		tt.Args("TEST", false).Rets(
			ui.T("TEST", ui.Bold, ui.FgWhite, ui.BgMagenta)),
		tt.Args("TEST", true).Rets(
			ui.Concat(
				ui.T("TEST", ui.Bold, ui.FgWhite, ui.BgMagenta),
				ui.T(" "))),
	})
}
