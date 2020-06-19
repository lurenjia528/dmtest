# go 操作dm数据库

## windows
windows上连接dm,安装dm时默认安装配置了odbc,所以可以直接连接

## linux
1. go get odbc 测试,缺少sql.sh即需要安装odbc
```bash
$go get github.com/alexbrainman/odbc
go: writing go.mod cache: mkdir /home/lurenjia/gopath/pkg/mod/cache/download/github.com/jackc/fake: permission denied
go: writing go.mod cache: mkdir /home/lurenjia/gopath/pkg/mod/cache/download/github.com/jackc/pgx/@v: permission denied
go: writing go.mod cache: open /home/lurenjia/gopath/pkg/mod/cache/download/github.com/denisenkom/go-mssqldb/@v/v0.0.0-20190707035753-2be1aa521ff4.mod496193015.tmp: permission denied
go: downloading github.com/alexbrainman/odbc v0.0.0-20200426075526-f0492dfa1575
go: downloading golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a
# github.com/alexbrainman/odbc/api
../../../../gopath/pkg/mod/github.com/alexbrainman/odbc@v0.0.0-20200426075526-f0492dfa1575/api/api_unix.go:14:11: fatal error: sql.h: No such file or directory
 // #include <sql.h>
           ^~~~~~~
compilation terminated.
```

2. 安装odbc 
```bash
 apt-get install unixodbc unixodbc-bin unixodbc-dev
```

3. 编写驱动信息

```bash
root@lurenjia:# cat /etc/odbcinst.ini
[DM8 ODBC DRIVER]
Description = ODBC DRIVER FOR DM8
Driver = /opt/dmdbms/drivers/odbc/libdodbc.so
root@lurenjia:# cat /etc/odbc.ini
[dm]
Description = DM ODBC DSN
Driver = DM8 ODBC DRIVER
SERVER = localhost
UID = SYSDBA
PWD = SYSDBA
TCP_PORT = 5236

```

4. 测试连接

```bash
 isql -v dm SYSDBA SYSDBA
+---------------------------------------+
| Connected!                            |
|                                       |
| sql-statement                         |
| help [tablename]                      |
| quit                                  |
|                                       |
+---------------------------------------+
SQL>
```
结果如上图,即为测试成功

可参考《DM程序员手册》