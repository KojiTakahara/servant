'use strict';

var ctrl = angular.module('cardCtrl', [
  'apiService',
  'wxCardDirective',
  'angularUtils.directives.dirPagination'
]);

/** カードTOP **/
ctrl.controller('cardController', ['$rootScope', '$scope', '$location', 'cardService', function($rootScope, $scope, $location, cardService) {
}]);

/** 検索 **/
ctrl.controller('cardSearchController', ['$rootScope', '$scope', '$stateParams', '$anchorScroll', 'cardService', function($rootScope, $scope, $stateParams, $anchorScroll, cardService) {
  $anchorScroll.yOffset = 0;
  $anchorScroll();
  $scope.cardSearch = function() {
    if (!$rootScope.searchCondition) {
      $anchorScroll.yOffset = 0;
      $anchorScroll();
      $scope.cardList = [];
      return;
    }
    cardService.search($rootScope.searchCondition).then(function(data) {
      if ($rootScope.searchCondition) {
        $scope.textSearch = $rootScope.searchCondition.text;
      }
      $anchorScroll.yOffset = 0;
      $anchorScroll();
      $scope.cardList = data;
    });
  };
  $scope.cardSearch();
}]);

/** エキスパンションリスト **/
ctrl.controller('cardExController', ['$rootScope', '$scope', '$stateParams', '$anchorScroll', 'cardService', function($rootScope, $scope, $stateParams, $anchorScroll, cardService) {
  $anchorScroll.yOffset = 0;
  $anchorScroll();
  cardService.getCardByExpansion($stateParams.expansion).then(function(data) {
    $scope.cardList = data;
  });
}]);

/** カード詳細 **/
ctrl.controller('cardDetailController', ['$rootScope', '$scope', '$stateParams', '$anchorScroll', 'cardService', function($rootScope, $scope, $stateParams, $anchorScroll, cardService) {
  $anchorScroll.yOffset = 0;
  $anchorScroll();
  cardService.getCardByExpansionAndNo($stateParams.expansion, $stateParams.no).then(function(data) {
    $scope.card = data;
  });
}]);
