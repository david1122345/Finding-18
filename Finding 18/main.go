package main

import "fmt"
import "math/rand"

func main(){
	turn:=1
	states:=[11]int8{}
	var num int;
	var difficulty int;
	fmt.Println("1| Easy")
	fmt.Println("2| Medium")
	fmt.Println("3| Hard")
	fmt.Println("4| Impossible")
	fmt.Print("Select a dificulty level(the number) ")
	fmt.Scan(&difficulty)
	var rand_num int;
	for(num!=12 && turn<13){
		rand_num=rand.Intn(3)+1
		if difficulty<=rand_num{
			states[play_random(states,turn)]=1;
		}else{
			states[playBot(states,turn)]=1;
		}
		if is_win(states,1)&& turn>4{
			display_board(states)
			fmt.Println("you loose :(")
			break
		}
		display_board(states)
		fmt.Print("Choose a number ")
		fmt.Scan(&num)
		states[num-1]=-1
		if is_win(states,-1)&& turn>4{
			display_board(states)
			fmt.Println("you win :)")
			break
		}
		turn+=2	
	}
}
func is_win(board [11]int8,player int8)bool{
	var i int
	var j int
	nums:=[]int{}
	for i=0;i<11;i++{
		if board[i]==player{
			nums=append(nums,i+1)
		}	
	}
	for i=0;i<len(nums);i++{
		for j=0;j<len(nums);j++{
			if nums[i]+nums[j]>=7 && nums[i]+nums[j]<=17 && i!=j{
				if board[17-nums[i]-nums[j]]==player &&(18-nums[i]-nums[j])!=nums[i] && (18-nums[i]-nums[j])!=nums[j]{
					return true
				}
			}
		}
	}
	return false
}
func display_board(board [11]int8){
	nums:=[]int{}
	var i int;
	for i=0;i<11;i++{
		if board[i]==1{
			nums=append(nums,i+1)
		}
	}
	fmt.Print("bot: ")
	fmt.Println(nums)
	nums=[]int{}
	for i=0;i<11;i++{
		if board[i]==-1{
			nums=append(nums,i+1)
		}
	}
	fmt.Print("you: ")
	fmt.Println(nums)
}
func playBot(board [11]int8,turn int)int{
	w:=find_win(board)
	if w!=12{
		return w
	}
	if turn==1{
		return 5
	}
	if turn==3{
		if board[6]==-1 || board[4]==-1{
			return 7
		}else{
			return 6
		}
	}
	if turn==5{
		sum:=0
		var i int;
		for i=0;i<11;i++{
			if board[i]==-1{
				sum+=i+1
			}
		}

		if sum>=7 && sum<18{
			if board[17-sum]==0{return 17-sum}
		}
		if board[4]==0{
			return 4
		}else if board[3]==0{
			return 3
		}
	}
	nums:=[]int{}
	var i int;
	for i=0;i<11;i++{
		if board[i]==1{
			nums=append(nums,i+1)
		}
	}
	if nums[1]+nums[0]>=7 && nums[1]+nums[0]<18{
		if board[17-nums[0]-nums[1]]==0{
			return 17-nums[0]-nums[1]
		}
	}
	if nums[2]+nums[0]>=7 && nums[0]+nums[2]<18{
		if board[17-nums[0]-nums[2]]==0{
			return 17-nums[0]-nums[2]
		}
	}
	if nums[1]+nums[2]>=7 && nums[1]+nums[2]<18{
		if board[17-nums[1]-nums[2]]==0{
			return 17-nums[2]-nums[1]
		}
	}
	for i=0;i<11;i++{
		if board[i]==-1{
			nums=append(nums,i+1)
		}
	}
	if nums[1]+nums[0]>=7 && nums[1]+nums[0]<18{
		if board[17-nums[0]-nums[1]]==0{
			return 17-nums[0]-nums[1]
		}
	}
	if nums[2]+nums[0]>=7 && nums[0]+nums[2]<18{
		if board[17-nums[0]-nums[2]]==0{
			return 17-nums[0]-nums[2]
		}
	}
	if nums[1]+nums[2]>=7 && nums[1]+nums[2]<18{
		if board[17-nums[1]-nums[2]]==0{
			return 17-nums[2]-nums[1]
		}
	}
	if board[3]==0{return 3}
	if board[4]==0{return 4}
	if board[6]==0{return 6}
	if board[7]==0{return 7}
	nums=[]int{}
	var j int
	for i=0;i<11;i++{
		if board[i]==-1{
			nums=append(nums,i+1)
		}	
	}
	for i=0;i<len(nums);i++{
		for j=0;j<len(nums);j++{
			if nums[i]+nums[j]>=7 && nums[i]+nums[j]<=17 && i!=j{
				if board[17-nums[i]-nums[j]]==0 &&(18-nums[i]-nums[j])!=nums[i] && (18-nums[i]-nums[j])!=nums[j]{
					return 17-nums[i]-nums[j]
				}
			}
		}
	}
	return 13// intetianal crash
}
func find_win(board [11]int8)int{
	var i int
	var j int
	nums:=[]int{}
	for i=0;i<11;i++{
		if board[i]==1{
			nums=append(nums,i+1)
		}	
	}
	for i=0;i<len(nums);i++{
		for j=0;j<len(nums);j++{
			if nums[i]+nums[j]>=7 && nums[i]+nums[j]<=17 && i!=j{
				if board[17-nums[i]-nums[j]]==0 &&(18-nums[i]-nums[j])!=nums[i] && (18-nums[i]-nums[j])!=nums[j]{
					return (17-nums[i]-nums[j])
				}
			}
		}
	}
	return 12
}
func play_random(board [11]int8,turn int)int{
	num:=rand.Intn(11-turn)
	var i int;
	for i=0;i<num;i++{
		if board[i]!=0{i++}
	}
	if board[i]==0{return i}
	for board[i]!=0{i++}
	return i
}