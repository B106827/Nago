# Nago

## ローカル開発環境
- Dockerを使用しています。

## 各種コマンド
|                                  |                                                                    |
|:---------------------------------|:-------------------------------------------------------------------|
| 起動                             | `$ docker-compose up` ( -d: daemon )                               |
| 停止                             | `$ docker-compsoe down`                                            |
| 再起動                           | `$ docker-compose restart api`                                     |
| 各コンテナにアクセス             | `$ docker-compose exec api sh`                                     |
| ログ表示                         | `$ docker-compose logs -f api`                                     |
| 滅びの呪文(諸々リセットしたい時) | `$ docker-compose down --rmi all --volumes --remove-orphans`       |

## docker-compose
| service     | link                  |
|:------------|:----------------------|
| front       | http://localhost:3001 |
| nginx+api   | http://localhost:80   |
| mysql       | http://localhost:3036 |
| maildev     | http://localhost:1080 |
