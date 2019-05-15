## 编码问题

如果爬取的数据不是 `utf-8` 编码，则爬虫在 `Download` 模块中自动会捕获 `header` 中 `charset` 字段，进而使用`iconv`包进行编码转换。