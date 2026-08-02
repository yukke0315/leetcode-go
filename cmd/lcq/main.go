// lcq は LeetCode の進捗と「3日後の再実装」サイクルを管理する CLI。
// 標準ライブラリのみで動く。
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const (
	progressFile = "progress.json"
	markdownFile = "PROGRESS.md"
	rootDir      = "neetcode150"
	dateLayout   = "2006-01-02"
)

// intervals は復習間隔（日）。
// 「何も見ずに解けた」が続くたびに次の段階へ進み、詰まったら先頭に戻る。
// 3日後の再実装がこのサイクルの起点。
var intervals = []int{3, 7, 21}

type Attempt struct {
	Date    string `json:"date"`
	Minutes int    `json:"minutes"`
	NoRef   bool   `json:"noRef"`  // 何も見ずに書けたか
	Solved  bool   `json:"solved"` // 最終的に通したか
	Note    string `json:"note,omitempty"`
}

type Problem struct {
	ID         int       `json:"id"`
	Slug       string    `json:"slug"`
	Title      string    `json:"title"`
	Category   string    `json:"category"`
	Difficulty string    `json:"difficulty"`
	Attempts   []Attempt `json:"attempts"`
}

type Progress struct {
	Problems []Problem `json:"problems"`
}

// streak は末尾から数えた「何も見ずに通した」連続回数。
func (p Problem) streak() int {
	n := 0
	for i := len(p.Attempts) - 1; i >= 0; i-- {
		a := p.Attempts[i]
		if !a.NoRef || !a.Solved {
			break
		}
		n++
	}
	return n
}

// mastered は復習サイクルを抜けたかどうか。
func (p Problem) mastered() bool {
	return p.streak() >= len(intervals)
}

// nextReview は次に再実装すべき日を返す。第2戻り値が false なら予定なし。
func (p Problem) nextReview() (time.Time, bool) {
	if len(p.Attempts) == 0 || p.mastered() {
		return time.Time{}, false
	}
	last, err := time.ParseInLocation(dateLayout, p.Attempts[len(p.Attempts)-1].Date, time.Local)
	if err != nil {
		return time.Time{}, false
	}
	i := p.streak()
	if i >= len(intervals) {
		i = len(intervals) - 1
	}
	return last.AddDate(0, 0, intervals[i]), true
}

func (p Problem) dir() string {
	return filepath.Join(rootDir, p.Category, fmt.Sprintf("%04d-%s", p.ID, p.Slug))
}

func load() (*Progress, error) {
	pr := &Progress{}
	b, err := os.ReadFile(progressFile)
	if errors.Is(err, os.ErrNotExist) {
		return pr, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, pr); err != nil {
		return nil, fmt.Errorf("%s の解析に失敗: %w", progressFile, err)
	}
	return pr, nil
}

func (pr *Progress) save() error {
	b, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(progressFile, append(b, '\n'), 0o644)
}

func (pr *Progress) find(slug string) *Problem {
	for i := range pr.Problems {
		if pr.Problems[i].Slug == slug {
			return &pr.Problems[i]
		}
	}
	return nil
}

func today() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

// --- new ---

