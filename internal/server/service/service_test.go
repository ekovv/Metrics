package service

import (
	"errors"
	"metrics/internal/server/domains/mocks"
	"reflect"
	"testing"
)

type storageMock func(c *mocks.Repository)

func TestService_SetMetric(t *testing.T) {

	type args struct {
		metric string
		name   string
		value  float64
	}

	tests := []struct { // добавляем слайс тестов
		name        string
		args        args
		storageMock storageMock
		wantErr     error
	}{
		{
			name: "invalidErr", // описываем каждый тест:
			args: args{
				metric: "sgnlsgls",
				name:   "wnejnw",
				value:  0,
			},
			storageMock: func(c *mocks.Repository) {},
			wantErr:     ErrInvalidMetric,
		},
		{
			name: "OK", // описываем каждый тест:
			args: args{
				metric: "gauge",
				name:   "wnejnw",
				value:  1.1,
			},
			storageMock: func(c *mocks.Repository) {
				c.Mock.On("Set", "wnejnw", 1.1).
					Return(nil).Times(1)
			},
			wantErr: nil,
		},
		{
			name: "invalidValue", // описываем каждый тест:
			args: args{
				metric: "gauge",
				name:   "wnejnw",
				value:  0,
			},
			storageMock: func(c *mocks.Repository) {},
			wantErr:     ErrInvalidValue,
		},
		{
			name: "invalidDataType", // описываем каждый тест:
			args: args{
				metric: "counter",
				name:   "wnejnw",
				value:  1.1,
			},
			storageMock: func(c *mocks.Repository) {}, //если код планирует дойти то делаю, если нет то нет
			wantErr:     ErrInvalidDataType,
		},
		{
			name: "OK#2", // описываем каждый тест:
			args: args{
				metric: "counter",
				name:   "wnejnw",
				value:  1.0,
			},
			storageMock: func(c *mocks.Repository) {
				c.Mock.On("Inc", "wnejnw", 1.0).Return(nil).Times(1)
			}, //если код планирует дойти то делаю, если нет то нет
			wantErr: nil,
		},
	}
	for _, tt := range tests { // цикл по всем тестам
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			s := NewService(repo)
			tt.storageMock(repo)

			if err := s.SetMetric(tt.args.metric, tt.args.name, tt.args.value); !errors.Is(err, tt.wantErr) {
				t.Errorf("Sum() = %v, want err %v", err, tt.wantErr)
			}
		})
	}

}

func TestService_GetAllMetrics(t *testing.T) {

	tests := []struct {
		name        string
		storageMock storageMock
		want        map[string]float64
	}{
		{
			name: "OK#1", // описываем каждый тест:
			storageMock: func(c *mocks.Repository) {
				c.Mock.On("Get").Return(map[string]float64{}).Times(1)
			}, //если код планирует дойти то делаю, если нет то нет
			want: map[string]float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			s := NewService(repo)
			tt.storageMock(repo)
			if got := s.GetAllMetrics(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestService_GetValueFromM(t *testing.T) {
//	type args struct {
//		name string
//	}
//
//	tests := []struct {
//		name        string
//		args        args
//		storageMock storageMock
//		want        float64
//		wantErr     bool
//	}{
//		{
//			name: "OK#1", // описываем каждый тест:
//			storageMock: func(c *mocks.Repository) {
//				c.Mock.On("GetOne").Return().Times(1)
//			}, //если код планирует дойти то делаю, если нет то нет
//			want: ,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			repo := mocks.NewRepository(t)
//			s := NewService(repo)
//			tt.storageMock(repo)
//			got, err := s.GetValueFromM(tt.args.name)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetValueFromM() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("GetValueFromM() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
