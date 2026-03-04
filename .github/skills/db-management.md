# スキル：データベース管理

## 構成ルール
- **Center DB**: ユーザーのDB割り振り（`user_db_mapping`）およびアクティブなマスターバージョンを管理。
- **Master DB**: `master_{version}` 形式で作成。マスターデータ（静的データ）を保持し、バージョンごとにDBを分離する。
- **User DB**: `user_db_{n}` 形式で作成。各ユーザーの個別データを保持。

## 初期化ルール
- `init.sh` (Linux) / `init.bat` (Windows) を使用してDBを構築する。
- **破壊的更新**: `init` スクリプト実行時は、既存のDBを全て `DROP` してから `CREATE` する。
- **SQL管理**: `/migrations/{center|master|user}/` 配下のSQLファイルを、ファイル名先頭の連番（例: 01_xxx.sql）順に実行する。
