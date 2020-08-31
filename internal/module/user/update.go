package user

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	img "image"
	"image/jpeg"
	"y-test/internal/constant/model"
	"y-test/platform/minio"
)

func (s *service) Update(ctx context.Context, u *model.User) (*model.User, error) {
	up, e := s.p.Update(ctx, u)
	if e != nil {
		return nil, e
	}

	return up, nil
}

func (s *service) UpdateAvatar(ctx context.Context, userid string, avatar img.Image) (*model.User, error) {
	u, err := s.p.FindByID(ctx, userid)
	if err != nil {
		logrus.Error("user not found : ", err)
		return nil, fmt.Errorf("not found user")
	}

	if avatar.Bounds().Dx() > 160 && avatar.Bounds().Dy() > 160 {
		avatar = s.r.Recize(avatar)
	}

	path := fmt.Sprintf("avatar/%s", userid)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, avatar, nil)
	if err != nil {
		logrus.Error("couldn't encode avatar : ", err)
		return nil, fmt.Errorf("couldn't encode avatar")
	}

	url, err := s.os.PutObject(ctx, &minio.ObjectOptions{
		File:   buf,
		Path:   path,
		Bucket: "user",
		Name:   fmt.Sprintf("avatar/%s.jpeg", userid),
		Size:   buf.Len(),
	})
	if err != nil {
		logrus.Error("couldn't save avatar : ", err)
		return nil, err
	}

	if u.AvatarLink == "" {
		u.AvatarLink = url
		if _, err := s.p.Update(ctx, u); err != nil {
			//todo figure out rollup for object storage
			logrus.Errorf("couldn't set avatar link to storage :", err)
			return nil, err
		}
	}

	return u, nil
}
