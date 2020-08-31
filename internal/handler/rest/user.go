package rest

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
	"image/jpeg"
	"net/http"
	"strconv"
	"y-test/internal/constant/model"
	"y-test/internal/module/user"
	"y-test/internal/security"
)

type UserHandler struct {
	s    user.Service
	auth *security.Jwt
}

func NewUserHandler(s user.Service, auth *security.Jwt) user.Handler {
	return &UserHandler{s: s, auth: auth}
}

func (uh *UserHandler) SignUp(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	log := logrus.WithContext(ctx)
	var s *model.SignUp
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		log.Error("Couldn't Decode r.Body, err=", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validator.Validate(s); err != nil {
		log.Error("Bad request, err=", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	cu, err := uh.s.Create(ctx, &model.User{Email: s.Email, Password: s.Password})
	if err != nil {
		log.Error("Couldn't update user, err=", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, tokenErr := uh.auth.CreateToken(cu.UserId)
	if tokenErr != nil {
		log.Error("Couldn't create token, err=", tokenErr)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)

	//i didn't add token_type and expires_in to response because it's test app
	err = json.NewEncoder(rw).Encode(struct {
		UserId      string `json:"userId"`
		AccessToken string `json:"accessToken"`
	}{
		UserId:      cu.UserId,
		AccessToken: token,
	})

	if err != nil {
		logrus.Error("Couldn't Encode, err=", err)
	}
}
func (uh *UserHandler) Update(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	log := logrus.WithContext(ctx)
	u, done := uh.decodeUser(rw, r, log)
	if done {
		return
	}
	u.UserId = mux.Vars(r)["id"]
	cu, err := uh.s.Update(ctx, u)
	if err != nil {
		log.Error("couldn't update user, err : ", err)
		//todo check by type
		if err.Error() == "not found" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(cu)
	if err != nil {
		logrus.Error("couldn't encode, err=", err)
	}
}

func (uh *UserHandler) UpdateAvatar(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	log := logrus.WithContext(ctx)
	avatar, _, err := r.FormFile("avatar")
	if err != nil {
		log.Error("couldn't get avatar from body :", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	img, err := jpeg.Decode(avatar)
	if err != nil {
		log.Error("couldn't convert avatar to image object :", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	userId := mux.Vars(r)["id"]
	defer avatar.Close()

	u, err := uh.s.UpdateAvatar(ctx, userId, img)
	if err != nil {
		log.Error("couldn't update avatar :", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(u)
}

func (uh *UserHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	log := logrus.WithContext(ctx)
	var s *model.SignUp
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		log.Error("Couldn't Decode r.Body, err=", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	cu, err := uh.s.ReadByCredentials(ctx, s.Email, s.Password)
	if err != nil {
		log.Error("Couldn't create user, err=", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, tokenErr := uh.auth.CreateToken(cu.UserId)
	if tokenErr != nil {
		log.Error("Couldn't create token, err=", tokenErr)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(struct {
		User        *model.User `json:"user"`
		AccessToken string      `json:"accessToken"`
	}{
		User:        cu,
		AccessToken: token,
	})

	if err != nil {
		logrus.Error("Couldn't Encode, err=", err)
	}
}

func (uh *UserHandler) Get(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	log := logrus.WithContext(ctx)
	userId := mux.Vars(r)["id"]
	u, err := uh.s.Read(ctx, userId)
	if err != nil {
		log.Error("couldn't get user from storage, err : ", err)
		//todo check by type
		if err.Error() == "not found" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(u)
	if err != nil {
		logrus.Error("Couldn't Encode, err=", err)
	}
}

func (uh *UserHandler) GetPageable(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	log := logrus.WithContext(ctx)
	values := mux.Vars(r)

	var convertErr error
	var pageSize int
	var pageNumber int

	pageSize, convertErr = strconv.Atoi(values["pageSize"])
	pageNumber, convertErr = strconv.Atoi(values["pageNumber"])
	if convertErr != nil {
		log.Error("couldn't convert pageSize or pageNumber to int, err:", convertErr)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := uh.s.ReadPageable(ctx, pageNumber, pageSize)
	if err != nil {
		logrus.Error("couldn't get users : ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(u)
	if err != nil {
		logrus.Error("couldn't e response encode")
	}
}

func (uh *UserHandler) decodeUser(rw http.ResponseWriter, r *http.Request, log *logrus.Entry) (*model.User, bool) {
	var u *model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Error("couldn't decode Body, err=", err)
		rw.WriteHeader(http.StatusBadRequest)
		return nil, true
	}
	return u, false
}
