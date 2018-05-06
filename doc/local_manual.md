# Docker を使ってローカルで環境を整える

```
$ git clone git@github.com:showwin/ISHOCON2.git
$ cd ISHOCON2
$ docker-compose build
$ docker-compose up
# app_1 と bench_1 のログに 'setup completed.' と出たら起動完了
```

## アプリケーション

```
$ docker exec -it ishocon2_app_1 /bin/bash
```

アプリケーションの起動は [マニュアル](https://github.com/showwin/ISHOCON2/blob/master/doc/manual.md) 参照

## ベンチマーカー

```
$ docker exec -it ishocon2_bench_1 /bin/bash
$ ./benchmark --ip app:443  # docker-compose.yml で link しているので app で到達できます
```
