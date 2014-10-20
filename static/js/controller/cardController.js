'use strict';

angular.module('cardCtrl', ['apiService', 'selecterForOptionWithNgRepeat', 'angularUtils.directives.dirPagination'])
.controller('cardController', ['$scope', '$http', '$location', 'cardService', function($scope, $http, $location, cardService) {
  $scope.categories = ['ルリグ', 'アーツ', 'シグニ', 'スペル'];
  $scope.realities = ['LR', 'LC', 'SR', 'R', 'C', 'ST', 'PR'];
  $scope.levels = [0, 1, 2, 3, 4];
  $scope.powers = [1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 11000, 12000, 13000, 14000, 15000];
  $scope.costs = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];


  $scope.setSelectedItem = function(value, index) {
    console.log(value);
    console.log(index);
  };

  var init = function() {
    cardService.getIllustrator().then(function(data) {
      $scope.illustrators = data;
    }, function() {
      // error
    });

    cardService.getConstraint().then(function(data) {
      $scope.constraints = data;
    }, function() {
      // error
    });

    cardService.getProduct().then(function(data) {
      $scope.products = data;
    }, function() {
      // error
    });

    cardService.getType().then(function(data) {
      $scope.types = data;
    }, function() {
      // error
    });

    var form = {expansion:'WX01'};
    cardService.search(form).then(function(data) {
      $scope.cardList = data;
    });

    setTimeout(function() {
      $("input").iCheck({
        checkboxClass: "icheckbox_square-yellow", //使用するテーマのスキンを指定する
        radioClass: "iradio_square-yellow" //使用するテーマのスキンを指定する
      });
      //$("select.hoge").selecter({customClass: 'fuga'});
      //$("select.piyo").selecter();
    }, 500);
  };
  init();


  $scope.setColor = function(color) {
    var result = '';
    switch (color){
    case 'white':
      result = '#F6BB42'; break;
    case 'red':
      result = '#DA4453'; break;
    case 'blue':
      result = '#3BAFDA'; break;
    case 'green':
      result = '#8CC152'; break;
    case 'black':
      result = '#967ADC'; break;
    case 'colorless':
      result = '#E6E9ED'; break;
    default:
      result = ''; break;
    }
    return "{backgroundColor: '" + result + "', borderColor: '" + result + "', height: '8px' }";
  };

  $scope.reset = function() {
    $scope.form = {};
    console.log($("select.hoge").length);
    setTimeout(function() {
      //$("select.hoge").selecter("destroy").selecter({customClass: 'fuga'});
      //$("select.piyo").selecter("destroy").selecter();
    }, 10);
  };



}]);