package routing

import (
	"reflect"
	"testing"
	"y-test/internal/module/user"
)

func TestUserInit(t *testing.T) {
	type args struct {
		handler user.Handler
	}
	tests := []struct {
		name string
		args args
		want user.Route
	}{
		{
			name: "success",
			args: args{
				handler: nil,
			},
			want: &userHandlers{
				handler: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandlers(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_userHandlers_Routers(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	mockedHandler := mock_user.NewMockHandler(ctrl)
//	type fields struct {
//		handler user.Handler
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   []*routers.Router
//	}{
//		{
//			name: "success",
//			fields: fields{
//				handler: mockedHandler,
//			},
//			want: []*routers.Router{
//				&routers.Router{
//					Method:  "GET",
//					URL:     "/test",
//					Handler: mockedHandler.Test,
//				},
//				&routers.Router{
//					Method:  "GET",
//					URL:     "/account/profile",
//					Handler: mockedHandler.ShowProfile,
//				},
//				&routers.Router{
//					Method:  "POST",
//					URL:     "/account/register",
//					Handler: mockedHandler.CreateNewAccount,
//				},
//				&routers.Router{
//					Method:  "POST",
//					URL:     "/account/login",
//					Handler: mockedHandler.SignUp,
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			uh := &userHandlers{
//				handler: tt.fields.handler,
//			}
//			got := uh.Routers()
//			for i, val := range got {
//				if !reflect.DeepEqual(val.Method, tt.want[i].Method) {
//					t.Errorf("userHandlers.Routers().Method = %v, want %v", got, tt.want)
//				}
//				if !reflect.DeepEqual(val.URL, tt.want[i].URL) {
//					t.Errorf("userHandlers.Routers().URL = %v, want %v", got, tt.want)
//				}
//			}
//		})
//	}
//}
