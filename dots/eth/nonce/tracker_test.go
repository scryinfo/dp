// Scry Info.  All rights reserved.
// license that can be found in the license file.

package nonce

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/db/gorms"
	"github.com/scryinfo/dp/dots/eth/client"
	"github.com/scryinfo/dp/dots/eth/nonce/mocks"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TrackerTestSuite struct {
	suite.Suite
	gorutines int
}

func (suite *TrackerTestSuite) SetupTest() {
	rand.Seed(time.Now().UnixNano())
	goroutinesCount := rand.Intn(100)
	suite.gorutines = goroutinesCount
}

func (suite *TrackerTestSuite) TestCreate() {
	// testing scenarios
	scenarios := []struct {
		desc                              string
		expError                          error
		configMockLineAndInjecter         func(mockLine *mocks.Line, mockInjecter *mocks.Injecter)
		patchCreateStorerWithDBAndMockSql func(db *sql.DB, mockSql sqlmock.Sqlmock)
	}{
		{
			desc: "happy path",
			configMockLineAndInjecter: func(mockLine *mocks.Line, mockInjecter *mocks.Injecter) {
				mockInjecter.On("GetByLiveId", dot.LiveId(client.ClientTypeId)).Return(&client.Client{}, nil)
				mockInjecter.On("GetByLiveId", dot.LiveId(gorms.TypeId)).Return(&gorms.Gorms{}, nil)
				mockLine.On("ToInjecter").Return(mockInjecter)
			},
			patchCreateStorerWithDBAndMockSql: func(db *sql.DB, mockSql sqlmock.Sqlmock) {
				createTableNoncesIFNotExistsSQLRegex := fmt.Sprintf("\\Q%v\\E", createTableNoncesIFNotExistsSQL())
				mockSql.ExpectExec(createTableNoncesIFNotExistsSQLRegex).WillReturnResult(sqlmock.NewResult(1, 0))

				createStorer = func(t *TrackerImp) error {
					if err := db.Ping(); err != nil {
						return err
					}
					storer, err := getStorer(db)
					if err != nil {
						return err
					}
					t.storer = storer
					return nil
				}
			},
		},
		{
			desc:     "Dot issue: Dot not found",
			expError: errors.New("ErrDotNotFound"),
			configMockLineAndInjecter: func(mockLine *mocks.Line, mockInjecter *mocks.Injecter) {
				mockInjecter.On("GetByLiveId", dot.LiveId(client.ClientTypeId)).Return(&client.Client{}, nil)
				mockInjecter.On("GetByLiveId", dot.LiveId(gorms.TypeId)).Return((*gorms.Gorms)(nil), errors.New("ErrDotNotFound"))
				mockLine.On("ToInjecter").Return(mockInjecter)
			},
		},
		{
			desc:     "Dot issue: internal error",
			expError: errors.New("ErrDotInternalError"),
			configMockLineAndInjecter: func(mockLine *mocks.Line, mockInjecter *mocks.Injecter) {
				mockInjecter.On("GetByLiveId", dot.LiveId(client.ClientTypeId)).Return(&client.Client{}, nil)
				mockInjecter.On("GetByLiveId", dot.LiveId(gorms.TypeId)).Return((*gorms.Gorms)(nil), errors.New("ErrDotInternalError"))
				mockLine.On("ToInjecter").Return(mockInjecter)
			},
		},
		{
			desc:     "DB issue: fail connecting DB",
			expError: errors.New("ErrConnectDB"),
			configMockLineAndInjecter: func(mockLine *mocks.Line, mockInjecter *mocks.Injecter) {
				mockInjecter.On("GetByLiveId", dot.LiveId(client.ClientTypeId)).Return(&client.Client{}, nil)
				mockInjecter.On("GetByLiveId", dot.LiveId(gorms.TypeId)).Return(&gorms.Gorms{}, nil)
				mockLine.On("ToInjecter").Return(mockInjecter)
			},
			patchCreateStorerWithDBAndMockSql: func(db *sql.DB, mockSql sqlmock.Sqlmock) {
				createTableNoncesIFNotExistsSQLRegex := fmt.Sprintf("\\Q%v\\E", createTableNoncesIFNotExistsSQL())
				mockSql.ExpectExec(createTableNoncesIFNotExistsSQLRegex).WillReturnError(errors.New("ErrConnectDB"))

				createStorer = func(t *TrackerImp) error {
					if err := db.Ping(); err != nil {
						return err
					}
					storer, err := getStorer(db)
					if err != nil {
						return err
					}
					t.storer = storer
					return nil
				}
			},
		},
		{
			desc:     "DB issue: fail inserting sql",
			expError: errors.New("ErrInsertSQL"),
			configMockLineAndInjecter: func(mockLine *mocks.Line, mockInjecter *mocks.Injecter) {
				mockInjecter.On("GetByLiveId", dot.LiveId(client.ClientTypeId)).Return(&client.Client{}, nil)
				mockInjecter.On("GetByLiveId", dot.LiveId(gorms.TypeId)).Return(&gorms.Gorms{}, nil)
				mockLine.On("ToInjecter").Return(mockInjecter)
			},
			patchCreateStorerWithDBAndMockSql: func(db *sql.DB, mockSql sqlmock.Sqlmock) {
				createTableNoncesIFNotExistsSQLRegex := fmt.Sprintf("\\Q%v\\E", createTableNoncesIFNotExistsSQL())
				mockSql.ExpectExec(createTableNoncesIFNotExistsSQLRegex).WillReturnError(errors.New("ErrInsertSQL"))

				createStorer = func(t *TrackerImp) error {
					if err := db.Ping(); err != nil {
						return err
					}
					storer, err := getStorer(db)
					if err != nil {
						return err
					}
					t.storer = storer
					return nil
				}
			},
		},
		{
			desc:     "DB issue: DB internal error",
			expError: errors.New("ErrDBInternalError"),
			configMockLineAndInjecter: func(mockLine *mocks.Line, mockInjecter *mocks.Injecter) {
				mockInjecter.On("GetByLiveId", dot.LiveId(client.ClientTypeId)).Return(&client.Client{}, nil)
				mockInjecter.On("GetByLiveId", dot.LiveId(gorms.TypeId)).Return(&gorms.Gorms{}, nil)
				mockLine.On("ToInjecter").Return(mockInjecter)
			},
			patchCreateStorerWithDBAndMockSql: func(db *sql.DB, mockSql sqlmock.Sqlmock) {
				createTableNoncesIFNotExistsSQLRegex := fmt.Sprintf("\\Q%v\\E", createTableNoncesIFNotExistsSQL())
				mockSql.ExpectExec(createTableNoncesIFNotExistsSQLRegex).WillReturnError(errors.New("ErrDBInternalError"))

				createStorer = func(t *TrackerImp) error {
					if err := db.Ping(); err != nil {
						return err
					}
					storer, err := getStorer(db)
					if err != nil {
						return err
					}
					t.storer = storer
					return nil
				}
			},
		},
	}

	// perform tests
	for _, scenario := range scenarios {
		// create tracker
		tracker := &TrackerImp{mu: &sync.Mutex{}}

		// mock Line and Injecter
		mockLine := &mocks.Line{}
		mockInjecter := &mocks.Injecter{}
		scenario.configMockLineAndInjecter(mockLine, mockInjecter)

		// monkey patch createStorer with DB and mockSql
		db, mockSql, err := sqlmock.New()
		if scenario.patchCreateStorerWithDBAndMockSql != nil {
			suite.Require().NoError(err)
			defer db.Close()
			defer func(origin func(t *TrackerImp) error) {
				createStorer = origin
			}(createStorer)
			scenario.patchCreateStorerWithDBAndMockSql(db, mockSql)
		}

		// execute code
		resErr := tracker.Create(mockLine)

		// testify results
		if scenario.expError != nil {
			suite.Require().Error(resErr)
			suite.EqualError(scenario.expError, resErr.Error())
		}

		// make sure all expectations were met
		mockLine.AssertExpectations(suite.T())
		mockInjecter.AssertExpectations(suite.T())
		suite.NoError(mockSql.ExpectationsWereMet())
	}
}

