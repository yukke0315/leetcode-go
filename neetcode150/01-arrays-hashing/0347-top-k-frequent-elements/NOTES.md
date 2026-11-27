# 347. Top K Frequent Elements (Medium)

https://leetcode.com/problems/top-k-frequent-elements/

## 何を問われているか（自分の言葉で）


## 気づくべきポイント
<!-- 「なぜこの解法を思いつけるか」の引き金になった観察を一行で -->


## 解法の型
<!-- ハッシュで数える / 両端から寄せる / 単調スタック など、抽象名で -->
cntをmap[int]int, len(nums)で用意。
  range numsでforループ
  　numsの要素をキーとしたcnt[n]をインクリメント
  並べ替え用のスライスsを[][]int, len(nums)+1で用意。
  range cntでforループ
  　値とキーを入れ替える(s[v] = append(s[v], n)
  結果格納用スライスresultを[]int, 0, kで用意。
  for i:=len(nums); i>=0; i--
  　range s[i]でforループ
  　　resultのスライスにappend
  　　kをデクリメント
      k==0ならreturn result
  return nil

## 詰まった点


## 再実装ログ

| 日付 | 分 | ノーカンニング | 通った | メモ |
|---|---|---|---|---|
| 2026-11-23 |  |  |  |  |
