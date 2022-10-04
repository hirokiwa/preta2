package repositoryImpl

import (
	"fmt"

	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure"
	"hackz.com/m/v2/infrastructure/dto"
	"strconv"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (repositoryImpl *UserRepositoryImpl) Findfollowee(followeeid string) ([]*model.User, error) {
	db := infrastructure.GetDB()
	var err error
	var follow []*dto.Follow = []*dto.Follow{}
	var user []*model.User = []*model.User{}
	if err := db.Model(&follow).Select("*").Joins("inner join `users` on follows.follower = users.userid").Where("follows.followee = ?", followeeid).Scan(&user).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return user, err

}

func (repositoryImpl *UserRepositoryImpl) Findfollower(followerid string) ([]*model.User, error) {
	db := infrastructure.GetDB()
	var err error
	var follow []*dto.Follow = []*dto.Follow{}
	var user []*model.User = []*model.User{}
	if err := db.Model(&follow).Select("*").Joins("inner join `users` on follows.followee = users.userid").Where("follows.follower = ?", followerid).Scan(&user).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return user, err

}
func (repositoryImpl *UserRepositoryImpl) FindUser(followerid string) (*model.User, error) {
	db := infrastructure.GetDB()
	var err error
	var user *model.User = &model.User{}
	if err := db.Where("userid = ?", followerid).First(&user).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return user, err

}
func (repositoryImpl *UserRepositoryImpl) CreateUser(input model.NewUser) (*model.User, error) {
	db := infrastructure.GetDB()
	var err error
	if err := db.Create(&dto.User{Userid: input.Userid, Name: input.Name}).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("user create Error!!!! err:%v\n", err)
		return &model.User{
			Userid: input.Userid,
			Name:   input.Name,
		}, err
	}

	return &model.User{
		Userid: input.Userid,
		Name:   input.Name,
	}, err
}

func (repositoryImpl *UserRepositoryImpl) CreateFollow(input model.NewFollow) (*model.User, error) {
	db := infrastructure.GetDB()
	var err error

	var folllower = &dto.User{}

	if err := db.First(&folllower, "userid=?", input.Follower).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("follower user Notfound:%v\n", err)
		return &model.User{
			Userid: input.Follower,
			Name:   folllower.Name,
		}, err
	}

	var folllowee = &dto.User{}

	if err := db.First(&folllowee, "userid=?", input.Followee).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("followee user Notfound:%v\n", err)
		return &model.User{
			Userid: input.Follower,
			Name:   folllower.Name,
		}, err
	}

	if err := db.Create(&dto.Follow{Followee: input.Followee, Follower: input.Follower}).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("followe err:%v\n", err)
		return &model.User{
			Userid: input.Follower,
			Name:   folllower.Name,
		}, err
	}

	return &model.User{
		Userid: input.Follower,
		Name:   folllower.Name,
	}, err
}

func (reppsitoryImpl *UserRepositoryImpl) FindFolloweeDiary(followerid string) ([]*model.UserDiary, error) {
	db := infrastructure.GetDB()
	var FolloweeDiary = []*dto.UserandDiary{}
	var follow []*dto.Follow = []*dto.Follow{}
	var UserDiary = []*model.UserDiary{}
	var Diary = []*model.Diary{}
	var err error
	if err := db.Model(&follow).Select("*").Joins("inner join `users` on follows.followee = users.userid").Joins("inner join `diaries` on follows.followee = diaries.userid").Joins("inner join `englishes` on diaries.diaryid = englishes.diaryid").Joins("inner join `emotions` on diaries.diaryid = emotions.diaryid").Where("follows.follower = ?", followerid).Order("diaries.userid desc").Scan(&FolloweeDiary).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	var i = 0
	var userdiaryindex = 0
	println(len(FolloweeDiary))


	if (len(FolloweeDiary) == 0) {

	} else {
		for i = 0; i < len(FolloweeDiary); i++ {
			if i > 0 {
				if FolloweeDiary[i].Userid != FolloweeDiary[i-1].Userid {
					fmt.Printf("(%%#v) %#v\n", UserDiary[userdiaryindex])
					Diary = Diary[0:0]
					UserDiary = append(UserDiary, &model.UserDiary{
						User: &model.User{
							Userid: FolloweeDiary[i-1].Userid,
							Name:   FolloweeDiary[i-1].Name,
						},
						Diary: Diary,
	
					})
					userdiaryindex += 1
					
				}
			}
			Diary = append(Diary, &model.Diary{
				Diaryid:     strconv.Itoa(FolloweeDiary[i].Diaryid),
				Word:        FolloweeDiary[i].Word,
				Imageurl:    FolloweeDiary[i].Imageurl,
				Englishword: FolloweeDiary[i].Englishword,
				CreatedAt:   FolloweeDiary[i].CreatedAt.String(),
				UpdatedAt:   FolloweeDiary[i].UpdatedAt.String(),
				Emotion: &model.Emotion{
					Diaryid:   strconv.Itoa(FolloweeDiary[i].Diaryid),
					Happy:     FolloweeDiary[i].Happy,
					Angry:     FolloweeDiary[i].Angry,
					Fear:      FolloweeDiary[i].Fear,
					Surprise:  FolloweeDiary[i].Surprise,
					Sad:       FolloweeDiary[i].Sad,
					CreatedAt: FolloweeDiary[i].CreatedAt.String(),
					UpdatedAt: FolloweeDiary[i].UpdatedAt.String(),
				},
				User: &model.User{
					Userid: FolloweeDiary[i].Userid,
					Name:   FolloweeDiary[i].Name,
				},
			})
			if i == len(FolloweeDiary)-1 {
				UserDiary = append(UserDiary, &model.UserDiary{
					User: &model.User{
						Userid: FolloweeDiary[i].Userid,
						Name:   FolloweeDiary[i].Name,
					},
					Diary: Diary,
				})
			}
		}
	}
	return UserDiary, err
}