func (suite *TrackerTestSuite) TestCreateAccount() {
	// testing scenarios
	scenarios := []struct {
		desc             string
		address          string
		expError         error
		configMockClient func(mockClient *mocks.EthClient)
		configMockSql    func(mockSql sqlmock.Sqlmock)
	}{
		{
			desc:    "happy path",
			address: "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", context.Background(), common.HexToAddress("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643")).Return(uint64(100), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce_count"}).AddRow("0"))
					insertNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", insertNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 100, true))
					mockSql.ExpectExec(insertNonceSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
				}
			},
		},
		{
			desc:     "input error: invalid input address",
			address:  "0xinvalidaddress",
			expError: ErrInvalidAddress,
		},
		{
			desc:     "ethClient issue: fail connecting Ethereum",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrConnectETHNetwork"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(0), errors.New("ErrConnectETHNetwork"))
			},
		},
		{
			desc:     "ethClient issue: Ethereum internal error",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrETHNetworkInternalError"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(0), errors.New("ErrETHNetworkInternalError"))
			},
		},
		{
			desc:     "DB issue: fail connecting DB",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrConnectDB"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(0), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnError(errors.New("ErrConnectDB"))
				}
			},
		},
		{
			desc:     "DB issue: fail inserting sql",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrInsertSQL"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(0), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnError(errors.New("ErrInsertSQL"))
				}
			},
		},
		{
			desc:     "DB issue: DB internal error",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrDBInternalError"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(0), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnError(errors.New("ErrDBInternalError"))
				}
			},
		},
		{
			desc:     "data conflict: account already exists",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: ErrAccountAlreadyExist,
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(0), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce_count"}).AddRow("1"))
				}
			},
		},
	}

	// perform tests
	for _, scenario := range scenarios {
		// create tracker
		tracker := &TrackerImp{mu: &sync.Mutex{}}

		// mock client
		mockClient := &mocks.EthClient{}
		tracker.ethClient = mockClient
		if scenario.configMockClient != nil {
			scenario.configMockClient(mockClient)
		}

		// mock sql
		db, mockSql, err := sqlmock.New()
		suite.Require().NoError(err)

		tracker.storer = &storer{dao: &dao{db}}
		defer db.Close()

		if scenario.configMockSql != nil {
			mockSql.MatchExpectationsInOrder(false)
			scenario.configMockSql(mockSql)
		}

		var wg sync.WaitGroup
		wg.Add(suite.gorutines)
		for i := 0; i < suite.gorutines; i++ {
			go func() {
				defer wg.Done()

				// execute code
				resErr := tracker.CreateAccount(scenario.address)

				// testify results
				if scenario.expError != nil {
					suite.EqualError(scenario.expError, resErr.Error())
				}
			}()
		}
		wg.Wait()

		// make sure all expectations were met
		mockClient.AssertExpectations(suite.T())
		suite.NoError(mockSql.ExpectationsWereMet())
	}
}

