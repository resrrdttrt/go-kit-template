package http

import (
	"fmt"
	"time"

	"go-kit-template/pkg/errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidUUID      = errors.New("invalid uuid")
	ErrInvalidRoleValue = errors.New("role must be enterprise, end_user or admin")
	ErrInvalidStatus    = errors.New("status must be active or inactive")
)

func errMissing(field string) error {
	return errors.Wrap(errors.ErrMalformedEntity, errors.New(fmt.Sprintf("missing field `%s`", field)))
}

type createUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

func (req createUserRequest) validate() error {
	if req.Name == "" {
		return errMissing("name")
	}
	if req.Username == "" {
		return errMissing("username")
	}
	if req.Password == "" {
		return errMissing("password")
	}
	if req.Email == "" {
		return errMissing("email")
	}
	if req.Phone == "" {
		return errMissing("phone")
	}
	if req.Role == "" {
		return errMissing("role")
	}
	if req.Role != "enterprise" && req.Role != "end_user" && req.Role != "admin" {
		return ErrInvalidRoleValue
	}
	if req.Status == "" {
		return errMissing("status")
	}
	if req.Status != "active" && req.Status != "inactive" {
		return ErrInvalidStatus
	}
	return nil
}

type updateUserRequest struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Status    string `json:"status"`
}

func (req updateUserRequest) validate() error {
	if req.ID == "" {
		return errMissing("user_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	if req.Role != "" && req.Role != "enterprise" && req.Role != "end_user" && req.Role != "admin" {
		return ErrInvalidRoleValue
	}
	if req.Status != "" && req.Status != "active" && req.Status != "inactive" {
		return ErrInvalidStatus
	}
	return nil
}

type getUserRequest struct {
	ID string `json:"id"`
}

func (req getUserRequest) validate() error {
	if req.ID == "" {
		return errMissing("user_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	return nil
}

type getGameRequest struct {
	ID string `json:"id"`
}

func (req getGameRequest) validate() error {
	if req.ID == "" {
		return errMissing("game_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	return nil
}

type createGameRequest struct {
	Name          string `json:"name"`
	Images        string `json:"images"`
	Type          string `json:"type"`
	ExchangeAllow bool   `json:"exchange_allow"`
	Tutorial      string `json:"tutorial"`
}

func (req createGameRequest) validate() error {
	if req.Name == "" {
		return errMissing("name")
	}
	if req.Images == "" {
		return errMissing("images")
	}
	if req.Type == "" {
		return errMissing("type")
	}
	if req.Tutorial == "" {
		return errMissing("tutorial")
	}
	return nil
}

type updateGameRequest struct {
	ID            string 
	Name          string `json:"name"`
	Images        string `json:"images"`
	Type          string `json:"type"`
	ExchangeAllow bool   `json:"exchange_allow"`
	Tutorial      string `json:"tutorial"`
}

func (req updateGameRequest) validate() error {
	if req.ID == "" {
		return errMissing("game_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	return nil
}

type statisticInTimeRequest struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (req statisticInTimeRequest) validate() error {
	if req.Start.IsZero() {
		return errMissing("start")
	}
	if req.End.IsZero() {
		return errMissing("end")
	}
	return nil
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req loginRequest) validate() error {
	if req.Username == "" {
		return errMissing("username")
	}
	if req.Password == "" {
		return errMissing("password")
	}
	return nil
}

type enterpriseRequest struct {
	Name     string `json:"name"`
	Field    string `json:"field"`
	Location string `json:"location"`
	GPS      string `json:"gps"`
	Status   string `json:"status"`
}

func (req enterpriseRequest) validate() error {
	if req.Name == "" {
		return errMissing("name")
	}
	return nil
}

type getEventIDRequest struct {
	ID string 
}

func (req getEventIDRequest) validate() error {
	if req.ID == "" {
		return errMissing("event_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	return nil
}

type createEventRequest struct {
	Name       string    `json:"name"`
	Images     string    `json:"images"`
	VoucherNum int       `json:"voucher_num"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	GameID     string    `json:"game_id"`
	UserID     string
}

func (req createEventRequest) validate() error {
	if req.Name == "" {
		return errMissing("name")
	}
	if req.Images == "" {
		return errMissing("images")
	}
	if req.VoucherNum == 0 {
		return errMissing("voucher_num")
	}
	if req.StartTime.IsZero() {
		return errMissing("start_time")
	}
	if req.EndTime.IsZero() {
		return errMissing("end_time")
	}
	if req.GameID == "" {
		return errMissing("game_id")
	} else {
		if _, err := uuid.Parse(req.GameID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	if req.UserID == "" {
		return errMissing("user_id")
	} else {
		if _, err := uuid.Parse(req.UserID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	return nil
}

type updateEventRequest struct {
	ID         string    
	Name       string    `json:"name"`
	Images     string    `json:"images"`
	VoucherNum int       `json:"voucher_num"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	GameID     string    `json:"game_id"`
	UserID     string
}

func (req updateEventRequest) validate() error {
	if req.ID == "" {
		return errMissing("event_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	return nil
}

type getVoucherByIDRequest struct {
	EventID string
	ID      string
}

func (req getVoucherByIDRequest) validate() error {
	if req.EventID == "" {
		return errMissing("event_id")
	} else {
		if _, err := uuid.Parse(req.EventID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	if req.ID == "" {
		return errMissing("voucher_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	return nil
}

type createVoucherRequest struct {
	Code        string    `json:"code"`
	Qrcode      string    `json:"qrcode"`
	Images      string    `json:"images"`
	Value       int       `json:"value"`
	Description string    `json:"description"`
	ExpiredTime time.Time `json:"expired_time"`
	Status      string    `json:"status"`
	EventID     string
}

func (req createVoucherRequest) validate() error {
	if req.Code == "" {
		return errMissing("code")
	}
	if req.Qrcode == "" {
		return errMissing("qrcode")
	}
	if req.Images == "" {
		return errMissing("images")
	}
	if req.Value == 0 {
		return errMissing("value")
	}
	if req.ExpiredTime.IsZero() {
		return errMissing("expired_time")
	}
	if req.Status == "" {
		return errMissing("status")
	}
	return nil
}

type updateVoucherRequest struct {
	ID          string
	Code        string    `json:"code"`
	Qrcode      string    `json:"qrcode"`
	Images      string    `json:"images"`
	Value       int       `json:"value"`
	Description string    `json:"description"`
	ExpiredTime time.Time `json:"expired_time"`
	Status      string    `json:"status"`
	EventID     string
}

func (req updateVoucherRequest) validate() error {
	if req.Code == "" {
		return errMissing("code")
	}
	if req.Qrcode == "" {
		return errMissing("qrcode")
	}
	if req.Images == "" {
		return errMissing("images")
	}
	if req.Value == 0 {
		return errMissing("value")
	}
	if req.ExpiredTime.IsZero() {
		return errMissing("expired_time")
	}
	return nil
}