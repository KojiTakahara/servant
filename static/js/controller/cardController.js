'use strict';

angular.module('cardCtrl', ['apiService'])
.controller('cardController', ['$scope', '$http', '$location', 'cardService', function($scope, $http, $location, cardService) {
  $scope.categories = ['ルリグ', 'アーツ', 'シグニ', 'スペル'];
  $scope.realities = ['LR', 'LC', 'SR', 'R', 'C', 'ST', 'PR'];
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

    var form = {expansion:'WX03'};
    cardService.search(form).then(function(data) {
      $scope.cardList = data;
    });

    $("input").iCheck({
      checkboxClass: "icheckbox_square-yellow", //使用するテーマのスキンを指定する
      radioClass: "iradio_square-yellow" //使用するテーマのスキンを指定する
    });
    setTimeout(function() {
      $("select").selecter();
      $("select.hoge").selecter({customClass: 'fuga'});
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
  };



}]);