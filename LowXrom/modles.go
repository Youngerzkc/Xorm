package main
import (
	"fmt"
	"errors"
	_"github.com/mattn/go-sqlite3" //导入驱动包
	"github.com/go-xorm/xorm"
	"log"
)
type Account struct{
	Id int64 
	Name string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`//乐观锁
}
var x *xorm.Engine
func  init()  {
	var  err  error
	x,err =xorm.NewEngine("sqlite3","./bank.db")
	if err!=nil{
		log.Fatalf("failed to create enige! %v",err)
	}
	if err=x.Sync(new(Account));err!=nil{
		log.Fatalf("falied to %v",err)
	}
}
func newAccount(name string,balance float64 )error{
  _,err:=x.Insert(&Account{Name:name,Balance:balance})
   return err
}
func getAccount(id int64)(*Account ,error ){
	a:=&Account{}
	has,err:=x.Id(id).Get(a)//a-->地址
	if err!=nil{
		return nil,err
	}else if !has{
		return nil,errors.New("Account not found")
	}
	return a,nil
}
func makeDeposit(id int64,balance float64)(*Account,error){
	a,err:=getAccount(id)
	if err!=nil{
		return nil,err
	}
		a.Balance+=balance
		_,err =x.Id(id).Update(a) //更新xrom数据问题，x.Update(a)---->不成功
		fmt.Println("+++++++++",err)
		return a,err
}
func  makeWithDraw(id int64,balance float64)(*Account,error){
	a,err:=getAccount(id)
	if err!=nil{
		return nil,err
	}
	if a.Balance < balance{
		return nil,errors.New("balance not enough")
	}
	a.Balance-=balance 
	_,err=x.Id(id).Update(a)
	return a,err
}
func  makeTransFer(id1 int64,balance float64,id2 int64) error  {
	a1,err:=getAccount(id1)
	if err!=nil{
		return err
	}
	a2 ,err:=getAccount(id2)
	if err!=nil{
		return err
	}
	if a1.Balance<balance {
		return errors.New("balance not enough")
	}
	a1.Balance-=balance
	a2.Balance+=balance
	if _,err:=x.Id(id1).Update(a1);err!=nil{   //发生错误，要进行事物回滚
		return err
	}else if _,err:=x.Id(id2).Update(a2);err!=nil{
		return err
	}
	return nil
}
func  getAccountAscId()(as []*Account,err error)  {
	err=x.Asc("id").Find(&as)//id 为数据库的字段 ，as err
	return as,err
}
func  deleteAccount(id int64) error {
	_,err:=x.Delete(&Account{Id:id})//删除所有符合条件的记录，可加where
	return err
}
