'use strict';

var app = angular.module('indexCtrl', []);
app.controller('indexController', ['$scope', '$http', '$location', '$anchorScroll', function($scope, $http, $location, $anchorScroll) {

  $scope.form = {};
  $scope.textSearch;
  $scope.levels = [0, 1, 2, 3, 4];
  $scope.powers = [1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 11000, 12000, 13000, 14000, 15000];
  $scope.costs = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];

  /**
   * トップページの場合はtrueを返す
   */
  $scope.isTopPage = function() {
    return $location.path() === '/';
  };

  $scope.handleKeydown = function(e) {
    if (e.which === 13) {
      $scope.textSearch = $scope.searchText;
      setTimeout(function(){

      $scope.search();
      },100);
    }
  };

  $scope.searchEvent = function() {
    $scope.textSearch = $scope.searchText;
    $scope.search();
  };

  $scope.search = function() {
    if ($scope.form) {
      if ($scope.form.constraintTemp) {
        $scope.form.constraint = $scope.form.constraintTemp.Type;
      }
      if ($scope.form.illusTemp) {
        $scope.form.illus = $scope.form.illusTemp.Name;
      }
      if ($scope.form.expansionTemp) {
        $scope.form.expansion = $scope.form.expansionTemp.Id;
      }
      if ($scope.form.typeTemp) {
        $scope.form.type = $scope.form.typeTemp.Name;
      }
    }
    $http({
      method: 'GET',
      url: '/api/search',
      params: $scope.form
    }).success(function(data, status, headers, config) {
      for (var i in data) {
        var cost = [];
        setCost(cost, data[i].CostWhite, '白');
        setCost(cost, data[i].CostRed, '赤');
        setCost(cost, data[i].CostBlue, '青');
        setCost(cost, data[i].CostGreen, '緑');
        setCost(cost, data[i].CostBlack, '黒');
        setCost(cost, data[i].CostColorless, '無');
        data[i].Cost = cost;
      }
      $scope.list = data;
      $location.hash('search');
      $anchorScroll();
      $location.hash('');
    });

    var setCost = function(cost, color, str) {
      for (var i = 0; i < color; i++) {
        cost.push(str);
      }
    };
  };

  $scope.setColor = function(color) {
    var result = 'color:';
    switch (color){
    case 'white':
      result = '#f0bf00'; break;
    case 'red':
      result = '#E92C1C'; break;
    case 'blue':
      result = '#1C8FE9'; break;
    case 'green':
      result = '#11BB02'; break;
    case 'black':
      result = '#921197'; break;
    case 'colorless':
      result = '#8B8B8B'; break;
    default:
      result = ''; break;
    }
    return "{color: '" + result + "'}";
  };

  $scope.reset = function() {
    $scope.form = {};
  };

}]);