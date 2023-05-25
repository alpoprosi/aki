package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func RecoveryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("recover: %w", err)

				c.Error(echo.ErrInternalServerError)
			}
		}()

		return next(c)
	}
}

func DBTransactionMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			txHandle := db.Begin()
			log.Print("beginning database transaction")

			defer func() {
				if r := recover(); r != nil {
					txHandle.Rollback()
				}
			}()

			c.Set("db_trx", txHandle)

			err := next(c)
			if err != nil {
				log.Printf("rolling back transaction during: %v", err)

				merr := txHandle.Rollback().Error
				if merr != nil {
					log.Errorf("rolling back transaction: %v", merr)
				}

				return err
			}

			log.Print("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				return fmt.Errorf("tx committing: %w", err)
			}

			return nil
		}
	}
}
