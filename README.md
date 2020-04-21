# 学习园地

## 开发环境

1. 安装model代码生成工具

   ```bash
   go install  github.com/volatiletech/sqlboiler
   go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
   ```

2. 修改生成配置`sqlboiler.toml`中的账号密码配置
3. 执行脚本`generate-table.sh`生成数据库操作代码
4. 使用`model/table`路径下生成的代码
