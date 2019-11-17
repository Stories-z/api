# 启动服务器方式

`$ cd src/main && go run main.go`

# 访问主资源方式

命令行内输入：

`$curl "http://localhost:8080/testuser"`

可获得json格式反馈：

`{`
  `"Current_user_url": "https://localhost:8080/testuser",`
  `"Emails_url": "https://localhost:8080/testuser/emails",`
  `"Repository_url": "https://localhost:8080/repos/{owner}/{repo}"`
`}`

# 访问子资源方式

命令行内输入：

`$curl "http://localhost:8080/repos/testuser/api"`

可获得json格式反馈：

`{`
  `"Owener": "testuser",`
  `"Name": "api"`
`}`

