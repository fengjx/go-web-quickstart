package meta

import (
    "github.com/fengjx/daox/sqlbuilder"
    "github.com/fengjx/daox/sqlbuilder/ql"


    "time"

)


// UserM 用户信息表
// auto generate by gen cmd tool
type UserM struct {
    ID string
    Username string
    Pwd string
    Salt string
    Nick string
    Utime string
    Ctime string
}

var UserMeta = UserM{
    ID: "id",
    Username: "username",
    Pwd: "pwd",
    Salt: "salt",
    Nick: "nick",
    Utime: "utime",
    Ctime: "ctime",
}




func (m UserM) IdIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.ID).In(args...)
}

func (m UserM) IdNotIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.ID).NotIn(args...)
}

func (m UserM) IdEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).EQ(val)
}

func (m UserM) IdNotEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).NotEQ(val)
}

func (m UserM) IdLT(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).LT(val)
}

func (m UserM) IdLTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).LTEQ(val)
}

func (m UserM) IdGT(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).GT(val)
}

func (m UserM) IdGTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).GTEQ(val)
}

func (m UserM) IdLike(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).Like(val)
}

func (m UserM) IdNotLike(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).NotLike(val)
}

func (m UserM) IdDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.ID)
}

func (m UserM) IdAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.ID)
}



func (m UserM) UsernameIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Username).In(args...)
}

func (m UserM) UsernameNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Username).NotIn(args...)
}

func (m UserM) UsernameEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Username).EQ(val)
}

func (m UserM) UsernameNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Username).NotEQ(val)
}

func (m UserM) UsernameLT(val string) sqlbuilder.Column {
	return ql.Col(m.Username).LT(val)
}

func (m UserM) UsernameLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Username).LTEQ(val)
}

func (m UserM) UsernameGT(val string) sqlbuilder.Column {
	return ql.Col(m.Username).GT(val)
}

func (m UserM) UsernameGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Username).GTEQ(val)
}

func (m UserM) UsernameLike(val string) sqlbuilder.Column {
	return ql.Col(m.Username).Like(val)
}

func (m UserM) UsernameNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Username).NotLike(val)
}

func (m UserM) UsernameDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Username)
}

func (m UserM) UsernameAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Username)
}



func (m UserM) PwdIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Pwd).In(args...)
}

func (m UserM) PwdNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Pwd).NotIn(args...)
}

func (m UserM) PwdEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).EQ(val)
}

func (m UserM) PwdNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).NotEQ(val)
}

func (m UserM) PwdLT(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).LT(val)
}

func (m UserM) PwdLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).LTEQ(val)
}

func (m UserM) PwdGT(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).GT(val)
}

func (m UserM) PwdGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).GTEQ(val)
}

func (m UserM) PwdLike(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).Like(val)
}

func (m UserM) PwdNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).NotLike(val)
}

func (m UserM) PwdDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Pwd)
}

func (m UserM) PwdAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Pwd)
}



func (m UserM) SaltIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Salt).In(args...)
}

func (m UserM) SaltNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Salt).NotIn(args...)
}

func (m UserM) SaltEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).EQ(val)
}

func (m UserM) SaltNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).NotEQ(val)
}

func (m UserM) SaltLT(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).LT(val)
}

func (m UserM) SaltLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).LTEQ(val)
}

func (m UserM) SaltGT(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).GT(val)
}

func (m UserM) SaltGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).GTEQ(val)
}

func (m UserM) SaltLike(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).Like(val)
}

func (m UserM) SaltNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).NotLike(val)
}

func (m UserM) SaltDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Salt)
}

func (m UserM) SaltAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Salt)
}



func (m UserM) NickIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Nick).In(args...)
}

func (m UserM) NickNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Nick).NotIn(args...)
}

func (m UserM) NickEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Nick).EQ(val)
}

func (m UserM) NickNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Nick).NotEQ(val)
}

func (m UserM) NickLT(val string) sqlbuilder.Column {
	return ql.Col(m.Nick).LT(val)
}

func (m UserM) NickLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Nick).LTEQ(val)
}

func (m UserM) NickGT(val string) sqlbuilder.Column {
	return ql.Col(m.Nick).GT(val)
}

func (m UserM) NickGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Nick).GTEQ(val)
}

func (m UserM) NickLike(val string) sqlbuilder.Column {
	return ql.Col(m.Nick).Like(val)
}

func (m UserM) NickNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Nick).NotLike(val)
}

func (m UserM) NickDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Nick)
}

func (m UserM) NickAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Nick)
}



func (m UserM) UtimeIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Utime).In(args...)
}

func (m UserM) UtimeNotIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Utime).NotIn(args...)
}

func (m UserM) UtimeEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).EQ(val)
}

func (m UserM) UtimeNotEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).NotEQ(val)
}

func (m UserM) UtimeLT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).LT(val)
}

func (m UserM) UtimeLTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).LTEQ(val)
}

func (m UserM) UtimeGT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).GT(val)
}

func (m UserM) UtimeGTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).GTEQ(val)
}

func (m UserM) UtimeLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).Like(val)
}

func (m UserM) UtimeNotLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).NotLike(val)
}

func (m UserM) UtimeDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Utime)
}

func (m UserM) UtimeAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Utime)
}



func (m UserM) CtimeIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Ctime).In(args...)
}

func (m UserM) CtimeNotIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Ctime).NotIn(args...)
}

func (m UserM) CtimeEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).EQ(val)
}

func (m UserM) CtimeNotEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).NotEQ(val)
}

func (m UserM) CtimeLT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).LT(val)
}

func (m UserM) CtimeLTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).LTEQ(val)
}

func (m UserM) CtimeGT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).GT(val)
}

func (m UserM) CtimeGTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).GTEQ(val)
}

func (m UserM) CtimeLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).Like(val)
}

func (m UserM) CtimeNotLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).NotLike(val)
}

func (m UserM) CtimeDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Ctime)
}

func (m UserM) CtimeAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Ctime)
}
