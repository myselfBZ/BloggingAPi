package handlers

import "github.com/myselfBZ/BloggingAPI/pkg/models"

type Handler struct{
    BlogStorage models.BlogStorage
    LikeStorage models.LikeStorage
    UserStorage models.UserStorage
}


func NewHandler(BlogStorage models.BlogStorage, Like models.LikeStorage, User models.UserStorage) *Handler{
    return &Handler{
        BlogStorage:BlogStorage, 
        LikeStorage: Like,
        UserStorage: User,
    }


}
