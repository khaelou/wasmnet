package pool

import (
	"fmt"
	"runtime"
)

func InitClient(payloadAmt int) {
	fmt.Println(`
                                    ,dPYb,
                                    IP''Yb
                                    I8  8I
                                    I8  8'
 gg,gggg,      ,ggggg,    ,ggggg,   I8 dP 
 I8P"  "Yb    dP"  "Y8gggdP"  "Y8gggI8dP  
 I8'    ,8i  i8'    ,8I i8'    ,8I  I8P   
,I8 _  ,d8' ,d8,   ,d8',d8,   ,d8' ,d8b,_ 
PI8 YY88888PP"Y8888P"  P"Y8888P"   8P'"Y88
 I8                                       
 I8                                       
 I8                                       
 I8                                       
 I8                                       
 I8                                                           									
 	`)

	runtime.GOMAXPROCS(4)

	fmt.Println("\t** CLIENT INIT **")
	fmt.Println()

	InitWorkload(payloadAmt) // Payload Amount
}
