package service

import "be-golang-chapter-26/LA-Chapter-26A-H/storage"

// NewInMemoryStorage menyediakan instance dari InMemoryStorage
func NewCachingData() *storage.CachingData {
	return &storage.CachingData{}
}

// NewDatabaseStorage menyediakan instance dari DatabaseStorage
func NewDatabaseStorage() *storage.DatabaseStorage {
	return &storage.DatabaseStorage{}
}
