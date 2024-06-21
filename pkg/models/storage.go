package models 


type BlogStorage interface{
    Create(*Blog) error 
    Delete(uint, uint) error
    Update(id uint, newB *Blog, userId uint) (*Blog, error)
    GetBlog(uint) (*Blog, error)
    GetBlogs() ([]Blog, error)
}


type LikeStorage interface{
    Like(uint, uint) error 
}

type UserStorage interface{
    Create(*User) (error)
    Update(string, uint) (*User ,error) 
    Get(string) (*User, error)
    Delete(uint) error 
}
