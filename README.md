# leetcode-go

Go で解く LeetCode（[NeetCode 150](https://neetcode.io/practice) の順序）。
解答置き場ではなく、**「何も見ずに書ける状態」を作るための間隔反復リポジトリ**。

## 方針

一度通ったコードは、覚えているとは言わない。同じ問題を、日を空けて、何も見ずに書き直せて初めて定着とみなす。

| 状態 | 次に再実装する日 |
|---|---|
| 初回 or 詰まった | 3 日後 |
| 何も見ずに通した ×1 | 7 日後 |
| 何も見ずに通した ×2 | 21 日後 |
| 何も見ずに通した ×3 | 定着。キューから外す |

この間隔管理を `lcq` コマンドが持つ。

## 目標

- 2028 年 5 月までに累計 350 問
- Medium が安定して解ける状態
- ペース: 週 5 問（新規）+ 期限の来た再実装

進捗は [PROGRESS.md](PROGRESS.md)。

## 構成

```
neetcode150/
  01-arrays-hashing/
    0217-contains-duplicate/
      solution.go       解答
      solution_test.go  テーブル駆動テスト
      NOTES.md          考え方・詰まった点・再実装ログ
cmd/lcq/                進捗管理 CLI（標準ライブラリのみ）
progress.json           状態
```

各問題は `package solution`。ディレクトリが分かれているのでパッケージ名は衝突しない。

## 使い方

```bash
go build -o lcq ./cmd/lcq

./lcq due                       # 今日再実装すべき問題
./lcq new 01-arrays-hashing 1 two-sum -diff Easy
go test ./neetcode150/01-arrays-hashing/0001-two-sum/
./lcq log two-sum -m 18 -noref  # 何も見ずに 18 分で通した
./lcq sync                      # PROGRESS.md を再生成
./lcq stats
```

全体テスト:

```bash
go test ./...
gofmt -l .
go vet ./...
```

## NOTES.md に書くこと

解法そのものより、**なぜその解法に辿り着けるか**を残す。3 日後の自分が読むのは「気づくべきポイント」の一行だけ。

## カテゴリ

| # | カテゴリ | 問題数 |
|---|---|---|
| 01 | Arrays & Hashing | 9 |
| 02 | Two Pointers | 5 |
| 03 | Sliding Window | 6 |
| 04 | Stack | 7 |
| 05 | Binary Search | 7 |
| 06 | Linked List | 11 |
| 07 | Trees | 15 |
| 08 | Tries | 3 |
| 09 | Heap / Priority Queue | 7 |
| 10 | Backtracking | 9 |
| 11 | Graphs | 13 |
| 12 | Advanced Graphs | 6 |
| 13 | 1-D Dynamic Programming | 12 |
| 14 | 2-D Dynamic Programming | 11 |
| 15 | Greedy | 8 |
| 16 | Intervals | 6 |
| 17 | Math & Geometry | 8 |
| 18 | Bit Manipulation | 7 |
