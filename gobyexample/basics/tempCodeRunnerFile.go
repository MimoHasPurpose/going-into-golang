select{
		case msg1:=<-c1:
			fmt.Println("received",msg1)
		case msg2:=<-c2:
			fmt.Println("received",msg2)
		}