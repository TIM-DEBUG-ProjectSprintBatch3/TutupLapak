package purchaseRepository

import (
	"context"

	Entity "github.com/TimDebug/FitByte/src/model/entities/purchase"
	"github.com/jackc/pgx/v5"
)

type IPurchaseRepository interface {
	InsertInto(tx pgx.Tx, ctx context.Context, entity Entity.Purchase) (string, error)
	// Create(ctx *fiber.Ctx, pool *pgxpool.Pool, activity Entity.Activity) (activityId string, err error)
	// GetValidCaloriesFactors(ctx *fiber.Ctx, pool *pgxpool.Pool, activityId, userId string) (*Entity.CaloriesFactor, error)
	// GetActivityByUserId(ctx *fiber.Ctx, pool *pgxpool.Pool, activityId, userId string) (string, error)
	// Update(ctx *fiber.Ctx, pool *pgxpool.Pool, activity Entity.Activity) error
	// Delete(ctx *fiber.Ctx, pool *pgxpool.Pool, activityId, userId string) error
}
