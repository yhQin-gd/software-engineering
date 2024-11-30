package image_r

import "gorm.io/gorm"

// IsImageFavoritedByUser checks if a user has favorited a specific image
func IsImageFavoritedByUser(db *gorm.DB, username string, url string) (bool, error) {
    var count int64
    err := db.Table("favoritedimage").Where("username = ? AND result = ?", username, url).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}
