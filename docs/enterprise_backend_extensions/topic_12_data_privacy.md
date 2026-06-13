# Topic 12: 数据隐私、审计与合规意识

## 适合插入位置

课程末尾或期末项目前。

## 核心问题

后端系统处理的不只是数据结构，也是用户数据。学生需要早一点建立隐私和审计意识。

## 必讲点

- 最小化收集：不需要就不收集
- 脱敏展示：手机号、邮箱、身份证号
- 敏感字段不进日志
- 管理员操作要审计
- 数据删除和数据备份存在冲突，需要制度设计

## 练习

指出下面日志的问题：

```text
login failed username=alice password=123456 token=eyJ...
```

要求学生改成：

```text
login failed username=alice reason=invalid_credentials request_id=req-001
```
