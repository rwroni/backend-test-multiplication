package sample

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"project-name/app/adapter/rest_api"
	"project-name/app/models/sql"
	"project-name/app/repository/biller"
	"project-name/app/repository/mbonboarding"
	"project-name/app/repository/scfigcn"
	"project-name/app/repository/userdt"
	"project-name/app/usecase/sample/request"
	"project-name/app/usecase/sample/response"
	"project-name/mocks/app/repository/biller"
	mocks2 "project-name/mocks/app/repository/mbonboarding"
	mocks3 "project-name/mocks/app/repository/userdt"
	"reflect"
	"testing"
)

func Test_sampleUseCase_Read(t *testing.T) {
	var responseError []response.ReadResponse
	sampleRepoMock := &mocks.RepositoryBiller{}

	type repoResponse struct {
		wantOut []sql.Biller
		wantErr error
	}

	tests := []struct {
		name         string
		repoResponse repoResponse
		wantOut      []response.ReadResponse
		wantHttpCode int
		wantCode     string
		wantErr      bool
	}{
		{
			name: "1. Success Test case",
			repoResponse: repoResponse{
				wantOut: []sql.Biller{
					{ID: 1, Name: "Test"},
				},
				wantErr: nil,
			},
			wantOut: []response.ReadResponse{
				{Name: "Test"},
			},
			wantCode:     "",
			wantHttpCode: 0,
			wantErr:      false,
		},
		{
			name: "2. Error #1 Validation",
			repoResponse: repoResponse{
				wantOut: []sql.Biller{},
				wantErr: errors.New("error"),
			},
			wantOut:      responseError,
			wantCode:     "",
			wantHttpCode: 0,
			wantErr:      true,
		},
	}

	var req request.SampleRequest
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sampleRepoMock.On("Read", mock.Anything).Return(tt.repoResponse.wantOut, tt.repoResponse.wantErr).Once()

			uc := &sampleUseCase{
				billerRepo: sampleRepoMock,
			}

			gotOut, gotHttpCode, gotCode, err := uc.Read(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Read() gotOut = %v, want %v", gotOut, tt.wantOut)
			}

			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("Read() gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}

			if gotCode != tt.wantCode {
				t.Errorf("Read() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func Test_sampleUseCase_HitExternalAPI(t *testing.T) {
	type fields struct {
		billerRepo       biller.RepositoryBiller
		scfiGcnRepo      scfigcn.RepositoryScfiGCN
		mbOnboardingRepo mbonboarding.RepositoryMBOnBoarding
		userDTRepo       userdt.RepositoryUserDT
		restAPI          rest_api.RestAPI
	}
	type args struct {
		in request.HitExternalAPIRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantOut      response.ReadResponse
		wantHttpCode int
		wantCode     string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &sampleUseCase{
				billerRepo:       tt.fields.billerRepo,
				scfiGcnRepo:      tt.fields.scfiGcnRepo,
				mbOnboardingRepo: tt.fields.mbOnboardingRepo,
				userDTRepo:       tt.fields.userDTRepo,
				restAPI:          tt.fields.restAPI,
			}
			gotOut, gotHttpCode, gotCode, err := uc.HitExternalAPI(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("HitExternalAPI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("HitExternalAPI() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("HitExternalAPI() gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("HitExternalAPI() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func Test_sampleUseCase_M32(t *testing.T) {
	type fields struct {
		billerRepo       biller.RepositoryBiller
		scfiGcnRepo      scfigcn.RepositoryScfiGCN
		mbOnboardingRepo mbonboarding.RepositoryMBOnBoarding
		userDTRepo       userdt.RepositoryUserDT
		restAPI          rest_api.RestAPI
	}
	type args struct {
		in request.MBOnBoardingRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantOut      response.MBOnBoardingResponse
		wantHttpCode int
		wantCode     string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &sampleUseCase{
				billerRepo:       tt.fields.billerRepo,
				scfiGcnRepo:      tt.fields.scfiGcnRepo,
				mbOnboardingRepo: tt.fields.mbOnboardingRepo,
				userDTRepo:       tt.fields.userDTRepo,
				restAPI:          tt.fields.restAPI,
			}
			gotOut, gotHttpCode, gotCode, err := uc.M32(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("M32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("M32() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("M32() gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("M32() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func Test_sampleUseCase_M37(t *testing.T) {
	type fields struct {
		billerRepo       biller.RepositoryBiller
		scfiGcnRepo      scfigcn.RepositoryScfiGCN
		mbOnboardingRepo mbonboarding.RepositoryMBOnBoarding
		userDTRepo       userdt.RepositoryUserDT
		restAPI          rest_api.RestAPI
	}
	type args struct {
		in request.MBOnBoardingRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantOut      response.MBOnBoardingResponse
		wantHttpCode int
		wantCode     string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &sampleUseCase{
				billerRepo:       tt.fields.billerRepo,
				scfiGcnRepo:      tt.fields.scfiGcnRepo,
				mbOnboardingRepo: tt.fields.mbOnboardingRepo,
				userDTRepo:       tt.fields.userDTRepo,
				restAPI:          tt.fields.restAPI,
			}
			gotOut, gotHttpCode, gotCode, err := uc.M37(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("M37() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("M37() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("M37() gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("M37() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func Test_sampleUseCase_QueryJoin(t *testing.T) {
	var request request.QueryJoinRequest
	mockRepoOnBoarding := &mocks2.RepositoryMBOnBoarding{}
	mockRepoUserDT := &mocks3.RepositoryUserDT{}

	type userDtResponse struct {
		wantOut sql.UserDT
		wantErr error
	}

	type onBoardingResponse struct {
		wantOut sql.MBOnboarding
		wantErr error
	}

	tests := []struct {
		name               string
		userDTResponse     userDtResponse
		onBoardingResponse onBoardingResponse
		wantOut            response.QueryJoinResponse
		wantHttpCode       int
		wantCode           string
		wantErr            bool
	}{
		{
			name: "#1. Success Case",
			userDTResponse: userDtResponse{
				wantErr: nil,
				wantOut: sql.UserDT{RowID: 1},
			},
			onBoardingResponse: onBoardingResponse{
				wantErr: nil,
				wantOut: sql.MBOnboarding{RowID: 1},
			},
			wantOut: response.QueryJoinResponse{},
			wantErr: false,
			wantHttpCode: 200,
			wantCode: "M1",
		},
		{
			name: "#2. Megative case",
			userDTResponse: userDtResponse{
				wantErr: nil,
				wantOut: sql.UserDT{RowID: 1},
			},
			onBoardingResponse: onBoardingResponse{
				wantErr: nil,
				wantOut: sql.MBOnboarding{RowID: 1},
			},
			wantOut: response.QueryJoinResponse{},
			wantErr: false,
			wantHttpCode: 200,
			wantCode: "M1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepoOnBoarding.On("FindByNIK", mock.Anything).Return(tt.onBoardingResponse.wantOut, tt.onBoardingResponse.wantErr).Once()
			mockRepoUserDT.On("FindByGCN", mock.Anything).Return(tt.userDTResponse.wantOut, tt.userDTResponse.wantErr).Once()

			uc := &sampleUseCase{
				mbOnboardingRepo: mockRepoOnBoarding,
				userDTRepo:       mockRepoUserDT,
			}
			gotOut, gotHttpCode, gotCode, err := uc.QueryJoin(request)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryJoin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("QueryJoin() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("QueryJoin() gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("QueryJoin() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
