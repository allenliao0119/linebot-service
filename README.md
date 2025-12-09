# LINE Bot Service

一個基於 Go 語言開發的 LINE 聊天機器人服務，提供簡單對話和 AI 對話功能。

## 專案概述

本專案採用分階段開發策略：

- **第一階段：Simple Chat（已完成）** - 基於關鍵字匹配的簡單對話系統
- **第二階段：AI Chat（規劃中）** - 整合 OpenAI API 的智能對話系統

## 功能特性

### ✅ 第一階段：Simple Chat（已實作）

- **關鍵字識別與回應**
  - 問候：「你好」、「hi」、「hello」
  - 感謝：「謝謝」、「感謝」、「thanks」
  - 再見：「再見」、「bye」、「goodbye」
  - 幫助：「幫助」、「help」
  - 預設回應：智能學習模式（回傳使用者訊息）

- **LINE Bot 整合**
  - Webhook 簽名驗證
  - 文字訊息處理
  - 粉絲跟隨事件歡迎訊息
  - 支援個人聊天、群組、聊天室多種訊息來源

- **HTTP 服務**
  - `/health` - 健康檢查端點
  - `/webhook` - LINE Webhook 端點

### 🚧 第二階段：AI Chat（開發中）

- OpenAI API 整合（架構已就緒，待實作）
- 智能對話功能
- 可配置的 AI 模型選擇

## 技術架構

### 技術棧

- **語言**: Go 1.25.1
- **Web 框架**: Gin v1.11.0
- **LINE SDK**: line-bot-sdk-go v8.18.0
- **配置管理**: envconfig + godotenv

### 目錄結構

```
linebot-service/
├── cmd/
│   └── server/
│       └── main.go              # 應用程式入口點
├── internal/
│   ├── config/
│   │   └── config.go            # 環境配置管理
│   ├── handler/
│   │   └── linebot_webhook.go   # Webhook 處理器
│   ├── service/
│   │   ├── chat.go              # ChatService 介面
│   │   └── linebot.go           # LINE Bot 服務
│   └── bot/
│       ├── simple_chat.go       # 簡單聊天模式
│       └── ai_chat.go           # AI 聊天模式（待實作）
├── go.mod
└── go.sum
```

### 架構設計

- **清潔架構**：分層設計（handler → service → bot）
- **介面驅動**：透過 `ChatService` 介面支援靈活切換聊天模式
- **依賴注入**：在 main.go 進行集中配置

## 快速開始

### 前置需求

- Go 1.25 或更高版本
- LINE Developers 帳號
- LINE Bot Channel（Messaging API）

### 安裝

1. Clone 專案

```bash
git clone <repository-url>
cd linebot-service
```

2. 安裝依賴

```bash
go mod download
```

3. 設定環境變數

建立 `.env` 檔案：

```env
# LINE Bot 設定（必需）
LINE_CHANNEL_SECRET=your_channel_secret
LINE_CHANNEL_ACCESS_TOKEN=your_channel_access_token

# 伺服器設定（選用）
SERVER_PORT=8080
SERVER_ENV=development

# 聊天模式（選用）
CHAT_MODE=simple

# OpenAI 設定（AI 模式使用）
# OPENAI_API_KEY=your_openai_api_key
# OPENAI_MODEL=gpt-4o-mini
```

### 執行

```bash
# 開發模式
go run cmd/server/main.go

# 生產模式
SERVER_ENV=production go run cmd/server/main.go

# 或建置後執行
go build -o bin/server cmd/server/main.go
./bin/server
```

### 本地測試

使用 ngrok 將本地服務暴露到外網：

```bash
ngrok http 8080
```

在 LINE Developers Console 設定 Webhook URL：
```
https://your-ngrok-url/webhook
```

## 環境變數說明

| 變數名 | 必需 | 預設值 | 說明 |
|--------|------|--------|------|
| `LINE_CHANNEL_SECRET` | ✅ | - | LINE 頻道密鑰 |
| `LINE_CHANNEL_ACCESS_TOKEN` | ✅ | - | LINE 頻道存取令牌 |
| `SERVER_PORT` | ❌ | 8080 | 伺服器埠號 |
| `SERVER_ENV` | ❌ | development | 環境模式（development/production） |
| `CHAT_MODE` | ❌ | simple | 聊天模式（simple/ai） |
| `OPENAI_API_KEY` | ⚠️ | - | OpenAI API Key（AI 模式必需） |
| `OPENAI_MODEL` | ❌ | gpt-4o-mini | OpenAI 模型名稱 |

## API 端點

### GET /health

健康檢查端點

**回應範例**：
```json
{
  "status": "healthy",
  "environment": "development"
}
```

### POST /webhook

LINE Webhook 端點，處理來自 LINE 平台的事件。

## 開發

### 切換聊天模式

透過環境變數切換：

```bash
# 簡單模式（預設）
CHAT_MODE=simple go run cmd/server/main.go

# AI 模式（待第二階段實作）
CHAT_MODE=ai OPENAI_API_KEY=sk-xxx go run cmd/server/main.go
```

### 專案約定

- 使用 Conventional Commits 格式提交訊息
- 遵循 Go 標準程式碼風格
- 使用依賴注入模式

## 未來規劃

- [ ] 完成 OpenAI API 整合（第二階段）
- [ ] 新增單元測試與整合測試
- [ ] 實作對話歷史記錄功能
- [ ] 支援更多 LINE 訊息類型（圖片、貼圖等）
- [ ] 新增日誌系統與監控

## 作者

Allen Liao
