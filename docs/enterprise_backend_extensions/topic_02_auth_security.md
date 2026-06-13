# Topic 02: 认证、密码与密钥管理

## 适合插入位置

Module 03 JWT 中间件之后。

## 要解决的问题

当前课程项目中有明文密码和硬编码 JWT secret，这是教学简化。拓展课要让学生知道真实项目为什么不能这么做。

## 核心概念

- 密码不加密保存，而是保存不可逆 hash。
- JWT secret 不写死在代码里，从环境变量或配置系统读取。
- Access Token 有过期时间。
- 鉴权和认证不同：认证回答“你是谁”，鉴权回答“你能做什么”。
- 不要在日志里打印密码、token、身份证号、手机号等敏感信息。

## 课堂 Demo

把 `project_user_center` 的注册逻辑从明文密码改成 bcrypt hash：

```go
hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
if err != nil {
	return nil, err
}
user.Password = string(hashed)
```

登录时：

```go
if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
	return "", ErrInvalidCredentials
}
```

## 练习

让学生找出当前项目中所有硬编码 secret/token/password，并按“教学可接受 / 必须改造”分类。

## 讨论题

为什么生产系统里“忘记密码”通常不能把原密码发给用户？
