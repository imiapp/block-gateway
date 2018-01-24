package server

import (
        "net/http"

        "github.com/fagongzi/log"
        "github.com/labstack/echo"
        "github.com/fagongzi/gateway/pkg/model"
)

func (server *AdminServer) newUser() echo.HandlerFunc{
	  log.Infof("test","test_hello_world")
	 return func(c echo.Context) error {
		var errstr string
		code := CodeSuccess
		user, err := model.UnMarshalUserFromReader(c.Request().Body())
		if nil != err {
			errstr = err.Error()
			code = CodeError
		} else {
			userFrDb, err := server.store.GetUser(user.Name)
			if nil != err {
				errstr = err.Error()
				code = CodeError
			}else{
				    if(userFrDb != nil && user.Name == userFrDb.Name){
	                  errstr = "The UserName had been registed!"
		                code = CodeError
                    }else{
						err := server.store.SaveUser(user)
						if nil != err {
							errstr = err.Error()
							code = CodeError
						}
                    }
			}
		}
		return c.JSON(http.StatusOK, &Result{
			Code:  code,
			Error: errstr,
	})
}
}



func (server *AdminServer) getUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		var errstr string
		code := CodeSuccess

		users, err := server.store.GetUsers()
		log.Info("==api_user.go","users",users)
		if err != nil {
			errstr = err.Error()
			code = CodeError
		}

		return c.JSON(http.StatusOK, &Result{
			Code:  code,
			Error: errstr,
			Value: users,
		})
	}
}

func (server *AdminServer) deleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var errstr string
		code := CodeSuccess

		name := c.Param("name")
		err := server.store.DeleteUser(name)

		if nil != err {
			errstr = err.Error()
			code = CodeError
		}

		return c.JSON(http.StatusOK, &Result{
			Code:  code,
			Error: errstr,
		})
	}
}


/**
*查询某个用户的详细信息 TODO
*/
func (server *AdminServer) getUserDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var errstr string
		code := CodeSuccess
    name := c.Param("name")
    userFrEtcd, err := server.store.GetUser(name)
		log.Info("==api_user.go","userFrEtcd",userFrEtcd)
		if nil != err {
			errstr = err.Error()
			code = CodeError
		}

		return c.JSON(http.StatusOK, &Result{
			Code:  code,
			Error: errstr,
			Value: userFrEtcd,
		})
	}
}

