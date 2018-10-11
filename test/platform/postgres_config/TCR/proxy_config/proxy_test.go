package proxy_config

import (
  "testing"
  "BigBang/internal/platform/postgres_config/client_config"
  "github.com/stretchr/testify/suite"
  "BigBang/internal/pkg/error_config"
)

const UUID1 = "uuid1"
const UUID2 = "uuid2"
const UUID3 = "uuid3"
const UUID4 = "uuid4"
const UUID5 = "uuid5"

var ProxyRecord1 = ProxyRecord{ID: 1, UUID: UUID1}
var ProxyRecord2 = ProxyRecord{ID: 2, UUID: UUID2}
var ProxyRecord3 = ProxyRecord{ID: 3, UUID: UUID3}
var ProxyRecord4 = ProxyRecord{ID: 4, UUID: UUID4}
var ProxyRecord5 = ProxyRecord{ID: 5, UUID: UUID5}

type ProxyTestSuite struct {
  suite.Suite
  ProxyExecutor ProxyExecutor
}

func (suite *ProxyTestSuite) SetupSuite() {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  suite.ProxyExecutor = ProxyExecutor{*postgresBigBangClient}
  suite.ProxyExecutor.DeleteProxyTable()
  suite.ProxyExecutor.CreateProxyTable()
}

func (suite *ProxyTestSuite) TearDownSuite() {
  suite.ProxyExecutor.DeleteProxyTable()
  suite.ProxyExecutor.C.Close()
}

func (suite *ProxyTestSuite) SetupTest() {
  suite.ProxyExecutor.ClearProxyTable()
}

func (suite *ProxyTestSuite) TestEmptyQueryForGetListOfProxyUUIDs() {
  listProxyUUDs := suite.ProxyExecutor.GetListOfProxyByCursor(0, 100)
  suite.Equal(0, len(*listProxyUUDs))
}

func (suite *ProxyTestSuite) TestUpsertProxyRecord() {
  defer func() {
    errPanic := recover();
    suite.Nil(errPanic)
  }()

  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord1)
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord2)
}

func (suite *ProxyTestSuite) TestEmptyQueryForGetProxyRecord() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoProxyUUIDExisting, message.ErrorCode)
    }
  }()
  suite.ProxyExecutor.GetProxyRecord(UUID1)
}

func (suite *ProxyTestSuite) TestNonEmptyQueryForGetProxyRecord() {
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord1)
  proxyRecord := suite.ProxyExecutor.GetProxyRecord(UUID1)
  suite.Equal(UUID1, proxyRecord.UUID)
}

func (suite *ProxyTestSuite) TestNonEmptyQueryForGetListOfProxyByCursorFirstQuery() {
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord1)
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord2)
  expectedListProxy := []ProxyRecord {ProxyRecord2, ProxyRecord1}
  listProxyUUDs := suite.ProxyExecutor.GetListOfProxyByCursor(0, 100)
  suite.Equal(expectedListProxy, *listProxyUUDs)
}

func (suite *ProxyTestSuite) TestNonEmptyQueryForGetListOfProxyByCursorInterQuery() {
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord1)
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord2)
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord3)
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord4)
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord5)
  expectedListProxy := []ProxyRecord {ProxyRecord4, ProxyRecord3}
  listProxyUUDs := suite.ProxyExecutor.GetListOfProxyByCursor(
    4, 2)
  suite.Equal(expectedListProxy, *listProxyUUDs)
}

func (suite *ProxyTestSuite) TestNonEmptyQueryForGetListOfProxyByCursorFinalQuery() {
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord1)
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord2)
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord3)
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord4)
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord5)
  expectedListProxy := []ProxyRecord {ProxyRecord3, ProxyRecord2, ProxyRecord1}
  listProxyUUDs := suite.ProxyExecutor.GetListOfProxyByCursor(
    3, 6)
  suite.Equal(expectedListProxy, *listProxyUUDs)
}

func (suite *ProxyTestSuite) TestVerifyNonExitingProxyUUID() {
  suite.False(suite.ProxyExecutor.VerifyProxyRecordExisting(UUID3))
}

func (suite *ProxyTestSuite) TestDeleteProxyRecord() {
  suite.ProxyExecutor.UpsertProxyRecord(&ProxyRecord1)
  suite.ProxyExecutor.DeleteProxyRecord(UUID1)
  suite.False(suite.ProxyExecutor.VerifyProxyRecordExisting(UUID1))
}

func TestProxyTestSuite(t *testing.T) {
  suite.Run(t, new(ProxyTestSuite))
}
