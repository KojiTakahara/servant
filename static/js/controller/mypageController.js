'use strict';

var app = angular.module('mypageCtrl', []);

app.controller('editDeckController', ['$scope', '$location', function($scope, $location) {
  $scope.cardNums = [1, 2, 3, 4];
  $scope.deck = {
    cards: []
  };
  for (var i = 0; i < 4; i++) {
    $scope.deck.cards.push({
      Name: 'カード名',
      Category: 'ルリグ',
      Color: 'red',
      Level: 4
    });
  }
}]);

app.controller('mypageController', ['$scope', '$location', function($scope, $location) {
  $scope.decks = [];

  $scope.editDeck = function() {
    alert('編集処理を実装してね');
  };

  $scope.deleteDeck = function(index) {
    alert('削除処理を実装してね ' + index);
  };

  $scope.createDeck = function() {
    $location.path('/mypage/deck/0');
  };

  var init = function() {
    for (var i = 0; i < 10; i++) {
      $scope.decks.push({
        Title: 'テストデッキ' + i,
        Id: 1000 + i,
        White: true,
        Red: true,
        Blue: true,
        Green: true,
        Black: true,
        Scope: 'SELECT',
        UpdatedAt: new Date()
      });
    }
  };
  init();

}]);