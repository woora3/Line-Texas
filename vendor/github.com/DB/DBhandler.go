/*
使用者是否在遊戲中
他手上有哪兩張牌
*/
package DB
import (
	"os"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
/*
傳入使用者MID
回傳使用者是否正在遊戲
*/
func UserGamming(MID string) bool{
	var GameID int
	GameID = 0;
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	db.QueryRow("SELECT GameID FROM sql6131889.GameAction WHERE MID = ? and Cancel = 0", MID ).Scan(&GameID)
	db.Close()
	if GameID == 0{
		return false
	}else{
		return true
	}
}

//Call WHEN PlayerToken ADD OR SUB
func AddPlayerToken(MID string,addtoken int){
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	db.QueryRow("UPDATE sql6131889.User SET UserToken=UserToken+? WHERE MID =?",addtoken,MID)
	db.Close()
}

//Call WHEN GAMETOKEN ADD OR SUB
func AddGameToken(RoomId int,addtoken int){
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	db.QueryRow("UPDATE sql6131889.Game SET GameToken=GameToken+? WHERE RoomID =?",addtoken,RoomId)
	db.Close()
}
