package redis

import (
	"backend/app/article"
	"backend/app/article/delivery/http/response"
	"backend/app/models"
	"backend/pkg/utils"
	"backend/platform/cache"

	// "backend/app/article/repository/postgres"

	"testing"

	mocket "github.com/Selvatico/go-mocket"

	"github.com/stretchr/testify/suite"
)

type RedisTestSuite struct {
	suite.Suite
	cache article.Cache
}

func TestArticleredis(t *testing.T) {
	suite.Run(t, new(RedisTestSuite))
}

func (suite *RedisTestSuite) SetupSuite() {
	suite.cache = suite.initializeRepository()

}
func (suite *RedisTestSuite) initializePaginationConfig() response.Index {

	data := []models.Article{}

	meta := utils.PaginationMeta{}

	return response.Index{
		Data: data,
		Meta: meta,
	}
}

func (suite *RedisTestSuite) initializeResponse() utils.PaginationConfig {

	var request utils.PaginationConfig

	return request
}

func (suite *RedisTestSuite) initializeRepository() article.Cache {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	db := cache.SetUpRedisForTesting()

	return New(db)
}

func (suite *RedisTestSuite) TestSet() {
	m := suite.initializeResponse()
	p := suite.initializePaginationConfig()

	suite.Run("successful set", func() {
		err := suite.cache.Set(m, p)
		suite.Nil(err)
	})

}

func (suite *RedisTestSuite) TestGet() {
	m := suite.initializeResponse()
	p := suite.initializePaginationConfig()

	suite.Run("successful set", func() {

		err := suite.cache.Set(m, p)
		suite.Nil(err)

		value, err := suite.cache.Get(m)
		suite.Nil(err)
		suite.NotNil(value)
	})
	suite.Run("Not found", func() {
		err := suite.cache.FlushAll()
		suite.Nil(err)

		value, err := suite.cache.Get(m)
		suite.NotNil(err)
		suite.Nil(value)
	})

}

func (suite *RedisTestSuite) TestFlush() {
	m := suite.initializeResponse()
	p := suite.initializePaginationConfig()

	suite.Run("successful Flush with data", func() {

		err := suite.cache.Set(m, p)
		suite.Nil(err)

		err = suite.cache.FlushAll()
		suite.Nil(err)

	})
	suite.Run("successful Flush without data", func() {

		err := suite.cache.FlushAll()
		suite.Nil(err)

	})

}
