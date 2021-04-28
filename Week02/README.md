
## 作业

Q：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

A：
1. 根据我的理解，sql.ErrNoRows其实不算错误，只需要在dao层返回表示找不到内容的提示符号即可，如nil；根据习惯，也可能返回not found之类比较友好点的提示，这样也可以区分不同的异常类型。
2. 其他sql.Err类型，应该往上层抛，为了定位问题，使用`errors.Wrapf`保存堆栈信息，并记入日志中。

第一次写go代码，确实有点吃力。

正常访问：http://localhost:9999/article?id=1
```
Id：%!s(int64=1)，Title：test
```

sql.ErrNoRows：http://localhost:9999/article?id=2
```
Article not found , id:2
```

