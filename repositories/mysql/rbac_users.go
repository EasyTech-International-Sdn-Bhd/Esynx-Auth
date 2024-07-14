package mysql

import (
	"errors"
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
)

type RbacUsersRepository struct {
	option *contracts.IRepository
}

func NewRbacUsersRepository(option *contracts.IRepository) *RbacUsersRepository {
	return &RbacUsersRepository{option}
}

func (r *RbacUsersRepository) Get(userName, password string) (*entities.RbacUsers, error) {
	var user entities.RbacUsers
	has, err := r.option.Db.Where("username = ? AND password = ? AND deleted = ?", userName, password, 0).Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	user.Password = "********" // shadow the password
	return &user, nil
}

func (r *RbacUsersRepository) GetByUserCode(userCode string) (*entities.RbacUsers, error) {
	var user entities.RbacUsers
	has, err := r.option.Db.Where("user_code = ? AND deleted = ?", userCode, 0).Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	user.Password = "********" // shadow the password
	return &user, nil
}

func (r *RbacUsersRepository) CreateUser(info models.CreateRbacUser) error {
	user, err := r.Get(info.Username, info.Password)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("user already exists")
	}
	newUser := entities.RbacUsers{
		Username:       info.Username,
		Password:       info.Password,
		ClientCompany:  info.ClientCompany,
		Metadata:       info.Metadata,
		BiDealer:       info.BiDealer,
		BiSubscription: info.BiSubscriptions,
		BiState:        info.BiState,
		BiIndustry:     info.BiIndustry,
	}
	newUser.ToCreate(info.CreatedBy)
	_, err = r.option.Db.InsertOne(&newUser)
	if err != nil {
		return err
	}
	return nil
}

func (r *RbacUsersRepository) UpdateUser(info models.UpdateRbacUser) error {
	user, err := r.GetByUserCode(info.UserCode)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user does not exist")
	}
	user.Password = info.Password
	user.Metadata = info.Metadata
	user.BiDealer = info.BiDealer
	user.BiState = info.BiState
	user.BiIndustry = info.BiIndustry
	user.ToUpdate(info.UpdatedBy)
	_, err = r.option.Db.Where("user_code = ?", info.UserCode).Update(&user)
	if err != nil {
		return err
	}
	return nil
}

func (r *RbacUsersRepository) DeleteUser(info models.DeleteRbacUser) error {
	user, err := r.GetByUserCode(info.UserCode)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user does not exist")
	}
	user.ToDelete(info.DeletedBy)
	_, err = r.option.Db.Where("user_code = ?", info.UserCode).Update(&user)
	if err != nil {
		return err
	}
	return nil
}
