package cache

import (
	"TestTask/internal/models"
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

type CacheService struct {
	cache *cache.Cache
}

func NewCacheService() *CacheService {
	log.Println("Initializing Cache Service...")
	return &CacheService{
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *CacheService) SetOrder(orderID int, order *models.Order) {
	key := fmt.Sprintf("order_%d", orderID)
	c.cache.Set(key, order, cache.DefaultExpiration)
	log.Printf("Order with ID %d has been cached under key '%s'", orderID, key)
}

func (c *CacheService) GetOrder(orderID int) (*models.Order, bool) {
	key := fmt.Sprintf("order_%d", orderID)
	order, found := c.cache.Get(key)
	if found {
		log.Printf("Order with ID %d found in cache under key '%s'", orderID, key)
		return order.(*models.Order), true
	}

	log.Printf("Order with ID %d not found in cache under key '%s'", orderID, key)
	return nil, false
}

func (c *CacheService) DeleteOrder(orderID int) {
	key := fmt.Sprintf("order_%d", orderID)
	c.cache.Delete(key)
	log.Printf("Order with ID %d has been removed from cache under key '%s'", orderID, key)
}

func (c *CacheService) SetOrders(key string, orders []models.Order) {
	c.cache.Set(key, orders, cache.DefaultExpiration)
	log.Printf("Orders have been cached under key '%s'", key)
}

func (c *CacheService) GetOrders(key string) ([]models.Order, bool) {
	orders, found := c.cache.Get(key)
	if found {
		log.Printf("Orders found in cache under key '%s'", key)
		return orders.([]models.Order), true
	}

	log.Printf("Orders not found in cache under key '%s'", key)
	return nil, false
}
