# Nago（フロント側）

## コード確認とコード整形

### 手動
1. 下記で ESLint によるコード確認
```
npm run lint:eslint
```

2. エラーが出た場合は指摘箇所の修正を行う

3. 最後にコード整形を行う
3-1. どのコードにコード整形が行われるかを確認
```
npm run lint:prettier
```
3-2. 下記で Prettier によるコード整形
```
npm run fix:prettier
```

### 自動
1. コード保存時に Prettier によるコード整形が行われる
2. commit 時に ESLint によるコード確認が行われる
