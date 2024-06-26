// Package markdown is markdown builder that includes to convert Markdown to HTML.
package markdown

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPlainText(t *testing.T) {
	t.Parallel()

	t.Run("success PlainText()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.PlainText("Hello")
		want := []string{"Hello"}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMarkdownHeader(t *testing.T) {
	t.Parallel()

	t.Run("success H1f()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H1f("%s", "Hello")
		want := "# Hello"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("success H2f()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H2f("%s", "Hello")
		want := "## Hello"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("success H3f()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H3f("%s", "Hello")
		want := "### Hello"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("success H4f()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H4f("%s", "Hello")
		want := "#### Hello"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("success H5f()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H5f("%s", "Hello")
		want := "##### Hello"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("success H6f()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H6f("%s", "Hello")
		want := "###### Hello"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMarkdownDetailsf(t *testing.T) {
	t.Parallel()

	t.Run("success Detailsf()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.Detailsf("Hello", "Good %s", "World")
		want := fmt.Sprintf("<details><summary>Hello</summary>%sGood World%s</details>", lineFeed(), lineFeed())
		got := m.body[0]

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMarkdownBulletList(t *testing.T) {
	t.Parallel()

	t.Run("success BulletList()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.BulletList("Hello", "World")
		want := []string{"- Hello", "- World"}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMarkdownNumberList(t *testing.T) {
	t.Parallel()

	t.Run("success NumberList()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.OrderedList("Hello", "World")
		want := []string{"1. Hello", "2. World"}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMarkdownCheckBoxf(t *testing.T) {
	t.Run("success CheckBoxf(); check [x]", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		set := []CheckBoxSet{
			{Text: "Hello", Checked: true},
			{Text: "World", Checked: false},
		}
		m.CheckBox(set)
		want := []string{
			"- [x] Hello",
			"- [ ] World",
		}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMarkdownBlockquote(t *testing.T) {
	t.Parallel()

	t.Run("success Blockquote()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.Blockquote(fmt.Sprintf("%s%s%s%s%s", "Hello", lineFeed(), "Good", lineFeed(), "World"))
		want := []string{
			"> Hello",
			"> Good",
			"> World",
		}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMarkdownCodeBlocks(t *testing.T) {
	t.Run("success CodeBlock()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.CodeBlocks(SyntaxHighlightGo, "Hello")
		want := []string{fmt.Sprintf("```go%sHello%s```", lineFeed(), lineFeed())}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestHorizontalRule(t *testing.T) {
	t.Parallel()

	t.Run("success HorizontalRule()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.HorizontalRule()
		want := []string{"---"}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestTableSetValidateColumns(t *testing.T) {
	t.Parallel()
	t.Run("success TableSet.ValidateColumns()", func(t *testing.T) {
		t.Parallel()

		set := TableSet{
			Header: []string{"Name", "Age"},
			Rows:   [][]string{{"David", "23"}},
		}

		err := set.ValidateColumns()
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})

	t.Run("failed TableSet.ValidateColumns(); invalid header", func(t *testing.T) {
		t.Parallel()

		set := TableSet{
			Header: []string{"Name", "Age"},
			Rows:   [][]string{{"David"}},
		}

		err := set.ValidateColumns()
		if err == nil {
			t.Error("expected error, but not occurred")
		}
	})
}

func TestMarkdownTable(t *testing.T) {
	t.Parallel()
	t.Run("success Table()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		set := TableSet{
			Header: []string{"Name", "Age"},
			Rows:   [][]string{{"David", "23"}},
		}
		m.Table(set)
		want := []string{
			fmt.Sprintf("| NAME  | AGE |%s|-------|-----|%s| David |  23 |%s",
				lineFeed(), lineFeed(), lineFeed()),
		}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestMarkdownBuildError(t *testing.T) {
	t.Parallel()

	t.Run("Error() return nil", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		if err := m.H1("sample").Build(); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})

	t.Run("Error() return error", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.Table(TableSet{
			Header: []string{"Name", "Age"},
			Rows:   [][]string{{"David"}},
		})
		if err := m.Build(); err == nil {
			t.Error("expected error, but not occurred")
		}
	})
}

func TestMarkdownLF(t *testing.T) {
	t.Parallel()
	t.Run("success Markdown.LF()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.LF()
		want := []string{"  "}
		got := m.body

		if !reflect.DeepEqual(want, got) {
			t.Errorf("value is mismatch want: %v, got: %v", want, got)
		}
	})
}

func TestMarkdownError(t *testing.T) {
	t.Parallel()

	t.Run("Error() return nil", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		if err := m.H1("sample").Error(); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})

	t.Run("Error() return error", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.Table(TableSet{
			Header: []string{"Name", "Age"},
			Rows:   [][]string{{"David"}},
		})
		if err := m.Error(); err == nil {
			t.Error("expected error, but not occurred")
		}
	})

	t.Run("Error() return error Custom Table", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.CustomTable(TableSet{
			Header: []string{"Name", "Age"},
			Rows:   [][]string{{"David"}},
		}, TableOptions{
			AutoWrapText: false,
		})
		if err := m.Error(); err == nil {
			t.Error("expected error, but not occurred")
		}
	})
}

func TestMarkdownCustomTable(t *testing.T) {
	t.Parallel()
	t.Run("success Table()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		set := TableSet{
			Header: []string{"Name", "Age"},
			Rows:   [][]string{{"David", "23"}},
		}
		m.CustomTable(set, TableOptions{
			AutoWrapText:      false,
			AutoFormatHeaders: false,
		})
		want := []string{
			fmt.Sprintf("| Name  | Age |%s|-------|-----|%s| David |  23 |%s",
				lineFeed(), lineFeed(), lineFeed()),
		}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}
