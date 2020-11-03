# Linux instructions

##### Bash scripting cheatsheet: https://devhints.io/bash

<hr>

##### Linux 指令
```
mkdir: 創建資料夾, 若帶 -p 參數, 要建立的目錄已經存在時, 就不會報錯

&: 如果在指令後面加上 &, 即表示指令在背景執行, 例如 script.sh &
&&: 用於分開兩個指令, 即第一道指令執行成功後, 才會執行第二道指令, 例如 mkdir dir && cd dir
|: 管線的符號, 將前一道指令的輸出, 作為第二道指令的輸入, 例如 kgp | grep redis
||: 表示前一道指令執行失敗後, 才會執行下一道指令, 例如 cat test.py || echo 'No file'
```

##### Linux user group 相關
```
// 新建 group
groupadd {group}

// 新建 user
useradd {username}  // 不會在/home下建立一個資料夾username
adduser {username}  // 會在/home下建立一個資料夾username

// 幫 user 增加 groups
adduser -G root,google-sudoers {user}

// 查看 user 的 group
groups {user}

今天若想要使用 ssh 連線至機器上,
可至 /home/{user}/.ssh 資料夾底下增加 authorized_keys 檔案,
內容為 xxx.pub, ex: ssh-rsa AAAAXXXXXXXXXXXXXB3TZTu1TNM6fPulE= devops@gmail.com

之後即可帶著 private key 連線進去: ssh -i files/keys/devops  {user}@{ip}
```

##### shell redirect 相關
```
file descriptor: 0 表示鍵盤輸入; 1 表示標準輸出; 2 表示錯誤輸出

> 默認為標準輸出重定向, 同 1>

2>&1: 標準錯誤輸出重定向到標準輸出

&>file: 把標準輸出和標準錯誤輸出, 都重定向到文件 file 中

/dev/null是一個特殊文件, 所有傳給它的東西它都丟棄掉
```