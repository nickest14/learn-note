# Nginx Note 
<br>

**Rewrite**
- last
  * 終止 rewrite, 發起新的內部請求, 重新開始匹配 location
- break
  * 終止 rewrite, 不發起新的並繼續此 location 的後續處理階段
- redirect
  * 返回客戶端 302, 使其暫時使用新網址再發送一次請求
- permanent
  * 返回客戶端 301, 使其使用新網址再發送一次請求，並讓客戶端知道原本的網址已經被永久轉移到新網址