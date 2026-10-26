package main

import (
	"os"
	"strings"
	"testing"
)

func exp(rows ...string) string {
	var b strings.Builder
	for _, r := range rows {
		b.WriteString(strings.TrimSuffix(r, "$"))
		b.WriteByte('\n')
	}
	return b.String()
}

var helloRows = []string{
	` _    _          _   _          $`,
	`| |  | |        | | | |         $`,
	`| |__| |   ___  | | | |   ___   $`,
	`|  __  |  / _ \ | | | |  / _ \  $`,
	`| |  | | |  __/ | | | | | (_) | $`,
	`|_|  |_|  \___| |_| |_|  \___/  $`,
	`                                $`,
	`                                $`,
}

var thereRows = []string{
	` _______   _                           $`,
	`|__   __| | |                          $`,
	`   | |    | |__     ___   _ __    ___  $`,
	`   | |    |  _ \   / _ \ | '__|  / _ \ $`,
	`   | |    | | | | |  __/ | |    |  __/ $`,
	`   |_|    |_| |_|  \___| |_|     \___| $`,
	`                                       $`,
	`                                       $`,
}

func TestTransformToAscii(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"empty", "", ""},
		{"only newline", "\\n", "\n"},
		{"hello newline", "Hello\\n", exp(helloRows...) + "\n"},
		{"hello", "hello", exp(
			` _              _   _          $`,
			`| |            | | | |         $`,
			`| |__     ___  | | | |   ___   $`,
			`|  _ \   / _ \ | | | |  / _ \  $`,
			`| | | | |  __/ | | | | | (_) | $`,
			`|_| |_|  \___| |_| |_|  \___/  $`,
			`                               $`,
			`                               $`,
		)},
		{"HeLlO", "HeLlO", exp(
			` _    _          _        _    ____   $`,
			`| |  | |        | |      | |  / __ \  $`,
			`| |__| |   ___  | |      | | | |  | | $`,
			`|  __  |  / _ \ | |      | | | |  | | $`,
			`| |  | | |  __/ | |____  | | | |__| | $`,
			`|_|  |_|  \___| |______| |_|  \____/  $`,
			`                                      $`,
			`                                      $`,
		)},
		{"Hello There", "Hello There", exp(
			` _    _          _   _                 _______   _                           $`,
			`| |  | |        | | | |               |__   __| | |                          $`,
			`| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  $`,
			`|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ $`,
			`| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ $`,
			`|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| $`,
			`                                                                             $`,
			`                                                                             $`,
		)},
		{"1Hello 2There", "1Hello 2There", exp(
			`     _    _          _   _                         _______   _                           $`,
			` _  | |  | |        | | | |                ____   |__   __| | |                          $`,
			`/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  $`,
			`| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ $`,
			`| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ $`,
			`|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| $`,
			`                                                                                         $`,
			`                                                                                         $`,
		)},
		{"braces", "{Hello There}", exp(
			`   __  _    _          _   _                 _______   _                           __    $`,
			`  / / | |  | |        | | | |               |__   __| | |                          \ \   $`,
			` | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  $`,
			`/ /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ $`,
			`\ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / $`,
			` | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  $`,
			`  \_\                                                                              /_/   $`,
			`                                                                                         $`,
		)},
		{"Hello\\nThere", "Hello\\nThere", exp(helloRows...) + exp(thereRows...)},
		{"Hello\\n\\nThere", "Hello\\n\\nThere", exp(helloRows...) + "\n" + exp(thereRows...)},
	}

	testData, _ := os.ReadFile("standard.txt")
	lines := strings.Split(string(testData), "\n")
	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			got := TransformToAscii(ts.input, lines)
			if got != ts.output {
				t.Fatalf("\ninput: %q\nGot:   %q\nWant:  %q", ts.input, got, ts.output)
			}
		})
	}
}
