function routeUser($routeProvider) {
    $routeProvider.when("/new/user",{
    	"templateUrl": "html/user/new.html",
    	"controller": UserCreateController
    }).when("/userManage",{
    	"templateUrl": "html/user/list.html",
    	"controller": UserController
    }).when("/user/getUserDetail/:userName",{
    	"templateUrl":"html/user/detail.html",
    	"controller":UserDetailController
    });
}

 


function UserController($scope, $routeParams, $http, $location, $route, Popup){

    $http.get("api/users").success(function (data) {
        $scope.users = data.value;
       var users = data.value;
        users = getRoleValue($scope,users)
        handleData($scope,users);
    });


    $scope.create = function() {
        $location.path("/new/user");
        $route.reload();
    }

        $scope.delete = function(name) {
        if("" == name || null == name){
    		Popup.notice('User Name can not be empty! ', 3000 , function () {
            });
    		return ;
    	}
    
        
        $http.delete('api/users/' + name).success(function(data){
                $location.path("/userManage");
                $route.reload();
        });
         
       
    }
   
    //查询用户信息详情 
    $scope.queryUserDetail = function(name) {
    	if("" == name || null == name){
    		Popup.notice('User Name can not be empty! ', 3000 , function () {
            });
    		return ;
    	}
        	$location.path("/user/getUserDetail/"+name);
    }
}

/**
 *  * 查询、展现某一个用户的详细信息
 *   * 
 *    * @param $scope
 *     * @param $routeParams
 *      * @param $http
 *       * @param $location
 *        * @param $route
 *         * @returns
 *          
 *          */
function UserDetailController($scope, $routeParams, $http, $location, $route){
    $http.get("api/users/detail/" + $routeParams.userName).success(function (data) {
        user = data.value;
        var users = new Array();
    	users.push(user);
    	getRoleValue($scope,users)
        $scope.user = users[0];
    });
  
        
    $scope.close = function(){
    	$location.path("/userManage");
    }
}

/**
 *  * 分页
 *   * @returns
 *    */
function handleData($scope,users){
	 
	$scope.data = users;
	 
	$scope.pageSize = 5;
	$scope.pages = Math.ceil($scope.data.length / $scope.pageSize); //分页数
	$scope.newPages = $scope.pages > 5 ? 5 : $scope.pages;
	$scope.pageList = [];
	$scope.selPage = 1;
	 
	$scope.setData = function () {
	$scope.users = $scope.data.slice(($scope.pageSize * ($scope.selPage - 1)), ($scope.selPage * $scope.pageSize)); 
	}
	$scope.users = $scope.data.slice(0, $scope.pageSize);
	 
	for (var i = 0; i < $scope.newPages; i++) {
	$scope.pageList.push(i + 1);
	}
	 
	$scope.selectPage = function (page) {
	 
	if (page < 1 || page > $scope.pages) return;
	 
	if (page > 2) {
	 
	var newpageList = [];
	for (var i = (page - 3) ; i < ((page + 2) > $scope.pages ? $scope.pages : (page + 2)) ; i++) {
	newpageList.push(i + 1);
	}
	$scope.pageList = newpageList;
	}
	$scope.selPage = page;
	$scope.setData();
	$scope.isActivePage(page);
	console.log("选择的页：" + page);
	};
	 
	$scope.isActivePage = function (page) {
	return $scope.selPage == page;
	};
 
	$scope.Previous = function () {
	$scope.selectPage($scope.selPage - 1);
	}
 
	$scope.Next = function () {
	$scope.selectPage($scope.selPage + 1);
}
}

function getRoleValue($scope,users){
	
        $scope.roles = [
		{key : "0", value : "Alliance Members"},
	        {key : "1", value : "Manage Members"}
	               ]; 
        for ( var i=0;i<users.length;i++) {
			if(users[i].rolecode == $scope.roles[0].key){
				users[i].rolevalue = $scope.roles[0].value
			}else{
				users[i].rolevalue = $scope.roles[1].value
			}
		}
	 return users;
}


function UserCreateController($scope, $routeParams, $http, $location, $route,Popup) {
    	$scope.roles = [
		{key : "0", value : "Consortium Members"},
                {key : "1", value : "User Members"}
                   ];
        $scope.newroles = $scope.roles[1].key;	
		$scope.add = function() {
	    	var name = $scope.userName;
	    	if("" == name || null == name){
	    		Popup.notice('User Name can not be empty!',3000, function () {
                    });
                    return ;
	    	}
	        var c = {
	            "username": $scope.userName,
	            "password": $scope.passWord,
	            "rolecode": $scope.newroles,
	        };
	        $http.post('api/newUser', c).success(function(data){
	            $location.path("/userManage");
	            $route.reload();
	        });
                }
}
