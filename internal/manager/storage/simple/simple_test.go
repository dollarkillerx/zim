package simple

import (
	"github.com/dollarkillerx/common/pkg/conf"
	"github.com/dollarkillerx/zim/pkg/utils"

	"log"
	"testing"
)

var storage *SimpleStorage

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	st, err := NewSimpleStorage(conf.PostgresConfiguration{
		Host:     "localhost",
		Port:     5432,
		User:     "root",
		Password: "root",
		DBName:   "zim",
		SSLMode:  false,
	}, conf.RedisConfiguration{
		Addr:     "localhost:6379",
		Db:       0,
		Password: "root",
	})
	if err != nil {
		panic(err)
	}

	storage = st
}

func TestSimpleStorage_SuperAdmin(t *testing.T) {
	superAdmin, err := storage.SuperAdminCreate()
	if err != nil {
		panic(err)
	}
	utils.PrintObject(superAdmin)

	reset, err := storage.SuperAdminReset(superAdmin.SupID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(reset)

	sa, err := storage.SuperAdminGetBySuperAdminID(superAdmin.SupID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(sa)

	err = storage.SuperAdminDel(sa.SupID)
	if err != nil {
		panic(err)
	}

	sa, err = storage.SuperAdminGetBySuperAdminID(superAdmin.SupID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(sa)
}

func TestSimpleStorage_Project(t *testing.T) {
	sa, err := storage.SuperAdminCreate()
	if err != nil {
		panic(err)
	}

	project, err := storage.ProjectCreate(sa.SupID, "测试项目")
	if err != nil {
		panic(err)
	}
	utils.PrintObject(project)

	projects, err := storage.ProjectList(sa.SupID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(projects)

	reset, err := storage.ProjectReset(project.SupID, project.ProjectID, "测试项目modify")
	if err != nil {
		panic(err)
	}
	utils.PrintObject(reset)

	err = storage.ProjectDelete(project.ProjectID)
	if err != nil {
		panic(err)
	}

	projects, err = storage.ProjectList(sa.SupID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(projects)
}

func TestSimpleStorage_User(t *testing.T) {
	projectID := "cfsre4hb1dspa5lbimu0"
	us1, err := storage.UserCreate(projectID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(us1)

	us2, err := storage.UserCreate(projectID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(us2)

	err = storage.UserRelevance(projectID, us1.UserID, us2.UserID)
	if err != nil {
		panic(err)
	}

	total, friends, err := storage.UserFriendsList(projectID, us1.UserID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(total)
	utils.PrintObject(friends)

	err = storage.UserUnRelevance(projectID, us1.UserID, us2.UserID)
	if err != nil {
		panic(err)
	}

	total, friends, err = storage.UserFriendsList(projectID, us1.UserID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(total)
	utils.PrintObject(friends)

	online, err := storage.UserOnline(projectID, []string{us1.UserID, us2.UserID})
	if err != nil {
		panic(err)
	}
	utils.PrintObject(online)

	err = storage.UserOnlinePing(projectID, []string{us2.UserID})
	if err != nil {
		panic(err)
	}

	online, err = storage.UserOnline(projectID, []string{us1.UserID, us2.UserID})
	if err != nil {
		panic(err)
	}
	utils.PrintObject(online)
}

func TestSimpleStorage_Group(t *testing.T) {
	projectID := "cfsre4hb1dspa5lbimu0"
	group, err := storage.GroupCreate(projectID)
	if err != nil {
		panic(err)
	}

	us1 := "cfsrsbhb1dspt4gnuho0"
	us2 := "cfsrsbhb1dspt4gnuhog"

	err = storage.GroupUserRelevance(projectID, group.GroupID, us1)
	if err != nil {
		panic(err)
	}

	err = storage.GroupUserRelevance(projectID, group.GroupID, us2)
	if err != nil {
		panic(err)
	}

	total, ds, err := storage.GroupUserList(projectID, group.GroupID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(total)
	utils.PrintObject(ds)

	err = storage.GroupUserUnRelevance(projectID, group.GroupID, us1)
	if err != nil {
		panic(err)
	}

	total, ds, err = storage.GroupUserList(projectID, group.GroupID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(total)
	utils.PrintObject(ds)

	err = storage.GroupDissolve(projectID, group.GroupID)
	if err != nil {
		panic(err)
	}

	total, ds, err = storage.GroupUserList(projectID, group.GroupID)
	if err != nil {
		panic(err)
	}
	utils.PrintObject(total)
	utils.PrintObject(ds)
}