func (suite *TrackerTestSuite) TestDeleteAccount() {
	// testing scenarios
	scenarios := []struct {
		desc             string
		address          string
		expError         error
		configMockClient func(mockClient *mocks.EthClient)
		configMockSql    func(mockSql sqlmock.Sqlmock)
	}{
		{
			desc:    "happy path",
			address: "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					deleteAllNoncesSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteAllNoncesSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(deleteAllNoncesSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
				}
			},
		},
		{
			desc:     "DB issue: fail connecting DB",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrConnectDB"),
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					deleteAllNoncesSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteAllNoncesSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(deleteAllNoncesSQLRegex).WillReturnError(errors.New("ErrConnectDB"))
				}
			},
		},
		{
			desc:     "DB issue: fail inserting sql",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrInsertSQL"),
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					deleteAllNoncesSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteAllNoncesSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(deleteAllNoncesSQLRegex).WillReturnError(errors.New("ErrInsertSQL"))
				}
			},
		},
		{
			desc:     "DB issue: DB internal error",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrDBInternalError"),
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					deleteAllNoncesSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteAllNoncesSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(deleteAllNoncesSQLRegex).WillReturnError(errors.New("ErrDBInternalError"))
				}
			},
		},
		{
			desc:     "data absent: account not exists",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: nil,
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					deleteAllNoncesSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteAllNoncesSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(deleteAllNoncesSQLRegex).WillReturnResult(sqlmock.NewResult(0, 0))
				}
			},
		},
	}

	// perform tests
	for _, scenario := range scenarios {
		// create tracker
		tracker := &TrackerImp{mu: &sync.Mutex{}}

		// mock client
		mockClient := &mocks.EthClient{}
		tracker.ethClient = mockClient
		if scenario.configMockClient != nil {
			scenario.configMockClient(mockClient)
		}

		// mock sql
		db, mockSql, err := sqlmock.New()
		suite.Require().NoError(err)

		tracker.storer = &storer{dao: &dao{db}}
		defer db.Close()
		if scenario.configMockSql != nil {
			mockSql.MatchExpectationsInOrder(false)
			scenario.configMockSql(mockSql)
		}

		var wg sync.WaitGroup
		wg.Add(suite.gorutines)
		for i := 0; i < suite.gorutines; i++ {
			go func() {
				defer wg.Done()

				// execute code
				resErr := tracker.DeleteAccount(scenario.address)

				// testify results
				if scenario.expError != nil {
					suite.EqualError(scenario.expError, resErr.Error())
				}
			}()
		}
		wg.Wait()

		// make sure all expectations were met
		mockClient.AssertExpectations(suite.T())
		suite.NoError(mockSql.ExpectationsWereMet())
	}
}

