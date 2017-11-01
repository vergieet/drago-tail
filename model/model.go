package model

type Model interface {
	Find(id int) struct{}
	Where() struct{}
	All() struct{}
}

func (u User)All()User{
	return u
}
