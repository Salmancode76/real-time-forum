package repository

import (
	"database/sql"
	"fmt"
	"log"
	"real-time-forum/internal/models/entities"
	"strings"
)

type PostModel struct {
	DB *sql.DB
}

func (p * PostModel) GetPostByUserID(userID string) ([]entities.Post, error){
	var Posts []entities.Post

	stmt :=`
			SELECT 
			p.id,
			p.title,
			p.body,
			p.created_at,
			Username,
			GROUP_CONCAT(c.name, ', ') AS category_names
		FROM post p
		INNER JOIN post_category pc ON p.id = pc.post_id
		INNER JOIN category c ON pc.category_id = c.id
		INNER JOIN User u on u.UserID = p.user_id
		where p.user_id = ?
		GROUP BY p.id;
	`

	rows,err:= p.DB.Query(stmt,userID)

	if err!= nil{
		log.Println("No Post found :",err)
		return Posts,err
	}

	for rows.Next(){
		var Post entities.Post
		var temp_category string
		rows.Scan(&Post.ID,&Post.Title,&Post.Content,&Post.Date,&Post.UserID,&temp_category)

		 cat := strings.Split((temp_category), ",")
	
		Post.Categories = cat

		Posts = append(Posts, Post)
		
	}

	return Posts,nil

}

func (p *PostModel) GetPostsByCategory(Category string)([]entities.Post, error){

	var Posts []entities.Post
		stmt :=`
		
		SELECT 
			p.id,
			p.title,
			p.body,
			p.created_at,
			Username,
			GROUP_CONCAT(c.name, ', ') AS category_names
		FROM post p
		INNER JOIN post_category pc ON p.id = pc.post_id
		INNER JOIN category c ON pc.category_id = c.id
		INNER JOIN User u on UserID = p.user_id
		WHERE p.id IN (
			SELECT DISTINCT pc1.post_id 
			FROM post_category pc1
			INNER JOIN category c1 ON pc1.category_id = c1.id
			WHERE c1.name = ?
		)
		GROUP BY p.id
		HAVING category_names LIKE '%' || ? || '%';


		
			`
	rows,err:= p.DB.Query(stmt,Category,Category)

		if err!=nil{
		log.Println("Can't get posts")
		return Posts,err
	}

	for rows.Next(){
			var Post entities.Post
		var temp_category string
		rows.Scan(&Post.ID,&Post.Title,&Post.Content,&Post.Date,&Post.UserID,&temp_category)

		 cat := strings.Split((temp_category), ",")
	
		Post.Categories = cat

		Posts = append(Posts, Post)
		
	}
	return Posts,nil

}

func (p * PostModel) GetPostComment(postID string)([]entities.Comment,error){
	var Comments []entities.Comment
	stmt :=`
	SELECT id,
       post_id,
       user_id,
       comment,
       username,
       created_at
  FROM comments
  Inner join  User  ON
   User.UserID = user_id
  where post_id = ?;
`

	rows,err := p.DB.Query(stmt,postID)

	if err!=nil{
		log.Println("Can't get comments")
		return Comments,err
	}
	for rows.Next(){
		var comment entities.Comment
		err = rows.Scan(&comment.ID,&comment.Postid,&comment.UserID,&comment.Comment,&comment.Username,&comment.Date)
		if err !=nil{
		return Comments,err
		}
		Comments = append(Comments, comment)
	}
	return Comments,nil
}

func (p *PostModel) FindPost(id string) (entities.Post,error) {
	var Post entities.Post
	var temp_category string

	stmt := `SELECT 
			post.id,
			post.title,
			post.body,
			post.created_at,
			User.Username,
			GROUP_CONCAT(category.name) AS categories
		FROM post
		INNER JOIN User
			ON User.UserID = post.user_id
		INNER JOIN post_category
			ON post_category.post_id = post.id
		INNER JOIN category
			ON category.id = post_category.category_id
		WHERE post.id = ?
		GROUP BY 
			post.id,
			post.title,
			post.body,
			post.created_at,
			User.Username;`

	row := p.DB.QueryRow(stmt, id)

	err := row.Scan(&Post.ID, &Post.Title, &Post.Content, &Post.Date, &Post.UserID, &temp_category)
	
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No post found with ID:", id)
		} else {
			log.Println("Failed to fetch the post:", err)
		}
		return Post,err
	}

	Post.Categories = strings.Split(temp_category, ",")

	log.Println(Post)
	return Post,nil
}

func (p * PostModel) FetchAllPost()([]entities.Post,error){
	var Posts []entities.Post
	stmt :=`
	
				SELECT 
			post.id,
			post.title,
			post.body,
			post.created_at,
			User.Username,
			GROUP_CONCAT(category.name) AS categories
		FROM post
		INNER JOIN User
			ON User.UserID = post.user_id
		INNER JOIN post_category
			ON post_category.post_id = post.id
		INNER JOIN category
			ON category.id = post_category.category_id
		GROUP BY 
			post.id,
			post.title,
			post.body,
			post.created_at,
			User.Username;
  `

  row,err := p.DB.Query(stmt)

  	if err != nil {
		return  Posts,fmt.Errorf("failed to fetch all post: %w", err)
	}


	for row.Next(){
				var Post entities.Post
		var temp_category string
		row.Scan(&Post.ID,&Post.Title,&Post.Content,&Post.Date,&Post.UserID,&temp_category)

		 cat := strings.Split((temp_category), ",")
	
		Post.Categories = cat

		Posts = append(Posts, Post)
		
	}

	return Posts,nil
}
func (p *PostModel) InsertComment (userID,comment,postID string)(error){
	stmt:=`INSERT INTO comments (
                         post_id,
                         user_id,
                         comment,
                         created_at
                     )
                     VALUES (
                     ?,
                     ?,
                     ?,
                     datetime('now')
                     );
	`
	_,err := p.DB.Exec(stmt,postID,userID,comment)

	if err!=nil{
		return fmt.Errorf("failed to insert Comment: %w", err)

	}
	return nil

}

func (p * PostModel) Insert(id,title,content string , categories[]string ) (error){
	
	stmt:=`INSERT INTO post (
                     title,
                     body,
                     created_at,
                     user_id
                 )
                 VALUES (
                     ?,
                     ?,
                     datetime('now'),
                     ?
                 );
`

	result,err := p.DB.Exec(stmt,title,content,id)

	if err != nil {
		return fmt.Errorf("failed to insert post: %w", err)
	}

	PostID,err:= result.LastInsertId()

	stmt2:= `INSERT INTO post_category (
                              post_id,
                              category_id
                          )
                          VALUES (
                              ?,
                              ?
                          );
`


	All_categories,err := p.GetAllCategories();

	if err !=nil{
		return err
	}
	for i :=0;i<len(categories);i++{
		fmt.Println(All_categories[categories[i]])



		_,err:= p.DB.Exec(stmt2,PostID,All_categories[categories[i]])
		if err !=nil{
			return fmt.Errorf("failed to insert cate post: %w", err)
		}

		
	}
	
	return nil


}


func (p *PostModel) GetAllCategories() (map[string]string, error) {
	stmt := `SELECT id,
       name
  FROM category;
`

	rows,err :=p.DB.Query(stmt)

		if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}
	defer rows.Close()

	categories := make(map[string]string)

	for rows.Next(){
		var id,name string
		rows.Scan(&id,&name)
		categories[name]=id
	}
	

	return categories,nil

}