func (suite *TrackerTestSuite) TestRetrieveNonceAt() {
	// testing scenarios
	scenarios := []struct {
		desc             string
		address          string
		expNonce         uint64
		expError         error
		configMockClient func(mockClient *mocks.EthClient)
		configMockSql    func(mockSql sqlmock.Sqlmock)
	}{
		{
			desc:     "happy path: only the one updated nonce in storage",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expNonce: 100,
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce_count"}).AddRow("1"))
					returnAndIncreaseUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", returnAndIncreaseUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(returnAndIncreaseUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("100"))

				}
			},
		},
		{
			desc:     "happy path: with recycled nonce(s) in storage",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expNonce: 99,
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce_count"}).AddRow("2"))
					deleteAndReturnMinRecycledNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteAndReturnMinRecycledNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(deleteAndReturnMinRecycledNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("99"))
				}
			},
		},
		{
			desc:     "DB issue: fail connecting DB",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrConnectDB"),
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnError(errors.New("ErrConnectDB"))
				}
			},
		},
		{
			desc:     "DB issue: fail inserting sql",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrInsertSQL"),
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnError(errors.New("ErrInsertSQL"))
				}
			},
		},
		{
			desc:     "DB issue: DB internal error",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrDBInternalError"),
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnError(errors.New("ErrDBInternalError"))
				}
			},
		},
		{
			desc:     "data absent: account not exists",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: ErrNonceNotExist,
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryAllNoncesCountSQLRegex := fmt.Sprintf("\\Q%v\\E", queryAllNoncesCountSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryAllNoncesCountSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce_count"}).AddRow("0"))
				}
			},
		},
	}

	// perform tests
	for _, scenario := range scenarios {
		// create tracker
		tracker := &TrackerImp{mu: &sync.Mutex{}}

		// mock client
		mockClient := &mocks.EthClient{}
		tracker.ethClient = mockClient
		if scenario.configMockClient != nil {
			scenario.configMockClient(mockClient)
		}

		// mock sql
		db, mockSql, err := sqlmock.New()
		suite.Require().NoError(err)

		tracker.storer = &storer{dao: &dao{db}}
		defer db.Close()
		if scenario.configMockSql != nil {
			mockSql.MatchExpectationsInOrder(false)
			scenario.configMockSql(mockSql)
		}

		var wg sync.WaitGroup
		wg.Add(suite.gorutines)
		for i := 0; i < suite.gorutines; i++ {
			go func() {
				defer wg.Done()

				// execute code
				resNonce, resErr := tracker.RetrieveNonceAt(scenario.address)

				// testify results
				if scenario.expError != nil {
					suite.EqualError(scenario.expError, resErr.Error())
				}
				suite.Equal(scenario.expNonce, resNonce)
			}()
		}
		wg.Wait()

		// make sure all expectations were met
		mockClient.AssertExpectations(suite.T())
		suite.NoError(mockSql.ExpectationsWereMet())
	}
}

