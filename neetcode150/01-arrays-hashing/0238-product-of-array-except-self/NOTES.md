# 238. Product Of Array Except Self (Medium)

https://leetcode.com/problems/product-of-array-except-self/

## 何を問われているか（自分の言葉で）


## 気づくべきポイント
<!-- 「なぜこの解法を思いつけるか」の引き金になった観察を一行で -->


## 解法の型
<!-- ハッシュで数える / 両端から寄せる / 単調スタック など、抽象名で -->
 answerを[]int, len(nums)で用意。
  bfr := make([]int, len(nums))用意
  bfr[0] = 1
  aft := make([]int, len(nums))用意
  aft[len(nums) - 1] = 1

  for i := 1; i < len(nums); i++でループ
  　前の積bfr[i]=bfr[i-1]*nums[i-1]

  for i := len(nums) - 2; i>= 0; i--で後ろからループ
  　後ろの積aft[i]=aft[i+1]*nums[i+1]

  range numsでforループ
  　answer[i] = bfr[i]*aft[i]

  return answer

  ２
  返り値格納用スライスanswerを[]int, len(nums)で用意。
  みている要素の左側の積格納用leftを[]int, len(nums)で用意。
  同右側用rightを[]int, len(nums)で用意。
  left[0]は1で初期化
  right[len(nums)-1]も1で初期化。

  for i:=1; i<len(nums); i++
  　left[i] = left[i-1]*nums[i-1]にする

  for i:=len(nums)-2; i>=0; i--
  　right[i] = right[i+1]*nums[i+1]にする

  for i:= range nums
  　answer[i] = left[i]*right[i]にする

  return answerする

## 詰まった点


## 再実装ログ

| 日付 | 分 | ノーカンニング | 通った | メモ |
|---|---|---|---|---|
| 2026-11-25 |  |  |  |  |
