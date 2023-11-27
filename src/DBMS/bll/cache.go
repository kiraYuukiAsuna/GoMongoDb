package bll

import (
	"DBMS/dal"
	"DBMS/dbmodel"
	"fmt"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"strconv"
	"time"
)

var DailyStatisticsInfo dbmodel.DailyStatisticsMetaInfoV1

type OnlineUserInfo struct {
	UserInfo          dbmodel.UserMetaInfoV1
	expired           bool
	LastHeartBeatTime time.Time
}

var OnlineUserInfoCache []OnlineUserInfo

func CronAutoSaveDailyStatistics() {
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)), cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	EntryID, err := c.AddFunc("@daily", func() {
		fmt.Println(time.Now(), "CronAutoSaveDailyStatistics...")
		timePoint := time.Now()
		year, mouth, day := timePoint.Date()
		DailyStatisticsInfo.Base.Id = primitive.NewObjectID()
		DailyStatisticsInfo.Base.Uuid = uuid.NewString()
		DailyStatisticsInfo.Base.ApiVersion = "V1"
		DailyStatisticsInfo.Name = strconv.Itoa(year) + "-" + mouth.String() + "-" + strconv.Itoa(day) + strconv.Itoa(int(DailyStatisticsInfo.ProjectQueryNumber))
		DailyStatisticsInfo.Day = strconv.Itoa(year) + "-" + mouth.String() + "-" + strconv.Itoa(day)
		DailyStatisticsInfo.Description = "Auto Generated by CronAutoSaveDailyStatistics."
		fmt.Println("TimerEvent: Time - " + timePoint.String() + " , Event - CronAutoSaveDailyStatistics performed.")
		DailyStatisticsInfo.ProjectQueryNumber += 1

		dal.CreateDailyStatistics(DailyStatisticsInfo, dal.GetDbInstance())
	})
	fmt.Println(time.Now(), EntryID, err)
	c.Start()
}

func CronHeartBeatValidationAndRefresh() {
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)), cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	EntryID, err := c.AddFunc("*/30 * * * * *", func() {
		fmt.Println(time.Now(), "CronHeartBeatValidationAndRefresh...")
		for idx, onlineUserInfo := range OnlineUserInfoCache {
			if time.Now().After(onlineUserInfo.LastHeartBeatTime) || onlineUserInfo.expired {
				onlineUserInfo.expired = true
				OnlineUserInfoCache = append(OnlineUserInfoCache[:idx], OnlineUserInfoCache[idx+1:]...)
				fmt.Println("User " + onlineUserInfo.UserInfo.Name + " HeartBeat expired")
			}
		}
	})
	fmt.Println(time.Now(), EntryID, err)
	c.Start()
}

//func TestCron() {
//	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)), cron.WithLogger(
//		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
//	i := 1
//	EntryID, err := c.AddFunc("*/5 * * * * *", func() {
//		fmt.Println(time.Now(), "每5s一次:", i)
//		time.Sleep(time.Second * 6)
//		i++
//	})
//	fmt.Println(time.Now(), EntryID, err)
//
//	c.Start()
//}
