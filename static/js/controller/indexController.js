'use strict';

var app = angular.module('indexCtrl', []);
app.controller('indexController', ['$scope', '$window', function($scope, $window) {

  /**
   * ログイン
   */
  $scope.login = function() {
    $window.location.href = '/api/twitter/login';
  };

}]);