func cmdNew(args []string) error {
	if len(args) < 3 {
		return errors.New("usage: lcq new <category> <id> <slug> [-title T] [-diff Easy|Medium|Hard]")
	}
	category, idStr, slug := args[0], args[1], args[2]

	fs := flag.NewFlagSet("new", flag.ExitOnError)
	title := fs.String("title", "", "問題タイトル")
	diff := fs.String("diff", "Medium", "難易度")
	if err := fs.Parse(args[3:]); err != nil {
		return err
	}

	var id int
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return fmt.Errorf("id が数値でない: %q", idStr)
	}
	if *title == "" {
		*title = titleFromSlug(slug)
	}

	pr, err := load()
	if err != nil {
		return err
	}
	if pr.find(slug) != nil {
		return fmt.Errorf("%s はすでに登録済み", slug)
	}

	p := Problem{ID: id, Slug: slug, Title: *title, Category: category, Difficulty: *diff}
	if err := os.MkdirAll(p.dir(), 0o755); err != nil {
		return err
	}

	rep := strings.NewReplacer(
		"{{TITLE}}", p.Title,
		"{{ID}}", fmt.Sprintf("%d", p.ID),
		"{{SLUG}}", p.Slug,
		"{{DIFF}}", p.Difficulty,
		"{{DATE}}", today().Format(dateLayout),
	)
	files := map[string]string{
		"solution.go":      solutionTmpl,
		"solution_test.go": testTmpl,
		"NOTES.md":         notesTmpl,
	}
	for name, tmpl := range files {
		path := filepath.Join(p.dir(), name)
		if _, err := os.Stat(path); err == nil {
			continue // 既存ファイルは上書きしない
		}
		if err := os.WriteFile(path, []byte(rep.Replace(tmpl)), 0o644); err != nil {
			return err
		}
	}

	pr.Problems = append(pr.Problems, p)
	if err := pr.save(); err != nil {
		return err
	}
	fmt.Printf("作成: %s\n", p.dir())
	return nil
}

func titleFromSlug(slug string) string {
	words := strings.Split(slug, "-")
	for i, w := range words {
		if w == "" {
			continue
		}
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, " ")
}

// --- log ---

func cmdLog(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: lcq log <slug> [-m 分] [-noref] [-fail] [-note \"...\"]")
	}
	slug := args[0]

	fs := flag.NewFlagSet("log", flag.ExitOnError)
	minutes := fs.Int("m", 0, "かかった分数")
	noRef := fs.Bool("noref", false, "何も見ずに書けた")
	fail := fs.Bool("fail", false, "通らなかった / 詰まった")
	note := fs.String("note", "", "メモ")
	date := fs.String("date", today().Format(dateLayout), "日付 YYYY-MM-DD")
	if err := fs.Parse(args[1:]); err != nil {
		return err
	}

	pr, err := load()
	if err != nil {
		return err
	}
	p := pr.find(slug)
	if p == nil {
		return fmt.Errorf("%s は未登録。先に lcq new を実行", slug)
	}
	if _, err := time.Parse(dateLayout, *date); err != nil {
		return fmt.Errorf("日付の形式が不正: %q", *date)
	}

	p.Attempts = append(p.Attempts, Attempt{
		Date:    *date,
		Minutes: *minutes,
		NoRef:   *noRef,
		Solved:  !*fail,
		Note:    *note,
	})
	if err := pr.save(); err != nil {
		return err
	}

	if next, ok := p.nextReview(); ok {
		fmt.Printf("記録: %s（連続 %d 回）→ 次回 %s\n",
			p.Slug, p.streak(), next.Format(dateLayout))
	} else {
		fmt.Printf("記録: %s → 定着完了。復習キューから外した\n", p.Slug)
	}
	return nil
}

// --- due ---

func cmdDue(args []string) error {
	fs := flag.NewFlagSet("due", flag.ExitOnError)
	within := fs.Int("in", 0, "何日先まで表示するか")
	if err := fs.Parse(args); err != nil {
		return err
	}

	pr, err := load()
	if err != nil {
		return err
	}
	limit := today().AddDate(0, 0, *within)

	type row struct {
		p    Problem
		when time.Time
	}
	var rows []row
	for _, p := range pr.Problems {
		next, ok := p.nextReview()
		if !ok || next.After(limit) {
			continue
		}
		rows = append(rows, row{p, next})
	}
	if len(rows) == 0 {
		fmt.Println("再実装の予定なし。新規問題へ進んでいい")
		return nil
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].when.Before(rows[j].when) })

	for _, r := range rows {
		mark := " "
		if !r.when.After(today()) {
			mark = "*" // 期限到来
		}
		fmt.Printf("%s %s  %-32s %s (%d回目)\n",
			mark, r.when.Format(dateLayout), r.p.Slug, r.p.Difficulty, len(r.p.Attempts)+1)
	}
	return nil
}

// --- stats ---

