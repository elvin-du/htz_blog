package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type articleModel struct{}

var _article articleModel

func ArticleModel() *articleModel {
	return &_article
}

func (this *articleModel) InfoById(id int) (*Article, error) {
	var a Article

	o := orm.NewOrm()
	err := o.Raw("SELECT id,title,content,nav_name,content_abstract,read_count, created_at FROM articles WHERE id = ?", id).QueryRow(&a)
	if nil != err {
		beego.Error(err)
		if orm.ErrNoRows == err {
			return &a, nil
		}

		return nil, err
	}
	return &a, nil
}

func (this *articleModel) Articles(nav string, page, num int) ([]*Article, error) {
	var as = []*Article{}

	o := orm.NewOrm()
	_, err := o.Raw("SELECT id,title,content,content_abstract,read_count,nav_name,created_at FROM articles WHERE nav_name=? ORDER BY created_at DESC LIMIT ?,?", nav, page*num, num).QueryRows(&as)
	if nil != err {
		beego.Error(err)
		if orm.ErrNoRows == err {
			return as, nil
		}

		return nil, err
	}

	return as, nil
}

func (this *articleModel) Home(page, num int) ([]*Article, error) {
	var as = []*Article{}

	o := orm.NewOrm()
	_, err := o.Raw("SELECT id,title,content,content_abstract,read_count,nav_name,created_at FROM articles ORDER BY created_at DESC LIMIT ?,?", page*num, num).QueryRows(&as)
	if nil != err {
		beego.Error(err)
		if orm.ErrNoRows == err {
			return as, nil
		}

		return nil, err
	}

	return as, nil
}

func (this *articleModel) Add(title, content, nav, ctime string) error {
	o := orm.NewOrm()
	_, err := o.Raw("INSERT articles(id,title,content,content_abstract,read_count,nav_name,created_at FROM articles ORDER BY created_at DESC LIMIT ?,?").Exec()
	if nil != err {
		beego.Error(err)
		if orm.ErrNoRows == err {
			return nil
		}

		return err
	}

	return nil
}
