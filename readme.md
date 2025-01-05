# 開發規格書

## 目的

開發一個命令行工具，將輸入的文字轉換成簡單的口語英文，並進一步轉換成 URL slug。

## 功能需求

1. 從 command line 接收輸入的文字。
2. 使用 OpenAI API 將文字轉換成簡單的口語英文並轉換成 slug。
   - Prompt: Convert this text into simple spoken English and then make it a URL slug. Adjust the number of words to simplify it. When responding, just give the result directly, no additional information is needed.

## 非功能需求

1. 使用 Go 語言實作。
2. 使用 openai 的套件與 OpenAI API 進行互動。
3. 提供簡單的命令行使用說明。

## 環境需求

1. Go 1.16 或更高版本。
2. 設定環境變數 `OPENAI_API_KEY` 以便訪問 OpenAI API。

## 使用說明

1. 編譯程式：

   ```sh
   go build -o to-slug main.go
   ```

2. 執行程式：

   ```sh
   ./to-slug "需要翻譯的文字"
   ```

3. 顯示使用說明：

   ```sh
   ./to-slug -h
   ```

## 錯誤處理

1. 如果 `OPENAI_API_KEY` 未設置，程式將輸出錯誤訊息並退出。
2. 如果 OpenAI API 請求失敗，程式將輸出錯誤訊息並返回空字串。

## 程式結構

- `main.go`: 主程式文件，包含主要邏輯和功能實現。
  - `init()`: 初始化函數，檢查環境變數並創建 OpenAI 客戶端。
  - `main()`: 主函數，解析命令行參數並調用翻譯和轉換函數。
  - `translateToSlug(text string) string`: 將文字翻譯為簡單的口語英文並轉換為 URL slug。

## 測試

1. 測試輸入不同的文字，檢查輸出是否符合預期。
2. 測試錯誤情況，例如未設置 `OPENAI_API_KEY` 或 API 請求失敗。
