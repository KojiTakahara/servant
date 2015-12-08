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
    $scope.deckForm.lrigWhite = undefined;
    $scope.deckForm.lrigRed = undefined;
    $scope.deckForm.lrigBlue = undefined;
    $scope.deckForm.lrigGreen = undefined;
    $scope.deckForm.lrigBlack = undefined;
    $scope.deckForm.mainWhite = undefined;
    $scope.deckForm.mainRed = undefined;
    $scope.deckForm.mainBlue = undefined;
    $scope.deckForm.mainGreen = undefined;
    $scope.deckForm.mainBlack = undefined;
  };
  $scope.deckSearch = function() {
    $scope.deckForm.offset = 0;
    $scope.decks = [];
    $scope.end = false;
    search();
  };
  var search = function() {
    var form = angular.copy($scope.deckForm);
    form.lrigWhite = form.lrigWhite === false ? undefined : form.lrigWhite;
    form.lrigRed = form.lrigRed === false ? undefined : form.lrigRed;
    form.lrigBlue = form.lrigBlue === false ? undefined : form.lrigBlue;
    form.lrigGreen = form.lrigGreen === false ? undefined : form.lrigGreen;
    form.lrigBlack = form.lrigBlack === false ? undefined : form.lrigBlack;
    form.mainWhite = form.mainWhite === false ? undefined : form.mainWhite;
    form.mainRed = form.mainRed === false ? undefined : form.mainRed;
    form.mainBlue = form.mainBlue === false ? undefined : form.mainBlue;
    form.mainGreen = form.mainGreen === false ? undefined : form.mainGreen;
    form.mainBlack = form.mainBlack === false ? undefined : form.mainBlack;
    cardService.getPublicDecks(form).then(function(data) {
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