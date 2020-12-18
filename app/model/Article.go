package model

//CREATE TABLE `articles` (
//  `id` int(10) NOT NULL AUTO_INCREMENT,
//  `category_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章分类id',
//  `title` varchar(20) NOT NULL COMMENT '标题',
//  `tags` text COMMENT 'tag列表',
//  `reads` int(10) unsigned DEFAULT '0' COMMENT '阅读数',
//  `comments` int(10) DEFAULT '0' COMMENT '评论数',
//  `content` text COMMENT '文章内容',
//  `updated_at` int(10) NOT NULL DEFAULT '0',
//  `created_at` int(10) NOT NULL DEFAULT '0',
//  PRIMARY KEY (`id`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章';
type Article struct {
	Base
	Id         int
	CategoryId int
	Title      string `gorm:"type:varchar;size:20"`
	Tags       string `gorm:"type:text"`
	Reads      int
	Comments   int
	Content    string `gorm:"type:text"`
}

func (Article) TableName() string {
	return "articles"
}
