package api

import (
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	ID         uint      `gorm:"primaryKey; autoIncrement" json:"id,omitempty"`
	IsVerified bool      `gorm:"default:0;column:isVerified" json:"isVerified,omitempty"`
	Name       string    `gorm:"type:VARCHAR(40);not null" json:"name,omitempty"`
	Email      string    `gorm:"type:VARCHAR(50);not null;unique" json:"email,omitempty"`
	Password   string    `gorm:"type:VARCHAR(60);not null" json:"password,omitempty"`
	Username   string    `gorm:"type:VARCHAR(25);not null;unique" json:"username,omitempty"`
	Gender string `gorm:"type:VARCHAR(1);default:'F'" json:"gender"`
	About string `gorm:"type:VARCHAR(200)" json:"about,omitempty"`
	ProfilePic string `gorm:"column:profilePic;default:null;type: VARCHAR(50)" json:"profilePic"`
	IsAuth bool `gorm:"-" json:"isAuth"`
	Channels []Channel `gorm:"many2many:userChannels;" json:"channelIds,omitempty"`
	Messages []Message `gorm:"foreignKey:senderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time `gorm:"column:createdAt" json:"createdAt,omitempty"`
	UpdatedAt  time.Time `gorm:"column:updatedAt" json:"updatedAt,omitempty"`
	
}

type Verification struct {
	ID        string `gorm:"primaryKey;type:VARCHAR(9) NOT NULL" json:"id"`
	UserId    uint
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ExpiresAt time.Time `gorm:"NOT NULL;column:expiresAt" json:"expiresAt"`
}

type Channel struct {
	ID uint `gorm:"primaryKey; autoIncrement" json:"id,omitempty"`
	IsGroup bool `gorm:"default:0;column:isGroup" json:"isGroup"`
	Messages []Message `gorm:"foreignKey:channelId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"messages"`
	Members []*User `gorm:"many2many:userChannels;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members,omitempty"`
}

type Message struct {
	ID string `gorm:"primaryKey" json:"id,omitempty"`
	FrontendId int `gorm:"-" json:"frontendId"`
	Text string `gorm:"type:VARCHAR(10000) NOT NULL" json:"text"`
	Sent bool `gorm:"default:0" json:"sent"`
	Received bool `gorm:"default:0" json:"received"`
	Read bool `gorm:"default:0" json:"read"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt,omitempty"`
	SenderId uint `json:"senderId" gorm:"column:senderId;"`
	Sender User `json:"-"`
	ChannelId uint `json:"channelId" gorm:"column:channelId;"`
	Channel Channel `json:"-"`
	Resend bool `gorm:"-" json:"resend"`
	Connection *websocket.Conn `gorm:"-" json:"-"`
	IsNewChannel bool `gorm:"-" json:"isNewChannel"`
}


type Acknowledge struct{
	MessageId string `json:"messageId"`
	FrontendId int `json:"frontendId"`
	CreatedAt time.Time `json:"createdAt"`
	SenderId uint `json:"senderId"`
	ChannelId uint `json:"channelId"`
	AckType string `json:"ackType"`
}

type UserStatus struct{
	Status string `json:"status"` // online offline typing <empty string>
	UserId uint `json:"userId"`
	ChannelId uint `json:"channelId"` 
}

type StreamData struct{
	Type string `json:"type"` // MSG ACK USR_STAT
	Message Message `json:"message,omitempty"`
	Acknowledge Acknowledge `json:"acknowledge,omitempty"`
	UserStatus UserStatus `json:"userStatus,omitempty"`
}