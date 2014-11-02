'use strict';

var ctrl = angular.module('cardCtrl', [
  'apiService',
  'wxCardDirective',
  'angularUtils.directives.dirPagination'
]);

/** カードTOP **/
ctrl.controller('cardController', ['$scope', '$location', 'cardService', function($scope, $location, cardService) {
  $scope.categories = ['ルリグ', 'アーツ', 'シグニ', 'スペル'];
  $scope.realities = ['LR', 'LC', 'SR', 'R', 'C', 'ST', 'PR'];
  $scope.levels = [0, 1, 2, 3, 4];
  $scope.powers = [1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 11000, 12000, 13000, 14000, 15000];
  $scope.costs = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];

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
    }, 500);
  };
  init();

  $scope.reset = function() {
    $scope.form = {};
  };
}]);

/** エキスパンションリスト **/
ctrl.controller('cardExController', ['$scope', '$stateParams', 'cardService', function($scope, $stateParams, cardService) {
  cardService.getCardByExpansion($stateParams.expansion).then(function(data) {
    $scope.cardList = data;
  });
}]);

/** カード詳細 **/
ctrl.controller('cardDetailController', ['$scope', '$stateParams', 'cardService', function($scope, $stateParams, cardService) {
  cardService.getCardByExpansionAndNo($stateParams.expansion, $stateParams.no).then(function(data) {
    $scope.card = data;
  });
}]);
