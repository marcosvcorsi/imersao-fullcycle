package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/mvcorsi/imersao-fullcycle/codepix-go/application/usecases"
	"github.com/mvcorsi/imersao-fullcycle/codepix-go/infra/repositories"
)

func TransactionUseCaseFactory(database *gorm.DB) usecases.TransactionUseCase {
	pixRepository := repositories.PixKeyRepositoryDb{Db: database}
	transactionRepository := repositories.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecases.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}

	return transactionUseCase
}
