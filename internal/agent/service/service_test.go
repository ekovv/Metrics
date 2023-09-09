package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"metrics/internal/agent/domains/mocks"
	"testing"
)

type clientMock func(c *mocks.MockClient)
type storageMock func(c *mocks.MockStorage)

func TestService_Send(t *testing.T) {
	tests := []struct {
		name        string
		clientMock  clientMock
		storageMock storageMock
		wantErr     error
	}{
		{
			name: "test1",
			clientMock: func(c *mocks.MockClient) {
				//c.EXPECT().Do(mock.Anything).Return(&http.Response{
				//	Body: io.NopCloser(bytes.NewBufferString("qfqfq")),
				//}, nil)
				c.EXPECT().Do(gomock.Any()).Return(nil, ErrInvalidRequest)
			},
			storageMock: func(c *mocks.MockStorage) {
				c.EXPECT().GetGauge().Return(map[string]float64{
					"qfqf": 1.1,
				}).Times(1)
				c.EXPECT().GetCounter().Return(map[string]int{
					"kkjr": 1,
				}) //выбирать тот для кого прогоняю тесты
			},
			wantErr: ErrInvalidRequest,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.NewMockClient(ctrl)
			storage := mocks.NewMockStorage(ctrl)
			tt.storageMock(storage)
			tt.clientMock(client)

			a := &Service{
				client:  client,
				storage: storage,
			}

			if err := a.Send(); !errors.Is(err, tt.wantErr) {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
