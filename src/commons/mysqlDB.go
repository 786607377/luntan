package commons

import (
    "database/sql"
    "fmt"
    _ "github.com/Go-SQL-Driver/MySQL"
)

//数据库操作的三个对象
var(
    db *sql.DB
    stmt *sql.Stmt
    rows *sql.Rows
)

//打开数据库链接
func openConn()(err error){
    db,err =sql.Open("mysql","root:mmk23887155@tcp(localhost:3306)/ego")
    if err!=nil{
        fmt.Println("连接失败",err )
        return
    }
    return nil
}

//关闭连接,首字母大写,需要跨包访问的
func CloseConn(){
    if rows!=nil{
        rows.Close()
    }
    if stmt!=nil{
        stmt.Close()
    }
    if db!=nil{
        db.Close()
    }
}
//执行DMl新增删除修改操作
func Dml(sql string,args ...interface{})(int64,error){
    err:=openConn()
    if err!=nil{
        fmt.Println("执行DML时出现错误,打开连接失败")
        return 0,err
    }
    //此处也是等号
    stmt,err=db.Prepare(sql)
    if err!= nil{
        fmt.Println("执行DML出现错误,预处理出现错误")
        return 0,err
    }
    //此处要有...标识切片,如果没有表示数组会报错
    result,err:=stmt.Exec(args...)
    if err!=nil {
        fmt.Println("执行DML时出现错误,执行错误")
        return 0,err
    }
    //受影响的行数
    count,err:=result.RowsAffected()
    if err !=nil{
        fmt.Println("执行DML时出现错误,获取受影响行数错误")
        return 0,err
    }
    defer CloseConn()
    return count,err
}

//执行DQL查询
func Dql(sql string,args ...interface{})(*sql.Rows,error){
    err:=openConn()
    if err!=nil{
        fmt.Println("执行DQL时出现错误,打开连接失败")
        return nil,err
    }
    stmt,err = db.Prepare(sql)
    if err !=nil{
        fmt.Println("执行DQL出现错误,预处理出现错误")
        return nil,err
    }
    rows,err = stmt.Query(args...)
    if err !=nil{
        fmt.Println("执行DQL出现错误,执行出现错误")
        return nil,err
    }
    //此处没有关闭以后调用函数时要记得关闭连接
    //defer CloseConn()
    return rows,err
}