func (suite *TrackerTestSuite) TestRestoreNonce() {
	// testing scenarios
	scenarios := []struct {
		desc             string
		address          string
		nonce            uint64
		expError         error
		configMockClient func(mockClient *mocks.EthClient)
		configMockSql    func(mockSql sqlmock.Sqlmock)
	}{
		{
			desc:    "happy path",
			address: "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			nonce:   99,
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 99))
					mockSql.ExpectQuery(queryNonceSQLRegex).WillReturnError(sql.ErrNoRows)
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("100"))
					insertNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", insertNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 99, false))
					mockSql.ExpectExec(insertNonceSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
				}
			},
		},
		{
			desc:     "input error: account not exists",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			nonce:    99,
			expError: ErrAccountNotExist,
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 99))
					mockSql.ExpectQuery(queryNonceSQLRegex).WillReturnError(sql.ErrNoRows)
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnError(sql.ErrNoRows)
				}
			},
		},
		{
			desc:     "input error: invalid recycled nonce",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			nonce:    101,
			expError: ErrInvalidRecycledNonce,
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 101))
					mockSql.ExpectQuery(queryNonceSQLRegex).WillReturnError(sql.ErrNoRows)
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("100"))
				}
			},
		},
		{
			desc:     "DB issue: fail connecting DB",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			nonce:    99,
			expError: errors.New("ErrConnectDB"),
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 99))
					mockSql.ExpectQuery(queryNonceSQLRegex).WillReturnError(sql.ErrNoRows)
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("100"))
					insertNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", insertNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 99, false))
					mockSql.ExpectExec(insertNonceSQLRegex).WillReturnError(errors.New("ErrConnectDB"))
				}
			},
		},
		{
			desc:     "DB issue: fail inserting sql",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			nonce:    99,
			expError: errors.New("ErrInsertSQL"),
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 99))
					mockSql.ExpectQuery(queryNonceSQLRegex).WillReturnError(sql.ErrNoRows)
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("100"))
					insertNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", insertNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 99, false))
					mockSql.ExpectExec(insertNonceSQLRegex).WillReturnError(errors.New("ErrInsertSQL"))
				}
			},
		},
		{
			desc:     "DB issue: DB internal error",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			nonce:    99,
			expError: errors.New("ErrDBInternalError"),
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 99))
					mockSql.ExpectQuery(queryNonceSQLRegex).WillReturnError(sql.ErrNoRows)
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("100"))
					insertNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", insertNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 99, false))
					mockSql.ExpectExec(insertNonceSQLRegex).WillReturnError(errors.New("ErrDBInternalError"))
				}
			},
		},
		{
			desc:     "data conflict: nonce already exists",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			nonce:    100,
			expError: ErrNonceAlreadyExist,
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643", 100))
					mockSql.ExpectQuery(queryNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("100"))
				}
			},
		},
	}

	// perform tests
	for _, scenario := range scenarios {
		// create tracker
		tracker := &TrackerImp{mu: &sync.Mutex{}}

		// mock client
		mockClient := &mocks.EthClient{}
		tracker.ethClient = mockClient
		if scenario.configMockClient != nil {
			scenario.configMockClient(mockClient)
		}

		// mock sql
		db, mockSql, err := sqlmock.New()
		suite.Require().NoError(err)

		tracker.storer = &storer{dao: &dao{db}}
		defer db.Close()
		if scenario.configMockSql != nil {
			mockSql.MatchExpectationsInOrder(false)
			scenario.configMockSql(mockSql)
		}

		var wg sync.WaitGroup
		wg.Add(suite.gorutines)
		for i := 0; i < suite.gorutines; i++ {
			go func() {
				defer wg.Done()

				// execute code
				resErr := tracker.RestoreNonce(scenario.address, scenario.nonce)

				// testify results
				if scenario.expError != nil {
					suite.EqualError(scenario.expError, resErr.Error())
				}
			}()
		}
		wg.Wait()

		// make sure all expectations were met
		mockClient.AssertExpectations(suite.T())
		suite.NoError(mockSql.ExpectationsWereMet())
	}
}

