# 機能
画像識別（犬顔か猫顔か）

![image](https://user-images.githubusercontent.com/89893576/135968794-a25dc8ee-a6a4-4589-9dad-fca3c4e2e1e0.png)

# 特徴
- ### バックエンド面
  - GraphQL(Go)
  - gRPC(Go)
  - gRPC(Python)
  - 機械学習 画像分類(Python)
- ### フロントエンド面
  - Next.js、appolo(graphQL)を使用

# アーキテクチャー
![graphQL_gRPC_ML](https://user-images.githubusercontent.com/89893576/135967805-08f3b52f-97d5-40d3-ba05-aab23cc7c710.jpg)


## 機能一覧
- アップロードした画像を犬顔なのか猫顔なのか識別
- 「web go言語」と「機械学習 python」とはgRPCにて通信