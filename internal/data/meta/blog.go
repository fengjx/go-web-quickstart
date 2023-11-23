package meta

import (
    "github.com/fengjx/daox/sqlbuilder"
    "github.com/fengjx/daox/sqlbuilder/ql"


    "time"

)


// BlogM 博客表
// auto generate by gen cmd tool
type BlogM struct {
    ID string
    UID string
    Title string
    Content string
    CreateTime string
    Utime string
    Ctime string
}

var BlogMeta = BlogM{
    ID: "id",
    UID: "uid",
    Title: "title",
    Content: "content",
    CreateTime: "create_time",
    Utime: "utime",
    Ctime: "ctime",
}




func (m BlogM) IdIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.ID).In(args...)
}

func (m BlogM) IdNotIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.ID).NotIn(args...)
}

func (m BlogM) IdEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).EQ(val)
}

func (m BlogM) IdNotEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).NotEQ(val)
}

func (m BlogM) IdLT(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).LT(val)
}

func (m BlogM) IdLTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).LTEQ(val)
}

func (m BlogM) IdGT(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).GT(val)
}

func (m BlogM) IdGTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).GTEQ(val)
}

func (m BlogM) IdLike(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).Like(val)
}

func (m BlogM) IdNotLike(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).NotLike(val)
}

func (m BlogM) IdDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.ID)
}

func (m BlogM) IdAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.ID)
}



func (m BlogM) UidIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.UID).In(args...)
}

func (m BlogM) UidNotIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.UID).NotIn(args...)
}

func (m BlogM) UidEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).EQ(val)
}

func (m BlogM) UidNotEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).NotEQ(val)
}

func (m BlogM) UidLT(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).LT(val)
}

func (m BlogM) UidLTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).LTEQ(val)
}

func (m BlogM) UidGT(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).GT(val)
}

func (m BlogM) UidGTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).GTEQ(val)
}

func (m BlogM) UidLike(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).Like(val)
}

func (m BlogM) UidNotLike(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).NotLike(val)
}

func (m BlogM) UidDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.UID)
}

func (m BlogM) UidAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.UID)
}



func (m BlogM) TitleIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Title).In(args...)
}

func (m BlogM) TitleNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Title).NotIn(args...)
}

func (m BlogM) TitleEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Title).EQ(val)
}

func (m BlogM) TitleNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Title).NotEQ(val)
}

func (m BlogM) TitleLT(val string) sqlbuilder.Column {
	return ql.Col(m.Title).LT(val)
}

func (m BlogM) TitleLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Title).LTEQ(val)
}

func (m BlogM) TitleGT(val string) sqlbuilder.Column {
	return ql.Col(m.Title).GT(val)
}

func (m BlogM) TitleGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Title).GTEQ(val)
}

func (m BlogM) TitleLike(val string) sqlbuilder.Column {
	return ql.Col(m.Title).Like(val)
}

func (m BlogM) TitleNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Title).NotLike(val)
}

func (m BlogM) TitleDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Title)
}

func (m BlogM) TitleAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Title)
}



func (m BlogM) ContentIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Content).In(args...)
}

func (m BlogM) ContentNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Content).NotIn(args...)
}

func (m BlogM) ContentEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Content).EQ(val)
}

func (m BlogM) ContentNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Content).NotEQ(val)
}

func (m BlogM) ContentLT(val string) sqlbuilder.Column {
	return ql.Col(m.Content).LT(val)
}

func (m BlogM) ContentLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Content).LTEQ(val)
}

func (m BlogM) ContentGT(val string) sqlbuilder.Column {
	return ql.Col(m.Content).GT(val)
}

func (m BlogM) ContentGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Content).GTEQ(val)
}

func (m BlogM) ContentLike(val string) sqlbuilder.Column {
	return ql.Col(m.Content).Like(val)
}

func (m BlogM) ContentNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Content).NotLike(val)
}

func (m BlogM) ContentDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Content)
}

func (m BlogM) ContentAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Content)
}



func (m BlogM) CreateTimeIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.CreateTime).In(args...)
}

func (m BlogM) CreateTimeNotIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.CreateTime).NotIn(args...)
}

func (m BlogM) CreateTimeEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.CreateTime).EQ(val)
}

func (m BlogM) CreateTimeNotEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.CreateTime).NotEQ(val)
}

func (m BlogM) CreateTimeLT(val int64) sqlbuilder.Column {
	return ql.Col(m.CreateTime).LT(val)
}

func (m BlogM) CreateTimeLTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.CreateTime).LTEQ(val)
}

func (m BlogM) CreateTimeGT(val int64) sqlbuilder.Column {
	return ql.Col(m.CreateTime).GT(val)
}

func (m BlogM) CreateTimeGTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.CreateTime).GTEQ(val)
}

func (m BlogM) CreateTimeLike(val int64) sqlbuilder.Column {
	return ql.Col(m.CreateTime).Like(val)
}

func (m BlogM) CreateTimeNotLike(val int64) sqlbuilder.Column {
	return ql.Col(m.CreateTime).NotLike(val)
}

func (m BlogM) CreateTimeDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.CreateTime)
}

func (m BlogM) CreateTimeAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.CreateTime)
}



func (m BlogM) UtimeIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Utime).In(args...)
}

func (m BlogM) UtimeNotIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Utime).NotIn(args...)
}

func (m BlogM) UtimeEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).EQ(val)
}

func (m BlogM) UtimeNotEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).NotEQ(val)
}

func (m BlogM) UtimeLT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).LT(val)
}

func (m BlogM) UtimeLTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).LTEQ(val)
}

func (m BlogM) UtimeGT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).GT(val)
}

func (m BlogM) UtimeGTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).GTEQ(val)
}

func (m BlogM) UtimeLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).Like(val)
}

func (m BlogM) UtimeNotLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).NotLike(val)
}

func (m BlogM) UtimeDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Utime)
}

func (m BlogM) UtimeAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Utime)
}



func (m BlogM) CtimeIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Ctime).In(args...)
}

func (m BlogM) CtimeNotIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Ctime).NotIn(args...)
}

func (m BlogM) CtimeEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).EQ(val)
}

func (m BlogM) CtimeNotEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).NotEQ(val)
}

func (m BlogM) CtimeLT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).LT(val)
}

func (m BlogM) CtimeLTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).LTEQ(val)
}

func (m BlogM) CtimeGT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).GT(val)
}

func (m BlogM) CtimeGTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).GTEQ(val)
}

func (m BlogM) CtimeLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).Like(val)
}

func (m BlogM) CtimeNotLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).NotLike(val)
}

func (m BlogM) CtimeDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Ctime)
}

func (m BlogM) CtimeAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Ctime)
}
