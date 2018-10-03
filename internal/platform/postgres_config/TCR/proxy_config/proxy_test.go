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
  listProxyUUDs := suite.ProxyExecutor.GetListOfProxyUUIDs()
  suite.Equal(0, len(*listProxyUUDs))
}

func (suite *ProxyTestSuite) TestUpsertProxyRecord() {
  defer func() {
    errPanic := recover();
    suite.Nil(errPanic)
  }()
  proxyRecord1 := ProxyRecord{UUID: UUID1}
  proxyRecord2 := ProxyRecord{UUID: UUID2}
  suite.ProxyExecutor.UpsertProxyRecord(&proxyRecord1)
  suite.ProxyExecutor.UpsertProxyRecord(&proxyRecord2)
}

func (suite *ProxyTestSuite) TestNonEmptyQueryForGetListOfProxyUUIDs() {
  proxyRecord1 := ProxyRecord{UUID: UUID1}
  proxyRecord2 := ProxyRecord{UUID: UUID2}
  suite.ProxyExecutor.UpsertProxyRecord(&proxyRecord1)
  suite.ProxyExecutor.UpsertProxyRecord(&proxyRecord2)
  expectedListProxyUUDs := []string {UUID1, UUID2,}
  listProxyUUDs := suite.ProxyExecutor.GetListOfProxyUUIDs()
  suite.Equal(expectedListProxyUUDs, *listProxyUUDs)
}

func (suite *ProxyTestSuite) TestVerifyNonExitingProxyUUID() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
       message := error_config.CreatedErrorInfoFromString(errPanic)
       suite.Equal(error_config.NoProxyUUIDExisting, message.ErrorCode)
    }
  }()
  suite.ProxyExecutor.VerifyProxyRecordExisting(UUID3)
}

func (suite *ProxyTestSuite) TestDeleteProxyRecord() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoProxyUUIDExisting, message.ErrorCode)
    }
  }()
  proxyRecord1 := ProxyRecord{UUID: UUID1}
  suite.ProxyExecutor.UpsertProxyRecord(&proxyRecord1)
  suite.ProxyExecutor.DeleteProxyRecord(UUID1)
  suite.ProxyExecutor.VerifyProxyRecordExisting(UUID1)
}

func TestProxyTestSuite(t *testing.T) {
  suite.Run(t, new(ProxyTestSuite))
}