func (reppsitoryImpl *UserRepositoryImpl) FindFollowerDiary(followeeid string) ([]*model.UserDiary, error) {
	db := infrastructure.GetDB()
	var FollowerDiary = []*dto.UserandDiary{}
	var follow []*dto.Follow = []*dto.Follow{}
	var UserDiary = []*model.UserDiary{}
	var Diary = []*model.Diary{}
	var err error
	if err := db.Model(&follow).Select("*").Joins("inner join `users` on follows.follower = users.userid").Joins("inner join `diaries` on follows.follower = diaries.userid").Joins("inner join `englishes` on diaries.diaryid = englishes.diaryid").Joins("inner join `emotions` on diaries.diaryid = emotions.diaryid").Where("follows.followee = ?", followeeid).Order("diaries.userid desc").Scan(&FollowerDiary).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	var i = 0
	var userdiaryindex = 0
	println(len(FollowerDiary))


	if (len(FollowerDiary) == 0) {

	} else {
		for i = 0; i < len(FollowerDiary); i++ {
			if i > 0 {
				if FollowerDiary[i].Userid != FollowerDiary[i-1].Userid {
					fmt.Printf("(%%#v) %#v\n", UserDiary[userdiaryindex])
					Diary = Diary[0:0]
					UserDiary = append(UserDiary, &model.UserDiary{
						User: &model.User{
							Userid: FollowerDiary[i-1].Userid,
							Name:   FollowerDiary[i-1].Name,
						},
						Diary: Diary,
	
					})
					userdiaryindex += 1
					
				}
			}
			Diary = append(Diary, &model.Diary{
				Diaryid:     strconv.Itoa(FollowerDiary[i].Diaryid),
				Word:        FollowerDiary[i].Word,
				Imageurl:    FollowerDiary[i].Imageurl,
				Englishword: FollowerDiary[i].Englishword,
				CreatedAt:   FollowerDiary[i].CreatedAt.String(),
				UpdatedAt:   FollowerDiary[i].UpdatedAt.String(),
				Emotion: &model.Emotion{
					Diaryid:   strconv.Itoa(FollowerDiary[i].Diaryid),
					Happy:     FollowerDiary[i].Happy,
					Angry:     FollowerDiary[i].Angry,
					Fear:      FollowerDiary[i].Fear,
					Surprise:  FollowerDiary[i].Surprise,
					Sad:       FollowerDiary[i].Sad,
					CreatedAt: FollowerDiary[i].CreatedAt.String(),
					UpdatedAt: FollowerDiary[i].UpdatedAt.String(),
				},
				User: &model.User{
					Userid: FollowerDiary[i].Userid,
					Name:   FollowerDiary[i].Name,
				},
			})
			if i == len(FollowerDiary)-1 {
				UserDiary = append(UserDiary, &model.UserDiary{
					User: &model.User{
						Userid: FollowerDiary[i].Userid,
						Name:   FollowerDiary[i].Name,
					},
					Diary: Diary,
				})
			}
		}
	}
	return UserDiary, err
}
