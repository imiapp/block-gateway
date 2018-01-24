angular.module('status', []).filter("status", function () {
    return function (input) {
        return input == 1 ? "UP" : "DOWN";
    };
});

/**
 *  gateway
 *
 *  Description
 */
var app = angular.module('gateway', ["status",'angular-popups']);
app.directive('jsonText', function () {
    return {
        restrict: 'A',
        require: 'ngModel',
        link: function (scope, element, attr, ngModel) {
            function into(input) {
                return JSON.parse(input);
            }
            function out(data) {
                return JSON.stringify(data);
            }
            ngModel.$parsers.push(into);
            ngModel.$formatters.push(out);
        }
    };
});

app.config(['$routeProvider', route]);
app.config(function (PopupProvider) {
	PopupProvider.title = 'Prompt';
	PopupProvider.okValue = 'Confirm';
	PopupProvider.cancelValue = 'Cancel';
});


// 
app.filter('textLengthSet', function() {
    return function(value, wordwise, max, tail) {
        if (!value) return '';

        max = parseInt(max, 10);
        if (!max) return value;
        if (value.length <= max) return value;

        value = value.substr(0, max);
        if (wordwise) {
        var lastspace = value.lastIndexOf(' ');
        if (lastspace != -1) {
        value = value.substr(0, lastspace);
        }
        }
    return value + (tail || ' …');//'...'可以换成其它文字
    };
});


function route($routeProvider) {
    routeServer($routeProvider);
    routeRouting($routeProvider);
    routeAPI($routeProvider);
    routeCluster($routeProvider);
    routeDashboard($routeProvider);
    routeProxy($routeProvider);
    routeUser($routeProvider);
}
