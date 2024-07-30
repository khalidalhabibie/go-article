package postgres_test

import (
	"backend/app/article"
	"backend/app/models"
	"backend/pkg/utils"
	"log"

	"backend/app/article/repository/postgres"
	"backend/platform/database"
	"testing"

	mocket "github.com/Selvatico/go-mocket"
	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/suite"
)

type PostgresTestSuite struct {
	suite.Suite
	repository  article.Repository
	transaction utils.TransactionManagerClient
}

func TestArticleMysql(t *testing.T) {
	suite.Run(t, new(PostgresTestSuite))
}

func (suite *PostgresTestSuite) SetupSuite() {
	suite.repository = suite.initializeRepository()
}
func (suite *PostgresTestSuite) initializeArticle() *models.Article {
	return &models.Article{}
}

func (suite *PostgresTestSuite) initializeRepository() article.Repository {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	db := database.SetupDBTests()

	suite.transaction = utils.TransactionManagerNew(db)

	return postgres.New(db)
}

func (suite *PostgresTestSuite) TestInsert() {
	m := suite.initializeArticle()

	suite.Run("successful creation", func() {
		mocket.Catcher.Reset().NewMock().WithError(nil)

		err := suite.repository.Insert(m, nil)
		suite.Nil(err)
	})

	suite.Run("failed creation with unknown databse error", func() {
		mocket.Catcher.Reset().NewMock().WithQueryException()

		err := suite.repository.Insert(m, nil)
		suite.Assert().NotNil(err)

	})

	suite.Run("success with transaction", func() {
		mocket.Catcher.Reset().NewMock().WithError(nil)

		tx := suite.transaction.NewTransaction()

		err := suite.repository.Insert(m, tx)
		log.Println("error ", err)

		suite.Nil(err)
	})

	suite.Run("with error response", func() {
		mocket.Catcher.Reset().NewMock().WithQueryException()

		tx := suite.transaction.NewTransaction()

		err := suite.repository.Insert(m, tx)
		log.Println("error ", err)

		suite.NotNil(err)
	})

}

func (suite *PostgresTestSuite) TestFindAll() {
	config := utils.NewPaginationConfig(20, 0, "id")
	suite.Run("1 result found", func() {

		reply := []map[string]interface{}{{"id": 1}}
		mocket.Catcher.Reset().NewMock().WithReply(reply)

		result, err := suite.repository.FindAll(config)
		suite.Nil(err)
		suite.Assert().Equal(1, len(result))
		suite.Assert().Equal(uint64(1), result[0].ID)
	})

	suite.Run("multiple result found", func() {
		reply := []map[string]interface{}{
			{"id": 1},
			{"id": 2},
			{"id": 3},
		}
		mocket.Catcher.Reset().NewMock().WithReply(reply)

		result, err := suite.repository.FindAll(config)
		suite.Nil(err)
		suite.Assert().Equal(3, len(result))
	})

	suite.Run("with unknown database error", func() {
		mocket.Catcher.Reset().NewMock().WithQueryException()
		result, err := suite.repository.FindAll(config)

		suite.Assert().NotNil(err)
		suite.Assert().Equal(fiber.ErrUnprocessableEntity, err)
		suite.Assert().Nil(result)
	})

	suite.Run("Record not found", func() {
		reply := []map[string]interface{}{} // empty response
		mocket.Catcher.Reset().NewMock().WithReply(reply)

		result, err := suite.repository.FindAll(config)
		suite.Assert().Nil(err)
		suite.Assert().Equal([]models.Article{}, result)
	})
}

func (suite *PostgresTestSuite) TestCount() {
	config := utils.NewDefaultPaginationConfig()
	suite.Run("1 topUp found", func() {

		reply := []map[string]interface{}{{"count": 1}}
		mocket.Catcher.Reset().NewMock().WithReply(reply)

		total, err := suite.repository.Count(config)
		suite.Nil(err)
		suite.Assert().Equal(int64(1), total)
	})

	suite.Run("3 topUp found", func() {
		reply := []map[string]interface{}{{"count": 3}}
		mocket.Catcher.Reset().NewMock().WithReply(reply)

		total, err := suite.repository.Count(config)
		suite.Nil(err)
		suite.Assert().Equal(int64(3), total)
	})

	suite.Run("with unknown database error", func() {
		mocket.Catcher.Reset().NewMock().WithQueryException()
		total, err := suite.repository.Count(config)

		suite.Assert().NotNil(err)
		suite.Assert().Equal(fiber.ErrUnprocessableEntity, err)
		suite.Assert().Equal(int64(0), total)
	})

	suite.Run("Record not found", func() {
		reply := []map[string]interface{}{{"count": 0}}
		mocket.Catcher.Reset().NewMock().WithReply(reply)

		total, err := suite.repository.Count(config)
		suite.Assert().Nil(err)
		suite.Assert().Equal(int64(0), total)
	})
}
