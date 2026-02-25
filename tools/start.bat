@echo off
chcp 65001 >nul
setlocal
cd /d %~dp0..

echo "=========================================="
echo "      ゲーム環境（サーバー+DB）起動"
echo "=========================================="

echo "[1/2] 以前の環境を掃除しています..."
docker compose down >nul 2>&1

echo "[2/2] 全てのサービスを起動中..."
echo "※ 初回は各データベースのダウンロードに時間がかかります。"
docker compose up -d --build

echo.
echo "------------------------------------------"
echo "✅ 起動成功！"
echo "・API: http://localhost:8080"
echo "・Postgres: localhost:5432"
echo "・Redis: localhost:6379"
echo "・Memcached: localhost:11211"
echo "------------------------------------------"
pause