package protoProductController

import (
	"context"
	"fmt"

	functionCallerInfo "github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/logger/helper"
	loggerZap "github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/logger/zap"
	"github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/service"
	"github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/service/proto/product"
	"github.com/samber/do/v2"
)

type ProtoProductController struct {
	productService service.ProductServiceInterface
	logger         loggerZap.LoggerInterface

	// Embed UnimplementedProductServiceServer to satisfy gRPC interface
	product.UnimplementedProductServiceServer
}

func NewProtoProductControllerInject(i do.Injector) (*ProtoProductController, error) {
	_productService := do.MustInvoke[service.ProductServiceInterface](i)
	_logger := do.MustInvoke[loggerZap.LoggerInterface](i)

	return &ProtoProductController{
		productService: _productService,
		logger:         _logger,
	}, nil
}

func (this ProtoProductController) GetProductDetailsWithId(ctx context.Context, request *product.ProductRequest) (*product.ProductsWithIdResponse, error) {
	result := product.ProductsWithIdResponse{}

	// todo; check cache ada profile user tidak?
	// belum tau kalian mau pakai cache middleware atau internal jadi saya kasih catatan saja-ad1ee

	// response, err := puc.userService.GetUserProfilesWithId(ctx, request.UserIds)
	response, err := this.productService.GetAll(ctx, request.ProductIds)
	if err != nil {
		fmt.Sprintf("GetUserDetailsWithId Error %s", err.Error())
		this.logger.Error(err.Error(), functionCallerInfo.UserControllerRegister)
		return nil, err
	}

	// Populate the response with the data from user profiles.
	for _, profile := range response {
		productResponse := &product.ProductWithIdResponse{
			// UserId:            profile.UserId,
			// Email:             profile.Email,
			// Phone:             profile.Phone,
			// BankAccountName:   profile.BankAccountName,
			// BankAccountHolder: profile.BankAccountHolder,
			// BankAccountNumber: profile.BankAccountNumber,
		}
		result.Products = append(result.Products, productResponse)
	}
	return &result, nil

	return nil, nil
}

// mustEmbedUnimplementedProductServiceServer
func (this ProtoProductController) mustEmbedUnimplementedProductServiceServer() {
	return
}
