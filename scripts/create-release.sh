#!/bin/bash

# 建立新版本的腳本
set -e

if [ -z "$1" ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 1.0.0"
    exit 1
fi

VERSION=$1

# 更新版本常數（可選）
# sed -i "s/version = \".*\"/version = \"$VERSION\"/" cmd/trafficmon/main.go

# 提交更改
git add .
git commit -m "Release version $VERSION" || true

# 建立標籤
git tag -a "v$VERSION" -m "Version $VERSION"

# 推送標籤（觸發 GitHub Actions）
git push origin "v$VERSION"

echo "Release v$VERSION created and pushed!"
