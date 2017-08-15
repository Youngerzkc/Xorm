package main

import (
	"fmt"
)
const promt=`Please enter of number of operation
1.Create new account
2.show detail of account
3.Deposite
4.Withdraw
5.Make transfer
6.List account By Id
7.List account By balances
8.Delete account
9.Exit`
func  main()  {
	fmt.Println("welcome to bank of xorm")
Exit:
	for{
		fmt.Println(promt)
		var num int 
		fmt.Scanf("%d\n",&num)
		switch num {
		case 1:
		fmt.Print("please enter <name> <balances>\n")
		var name string
		var balances float64
		fmt.Scanf("%s %f\n",&name,&balances)
		if err:=newAccount(name, balances);err!=nil{
			fmt.Println(err)
		 }
		case 2:
		fmt.Print("please enter <id>\n")
		var id int64
		fmt.Scanf("%d",&id)
		a,err:=getAccount(id)
		if err!=nil{
			fmt.Println("+++++++++++")
			fmt.Println(err)
		}else{
			fmt.Println("---------------")
			fmt.Printf("%#v",a)//打印结构体所有信息
		}
		case 3:
		fmt.Println("please enter <id> <balance> ")
		var id int64
		var balance float64
		fmt.Scanf("%d %f\n",&id,&balance)
		a,err:=makeDeposit(id, balance)
		if err!=nil{
			fmt.Println(err)
		}else{
			fmt.Printf("%#v", a)
		}
		case 4:
		fmt.Println("please enter <id> <balance> ")
		var id int64
		var balance float64
		fmt.Scanf("%d %f\n",&id,&balance)
		a,err:=makeWithDraw(id, balance)
		if err!=nil{
			fmt.Println(err)
		}else{
			fmt.Printf("%#v", a)
		}
		case 5:
		fmt.Println("please enter <id1> <balance> <id2> ")
		var id1 int64
		var id2 int64
		var balance float64
		fmt.Scanf("%d %f %d\n",&id1,&balance,&id2)
		err:=makeTransFer(id1, balance ,id2)
		if err!=nil{
			fmt.Println(err)
		}
		case 6:
		as,err:=getAccountAscId()
		if err!=nil{
			fmt.Println(err)
		}else{
			for i,a:=range as{
				fmt.Printf("%d:%#v",i,a)
			}
		}
		case 7:
		case 8:
 		 fmt.Println("delete Account")
		 fmt.Println("please enter <id>")
		 var  id int64
		 fmt.Scanf("%d\n",&id)
		 err:=deleteAccount(id)
		 if err!=nil{
			 fmt.Println(err)
		 }		 
		 case 9:	
		  break Exit	//break??
		}
	}
}