func cmdStats(args []string) error {
	pr, err := load()
	if err != nil {
		return err
	}
	var attempted, mastered, totalMin int
	byDiff := map[string]int{}
	for _, p := range pr.Problems {
		if len(p.Attempts) > 0 {
			attempted++
			byDiff[p.Difficulty]++
		}
		if p.mastered() {
			mastered++
		}
		for _, a := range p.Attempts {
			totalMin += a.Minutes
		}
	}
	fmt.Printf("登録        : %d 問\n", len(pr.Problems))
	fmt.Printf("着手済み    : %d 問 (Easy %d / Medium %d / Hard %d)\n",
		attempted, byDiff["Easy"], byDiff["Medium"], byDiff["Hard"])
	fmt.Printf("定着済み    : %d 問\n", mastered)
	fmt.Printf("累計時間    : %d 時間 %d 分\n", totalMin/60, totalMin%60)
	fmt.Printf("NeetCode150 : %d / 150\n", mastered)
	return nil
}

// --- sync ---

func cmdSync(args []string) error {
	pr, err := load()
	if err != nil {
		return err
	}
	sort.Slice(pr.Problems, func(i, j int) bool {
		if pr.Problems[i].Category != pr.Problems[j].Category {
			return pr.Problems[i].Category < pr.Problems[j].Category
		}
		return pr.Problems[i].ID < pr.Problems[j].ID
	})

	var b strings.Builder
	b.WriteString("# 進捗\n\n")
	fmt.Fprintf(&b, "更新日: %s\n\n", today().Format(dateLayout))
	b.WriteString("| # | 問題 | 難易度 | 挑戦 | 連続ノーカンニング | 次回 |\n")
	b.WriteString("|---|---|---|---|---|---|\n")
	for _, p := range pr.Problems {
		next := "—"
		if p.mastered() {
			next = "定着"
		} else if t, ok := p.nextReview(); ok {
			next = t.Format(dateLayout)
		}
		fmt.Fprintf(&b, "| %d | [%s](%s) | %s | %d | %d | %s |\n",
			p.ID, p.Title, filepath.ToSlash(p.dir()), p.Difficulty,
			len(p.Attempts), p.streak(), next)
	}
	if err := os.WriteFile(markdownFile, []byte(b.String()), 0o644); err != nil {
		return err
	}
	if err := pr.save(); err != nil {
		return err
	}
	fmt.Printf("%s を更新\n", markdownFile)
	return nil
}

// --- main ---

func usage() {
	fmt.Fprint(os.Stderr, `lcq — LeetCode 進捗管理

  lcq new <category> <id> <slug> [-title T] [-diff D]   問題フォルダを作る
  lcq log <slug> [-m 分] [-noref] [-fail] [-note "..."]  挑戦を記録する
  lcq due [-in 日数]                                     再実装すべき問題を出す
  lcq stats                                              集計
  lcq sync                                               PROGRESS.md を再生成
`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	cmds := map[string]func([]string) error{
		"new":   cmdNew,
		"log":   cmdLog,
		"due":   cmdDue,
		"stats": cmdStats,
		"sync":  cmdSync,
	}
	run, ok := cmds[os.Args[1]]
	if !ok {
		usage()
		os.Exit(2)
	}
	if err := run(os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

const solutionTmpl = `package solution

// {{ID}}. {{TITLE}} ({{DIFF}})
// https://leetcode.com/problems/{{SLUG}}/
//
// 計算量: O(?) time / O(?) space
func solve() {
	panic("not implemented")
}
`

const testTmpl = `package solution

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		// TODO: 入力と期待値のフィールドを足す
	}{
		{name: "example 1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Skip("TODO: テストを書く")
		})
	}
}
`

const notesTmpl = `# {{ID}}. {{TITLE}} ({{DIFF}})

https://leetcode.com/problems/{{SLUG}}/

## 何を問われているか（自分の言葉で）


## 気づくべきポイント
<!-- 「なぜこの解法を思いつけるか」の引き金になった観察を一行で -->


## 解法の型
<!-- ハッシュで数える / 両端から寄せる / 単調スタック など、抽象名で -->


## 詰まった点


## 再実装ログ

| 日付 | 分 | ノーカンニング | 通った | メモ |
|---|---|---|---|---|
| {{DATE}} |  |  |  |  |
`
