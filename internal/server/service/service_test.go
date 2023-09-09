package service

import (
	"errors"
	"metrics/internal/server/domains/mocks"
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
