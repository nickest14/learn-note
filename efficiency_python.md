# Efficiency Python Note
<hr>

**1. Python Thinking**
- Item 1 : 知道自己是在使用哪個版本的 python
- Item 2 : 遵循 <font color=#ff0000>PEP8</font> 程式碼風格
- Item 3 : 了解 <font color=#ff0000>bytes</font> 及 <font color=#ff0000> str </font> 的差別
    ```
    >>> "nick".encode("utf-8")
    b'nick'
    >>> b"nick".decode("utf-8")
    'nick'
    ```

- Item 4 : 使用 F-Strings 取代 C-style format strings
- Item 5 : 寫 helper functions 取代複雜的賦值表達式
    ```
    # bad
    red = int(my_values.get('red', [''])[0] or 0)
    
    # better
    def get_first_int(values, key, default=0):
        found = values.get(key, [''])
        if found[0]:
            return int(found[0])
        return default
    ```

- Item 6 : 使用 multiple assignment unpacking 取代 indexing
- Item 7 : 使用 enumerate 取代 range
    ```
    snacks = [('bacon', 350), ('donut', 240), ('muffin', 190)]

    # bad
    for i in range(len(snacks)):
        item = snacks[i]
        name = item[0]
        calories = item[1]
        print(f'#{i+1}: {name} has {calories} calories')

    # better
    for rank, (name, calories) in enumerate(snacks, 1):
        print(f'#{rank}: {name} has {calories} calories')
    ```
    * enumurate 可讓你達到for loop 的效果並同時拿到 index位置
    * 可以指定 enumurate 第2個參數當作起始值, 預設是0
- Item 8 : 使用 zip 來處理平行的遍歷
    ```
    names = ['Cecilia', 'Lise', 'Marie']
    counts = [len(n) for n in names]
    for name, count in zip(names, counts):
        print(name, count)
    ```
    * 若是長度不平衡則可以使用 itertools 的 zip_longest function

- Item 9 : 在 for 或 while loop 後避免使用 else block
    * 不建議使用是因為不直覺, 且易造成混淆
- Item 10 : 用賦值表達式避免重複
    * Python 3.8 後可使用 walrus operator
    ```
    # old
    count = dict.get('lemon', 0)
    if count:
        ...
    else:
        ...

    #  new
    if count := dict.get('lemon', 0):
        ...
    else:
        ...
    ``` 
**2. List and Dictionaries**
- Item 11 : 知道如何使用 slice
- Item 12 : 在一個表達式中避免同時用 striding 及 slicing
  * start, end 及 stride 同時使用易造成混淆, 建議分開使用(用兩次表達式)

**3. Functions**