func (suite *TrackerTestSuite) TestResetNonceAt() {
	// testing scenarios
	scenarios := []struct {
		desc             string
		address          string
		expError         error
		configMockClient func(mockClient *mocks.EthClient)
		configMockSql    func(mockSql sqlmock.Sqlmock)
	}{
		{
			desc:    "happy path",
			address: "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", context.Background(), common.HexToAddress("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643")).Return(uint64(100), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("100"))

					mockSql.ExpectBegin()
					updateTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", updateTheOnlyUpdatedNonceSQL(100, "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(updateTheOnlyUpdatedNonceSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
					deleteRecycledNoncesSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteRecycledNoncesSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(deleteRecycledNoncesSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
					mockSql.ExpectCommit()
				}
			},
		},
		{
			desc:     "input error: account not exists",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: ErrAccountNotExist,
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", context.Background(), common.HexToAddress("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643")).Return(uint64(100), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnError(sql.ErrNoRows)
				}
			},
		},
		{
			desc:     "ethClient issue: fail connecting Ethereum",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrConnectETHNetwork"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(0), errors.New("ErrConnectETHNetwork"))
			},
		},
		{
			desc:     "ethClient issue: Ethereum internal error",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrETHNetworkInternalError"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(0), errors.New("ErrETHNetworkInternalError"))
			},
		},
		{
			desc:     "DB issue: fail connecting DB when rollback",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrConnectDB"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", context.Background(), common.HexToAddress("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643")).Return(uint64(100), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("200"))

					mockSql.ExpectBegin()
					updateTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", updateTheOnlyUpdatedNonceSQL(100, "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(updateTheOnlyUpdatedNonceSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
					deleteRecycledNoncesSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteRecycledNoncesSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(deleteRecycledNoncesSQLRegex).WillReturnError(errors.New("ErrConnectDB"))
					mockSql.ExpectRollback().WillReturnError(errors.New("ErrConnectDB"))
				}
			},
		},
		{
			desc:     "DB issue: fail connecting DB when commit",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrConnectDB"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", context.Background(), common.HexToAddress("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643")).Return(uint64(100), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("200"))

					mockSql.ExpectBegin()
					updateTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", updateTheOnlyUpdatedNonceSQL(100, "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(updateTheOnlyUpdatedNonceSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
					deleteRecycledNoncesSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteRecycledNoncesSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(deleteRecycledNoncesSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
					mockSql.ExpectCommit().WillReturnError(errors.New("ErrConnectDB"))
				}
			},
		},
		{
			desc:     "DB issue: fail inserting sql",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrInsertSQL"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", context.Background(), common.HexToAddress("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643")).Return(uint64(100), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("200"))

					mockSql.ExpectBegin()
					updateTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", updateTheOnlyUpdatedNonceSQL(100, "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(updateTheOnlyUpdatedNonceSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
					deleteRecycledNoncesSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteRecycledNoncesSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(deleteRecycledNoncesSQLRegex).WillReturnError(errors.New("ErrInsertSQL"))
					mockSql.ExpectRollback().WillReturnError(errors.New("ErrInsertSQL"))
				}
			},
		},
		{
			desc:     "DB issue: DB internal error",
			address:  "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643",
			expError: errors.New("ErrDBInternalError"),
			configMockClient: func(mockclient *mocks.EthClient) {
				mockclient.On("PendingNonceAt", context.Background(), common.HexToAddress("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643")).Return(uint64(100), nil)
			},
			configMockSql: func(mockSql sqlmock.Sqlmock) {
				for i := 0; i < suite.gorutines; i++ {
					queryTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", queryTheOnlyUpdatedNonceSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectQuery(queryTheOnlyUpdatedNonceSQLRegex).WillReturnRows(sqlmock.NewRows([]string{"nonce"}).AddRow("200"))

					mockSql.ExpectBegin()
					updateTheOnlyUpdatedNonceSQLRegex := fmt.Sprintf("\\Q%v\\E", updateTheOnlyUpdatedNonceSQL(100, "0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(updateTheOnlyUpdatedNonceSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
					deleteRecycledNoncesSQLRegex := fmt.Sprintf("\\Q%v\\E", deleteRecycledNoncesSQL("0x99ad45666e02a89eaf06780d6cbd7ac9bb43f643"))
					mockSql.ExpectExec(deleteRecycledNoncesSQLRegex).WillReturnResult(sqlmock.NewResult(1, 1))
					mockSql.ExpectCommit().WillReturnError(errors.New("ErrDBInternalError"))
				}
			},
		},
	}

	// perform tests
	for _, scenario := range scenarios {
		// create tracker
		tracker := &TrackerImp{mu: &sync.Mutex{}}

		// mock client
		mockClient := &mocks.EthClient{}
		tracker.ethClient = mockClient
		if scenario.configMockClient != nil {
			scenario.configMockClient(mockClient)
		}

		// mock sql
		db, mockSql, err := sqlmock.New()
		suite.Require().NoError(err)

		tracker.storer = &storer{dao: &dao{db}}
		defer db.Close()
		if scenario.configMockSql != nil {
			mockSql.MatchExpectationsInOrder(false)
			scenario.configMockSql(mockSql)
		}

		var wg sync.WaitGroup
		wg.Add(suite.gorutines)
		for i := 0; i < suite.gorutines; i++ {
			go func() {
				defer wg.Done()

				// execute code
				resErr := tracker.ResetNonceAt(scenario.address)

				// testify results
				if scenario.expError != nil {
					suite.EqualError(scenario.expError, resErr.Error())
				}
			}()
		}
		wg.Wait()

		// make sure all expectations were met
		mockClient.AssertExpectations(suite.T())
		suite.NoError(mockSql.ExpectationsWereMet())
	}
}

func TestTrackerTestSuite(t *testing.T) {
	suite.Run(t, new(TrackerTestSuite))
}