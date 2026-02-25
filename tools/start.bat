@echo off
@rem 文字化け防止
chcp 65001 >nul
setlocal
cd /d %~dp0..

echo ==========================================
echo [1/3] 古いサーバー情報をクリアしています...
docker stop game-api >nul 2>&1
docker rm game-api >nul 2>&1

echo [2/3] 最新のプログラムを読み込んでいます...
docker build -f docker/Dockerfile -t my-game-server .

echo [3/3] サーバーを起動しています...
docker run -d -p 8080:8080 --name game-api my-game-server

echo.
echo ==========================================
echo 完了しました！
echo.
echo ブラウザで以下のURLを開いてください:
echo http://localhost:8080
echo ==========================================
echo.
pause