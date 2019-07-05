# Redis Note
<br>

## Redis 筆記
#### Reference [Triton Redis slide](https://github.com/TritonHo/slides/blob/master/Taipei%202019-06%20talk/redis-2019.pdf?fbclid=IwAR0e2eyuy4kFkYIwGAZEITK3kyirEKbHMvKSldpPPLrm4GB0pAbR4Lv2nRg)

**吃飯定理:**
- 便宜, 好吃, 不用排隊 正常餐廳只能滿足其中兩項
- 同時滿足三項者, 最終一定虧本倒閉

**Redis為單線程**
- Perdidtence
  * 在當機後不會引發資料流失
- Low latency
  * 資料庫能用極短時間完成單一工作
-  以上兩者最多只能要一個
-  Redis 是追求 Low latency, 想單用redis zero data loss, <font color="#ff0000">impossible</font>
#### Redis 的預設, 是每一萬個 write 才會寫入 hardisk , 若 redis 當機,一定會有 data loss 的  
- 使用 Redis, 用作caching, 資料同時存放於主資料庫
- 儲存沒了也死不了的 Hot Data

**Single Key Consistency** (Redis cluster)
- 不同的keyvalue 會放到不同的redis機器
  * 除非只跑single node
- 要使用樂觀鎖時, 無可避免要用上Hash
  * Data 來存資料
  * LastUpdate TS 來存最後改動時間

**Cache 常犯錯誤**
- 只用local cache
  * 指 Application Server(AS)上的local memory
  * 缺點: 改動沒辦法反映到所有的伺服器上 & 新開的AS, 其local cache 是全空的
  * 指 Application Server(AS)上的local memory
- 只用[一般]的caching
  * 1. 從Redis 拿資料, 如有則直接回傳
  * 2. 從主資料拿資料X
  * 3. 把資料x 放回Redis
  * 4. 回傳
  * <font color="#ff0000">看起來很正常, 但在高流量下實際上必須要</font>
  * 1. 從Redis 拿資料, 如有則直接回傳
  * 2. **拿到資料x的鎖 (在離開時釋放)**
  * 3. **再次從Redis 拿資料x, 如有則直接釋放**
  * 4. 從主資料拿資料X
  * 5. 把資料x 放回Redis
  * 6. 回傳    

- 沒使用consistency hash
  * 別使用mod來決定某一key value 位置, 用這方法, 當系統繁忙要加開redis時, 會讓caching 全滅
  * 使用 consistent hash
- 沒對hot data預熱
  * 所某一排行榜需5s 才能生產出來, 一旦cache miss, 一堆人需要等待這份資料 => slow
  * 寫一個 crontab, 在cache miss前到DB先拿資料放到redis
- 沒設定合理的TTL
  * 瘋狂的TTL, 最終會讓redis 存太多過期資料, 觸發cache eviction
  * 在 peak hour 的寫入觸發, 須先清出空間才能寫入, 引發超高latency

**把Redis 當成簡陋版 lock server**
- 專業應用 => Zookeeper / etcd

**Anti-pattern: Barrier**
- 用在 Redis 會害你失掉system robustness

**Barrier vs lock**
- 雖然兩者都適用SETNX, 目的完全不同
- SETNX 回答0時
  - Barrier是直接return
  - Lock 是 sleeping再重試
- SETNX 的 TTL
  - Barrier是長時間的(資料更新的隔距)
  - Lock 是短時間的 (critical zone的最大執行時間)

**Ratelimiting**
- Nginx 的 ratelimiting 是以API為單位的, 對一般系統其實很夠用了
- 當量級到C50k, redis系統很大機率會陣亡, 可以善用 local buffering, 減輕redis工作量, 若是application server 拿不到資料, 便乾脆拒絕所有同類request 1s

**data snapshotting**
- 一個持續改動中的排行榜, 沒有snapshotting, 剛好第10名的物品變第11, 去拿page2的資料, 就會看到重複的資料
- <font color="#ff0000">Snapshotting with redis V1</font>
  - 建立crontab, 每一分從資料庫建立新的snapshot, 丟到redis中, ex: DATAxxx-20190710:1200
- <font color="#ff0000">Snapshotting with redis V2</font>
  - 若資料長期沒有更動, V1只會建立大量重複的snapshots, 如果一份snapshots 跟之前相同, 他只需把該ts存起來, 而不用存相同的資料
- snapshot 是永遠不會改動的資料, 其localcache TTL 應跟 Redis 相同

