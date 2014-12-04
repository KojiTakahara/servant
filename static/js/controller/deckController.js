'use strict';

var app = angular.module('deckCtrl', [
  'apiService',
  'wxCardDirective',
  'infinite-scroll',
]);
app.controller('deckController', ['$scope', '$timeout', 'cardService', function($scope, $timeout, cardService) {
  $scope.decks = [];
  $scope.end = false;
  $scope.deckForm = {
    limit: 100,
    offset: 0
  };
  $scope.more = function() {
    $scope.deckForm.offset = $scope.deckForm.offset + 100;
    search();
  };
  $scope.deckFormReset = function() {
    $scope.deckForm.white = undefined;
    $scope.deckForm.red = undefined;
    $scope.deckForm.blue = undefined;
    $scope.deckForm.green = undefined;
    $scope.deckForm.black = undefined;
  };
  $scope.deckSearch = function() {
    $scope.deckForm.offset = 0;
    $scope.decks = [];
    $scope.end = false;
    search();
  };
  var search = function() {
    cardService.getPublicDecks($scope.deckForm).then(function(data) {
      if (data.length !== 0) {
        $scope.decks = $scope.decks.concat(data);
      } else {
        $scope.end = true;
      }
    });
  };
  var init = function() {
    search();
  };
  init();
}]);
app.controller('deckDetailController', ['$scope', '$stateParams', 'cardService', function($scope, $stateParams, cardService) {
  $scope.deck = null;
  $scope.alerts = [];
  $scope.closeAlert = function(index) {
    $scope.alerts.splice(index, 1);
  };
  $scope.sort = function(list, predicate) {
    if (list === 'lrig') {
      $scope.deck.lrigPredicate = predicate;
      $scope.deck.lrigReverse = !$scope.deck.lrigReverse;
    }
    if (list === 'main') {
      $scope.deck.mainPredicate = predicate;
      $scope.deck.mainReverse = !$scope.deck.mainReverse;
    }
  };
  var init = function() {
    cardService.getDeckById($stateParams.id).then(function(res) {
      if ((!res.Lrig && !res.Main) || res.Scope === 'PRIVATE') {
        $scope.alerts.push({ type: 'danger', msg: 'データが存在しないか、デッキが公開状態ではありません。' });
        $scope.error = true;
      } else {
        $scope.deck = res;
      }
    }, function(err) {
      $scope.alerts.push({ type: 'danger', msg: 'データが存在しないか、デッキが公開状態ではありません。' });
      $scope.error = true;
    });
  };
  init();

}]);